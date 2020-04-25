package strutil

import (
	"fmt"
	"strconv"
)

func String2Intv1(n int) string {
	return strconv.Itoa(n)
}

func String2Intv2(n int) string {
	return fmt.Sprintf("%d", n)
}
