package main

import "fmt"

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	dummy := &ListNode{} // 用哨兵节点简化代码逻辑
	node := dummy        // node 指向新链表的末尾

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			node.Next = list2
			list2 = list2.Next
		} else {
			node.Next = list1
			list1 = list1.Next
		}
		node = node.Next
	}
	if list1 != nil {
		node.Next = list1
	} else {
		node.Next = list2
	}
	return node
}

func main() {
	//node1 := &ListNode{
	//	Val:  4,
	//	Next: nil,
	//}
	node2 := &ListNode{Val: 3,
		Next: nil,
	}
	node3 := &ListNode{Val: 1,
		Next: node2,
	}

	//node4 := &ListNode{Val: 4,
	//	Next: nil,
	//}
	node5 := &ListNode{Val: 2,
		Next: nil,
	}
	node6 := &ListNode{Val: 1,
		Next: node5,
	}
	lists := mergeTwoLists(node6, node3)

	for lists.Next != nil {
		fmt.Println(lists)
		lists = lists.Next

	}

}
