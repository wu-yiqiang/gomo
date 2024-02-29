package strategy

// 定义 LRU 结构体
type Cache struct {
	// 缓存的当前容量
	Size int
	// 缓存的容量限制
	Limit int
	// 缓存数据的 map
	CacheMap map[string]*DoubleLinkList
	// 定义头、尾指针
	Head, Tail *DoubleLinkList
}

// 定义双链表
type DoubleLinkList struct {
	Key       string
	Value     int
	Next, Pre *DoubleLinkList
}

// New is the Constructor of Cache
func New() *Cache {
	return &Cache{
		Size:     1,
		Limit:    5,
		CacheMap: map[string]*DoubleLinkList{"1": &DoubleLinkList{Key: "1", Value: 1, Pre: nil, Next: nil}},
		Head:     nil,
		Tail:     nil,
	}
}
