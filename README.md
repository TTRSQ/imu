# imu
id meet upper

## usage
```
package main

import (
	"fmt"

	"github.com/TTRSQ/imu"
)

func main() {
	im := imu.NewMeetUpper(2)

	// return true when id exist before.
	fmt.Println(im.Apply(imu.Item{
		ID:   "1",
		Data: nil,
	})) // false
	fmt.Println(im.Apply(imu.Item{
		ID:   "2",
		Data: nil,
	})) // false
	fmt.Println(im.Apply(imu.Item{
		ID:   "3",
		Data: nil,
	})) // false
	fmt.Println(im.Apply(imu.Item{
		ID:   "2",
		Data: nil,
	})) // true, and get nil
	fmt.Println(im.Apply(imu.Item{
		ID:   "4",
		Data: nil,
	})) // false

	// same id come after 2 * poolSize return false
	fmt.Println(im.Apply(imu.Item{
		ID:   "1",
		Data: "hogehoge",
	})) // false
	fmt.Println(im.Apply(imu.Item{
		ID:   "1",
		Data: nil,
	})) // true and get hogehoge
}

```