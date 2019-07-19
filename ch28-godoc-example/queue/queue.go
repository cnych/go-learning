package queue

// 一个队列，可以存放任何类型的数据
type Queue []interface{}

// 往队列当中添加一个元素
func (q *Queue) Push(n interface{}) {
	*q = append(*q, n)
}

// 删除队列当中的第一个元素
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}
