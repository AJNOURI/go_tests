package main

import (
	"crypto/tls"
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

// +build this,is,too,late

type T struct {
	A int `not a canonical tag`
}

func foo() []byte       { return []byte{} }
func bar(mu sync.Mutex) {} // Locks have to be passed as pointers

func main() {
	fmt.Sprintf("%s", 1) // Wrong verb
	fmt.Sprint("%s", 1)  // Wrong method

	var a int
	a = a // Useless assignment

	_ = tls.Certificate{nil, nil, nil, nil, nil} // These fields should be keyed

	if foo != nil { // We meant to check foo() != nil
	}

	x := uintptr(0)
	_ = unsafe.Pointer(x) // Misuse of unsafe.Pointer

	for _, x := range []int{1, 2, 3} {
		go func() {
			fmt.Println(x) // Closing over x wrong, will probably produce "3 3 3"
		}()
	}

	var b int32
	b = atomic.AddInt32(&b, 1) // We shouldn't assign to b

	var c bool
	if c == false && c == true { // Always false
	}

	if c || c { // Redundant
	}

	return
	fmt.Println("unreachable") // Unreachable code
}
