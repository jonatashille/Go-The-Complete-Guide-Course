package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	var hobbies [3]string = [3]string{"sports", "spiders", "eat"}

	fmt.Print("1) ")
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	fmt.Print("2.1) ")
	fmt.Println(hobbies[0])

	//		- The second and third element combined as a new list
	slicedHobbies := hobbies[1:]
	fmt.Print("2.2) ")
	fmt.Println(slicedHobbies)

	// 3) Create a slice based on the first element that contains the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	slicedHobbies2 := hobbies[:2]
	slicedHobbies3 := hobbies[0:2]
	fmt.Print("3.1) ")
	fmt.Println(slicedHobbies2)
	fmt.Print("3.2) ")
	fmt.Println(slicedHobbies3)

	// 4) Re-slice the slice from (3) and change it to contain the second and last element of the original array.
	slicedHobbies3 = hobbies[1:3]
	fmt.Print("4) ")
	fmt.Println(slicedHobbies3)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	var goals []string = []string{"learn go", "update my knowledge"}
	fmt.Print("5) ")
	fmt.Println(goals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "gain new skills"
	goals = append(goals, "be a better programmer")
	fmt.Print("6) ")
	fmt.Println(goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a dynamic list of products (at least 2 products).
	product1 := Product{
		title: "Bike",
		id:    1,
		price: 4999.99,
	}

	product2 := Product{
		title: "Hemelt",
		id:    2,
		price: 379.49,
	}

	var products []Product = []Product{product1, product2}

	fmt.Print("7.1) ")
	fmt.Println(products)
	//		Then add a third product to the existing list of products.

	product3 := Product{
		title: "Tire",
		id:    3,
		price: 157.86,
	}

	products = append(products, product3)
	fmt.Print("7.2) ")
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
