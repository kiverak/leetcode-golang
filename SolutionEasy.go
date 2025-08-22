package main

import "container/list"

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if j, found := indexMap[diff]; found {
			return []int{j, i}
		}
		indexMap[num] = i
	}
	return nil
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var stack []rune
	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			lastOpen := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (char == ')' && lastOpen != '(') ||
				(char == ']' && lastOpen != '[') ||
				(char == '}' && lastOpen != '{') {
				return false
			}
		}
	}
	return len(stack) == 0
}

func isValid2(s string) bool {
	stack := []rune{}
	hash := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range s {
		if match, found := hash[char]; found {
			if len(stack) > 0 && stack[len(stack)-1] == match {
				stack = stack[:len(stack)-1] // Pop
			} else {
				return false // Invalid
			}
		} else {
			stack = append(stack, char) // Push opening bracket
		}
	}
	return len(stack) == 0
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	op := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			op.Next = list1
			list1 = list1.Next
		} else {
			op.Next = list2
			list2 = list2.Next
		}
		op = op.Next
	}
	if list1 != nil {
		op.Next = list1
	}
	if list2 != nil {
		op.Next = list2
	}
	return dummy.Next
}

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	n1 := 1
	n2 := 2
	sum := 0

	for i := 3; i <= n; i++ {
		sum = n1 + n2
		n1 = n2
		n2 = sum
	}
	return sum
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)
	level := 0
	for queue.Len() > 0 {
		n := queue.Len()
		for i := 0; i < n; i++ {
			e := queue.Front()
			queue.Remove(e)
			node := e.Value.(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		level++
	}
	return level
}
