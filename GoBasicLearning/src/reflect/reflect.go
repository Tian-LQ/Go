package reflect

// TODO reflect.TypeOf VS reflect.ValueOf
// 1.reflect.TypeOf 返回类型(reflect.Type)
// 2.reflect.ValueOf 返回值(reflect.Value)
// 3.可以从(reflect.Value)获得类型
// 4.通过kind来判断类型

// TODO Struct Tag
// 序列化与反序列化相关的tag

type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// TODO DeepEqual
// 我们知道slice和map类型均为引用类型，因此其只能与nil进行比较
// reflect提供了DeepEqual方法帮助我们相互比较slice对象(map对象)

//  Elem() 获取指针指向的值
