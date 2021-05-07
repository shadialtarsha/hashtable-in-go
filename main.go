package main

import (
	"fmt"
	"hash-tables-in-go/hashtable"
)

func main() {
	hashTable := hashtable.Init()
	hashTable.Put("1", "shadi")
	hashTable.Put("4", "phelep")
	hashTable.Put("10", "Jodel")

	if v, ok := hashTable.Find("10"); ok {
		fmt.Printf("The value is %s\n", v)
	} else {
		fmt.Println("Key not found")
	}

	hashTable.Delete("10")

	if v, ok := hashTable.Find("10"); ok {
		fmt.Printf("The value is %s\n", v)
	} else {
		fmt.Println("Key not found")
	}
}
