package CodeTop

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getValues(head *ListNode) []int {
	if head == nil {
		return nil
	}
	fast, res := head, []int{}
	for fast != nil {
		res = append(res, fast.Val)
		fast = fast.Next
	}

	return res
}
