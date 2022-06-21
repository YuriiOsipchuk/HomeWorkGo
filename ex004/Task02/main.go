package main

import  (
  "fmt"
  "strings"
  "strconv"
)

func main() {
  input := "1"
  // input := "1 2 3 4 5"
  input := "1 9 3 4 -5"
  var result string
  var min int32
  var max int32
  InputArrString := strings.Split(input, " ")


  for i := 0; i < len(InputArrString); i++ {

    tempIntVal, _ := strconv.Atoi(InputArrString[i])

    if i == 0 {
      max = int32(tempIntVal)
      min = int32(tempIntVal)
    }

    if int32(tempIntVal) > max {
      max = int32(tempIntVal)
    }

    if int32(tempIntVal) < min {
      min = int32(tempIntVal)
    }
    
  } 

  if len(InputArrString) > 1 {
    result = strconv.FormatInt(int64(max), 10) + " " + strconv.FormatInt(int64(min), 10)
  } else {
    result = strconv.FormatInt(int64(max), 10)
  }

  fmt.Println(result)
}
