package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	fullName   string
	age        int
	idNumber   string
	occupation string
	address    string
}

func (person Person) ToString() string {
	var strAge string = strconv.Itoa(person.age)
	var personToString string = strings.Join(
		[]string{
			"\nHello ", person.fullName, "\n",
			"You are ", strAge, "\n",
			"with ID number ", person.idNumber, "\n",
			"Your occupation is ", person.occupation, "\n",
			"You live at ", person.address, "\n",
		},
		"")
	return personToString
}

func stripSuffix(item string) string {
	return strings.ReplaceAll(item, "\r\n", "")
}

func printErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
	return
}

func main() {
	// Global Error Declaration
	var err error

	fmt.Println("Hello User!\nLets capture your personal details!")
	var stdin *bufio.Reader = bufio.NewReader(os.Stdin)

	var fullName string
	fmt.Print("Enter your full name: ")
	fullName, err = stdin.ReadString('\n')
	fullName = stripSuffix(fullName)
	printErr(err)

	var age int
	var str_age string
	fmt.Print("Enter your age: ")
	str_age, err = stdin.ReadString('\n')
	printErr(err)
	conv := strings.Replace(str_age, "\r\n", "", 1)
	age, err = strconv.Atoi(conv)
	printErr(err)

	var idNumber string
	fmt.Print("Enter your ID number: ")
	idNumber, err = stdin.ReadString('\n')
	idNumber = stripSuffix(idNumber)
	printErr(err)

	fmt.Print("Enter your occupation: ")
	var occupation string
	occupation, err = stdin.ReadString('\n')
	occupation = stripSuffix(occupation)
	printErr(err)

	var address string
	fmt.Print("Enter your address: ")
	address, err = stdin.ReadString('\n')
	address = stripSuffix(address)
	printErr(err)

	fmt.Println("Compiling your personal details...")
	var sleep_time int = 5
	time.Sleep(1 * time.Second)
	fmt.Printf("Results ready in: (5)")
	for i := sleep_time; i >= 0; i-- {
		fmt.Print("\b\b")
		fmt.Printf("%d)", i)
		time.Sleep(1 * time.Second)
	}

	var person = Person{fullName: fullName, age: age, idNumber: idNumber, occupation: occupation, address: address}
	fmt.Printf(person.ToString())
}
