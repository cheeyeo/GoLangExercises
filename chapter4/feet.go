package main

import "fmt"

func main(){
  fmt.Println("Enter distance in feet:")
  var Feet float64
  fmt.Scanf("%f", &Feet)

  Meters := Feet * 0.3048
  fmt.Println("Distance in meters", Meters)
}
