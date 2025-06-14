package main

import "fmt"

const LoginToken string = "gfjhkghghkjsd"

func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("variables is of type: %T \n", username)

	var smallVal bool = true
	fmt.Println(smallVal)
	fmt.Printf("variables is of type: %T \n", smallVal)

	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variables is of type: %T \n", anothervariable)

	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variables is of type: %T \n", LoginToken)
	//:= short variable declaration operator in go
}
