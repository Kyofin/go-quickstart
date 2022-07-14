package main

func main() {
	i := 0
HERE:
	println(i)
	i++
	goto HERE
}
