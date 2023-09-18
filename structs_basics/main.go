package main

import "fmt"

func main() {

	title1, author, year := "the devine comedy", "dante aligheri", 2022
	title2, author2, year2 := "mcbeth", "willian shakesspehere", 1606

	fmt.Println("Book 1", title1, author, year)
	fmt.Println("book 2", title2, author2, year2)

	type book struct {
		title  string
		author string
		year   int
	}
	myBook := book{"test", "test", 1320}
	bestBook := book{title: "animal farm", author: "george", year: 2023}

	fmt.Println(myBook, bestBook)

	lastBook := book{title: "anna karenia"}

	fmt.Println(lastBook.title)

	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "rohit",
		lastName:  "degamwar",
		age:       48,
	}

	fmt.Println(diana)

	type Book struct {
		string
		float32
		bool
	}

	b1 := Book{"test", 10.2, false}

	fmt.Println(b1)

	//embedded structs

}
