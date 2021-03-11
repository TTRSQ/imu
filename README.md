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
	fmt.Println(im.Apply("1")) // false
	fmt.Println(im.Apply("2")) // false
	fmt.Println(im.Apply("3")) // false
	fmt.Println(im.Apply("2")) // true
	fmt.Println(im.Apply("4")) // false

	// same id come after 2 * poolSize return false
	fmt.Println(im.Apply("1")) // false
	fmt.Println(im.Apply("1")) // true
}

```