package main

import "flag"

func main() {
	filename := flag.String("file", "in.txt", "File to read from")
	flag.Parse()

	println(*filename)
}
