package main

import "fmt"

func main() {

	// Iterating arrays using range
	// The for loop can be used to iterate over elems of array
	arr := [...]float64{2.0, 4.2, 6.4, 50, 100}

	// Here I'm using for loop to iterate over the elements of the array
	// starting from index 0 to length of the array - 1.
	for a := 0; a < len(arr); a++ {
		fmt.Printf("%dth element of a is %4f\n", a, arr[a])
	}
	fmt.Println()

	// Example 2:
	names := [5]string{"Dan", "Paul", "Nunez", "Mane"}
	for d := 0; d < len(names); d++ {
		fmt.Printf("Name: %s\n", names[d])
	}
	fmt.Println()

	// Another huge example!
	floatNumbers := [...]float64{12.3, 37.3, 30.1, 20.3}
	sumOfNumbers := float64(0)
	fmt.Println(floatNumbers)

	// Range returns both the index and value
	// Returns both the index and value
	for k, v := range floatNumbers {
		fmt.Printf("%d the element of a is %.2f \n", k, v)
		sumOfNumbers += v
	}
	fmt.Println("The sum of numbers is: ", sumOfNumbers)

}
