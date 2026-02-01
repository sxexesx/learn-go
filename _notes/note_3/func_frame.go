package main

func main() {
	a := 1
	println("before", &a)
	foo()
	println("after", &a)
}

//go:noinline
func foo() {
	bara()
	baza()
}

//go:noinline
// func bar() {
// 	arr := [1500]byte{}
// 	println(&arr)
//}

//go:noinline
func bara() *int {
	a := 100
	return &a
}

//go:noinline
func baza() int {
	c := 100
	return c
}
