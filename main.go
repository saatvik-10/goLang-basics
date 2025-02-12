package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/saatvik-10/goLang-basics/helpers"
)

type User struct {
	FirstName string
	LastName  string
}

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Person struct {
	Name    string `json:"Name"`
	Age     string `json:"Age"`
	Has_Dog bool   `json:"has_dog"`
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

	isTrue := true

	if isTrue {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	animals := []string{"dog", "cat", "fish"}

	for _, animal := range animals {
		fmt.Println(animal)
	}

	animalss := make(map[string]string)
	animalss["dog"] = "Shane"
	animalss["cat"] = "Mittens"

	for animalType, animal := range animalss {
		fmt.Println(animalType, animal)
	}

	firstLine := "Once upon a time in Mumbai dobara"

	for i, char := range firstLine {
		fmt.Println(i, ":", char) //string is a slice of bytes
	}

	type Users struct {
		FirstName string
		LastName  string
		Age       string
	}

	var user []Users
	user = append(user, Users{"Shane", "Lee", "25"})
	user = append(user, Users{"Bond", "Miller", "35"})
	user = append(user, Users{"Mark", "Sam", "21"})

	for _, l := range user {
		fmt.Println(l.FirstName, l.LastName, l.Age)
	}

	dog := Dog{
		Name:  "Fido",
		Breed: "Poodle",
	}

	PrintInfo(&dog)

	var myVar helpers.SomeType
	myVar.Name = "Shane"
	myVar.Age = 25

	fmt.Println(myVar)

	intChan := make(chan int)
	defer close(intChan) //execute whtever comes after this line when the function is done executing

	go calcVal(intChan)

	num := <-intChan

	fmt.Println(num)

	myJson := `
	[
		{
			"Name": "Shane",
			"Age": "25",
			"has_dog": true
		},
		{
			"Name": "Bond",
			"Age": "35",
			"has_dog": false
		}
	]
	`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		fmt.Println("Error unmarshalling JSON", err)
	}

	fmt.Printf("unmarshalled: %v", unmarshalled)

	//write json from a struct
	var mySlices []Person

	var m1 Person
	m1.Name = "Shane"
	m1.Age = "25"
	m1.Has_Dog = true

	mySlices = append(mySlices, m1)

	var m2 Person
	m2.Name = "Shane"
	m2.Age = "25"
	m2.Has_Dog = true

	mySlices = append(mySlices, m2)

	newJson, err := json.MarshalIndent((mySlices), "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON", err)
	}
	fmt.Println(string(newJson))
}

const numPool = 10

func calcVal(intChan chan int) {
	RandomNumber := helpers.RandNum(numPool)
	intChan <- RandomNumber
}

func PrintInfo(a Animal) {
	fmt.Println(a.Says())
	fmt.Println(a.NumberOfLegs())
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func saySomething() (string, string) {
	return "something", "something else"
}

func changeUsingPointer(s *string) {
	newVal := "new value"
	*s = newVal
}
