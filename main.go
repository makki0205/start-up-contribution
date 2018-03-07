package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please input github userID.\n\n./contribute [github userID]\n")
		os.Exit(0)
	}
	contributeCount := getContribute(args[1])
	
	fmt.Printf("\n ʕ◔ϖ◔ʔ  << Today Contribution is %v\n\n", contributeCount)

}

func getContribute(id string) string {
	doc, err := goquery.NewDocument("https://github.com/users/" + id + "/contributions")
	if err != nil {
		panic(err)
	}

	count, _ := doc.Find("rect").Last().Attr("data-count")
	return count
}