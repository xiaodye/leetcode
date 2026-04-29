package leetcode

// LRUCache 结构体
type LRUCache struct {
	capacity int
	cache    map[int]*DLinkedNode // 哈希表：key -> 链表节点
	head     *DLinkedNode         // 链表头（虚拟头节点，最久未使用在 head.Next）
	tail     *DLinkedNode         // 链表尾（虚拟尾节点，最近使用在 tail.Prev）
}

type DLinkedNode struct {
	key  int
	val  int
	prev *DLinkedNode
	next *DLinkedNode
}

// Constructor 初始化 LRU 缓存
func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

// 将一个节点移动到链表末尾（最近使用）
func (lru *LRUCache) moveToTail(node *DLinkedNode) {
	// 先从原位置移除
	node.prev.next = node.next
	node.next.prev = node.prev
	// 插入到 tail 之前
	node.prev = lru.tail.prev
	node.next = lru.tail
	lru.tail.prev.next = node
	lru.tail.prev = node
}

// 在尾部添加新节点（最近使用）
func (lru *LRUCache) addToTail(node *DLinkedNode) {
	node.prev = lru.tail.prev
	node.next = lru.tail
	lru.tail.prev.next = node
	lru.tail.prev = node
}

// 移除头部节点（最久未使用）
func (lru *LRUCache) removeHead() *DLinkedNode {
	if lru.head.next == lru.tail {
		return nil
	}
	node := lru.head.next
	lru.head.next = node.next
	node.next.prev = lru.head
	return node
}

// Get 获取值，并将该键移到最近使用
func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		lru.moveToTail(node)
		return node.val
	}
	return -1
}

// Put 插入或更新键值对
func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.cache[key]; ok {
		// 已存在：更新值并移到尾部
		node.val = value
		lru.moveToTail(node)
		return
	}
	// 新节点
	newNode := &DLinkedNode{key: key, val: value}
	lru.cache[key] = newNode
	lru.addToTail(newNode)
	if len(lru.cache) > lru.capacity {
		// 删除头节点（最久未使用）
		headNode := lru.removeHead()
		delete(lru.cache, headNode.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
