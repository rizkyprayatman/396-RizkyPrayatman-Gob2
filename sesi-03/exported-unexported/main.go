package main

import "exported-unexported/helpers"

func main() {
	helpers.Greet()

	var person = helpers.Person{}
	person.Invokegreet()
}
