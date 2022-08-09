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

var ageKirillPointer *int = &kirill.age
var contactPhonePointer *int = &kirill.contact.phoneNumber

func main() {
	*ageKirillPointer = 19
	*contactPhonePointer = 12398731987
	fmt.Println(kirill.age, kirill.name, kirill.contact.phoneNumber)
}
