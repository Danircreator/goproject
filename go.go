package main

import "fmt"


func main() {
  my_print(6,4,1,4,2,123,1)
}

func my_print(nums ...int){
  
  for _, nums := range nums{
  fmt.Println(nums)
  }

}