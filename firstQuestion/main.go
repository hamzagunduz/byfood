package main

func main() {
	divideByTwo(9)
}

func divideByTwo(number int) {
	var half int
	if number/2 > 1 {
		half = number / 2
		divideByTwo(half)
	}
	println(number)
}
