package main

import (
	"fmt"
	"io"
	"os"
)

/* Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
Defer is commonly used to simplify functions that perform various clean-up actions.
*/
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
