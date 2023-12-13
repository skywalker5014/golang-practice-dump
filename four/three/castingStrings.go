package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"time"
	"unicode/utf8"
)

var pl = fmt.Println

func main(){
	cV3 := "5000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.Type(cV4))





// in go characteres are called rune 

stt := "asjdfd"

pl("rune count", utf8.RuneCountInString)

// %d integer, %#U for rune unicode, %c character, %f float -- %.2f 2 decimal places in float 
// %t boolean , %s string  , %o base 8 , %x base 16, %v guesses based on data type
// %T type of supplied value

// time 
now := time.Now()

nownow := now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()

randnum := rand.Intn(50)  // print random number upto 50 not includuing 50


//math library got everyting from algebra to trigonometry and calculus 


returnnedValue := fmt.Sprintf("%9.f", 3.212234 ) // this will format the string and also return it while printing it 


//array
var arr1 [5]int // array of integers with 5 space/size
arr[0] = 1 
arr2 := [5]int{1,2,3,4,5}

len(arr2)

for i, v := range arr2 {
	fmt.Printf("%d : %d", i , v)
}

arr3 := [2][2]int{
	{1,1},
	{2,3}
}

for i := 0, i <  2 ; i++ {
	for j := 0, j < 2; j++{
		pl(arr3[i][j])
	}
}

//slices
str1 := "abcde"

rrr := rune(str1)

for _, v := range rrr {
	pl(v)
}

sl1 := make([]string, 6) //? 


func divide(a int, b int) (ans int, err error){
	if y == 0{
		return 0, fmt.Errorf("YOU CANT DIVIDE BY ZERO")
	} else {
		return a/b, nil
	}
}


func sum(nums ...int) int {
	sm := 0
	for _, num:= range nums{
		sum+= num
	}
	return sum
}

func getarraysum(arr []int) int {
	sum:=0
	for _, val:=range arr{
		sum += val
	}
	return sum
}



//file io

f, err := os.Create("filename.txt")
if err != nil {
	log.Fatal(err)
}
defer f.Close() //as long as we are inside main func file is open and once we are out of main func file is closed
primearr := []int{1,2,3,4,5}

var primearr2 []string
for _, val:= range primearr{
	primearr2 = append(primearr2, strconv.Itoa(val))
}
for _, num := range primearr2 {
	_, err := f.WriteString(num + "\n")
	if err != nil {
		log.Fatal(err)
	}
}


//map 
// var mapname map[keytype]valuetype
var myMap map[string]string
mymap2 = make(map[string]string)
mymap3 := make(map[string]string)
myMap["one"] = "value1"

mymap4 := map[int]string{1: "one", 2: "two"}

fmt.Println(myMap["one"])

//generics 

type Myconstraints interface{
	int | float64
}


func getsum[T Myconstraint](x T, y T) T{
	return x + y
}

fmt.Println(getsum(5,4))
fmt.Println(getsum(5.4,4.5))


//structs :: store values with many different data types 

type customer struct {
	name string
	address string
	number int64
}

var ts customer
ts.name = "john cena"
ts.address = "somewheree"
ts.number = 999

newperson := customer{"tom","lalaland",6767}

//struct composition

type business struct {
	role string
	salary float32
	customer
}

func (b business) info(){
	fmt.Printf(b.role, b.salary, b.customer.name, b.customer.address)
}

con1 := person{ "ron", "gogoland", 8989,}
bus1 := business{"developer", 345.23, con1,}

//defined type

type Tsp float64
type TBs float64
type ML float64

func tsptoMl(tsp Tsp) ML{
	return ML(tsp * 14.55)
}


//interfaces 

type Animal interface{
	Sound()
	Sound2()
}

type Cat string 

func (c Cat ) Attack(){
	fmt.Println("cat attack")
}

func (c Cat) Sound(){
	fmt.Println("meaow")
}

func (c Cat) Sound2(){
	fmt.Println("prrrrrr")
}


var kitty Animal
kitty = Cat("kitty")

kitty.Sound()
kitty.Sound2()

//concurrency :: multiple blocks of code share the same execution time at a time to execute faster and not wait for each other the complete 
// this is multithreading but in go its called goroutines (the order of execution of these routines cant be decided)










}

