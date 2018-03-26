package main

import (
	"fmt"

	"bst"
)

func main() {
	//b := bst.New()
	//var b *bst.BSTree
	b := bst.New()
	fmt.Println(b)

	b.Put(1, 10)
	b.Put(5, "aa")
	b.Put(100, "aa")
	b.Put(5, "5")
	b.Put(4, "4")
	b.Put(22, "22")
	b.Put(42, "42")
	b.Put(92, "92")
	b.Put(30, "30")
	b.Put(21, "21")

	fmt.Println("min: ", b.Min())
	fmt.Println("max: ", b.Max())
	b.Delete(30)
	b.Traverse()
	// b.DeleteMin()
	// b.DeleteMax()
	// fmt.Println("min: ", b.Min())
	// fmt.Println("max: ", b.Max())
	// b.DeleteMax()
	// fmt.Println("size: ", b.Size())

	// // b.Delete(1)
	// b.Delete(1)
	// fmt.Println("min: ", b.Max().Key())
	// fmt.Println(b.Get(100), b.Size())

	// fmt.Println(b.Get(5))
	//fmt.Println(b.Min())
}
