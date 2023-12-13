package greetings
import ("fmt" 
		"errors"
	)
func Hello(name string) (string, error){
	if name == ""{
		return "", errors.New("empty name")
	}
	//var message string
	//message = fmt.Sprintf("hi, %v, welcome", name)
	message := fmt.Sprintf("Hi, %v. welcome", name)
	return message, nil
}