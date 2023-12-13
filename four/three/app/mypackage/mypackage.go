package stuff

import (
	"strconv"
)

var Name string = "john"

func IntArrToStrArr(intArr []int) []string{
	var strArr []string
	
	for _, v := range intArr {
		strArr = append(strArr, strconv.Itoa(v))
	}
	return strArr
}