package main

import (
	"fmt"
	"time"
)

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int
	lended int
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkOut.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMemberAudit(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}

}
func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}
func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of the library")
		return false
	}
	if book.lended == book.total {
		fmt.Println("bro....No more books to lend")
		return false
	}
	book.lended += 1
	library.books[title] = book
	member.books[title] = LendAudit{checkOut: time.Now().Round()}
	return true
}
func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("member did not check out this book")
		return false
	}
	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}
func main() {

}
