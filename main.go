package main

import (
	"container/list"
	"errors"
	"fmt"
	"strconv"
)

const left = 0
const center = 1
const right = 2

var towers = []*list.List{
	list.New(), // left tower
	list.New(), // center tower
	list.New(), // right tower
}

func main() {
	fmt.Println("left -> right pool")
	initTower(towers[left], 3)
	printTower(towers[left], towers[center], towers[right])
	//// 2
	//// 3 _ 1
	popPush(towers[left], towers[right])
	////
	//// 3 2 1
	popPush(towers[left], towers[center])
	////   1
	//// 3 2 _
	popPush(towers[right], towers[center])
	////   1
	//// _ 2 3
	popPush(towers[left], towers[right])
	////
	//// 1 2 3
	popPush(towers[center], towers[left])
	////     2
	//// 1 _ 3
	popPush(towers[center], towers[right])
	////     1
	////     2
	//// _ _ 3
	popPush(towers[left], towers[right])

	printTower(towers[left], towers[center], towers[right])
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
	n, err := pop(po)
	if err == nil {
		push(pu, n)
	}
}

func pop(l *list.List) (int, error) {
	if l.Len() == 0 {
		return 0, errors.New("towers broken")
	}
	n := l.Remove(l.Back())
	t, ok := n.(int)
	if ok != false {
		return t, nil
	}
	return 0, errors.New("cannot convert int")
}

func push(l *list.List, n int) {
	l.PushBack(n)
}
