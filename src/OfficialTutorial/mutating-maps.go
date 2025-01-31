/*
修改映射
在映射 m 中插入或修改元素：

m[key] = elem
获取元素：

elem = m[key]
删除元素：

delete(m, key)
通过双赋值检测某个键是否存在：

elem, ok = m[key]
若 key 在 m 中，ok 为 true ；否则，ok 为 false。

若 key 不在映射中，那么 elem 是该映射元素类型的零值。

同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。

注 ：若 elem 或 ok 还未声明，你可以使用短变量声明：

elem, ok := m[key]
*/
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The Value", m["Answer"])

	m["Answer2"] = 48
	fmt.Println("The Value", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The Value", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The Value", v, "Present?", ok)

	vv, ok := m["Answer2"]
	fmt.Println("The Value", vv, "Present?", ok)
}
