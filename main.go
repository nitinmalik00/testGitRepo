package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateId(name string, age int) string {
	hash := name + fmt.Sprintf("%d", age) + "#"
	random := rand.Intn(266)
	hash += fmt.Sprintf("%d", random)
	fmt.Println("Resulting hash: ", hash)
	return hash
}

func main() {
	var name string
	var age int

	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Println("Generating your unique ID (Aadhar).....")
	UID := generateId(name, age)
	fmt.Println("Successfully generated!!!")

	dt := time.Now()
	fmt.Printf("%s your age as on %d %s %d is %d and your ID is %s\n", name, dt.Day(), dt.Month(), dt.Year(), age, UID)

	fmt.Println("Logging out...")
}
