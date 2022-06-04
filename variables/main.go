package main

import "fmt"

func main() {

	var camelCaseVar string = "camelCase"
	var stringExample string
	const constIntegerExample int = 50
	var varInteger int = constIntegerExample
	stringExample = "example"
	var scanVar string

	// syntatic sugar
	syntaticSugar := "Only used for variables, not constants. Also does not work if you want to explictly declare types"

	fmt.Println()

	fmt.Println("Example", camelCaseVar, "Variable")
	fmt.Printf("Constante %v of values:%v\n", stringExample, constIntegerExample)
	fmt.Printf("Operation %s %v-%c=%v\n", stringExample, varInteger, constIntegerExample, varInteger-constIntegerExample)

	fmt.Println("syntatic sugar... ")
	fmt.Println(syntaticSugar)

	fmt.Print("Input/output processing Example ...: ")

	// pointer '&' that references addresses in memory like a address of a variables
	fmt.Scan(&scanVar)

	fmt.Println("printing input...\n", scanVar)
	fmt.Println("Printing the 'scanVar' variable memory location", &scanVar)
}
