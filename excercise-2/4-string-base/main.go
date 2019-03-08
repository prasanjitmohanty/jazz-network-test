package main

import "fmt"

func main() {
	fmt.Println(fromDeci(32,182))
}

 func fromDeci(base uint, inputNum uint) string{ 
	var s string
	var result string
	
  
	for inputNum > 0{
		s+=reVal(inputNum % base)
		inputNum /= base
	}
	
	for _,v := range s {
		result = string(v) + result
	  }
    
    return result
} 

func reVal(num uint) string {
	if num >= 0 && num <= 9 {
		return (string)(num + 48)
	} else {
		return (string)(num - 10 + 65)
	}

}
