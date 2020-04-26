//Nicholas Larsen
//4/25/2020
//finds out the commission of a total price
package main

import "fmt"

func commission(small float64, medium float64, large float64)float64{
var total float64
fmt.Println("what is the total value of your home?")
fmt.Scanln(&total)
//asks for the total value
if total<50000{
  return small*total
}else if total>50000 && total<=250000{
  return medium*total
}else{
  return large*total
  //returns whichever value based on the total
}
}
func main () {
small:=.04
medium:=.06
large:=.0675
//commission rates
fmt.Println("your commission total is",(commission(small, medium, large)))
//your commission total
}