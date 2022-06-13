package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// cara ke 1
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	// cara ke 2
	utc := time.Date(2022, 06, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	//cara ke 3
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2022-06-10")
	fmt.Println(parse)

}
