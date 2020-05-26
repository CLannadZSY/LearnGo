/*
映射的文法（续）
若顶级类型只是一个类型名，你可以在文法的元素中省略它。
*/

package main

import "fmt"

func main1() {
	type Vertex struct {
		Lat float64
		Lbb int64
		Lcc string
	}

	var m = map[string]Vertex{
		"Google": {1.2, 222, "google.com"},
		"Bing":   {1.1, 333, "bing.com"},
	}
	fmt.Println(m)
}

func main2() {
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}

func main() {
	main1()
	main2()
}
