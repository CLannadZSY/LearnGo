package main

import "fmt"

func main() {
	i, j := 100, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	fmt.Println(i, p, *p)
	*p = 21        // 通过指针设置 i 的值
	fmt.Println(i) //查看 i 的值
	c := p
	fmt.Println(c, *c)
	p = &j         // 指向 j
	*p = *p / 37   // 通过指针 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}
