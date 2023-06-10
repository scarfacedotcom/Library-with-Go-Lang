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

func printMemberAudits(library *Library) {
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
	member.books[title] = LendAudit{checkOut: time.Now()}
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
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["crypto in Rwanda"] = BookEntry{
		total:  4,
		lended: 0,
	}

	library.books["Blockchain book"] = BookEntry{
		total:  8,
		lended: 0,
	}

	library.books["Crypto Alphas"] = BookEntry{
		total:  5,
		lended: 0,
	}

	library.books["Dawn of a new era"] = BookEntry{
		total:  8,
		lended: 0,
	}

	library.members["scarface"] = Member{"scarface", make(map[Title]LendAudit)}
	library.members["jay"] = Member{"jay", make(map[Title]LendAudit)}
	library.members["nancy"] = Member{"nancy", make(map[Title]LendAudit)}

	fmt.Println("\nInitial:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	member := library.members["scarface"]
	checkedOut := checkoutBook(&library, "Blockchain book", &member)
	fmt.Println("\nCheck out a book:")
	if checkedOut {
		printLibraryBooks(&library)
		printMemberAudits(&library)

	}
	returned := returnBook(&library, "Blockchain book", &member)
	fmt.Println("\n Book has been returned")
	if returned {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}
}
