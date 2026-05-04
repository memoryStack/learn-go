package main

import "fmt"

type User struct {
	name string
	age int
}

func (u *User) changeName(name string) {
	u.name = name
}

func (u User) printAge() {
	fmt.Println(u.age)
}

type Duration int64

func (d *Duration) hours() Duration {
	return *d / 3600
}

func main() {
	anuj := User{
		name: "Anuj",
		age: 20,
	}

	anuj.changeName("Hickup")
	anuj.printAge()

	// fmt.Println (Duration(10000).hours())
	duration := Duration(10000)
	fmt.Println (duration.hours())

}
