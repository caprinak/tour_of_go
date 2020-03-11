package main

import ("fmt"
		"strings"
)
func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

	//func WordCount(s string) map[string]int {
		//wc := make(map[string]int)

		s1 := "I love you"
		fmt.Println(strings.Fields(s1))
		arr := strings.Fields(s1)
		fmt.Println(arr[1])

		//for _, w := range strings.Fields(s) {
			//wc[w]++
	
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
