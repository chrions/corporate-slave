package fyf

type OrderedStream struct {
	m   map[int]string
	ptr int
	len int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		len: n,
		m:   map[int]string{},
		ptr: 1,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	ret := make([]string, 0, this.len)
	this.m[idKey] = value

	for i := this.ptr; i <= this.len; i++ {
		if v, ok := this.m[i]; ok {
			ret = append(ret, v)
			this.ptr = i + 1
		} else {
			break
		}
	}

	return ret
}
