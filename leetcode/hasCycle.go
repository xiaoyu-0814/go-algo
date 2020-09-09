package leetcode

//141. 环形链表
/*
给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。



示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。


示例 2：

输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。


示例 3：

输入：head = [1], pos = -1
输出：false
解释：链表中没有环。
*/

//1：利用map存储遍历过的节点，如果在遍历的同时在map取出相同的节点，则为true
//空间复杂度为O(n)
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	hashMap := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := hashMap[head]; ok {
			return true
		}
		hashMap[head] = struct{}{}
		head = head.Next
	}

	return false
}

//快慢指针
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
