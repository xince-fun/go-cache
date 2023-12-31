package gocache

// A ByteView holds an immutable view of bytes.
type ByteView struct {
	b []byte
}

/*
在我们的lru.Cache实现中，要求被缓存对象必须实现 Value 接口
*/

// Len returns the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

/*
b 是只读的,使用 ByteSlice() 方法返回一个拷贝，防止缓存值被外部程序修改
*/

// ByteSlice returns a copy of the data as a byte slice.
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String returns the data as a string, making a copy if necessary.
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
