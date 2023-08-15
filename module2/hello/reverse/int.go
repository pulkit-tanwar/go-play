package reverse

import (
	"fmt"
	"strconv"
)

// Int returns the decimal reversal of the integer i.
func Int(i int) int {
	ans, err := strconv.Atoi(String(strconv.Itoa(i)))
	if err != nil {
		fmt.Println("error: ", err)
	}
	return ans
}
