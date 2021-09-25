package imu

import (
	"sync"
)

type Item struct {
	ID   string
	Data interface{}
}

type MeetUpper interface {
	// apply .. returns sameID Item
	Apply(item Item) (interface{}, bool)
}

type meetUpper struct {
	poolIdx  int
	poolSize int
	pools    [3]pool
}

func NewMeetUpper(poolSize int) MeetUpper {
	meetUpper := meetUpper{
		pools: [3]pool{
			{maxCount: poolSize},
			{maxCount: poolSize},
			{maxCount: poolSize},
		},
		poolSize: poolSize,
	}
	return &meetUpper
}

type pool struct {
	count    int
	maxCount int
	maps     sync.Map
}

// apply .. returns sameID Item
func (imu *meetUpper) Apply(item Item) (interface{}, bool) {
	data, exist, isFull := imu.pools[imu.poolIdx].add(item)
	beforeData, beforeExist := imu.pools[(imu.poolIdx+2)%3].maps.Load(item.ID)
	if isFull {
		imu.poolIdx = (imu.poolIdx + 1) % 3
		// poolIdx = 0 ~ 2
		// 初めて 1になった時に2を初期化する
		imu.pools[(imu.poolIdx+1)%3] = pool{maxCount: imu.poolSize}
	}
	if exist {
		return data, true
	}
	if beforeExist {
		return beforeData, true
	}
	return nil, false
}

// add .. return is ( exist, isFull ).
func (p *pool) add(item Item) (interface{}, bool, bool) {
	data, exist := p.maps.Load(item.ID)
	if !exist {
		p.maps.Store(item.ID, item.Data)
		p.count++
	}
	return data, exist, p.count >= p.maxCount
}
