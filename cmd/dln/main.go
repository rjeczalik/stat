package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func die(v interface{}) {
	fmt.Fprintln(os.Stderr, v)
	os.Exit(1)
}

func main() {
	var r io.Reader = os.Stdin
	switch len(os.Args) {
	case 1:
	case 2:
		f, err := os.Open(os.Args[1])
		if err != nil {
			die(err)
		}
		defer f.Close()
		r = f
	default:
		die("dln: invalid arguments")
	}
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return
	}
	prev, err := strconv.ParseUint(s.Text(), 10, 64)
	if err != nil {
		die(err)
	}
	for s.Scan() {
		cur, err := strconv.ParseUint(s.Text(), 10, 64)
		if err != nil {
			die(err)
		}
		fmt.Println(cur - prev)
		prev = cur
	}
	if err := s.Err(); err != nil {
		die(err)
	}
}
