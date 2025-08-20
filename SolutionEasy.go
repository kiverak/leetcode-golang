package main

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
