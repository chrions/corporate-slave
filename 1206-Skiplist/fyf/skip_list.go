package fyf

import (
	"math"
	"math/rand"
)

const (
	//最高层数
	MAX_LEVEL = 16
)

//跳表节点结构体
type skipListNode struct {
	//跳表保存的值
	v interface{}
	//用于排序的分值
	score int
	//层高
	level int
	//每层前进指针
	forwards []*skipListNode
}

//新建跳表节点
func newSkipListNode(v interface{}, score, level int) *skipListNode {
	return &skipListNode{v: v, score: score, forwards: make([]*skipListNode, level, level), level: level}
}

//跳表结构体
type Skiplist struct {
	//跳表头结点
	head *skipListNode
	//跳表当前层数
	level int
	//跳表长度
	length int
}

//实例化跳表对象
func Constructor() Skiplist {
	//头结点，便于操作
	head := newSkipListNode(0, math.MinInt32, MAX_LEVEL)
	return Skiplist{head, 1, 0}
}

//插入节点到跳表中
func (this *Skiplist) Add(num int) {

	//查找插入位置
	cur := this.head
	//记录每层的路径
	update := [MAX_LEVEL]*skipListNode{}
	i := MAX_LEVEL - 1
	for ; i >= 0; i-- {
		for nil != cur.forwards[i] {

			if cur.forwards[i].score >= num {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if nil == cur.forwards[i] {
			update[i] = cur
		}
	}

	//通过随机算法获取该节点层数
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	//创建一个新的跳表节点
	newNode := newSkipListNode(nil, num, level)

	//原有节点连接
	for i := 0; i <= level-1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	//如果当前节点的层数大于之前跳表的层数
	//更新当前跳表层数
	if level > this.level {
		this.level = level
	}
}

//查找
func (this *Skiplist) Search(target int) bool {
	if this.length == 0 {
		return false
	}

	cur := this.head
	for i := this.level - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == target {
				return true
			} else if cur.forwards[i].score > target {
				break
			}
			cur = cur.forwards[i]
		}
	}

	return false
}

//删除节点
func (this *Skiplist) Erase(num int) bool {

	//查找前驱节点
	cur := this.head
	//记录前驱路径
	update := [MAX_LEVEL]*skipListNode{}
	for i := this.level - 1; i >= 0; i-- {
		update[i] = this.head
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == num {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}
	flag := false
	cur = update[0].forwards[0]
	for i := cur.level - 1; i >= 0; i-- {
		if update[i] == this.head && cur.forwards[i] == nil {
			this.level = i
		}

		if nil == update[i].forwards[i] {
			update[i].forwards[i] = nil
		} else {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
			flag = true
		}
	}

	this.length--

	return flag
}
