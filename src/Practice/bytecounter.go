package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("Hello World"))
	fmt.Println(c)
	c = 0
	var name = "Tom"
	fmt.Fprintf(&c, "Hello %s", name)
	fmt.Println(c)
}
