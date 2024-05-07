package main

import "fmt"

type Tree struct {
	v int
	l *Tree
	r *Tree
}

func NewTree(arr []int) *Tree {
	if len(arr) == 0 {
		return nil
	}

	root := &Tree{v: arr[0]}
	queue := []*Tree{root}

	i := 1
	for len(queue) > 0 && i < len(arr) {
		curr := queue[0]
		queue = queue[1:]

		if i < len(arr) && arr[i] != -1 {
			curr.l = &Tree{v: arr[i]}
			queue = append(queue, curr.l)
		}
		i++

		if i < len(arr) && arr[i] != -1 {
			curr.r = &Tree{v: arr[i]}
			queue = append(queue, curr.r)
		}
		i++
	}

	return root

}

func (t *Tree) Sum() int {
	num := t.v

	if t.l != nil {
		num += t.l.Sum()
	}

	if t.r != nil {
		num += t.r.Sum()
	}

	return num
}

func main() {
	// 创建一棵树：
	//       1
	//      / \
	//     2   3
	//    / \   \
	//   4   5   6
	arr := []int{1, 2, 3, 4, 5}
	root := NewTree(arr)

	// 调用 Sum() 方法并打印结果
	sum := root.Sum()
	fmt.Println("The sum of all tree nodes is:", sum)
}
