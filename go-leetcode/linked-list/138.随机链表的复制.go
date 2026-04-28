package leetcode

// Node 链表节点定义
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 建立映射，源节点 -> 新节点
	m := make(map[*Node]*Node)

	// 复制各节点，并建立 “原节点 -> 新节点” 的 Map 映射
	curr := head
	for curr != nil {
		m[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	// 指针重置
	curr = head

	// 构建新链表的 next 和 random 指向
	for curr != nil {
		// 复制节点之间的前后关系，注意最后一个节点
		m[curr].Next = m[curr.Next] // 如果 curr.Next == nil，m[nil] 返回 nil（Go map 的零值）

		// 复制随机指针
		m[curr].Random = m[curr.Random]

		curr = curr.Next
	}

	// 返回新链表的头节点
	return m[head]
}
