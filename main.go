package main

import "fmt"

func start() {
	shouldContinue := true
	choice := -1

END:
	for shouldContinue {
		fmt.Println("Please choose one of the following option to continue")
		fmt.Println("1. Get All Book")
		fmt.Println("2. Search book by name")
		fmt.Println("3. Search book by author")
		fmt.Println("4. Add a book to the library") //only for admin access
		fmt.Println("5. Borrow a book")             // only for member access
		fmt.Println("6. Return a book")             // only for  member access

		fmt.Print("Enter your choice : ")
		fmt.Scan(&choice)
		fmt.Printf("Your choice is %d\n", choice)
		switch choice {
		case 1:
			//getAllBook()
		case 2:
			//searchByName()
		case 3:
			//searchByAuthor()
		case 4:
			//addBook()
		case 5:
			//borrowBook()
		case 6:
			//returnBook()
		default:
			shouldContinue = false
			break END
		}
	}
}

func main() {
	fmt.Println("Welcome to the library app")
	start()
	fmt.Println("Thanks for visiting the library, Happy studying!!")

}
