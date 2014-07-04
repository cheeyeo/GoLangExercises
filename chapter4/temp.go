package main

import "fmt"

func main(){
  fmt.Println("Enter a temperature in Farenheit:")
  var F float64
  fmt.Scanf("%f", &F)

  C := (F - 32) * 5/9
  fmt.Println(" Celsius", C)
}
