package main

import (
	"container/list"
	"fmt"
	"strconv"
)


func main() {
	left := list.New()
	center := list.New()
	right := list.New()

	fmt.Println("left -> right pool")
	initTower(left, 3)
	printTower(left, center, right)
	//// 2
	//// 3 _ 1
	popPush(left, right)
	////
	//// 3 2 1
	popPush(left, center)
	////   1
	//// 3 2 _
	popPush(right, center)
	////   1
	//// _ 2 3
	popPush(left, right)
	////
	//// 1 2 3
	popPush(center, left)
	////     2
	//// 1 _ 3
	popPush(center, right)
	////     1
	////     2
	//// _ _ 3
	popPush(left, right)

	printTower(left, center, right)
}

func initTower(left *list.List, n int) {
	for n > 0 {
		left.PushBack(n)
		n--
	}
}

func printTower(left, center, right *list.List) {
	fmt.Println("--------------------------------")
	printOneTower(left)
	printOneTower(center)
	printOneTower(right)
	fmt.Println("--------------------------------")
}

func printOneTower(l *list.List) {
	var output string
	for e := l.Front(); e != nil; e = e.Next() {
		num := e.Value.(int)
		output += strconv.Itoa(num) + " "
	}
	fmt.Println(output)
}

func popPush(po, pu *list.List) {
	n := pop(po)
	if n == 0 {
		return
	}
	push(pu, n)
}

func pop(l *list.List) int {
	if l.Len() == 0 {
		return 0
	}
	n := l.Remove(l.Back())
	t, ok := n.(int)
	if ok != false {
		return t
	}
	return 0
}

func push(l *list.List, n int) {
	l.PushBack(n)
}
