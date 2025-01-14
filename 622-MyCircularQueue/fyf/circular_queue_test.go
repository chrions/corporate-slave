package fyf

import "testing"

func TestQueue(t *testing.T) {
	c := Constructor(4)
	t.Log(c.EnQueue(1))
	t.Log(c.EnQueue(2))
	t.Log(c.EnQueue(3))
	t.Log(c.EnQueue(4))
	t.Log(c.Rear())
	t.Log(c.IsFull())
	t.Log(c.DeQueue())
	t.Log(c.EnQueue(5))
	t.Log(c.Rear())
}
