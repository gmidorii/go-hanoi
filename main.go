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

var count int

func main() {
	fmt.Println("left -> right pool")

	// 3
	initTower(towers, 3)
	fmt.Println("Init: 3")
	printTower(towers)
	//towerCalc(towers, towers[left].Len())

	// 4
	initTower(towers, 4)
	fmt.Println("Init: 4")
	printTower(towers)
	towerCalc(towers, towers[left].Len())

	// 5
	initTower(towers, 5)
	fmt.Println("Init: 5 ------------------------------------------")
	printTower(towers)
	towerCalc(towers, towers[left].Len())
}

func towerCalc2(towers []*list.List) {
	if towers[center].Len() == 0 && towers[right].Len() == 0 {
		if towers[left].Len()%2 == 0 {
			move(towers[left], towers[center])
		} else {
			move(towers[left], towers[right])
		}
	}

	if towers[center].Back().Value.(int) - towers[right].Back().Value.(int) == 1 {
		move(towers[right], towers[center])
	}
	if towers[right].Back().Value.(int) - towers[center].Back().Value.(int) == 1 {
		move(towers[center], towers[right])
	}

	if towers[center].Back().Value.(int) > towers[right].Back().Value.(int){
		move(towers[right], towers[left])
	}

}

func towerCalc(towers []*list.List, n int) {
	if n == towers[right].Len() || count > 32 {
		return
	}
	count++
	lenl := towers[left].Len()
	lenc := towers[center].Len()
	lenr := towers[right].Len()
	var len int
	var start *list.List
	var middle *list.List
	var goal = towers[right]
	if lenl > lenc && lenl > lenr {
		fmt.Println("s:LEFT")
		start = towers[left]
		len = lenl
		goal = towers[right]
		middle = towers[center]
	} else if lenc > lenr {
		fmt.Println("s:CENTER")
		start = towers[center]
		len = lenc
		goal = towers[right]
		middle = towers[left]
	} else {
		fmt.Println("s:RIGHT")
		start = towers[right]
		len = lenr
		goal = towers[center]
		middle = towers[left]
	}

	if len%2 == 0 {
		move(start, middle)
		move(start, goal)
		move(middle, goal)
		move(start, middle)
		move(goal, start)
		move(goal, middle)
		move(start, middle)
		move(start, goal)
	} else {
		move(start, goal)
		move(start, middle)
		move(goal, middle)
		move(start, goal)
		move(middle, start)
		move(middle, goal)
		move(start, goal)
	}

	if middle.Len() == 0 && towers[right].Len() == 0 {
		move(middle, towers[right])
	}
	towerCalc(towers, n)
}

func initTower(towers []*list.List, n int) {
	for towers[left].Len() != 0 {
		towers[left].Remove(towers[left].Back())
	}
	for towers[center].Len() != 0 {
		towers[center].Remove(towers[center].Back())
	}
	for towers[right].Len() != 0 {
		towers[right].Remove(towers[right].Back())
	}

	for n > 0 {
		towers[left].PushBack(n)
		n--
	}
}

func printTower(tower []*list.List) {
	fmt.Println("--------------------------------")
	fmt.Print("l: ")
	printOneTower(tower[left])
	fmt.Print("c: ")
	printOneTower(tower[center])
	fmt.Print("r: ")
	printOneTower(tower[right])
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

func move(po, pu *list.List) {
	n, err := pop(po)
	if err == nil {
		push(pu, n)
	}
	printTower(towers)
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

func threeTower() {
	//// 2
	//// 3 _ 1
	move(towers[left], towers[right])
	////
	//// 3 2 1
	move(towers[left], towers[center])
	////   1
	//// 3 2 _
	move(towers[right], towers[center])
	////   1
	//// _ 2 3
	move(towers[left], towers[right])
	////
	//// 1 2 3
	move(towers[center], towers[left])
	////     2
	//// 1 _ 3
	move(towers[center], towers[right])
	////     1
	////     2
	//// _ _ 3
	move(towers[left], towers[right])
	printTower(towers)
}
