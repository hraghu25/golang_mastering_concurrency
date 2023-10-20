package main

import "fmt"

// Person struct which hold field names or properties of a type Person
type Person struct {
	FirstName  string
	SecondName string
}

// Method receiver for a pointer Person
func (person *Person) FullName() string {
	return fmt.Sprintf("Full Name is %s %s", person.FirstName, person.SecondName)
}

func main() {
	// Creating an unbuffered channel.
	resultChannel := make(chan string, 139)

	// Declare actual value to a variable
	p := &Person{
		FirstName:  "Raghavendra",
		SecondName: "Hiremath",
	}

	// Execute goroutine in the background.
	go func() {
		result := p.FullName()
		resultChannel <- result
		close(resultChannel)
	}()

	// Validate the channel is open or closed
	// if ok = True, channel is open and it has value
	// if ok = False, channel is closed and the value is empty.
	result, ok := <-resultChannel
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("Channel is closed, no result available")
	}

	fmt.Print("Press Enter to exit....")
	fmt.Scanln()
}
