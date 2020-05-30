package main

func main() {
	Operator()
}

func Operator() {
	// x << n --> x * 2 ^ n
	println(1 << 1) // 1 * 2 ^ 1 = 2
	println(1 << 2) // 1 * 2 ^ 2 = 4
	println(1 << 5) // 1 * 2 ^ 5 = 32

	//但是5/4的结果是
	//1，因为整数除法会向着0方向截断余数。
	// x >> n --> x / 2 ^ n
	println(100 >> 1) // 100 / 2 ^ 1 = 50
	println(100 >> 2) // 100 / 2 ^ 2 = 25

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	println(x)
	println(y)
}
