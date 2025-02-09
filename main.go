package main

import (
	"fmt"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	fmt.Println("Hello, World!")

	var whatToSay string
	var i int

	whatToSay = "GoodBye, cruel world!"
	fmt.Println(whatToSay)

	i = 10
	fmt.Println(i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println(whatWasSaid, theOtherThingThatWasSaid)

	changeUsingPointer(&whatToSay)
	fmt.Println(whatToSay)

	myMap := make(map[string]string)

	myMap["dog"] = "Shane"
	myMap["cat"] = "Mittens"

	fmt.Println(myMap["dog"])
	fmt.Println(myMap["cat"])

	myMapp := make(map[string]User)

	me := User{
		FirstName: "Shane",
		LastName:  "Lee",
	}

	myMapp["me"] = me

	fmt.Println(myMapp["me"])
	fmt.Println(myMapp["me"].FirstName)

	var mySlice []int
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 2)

	sort.Ints(mySlice)

	fmt.Println(mySlice)

	nums := []int{1, 2, 3, 4, 5} //int, string

	fmt.Println(nums)
	fmt.Println(nums[0:4])
}

func saySomething() (string, string) {
	return "something", "something else"
}

func changeUsingPointer(s *string) {
	newVal := "new value"
	*s = newVal
}
