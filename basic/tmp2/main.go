package main

import "fmt"

func main(){
	err := generr()
	fmt.Printf("%s\n", err)
	fmt.Println("end of program")
}

func generr()error{
	return fmt.Errorf("this is error")
}
