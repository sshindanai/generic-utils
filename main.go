package main

import (
	"fmt"

	"github.com/sshindanai/generic-utils/gutils"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	gutils.ForEach(arr, func(v int) {
		fmt.Println(v)
	})
}
