package containers

// 所有容器都实现Container接口
type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
	String() string
}
