package main

func main() {
	_ = calculate()
}

func calculate() *int {
	x := 2
	f := x * 2
	return &f
}
