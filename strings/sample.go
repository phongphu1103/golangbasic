package main

import (
	"fmt"
	"unicode/utf8"
	"net/url"
)

func main() {
	// declare multiple variables at once
	b, c := 1, 2
	fmt.Println(b)
	fmt.Println(c)
	// declare a constant value ~ hello
	const str = "สวัสดี"
	fmt.Println("Length: ", len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}

	fmt.Println()
	fmt.Println("Rune count: ", utf8.RuneCountInString(str))

	for idx, runeValue := range str {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(str); i += w {
		runeValue, width := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
	}

	s := "Ann' 1=1--"
	fmt.Println("http://example.com/"+url.QueryEscape(s))

	queryParams := url.Values{}
	queryParams.Add("title", "title has contain speci@l ch@r@cter")
	queryParams.Add("author", "zodi@c")
	fmt.Println("http://example.com/"+queryParams.Encode())
}