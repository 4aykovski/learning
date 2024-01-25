package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("https://series.naver.com/comic/detail.series?productNo=6030941?123123")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(u.Host)

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		log.Fatalf(err.Error())
	}

	productNo := q["productNo"]

	fmt.Println(productNo)
}
