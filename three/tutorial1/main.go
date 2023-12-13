 package main
 import "fmt"
 import "unicode/utf8"

 func main(){
	fmt.Println("hello world")

	var intNum int = 323
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 123.456
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result int32 = int32(floatNum32) + intNum32
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myStr string = "hello \n world"
	var mystr string = `hello
	world`
	fmt.Println(myStr)
	fmt.Println(mystr)

	fmt.Println(len("test"))
	fmt.Println(len("A"))
	fmt.Println(utf8.RuneCountInString("A"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var mybool bool = true
	fmt.Println(mybool)

	// myvar := "string"
	// myvar2 := 32
	// myvar3 := 11.09
	// myvar4 := false

	// var one, two string = "abc", "dac"
	// three, four := 2, 3

	// const myconst string = "constant string"
	// const myconst2 int = 44

	too := 2
	var one int = 2

	test, tet := division( one , too)
	fmt.Printf("remainder is %v and result is %v", tet, test)

 }

 func printMe(){
	fmt.Println("hello from outside function")
 }

 func two(test string, test2 int){
		fmt.Println(test, test2)
 }

 func division(nume int, deno int) (int,int ){
	var result int = nume/deno
	remainder := nume%deno
	return result, remainder
 }