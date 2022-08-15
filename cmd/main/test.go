package main

// Расширение для автозаполнения (установить): Tabnine AI Autocomplete for Javascript, Python, Typescript, PHP, Go, Java, Ruby & more
import "fmt"

type phone struct {
	phoneNumber int
	region      string
}

type human struct {
	name    string
	age     int
	man     bool
	contact phone
}

var kirill human = human{
	name: "Kirill",
	age:  18,
	man:  true,
	contact: phone{
		phoneNumber: 9867233,
		region:      "Vologda",
	},
}

func structPlay() {
	//игрульки со структурами
	*ageKirillPointer = 19
	*contactPhonePointer = 12398731987
	if kirill.age == 19 {
		fmt.Println(kirill.age, kirill.name, kirill.contact.phoneNumber)
	} else {
		panic("слава вове putinu")
	}
}

var numbers = [...]string{"Kirill", "Kudryavcev"}
var numbers2 [5]string = [5]string{"1", "2", "3", "4"}
var numbers3 = [5]string{"4", "3", "2", "1"}

func massivePerebor(massive []string) {
	defer fmt.Println(numbers2, numbers3)
	for index, value := range massive {
		fmt.Println(value, index)
	}
}

var ageKirillPointer *int = &kirill.age
var contactPhonePointer *int = &kirill.contact.phoneNumber

func main() {
	massivePerebor(numbers2[:])
}
