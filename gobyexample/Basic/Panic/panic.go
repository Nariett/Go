// package main

//	func main() {
//		panic("a problem")
//		_, err := os.Create("/tmp/file")
//		if err != nil {
//			panic(err)
//		}
//	}
package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {

	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
