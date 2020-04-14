package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

func (node *ListNode) Print() {
	//	debug method
	var num string
	cur := node;
	for {
		digit := strconv.Itoa(cur.Val)
		num = digit + num
		if cur.Next != nil {
			cur = cur.Next
		} else {
			fmt.Println(num)
			break;
		}
	}
}

func GenerateNum(num string) *ListNode {
	byteArr := strings.Split(num, "");
	var head *ListNode
	var cur *ListNode = nil
	begin := len(byteArr) - 1
	for i := begin; i >= 0; i-- {
		value, _ := strconv.Atoi(byteArr[i])
		node := ListNode{
			Val:  value,
			Next: nil,
		}
		if cur != nil {
			cur.Next = &node
		}
		cur = &node
		if i == begin {
			head = &node
		}
	}
	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode = nil
	l1cur := l1
	l2cur := l2
	cur := res
	cf := 0
	for {
		if l1cur.Next == nil || l2cur.Next == nil {
			// 最后一位相加 ugly
			added := l1cur.Val + l2cur.Val + cf
			node := &ListNode{
				Val:  added % 10,
				Next: nil,
			}
			if cur == nil {
				cur = node
				res = node
			}else{
				cur.Next = node
			}


			cf = added / 10
			cur = node

			if l1cur.Next == nil && l2cur.Next == nil {
				if (cf == 1) { //overflow
					cur.Next = &ListNode{
						Val:  cf,
						Next: nil,
					}
				}
			} else {
				remain := l1cur
				if l1cur.Next == nil {
					remain = l2cur
				}
				// golang尾递归?
				last := addTwoNumbers(&ListNode{
					Val: cf,
					Next: nil,
				}, remain.Next)

				if last != nil {
					cur.Next = last
				}

			}
			break;
		} else {
			added := l1cur.Val + l2cur.Val + cf
			cf = added / 10
			node := &ListNode{
				Val:  added % 10,
				Next: nil,
			}
			if cur == nil { //head node
				res = node
			} else {
				cur.Next = node
			}
			cur = node
			l1cur = l1cur.Next
			l2cur = l2cur.Next
		}
	}
	return res;
}
func main() {
	l1 := GenerateNum("923")
	l2 := GenerateNum("1177")
	res := addTwoNumbers(l1, l2)
	res.Print()
}
