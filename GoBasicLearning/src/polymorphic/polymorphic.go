package polymorphic

// TODO 空接口与断言
// 1.空接口可以表示任何类型
// 2.通过断言来将空接口转换为指定类型
// v, ok := obj.(int)	// ok == true 表示转换成功

// TODO Go接口最佳实践
// 1.倾向于使用小的接口定义(很多接口只包含一个方法)

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}

// 2.较大的接口定义，可以由多个小接口定义组合而成

type ReadWriter interface {
	Reader
	Writer
}

// 3.只依赖于必要功能的最小接口
// TODO 这里针对第三点做下说明：我们在client端使用接口时，一定要依赖于最小的接口，从而提高方法的复用性

func StoreData(writer Writer) error {
	// xxx
	return nil
}
