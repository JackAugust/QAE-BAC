package EA

type Queue struct {
	Queue   []interface{}
	MaxSize int
	Len     int
	Cur     int
	Head    int
}

func NewQueue(maxSize ...int) *Queue {
	size := 2
	if len(maxSize) != 0 {
		size = maxSize[0]
	}
	return &Queue{
		Queue:   make([]interface{}, size),
		MaxSize: size,
		Len:     0,
		Cur:     0,
		Head:    0,
	}
}

// 入队
// Enqueue 循环添加数据
func (q *Queue) Enqueue(record interface{}) {
	if q.IsFull() {
		// panic("队列已满")
		// 队列已满时，从头开始覆盖
		q.Head = (q.Head + 1) % q.MaxSize
		q.Cur = q.Cur % q.MaxSize
		q.Len--
	}
	q.Queue[q.Cur] = record
	q.Cur = (q.Cur + 1) % q.MaxSize
	// q.Len = (q.Len + 1) % q.MaxSize
	q.Len++
}

// 出队
// func (q *Queue) Dequeue() interface{} {
// 	if q.IsEmpty() {
// 		return nil
// 	}
// 	record := q.Queue[q.Head]
// 	q.Queue[q.Head] = nil
// 	fmt.Println("record is ", record)
// 	q.Head = (q.Head + 1) % q.MaxSize
// 	q.Len = (q.Len - 1) % q.MaxSize
// 	return record
// }

// 测试的时候用这个
// 获取队头操作
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	record := q.Queue[q.Head]
	q.Head = (q.Head + 1) % q.Len
	return record
}

// 判断是否为空
func (q *Queue) IsEmpty() bool {
	return q.Len == 0
}
func (q *Queue) IsFull() bool {
	return q.Len == q.MaxSize
}
