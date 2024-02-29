package strategy

func (c *Cache) Get(key string) (value int, ok bool) {
	if node, ok := c.CacheMap[key]; ok {
		// 如果存在, 从缓存中获取该 key 的 value
		// 将该结点移动都 head 结点后面，head、tail 结点不存数据
		c.moveNodeToHead(node)
		return node.Value, true
	}
	return
}

func (c *Cache) moveNodeToHead(node *DoubleLinkList) {
	// 将节点从当前队列中清除
	no := c.removeNode(node)
	// 添加节点到当前队列中
	no.Next = c.Head
}

func (c *Cache) removeNode(node *DoubleLinkList) *DoubleLinkList {
	// 查找到当前的节点
	pre := node.Pre
	next := node.Next
	pre.Next = next
	// 生成一个新节点数据
	return &DoubleLinkList{
		Next:  nil,
		Pre:   nil,
		Value: node.Value,
		Key:   node.Key,
	}
}

func (c *Cache) Add(newNode *DoubleLinkList) {
	key := newNode.Key
	_, ok := c.CacheMap[key]
	if c.Size > c.Limit {
		// 删除最久未访问的节点
	}
	if ok {
		return
	} else {
		newNode.Next = c.Head
		c.Size++
	}
}

func (c *Cache) FormatePrint() {

}
