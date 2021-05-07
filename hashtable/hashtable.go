package hashtable

import (
	"fmt"
	"hash-tables-in-go/linkedlist"
	"hash/maphash"
)

const arraySize = 33

var hashFun maphash.Hash

type Hashtable struct {
	array [arraySize]*linkedlist.Linkedlist
}

func Init() *Hashtable {
	hashFun.SetSeed(maphash.MakeSeed())

	ht := &Hashtable{}
	for i := range ht.array {
		ht.array[i] = linkedlist.Init()
	}
	return ht
}

func hash(key string) int {
	hashFun.WriteString(key)
	hashed := hashFun.Sum64()
	hashFun.Reset()

	index := hashed % uint64(arraySize)
	return int(index)
}

func (h *Hashtable) Find(key string) (string, bool) {
	index := hash(key)
	return h.array[index].Search(key)
}

func (h *Hashtable) Put(key, value string) {
	if _, ok := h.Find(key); ok {
		fmt.Println("Key already exists")
		return
	}
	index := hash(key)
	h.array[index].Insert(key, value)
}

func (h *Hashtable) Delete(key string) {
	if _, ok := h.Find(key); !ok {
		fmt.Println("Key doesn't exist")
		return
	}
	index := hash(key)
	h.array[index].Delete(key)
}
