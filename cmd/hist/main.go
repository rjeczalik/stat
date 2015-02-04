package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type pair struct {
	s string
	n int
}

type byvalue []pair // implements counting set ordered by value

func (b byvalue) Len() int           { return len(b) }
func (b byvalue) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byvalue) Less(i, j int) bool { return b[i].s < b[j].s }

func (b byvalue) Search(s string) int {
	return sort.Search(len(b), func(i int) bool { return b[i].s >= s })
}

func (b *byvalue) Add(s string) int {
	switch i := b.Search(s); {
	case i == len(*b):
		*b = append(*b, pair{s: s, n: 1})
		return 1
	case (*b)[i].s == s:
		(*b)[i].n++
		return (*b)[i].n
	default:
		*b = append(*b, pair{})
		copy((*b)[i+1:], (*b)[i:])
		(*b)[i] = pair{s: s, n: 1}
		return 1
	}
}

type bycount []pair

func (b bycount) Len() int           { return len(b) }
func (b bycount) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b bycount) Less(i, j int) bool { return b[i].n >= b[j].n }
func (b bycount) Sort()              { sort.Sort(b) }

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
		die("hist: invalid arguments")
	}
	max := 0
	hist := byvalue{}
	s := bufio.NewScanner(r)
	for s.Scan() {
		if n := hist.Add(s.Text()); n > max {
			max = n
		}
	}
	if err := s.Err(); err != nil {
		die(err)
	}
	if len(hist) == 0 {
		return
	}
	bycount(hist).Sort()
	format := "%" + strconv.Itoa(len(strconv.Itoa(max))) + "s\t%d\t"
	n, _ := fmt.Printf(format, hist[0].s, hist[0].n)
	n = 70 - n
	fmt.Println(strings.Repeat("#", n))
	for i := 1; i < len(hist); i++ {
		fmt.Printf(format, hist[i].s, hist[i].n)
		if r := n * hist[i].n / max; r != 0 {
			fmt.Print(strings.Repeat("#", r))
		}
		fmt.Println()
	}
}
