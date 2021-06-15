package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
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

//3ayzha mn 8er 3nawen 3shan printall
func printProduct(aproduct Product) {
	fmt.Printf("%12v| %12v| %12v|\n", aproduct.name, aproduct.price, aproduct.quantity)
}

func printLabels() {
	fmt.Printf("%12v| %12v| %12v|\n", "name", "price", " quantity")
	fmt.Println("    --------------------------------------")
}

//zwd feha el 3nawen fo2y
func printProductWithLabels(aproduct Product) {
	printLabels()
	fmt.Printf("%12v| %12v| %12v|\n", aproduct.name, aproduct.price, aproduct.quantity)
}

func printAllProducts(productArr []Product) {

	for p := range productArr {

		printProduct(productArr[p])
	}
}

func addProduct() Product {
	var newProduct Product
	newProduct.name = inputString("enter Product name:")
	newProduct.price = float32(inputNumber("enter Product price:"))
	newProduct.quantity = int(inputNumber("enter Product quantity:"))
	fmt.Printf("%v is added successfully\n", newProduct.name)
	return newProduct
}
func findExistingproduct(productArr []Product) int {
	var existingProduct string
	var wantedProduct int
	existingProduct = inputString("enter the name of product that you are looking for:")
	for e := range productArr {
		if existingProduct == productArr[e].name {
			fmt.Printf("%v is available:\n", existingProduct)
			wantedProduct = e
			printProductWithLabels(productArr[e])
		}

	}

	return wantedProduct
}

func editProduct(p Product) Product {
	var i int
	var data string
	var NPrice float64
	var NQuantity int
	fmt.Println("1- Edit product Name ")
	fmt.Println("2- Edit product price ")
	fmt.Println("3- Edit product quantity ")
	i = int(inputNumber("choose what you want to edit (1 to 3)."))
	if i == 1 {
		data = inputString("Enter the new name of the product :")
		p.name = data
	} else if i == 2 {
		NPrice = inputNumber("Enter the new price of the product :")
		p.price = float32(NPrice)
	} else if i == 3 {
		NQuantity = int(inputNumber("Enter the new quantity of the product :"))
		p.quantity = NQuantity
	}
	fmt.Print("new data after editing is:\n")
	printProductWithLabels(p)

	return p
}

func deleteproduct(s []Product, index int) []Product {
	fmt.Println("product is deleted successfully")
	return append(s[:index], s[index+1:]...)
}

func main() {

	aproduct1 := Product{"cheese", 3.4, 5}
	aproduct6 := Product{"Mushroom", 8.9, 10}
	aproduct11 := Product{"milk", 13.14, 15}
	aproduct16 := Product{"water", 18.19, 20}
	productArray := []Product{aproduct1, aproduct6, aproduct11, aproduct16}

	var repeatAnswer byte
	//for repeatAnswer != 'n' && repeatAnswer != 'N' {
	for {

		menuChoice := chooseFromMenu()
		//menuChoice := 1
		if menuChoice == 6 {
			break
		}
		switch menuChoice {
		case 1:
			{
				printLabels()
				printAllProducts(productArray)
			}
		case 2:
			{
				productArray = append(productArray, addProduct())
			}
		case 3:
			findExistingproduct(productArray)
		case 4:
			var itemID int = findExistingproduct(productArray)
			productArray[itemID] = editProduct(productArray[itemID])
		case 5:
			pID := findExistingproduct(productArray)
			deleteAnswer := inputByte("press N to cancel, any other key to continue:")
			if deleteAnswer == 'n' || deleteAnswer == 'N' {
				fmt.Println("operation is cancelled.")
			} else {
				productArray = deleteproduct(productArray, pID)
			}

		}
		fmt.Println("-------------------------------------------")
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
