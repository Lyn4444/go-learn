package main

import "fmt"

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//
//
// 示例 2：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//
//
//
// 提示：
//
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
//
// Related Topics 递归 链表 数学 👍 10567 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode = &ListNode{Val: -1}

	head := res

	carry := 0

	for {
		var tmp *ListNode
		if l1 != nil || l2 != nil {
			if l1 != nil && l2 != nil {
				tmp = &ListNode{Val: (carry + l1.Val + l2.Val) % 10}
				carry = (carry + l1.Val + l2.Val) / 10
				l1 = l1.Next
				l2 = l2.Next
			} else if l1 != nil {
				tmp = &ListNode{Val: (carry + l1.Val) % 10}
				carry = (carry + l1.Val) / 10
				l1 = l1.Next
			} else {
				tmp = &ListNode{Val: (carry + l2.Val) % 10}
				carry = (carry + l2.Val) / 10
				l2 = l2.Next
			}
			head.Next = tmp
			head = head.Next
		} else {
			break
		}
	}

	if carry != 0 {
		head.Next = &ListNode{Val: carry}
	}

	return res.Next
}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}
	return head
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// 创建两个链表 l1 和 l2
	l1 := createList([]int{2, 4, 3})
	l2 := createList([]int{5, 6, 4})

	// 添加两个链表
	result := addTwoNumbers(l1, l2)

	// 打印结果链表
	printList(result)
}

//leetcode submit region end(Prohibit modification and deletion)
