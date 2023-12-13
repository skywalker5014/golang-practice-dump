package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

var pl = fmt.Println

func main(){
	pl("what is the name")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("hello", name)
	} else {
		log.Fatal(err)
	}
}