package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
	"github.com/google/uuid"
)

func main() {
	format_date := "2006-01-02"
	str := "Black Panther"
	leng := len(str)
	fmt.Println(str + "\n")
	fmt.Println(leng)

	local, err := time.LoadLocation("Asia/Saigon")
	if err != nil {
		fmt.Println(err)
	}
	current_time := time.Now().In(local)
	next_dates := current_time.AddDate(0, 0, 6)
	fmt.Println(math.Pi)
	fmt.Println(current_time.Format(format_date))
	fmt.Println(next_dates.Format(format_date))

	uuidWithHyphen := uuid.New()
    fmt.Println(uuidWithHyphen)

	//Floyd's triangle
	var rows int
	t := 1

	fmt.Print("Enter rows: ")
	fmt.Scan(&rows)

	for i := 1; i <= rows; i++ {
		for k := 1; k <= i; k++ {
			if t < 10 {
				str = "0" + strconv.Itoa(t)
				fmt.Printf(" %v", str)
			} else {
				fmt.Printf(" %d", t)
			}
			t++
		}

		fmt.Println("")
	}

	// declaring array
	// var x []int
	// declaring array without length
	x := [...]int{2, 4, 6, 8, 10}
	var total int = 0
	for _, value := range x {
		total = total + value
	}
	fmt.Println("total:", total)

	//Fibonacci
	t1 := 0
	t2 := 1
	nextTerm := 0
	fmt.Print("Fibonacci Series:")
	for i := 1; i <= rows; i++ {
		if i == 1 {
			fmt.Print(" ", t1)
			continue
		}
		if i == 2 {
			fmt.Print(" ", t2)
			continue
		}
		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
		fmt.Print(" ", nextTerm)
	}

}
