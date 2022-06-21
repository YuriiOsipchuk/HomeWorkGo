package main

import  "fmt"

func main() {
  // arr := []int{3, 4, 4, 3, 6, 3}
  arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
  var result []int
  var check bool

  for i := 0; i < len(arr); i++ {
    if i == 0 {
      result = append(result, arr[i])
    } 
    
    for j := 0; j < len(result); j++ {
      if arr[i] == result[j] {
        check = false
        break
      } else {
        check = true
      }
    }

    if check {
      result = append(result, arr[i])
    }

  }
  
  fmt.Println(arr)
  fmt.Println(result)
}
