package main

import (
	"fmt"
	"net/url"
)

type Values map[string][]string

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func main() {
	m := url.Values{"lang": {"en"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")
	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1" (first value)
	fmt.Println(m["item"])     // "[1 2]" (direct map access)
	m = nil
	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3")         // panic: assignment to entry in nil map

}
