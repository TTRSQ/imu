package imu

import (
	"sync"
)

type meetUpper struct {
	poolIdx  int
	poolSize int
	pools    [3]pool
}

func NewMeetUpper(poolSize int) meetUpper {
	meetUpper := meetUpper{
		pools: [3]pool{
			{maxCount: poolSize},
			{maxCount: poolSize},
			{maxCount: poolSize},
		},
		poolSize: poolSize,
	}
	return meetUpper
}

type pool struct {
	count    int
	maxCount int
	maps     sync.Map
}

// apply .. returns meets
func (imu *meetUpper) Apply(id string) bool {
	exist, isFull := imu.pools[imu.poolIdx].add(id)
	_, beforeExist := imu.pools[(imu.poolIdx+2)%3].maps.Load(id)
	if isFull {
		imu.poolIdx = (imu.poolIdx + 1) % 3
		// poolIdx = 0 ~ 2
		// 初めて 1になった時に2を初期化する
		imu.pools[(imu.poolIdx+1)%3] = pool{maxCount: imu.poolSize}
	}
	return exist || beforeExist
}

// add .. return is ( exist, isFull ).
func (p *pool) add(id string) (bool, bool) {
	_, exist := p.maps.Load(id)
	if !exist {
		p.maps.Store(id, true)
		p.count++
	}
	return exist, p.count == p.maxCount
}
