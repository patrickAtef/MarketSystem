package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id       int
	name     string
	price    float32
	quantity int
}

func inputNumber(message string) float64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(" %v ", message)
	input, _ := reader.ReadString('\n')
	num, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if num != 0 {
		return num
	} else {
		fmt.Print("invalid entry!!")
		return inputNumber(message)
	}

}

func inputString(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(" %v ", message)
	input, _ := reader.ReadString('\n')
	return input[:(len(input))-2]

}

func inputByte(message string) byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(" %v ", message)
	input, _ := reader.ReadByte()
	return input

}

func printProduct(aproduct Product) {
	fmt.Printf("%4v| %12v| %12v| %12v|\n", aproduct.id, aproduct.name, aproduct.price, aproduct.quantity)
}

func printAllProducts(productArr []Product) {

	for p := range productArr {

		printProduct(productArr[p])
	}
}

func addProduct(newID int) Product {
	var newProduct Product
	newProduct.id = newID
	newProduct.name = inputString("enter Product name:")
	newProduct.price = float32(inputNumber("enter Product price:"))
	newProduct.quantity = int(inputNumber("enter Product quantity:"))
	return newProduct
}

func main() {

	aproduct1 := Product{1, "cheese", 3.4, 5}
	aproduct6 := Product{2, "Mushroom", 8.9, 10}
	aproduct11 := Product{3, "milk", 13.14, 15}
	aproduct16 := Product{4, "water", 18.19, 20}
	productArray := []Product{aproduct1, aproduct6, aproduct11, aproduct16}

	var repeatAnswer byte
	//for repeatAnswer != 'n' && repeatAnswer != 'N' {
	for {

		menuChoice := chooseFromMenu()
		//menuChoice := 1

		switch menuChoice {
		case 1:
			{
				fmt.Printf("%4v| %12v| %12v| %12v|\n", " id", "name", "price", " quantity")

				printAllProducts(productArray)

			}
		case 2:
			{
				newSize := len(productArray) + 1
				productArray = append(productArray, addProduct(newSize))
				printAllProducts(productArray)
			}

		case 3:
			fmt.Print("3. find existing product\n")
		case 4:
			fmt.Print("4. edit existing product\n")
		case 5:
			fmt.Print("5. delete existing product\n")
		case 6:
			fmt.Print("6. exit\n")
		}
		fmt.Println("---------------------------------------")
		repeatAnswer = inputByte("press N to exit, any other key to go back to main menu:")
		if repeatAnswer == 'n' || repeatAnswer == 'N' {
			break
		}

	}
	//byd5lny f infinte loop !!
	//fmt.Print("press N to exit, any other key to go back to main menu:")
	//fmt.Scanln(repeatAnswer)

	//}

}
func chooseFromMenu() int {
	//ay 7ga test
	var input int
	fmt.Print("choose from (1,2,3,4,5,6) to perform one of the following operations:\n")
	fmt.Print("1. view all products\n")
	fmt.Print("2. add new product\n")
	fmt.Print("3. find existing product\n")
	fmt.Print("4. edit existing product\n")
	fmt.Print("5. delete existing product\n")
	fmt.Print("6. exit\n")
	for {
		input = int(inputNumber("enter your choice :"))
		if input == 1 || input == 2 || input == 3 || input == 4 || input == 5 || input == 6 {
			break
		} else {
			fmt.Println("you can only choose from 1 to 6.")

		}
	}
	return input

}
