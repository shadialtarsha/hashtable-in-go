package linkedlist

import (
	"fmt"
	"strings"
)

type Linkedlist struct {
	head *node
}

func Init() *Linkedlist {
	return &Linkedlist{}
}

func (l *Linkedlist) Insert(key string, value string) *Linkedlist {
	n := node{key: key, value: value}
	n.nextNode = l.head
	l.head = &n
	return l
}

func (l *Linkedlist) Search(key string) (string, bool) {
	currentNode := l.head

	for currentNode != nil {
		if currentNode.key == key {
			return currentNode.value, true
		}
		currentNode = currentNode.nextNode
	}

	return "", false
}

func (l *Linkedlist) Delete(key string) *Linkedlist {
	if l.head.key == key {
		l.head = l.head.nextNode
		return l
	}

	previousNode := l.head
	for previousNode != nil && previousNode.nextNode != nil {
		if previousNode.nextNode.key == key {
			previousNode.nextNode = previousNode.nextNode.nextNode
		}
		previousNode = previousNode.nextNode
	}

	return l
}

func (l *Linkedlist) Print() {
	var sb strings.Builder
	n := l.head

	for n != nil {
		sb.WriteString(n.value)
		if n.nextNode != nil {
			sb.WriteString("->")
		}
		n = n.nextNode
	}

	fmt.Println(sb.String())
}
