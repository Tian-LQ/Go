package main

func f1() {
	defer func() {
		println(1)
	}()
	defer func() {
		println(2)
	}()
}

func main() {
	f1()
	println(3)
}
