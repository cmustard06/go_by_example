package main

import (
	"fmt"
	"github"
	"log"
)

func main() {
	//terms := []string{"python","scanner"}
	//res, err := github.SearchIssues(terms)
	//if err!=nil{
	//	fmt.Errorf("error is %s", err.Error())
	//	os.Exit(1)
	//}
	//fmt.Printf("res:%#v\n", res)
	//fmt.Println(res.Items[0].Title,res.Items[0].Body,res.Items[0].User.Login)

	var issues []string = []string{"repo:golang/go", "is:open", "json", "decoder"}
	result, err := github.SearchIssues(issues)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues :\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

}
