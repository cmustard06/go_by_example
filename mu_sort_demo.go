package main

import (
	"sort"
	"fmt"
)

/*
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
*/

type byLength []string

func (s byLength) Len() int{
	return len(s)
}

func (s byLength) Swap(i, j int){
	s[i],s[j] = s[j],s[i]
}

func (s byLength) Less(i,j int) bool{
	return len(s[i])<len(s[j])

}

func main(){
	fruits := []string{"peach","banana","kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

}
