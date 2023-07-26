package main

var (
	b bool
	c int     = 10
	d string  = "Hello World"
	e float32 = 3.14
)

func main() {

	inside := "Inside"

	println(inside)
	println(b)
	println(c)
	println(d)
	println(e)

	inside = "Inside 2"

	println(inside)
}
