package main

import "github.com/tirava/hash-map/pkg/hashmap"

func main() {

	h := hashmap.NewHashMap()

	h.Set(5, "test5")
	h.Set(15, "test15")
	h.Del(5)
	h.Set(3, "test3")
	h.Set(3, "test3")
	h.Set(3, "test3")
	h.Set(3, "test3")
	h.Set(3, "test3")
	h.Set(3, "test30")
	h.Set(6, "test6")
	h.Set(7, "test7")
	h.Set(8, "test8")
	h.Set(9, "test9")
	h.Set(10, "test10")
	h.Set(11, "test11")
	h.Set(12, "test12")
	h.Set(13, "test13")

	//fmt.Println(h.Get(3))
	//fmt.Println(h.Get(15))
	//fmt.Println(h.Get(3))
	//fmt.Println(h.Get(0))

	//fmt.Println("===============")
	//h.Extend()

	//fmt.Println(h.Get(5))
	//fmt.Println(h.Get(15))
	//fmt.Println(h.Get(3))
	//fmt.Println(h.Get(0))

	h.Del(6)

}
