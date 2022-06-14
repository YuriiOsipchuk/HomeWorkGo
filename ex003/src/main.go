package main

import (
  "fmt"
  "math"
)

func main() {
  const pricePear  float64 = 7
  const priceApple float64 = 5.99

  var (
    countPear        int32 = 0
    countApple       int32 = 0
    cashUA           float64 = 23
    tempCashUA       float64 = 0
  )


  fmt.Printf("Одне яблуко коштує %v грн. Ціна однієї груші - %v грн.\n", priceApple, pricePear)
  fmt.Printf("Ми маємо %v грн.\n", cashUA)
  fmt.Println("")
  fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")

  countPear  = 8
  countApple = 9
  tempCashUA =  float64(countPear) * pricePear
  tempCashUA += float64(countApple) * priceApple

  // fmt.Printf("%v яблук та %v груш коштує %v грн.\n\n", countApple, countPear, math.Floor(tempCashUA))
  fmt.Printf("%v яблук та %v груш коштує %v грн.\n\n", countApple, countPear, math.Round(tempCashUA))
  // fmt.Printf("%v яблук та %v груш коштує %v грн.\n\n", countApple, countPear, math.Ceil(tempCashUA))

  fmt.Println("2. Скільки груш ми можемо купити?")
  countPear = int32(cashUA/pricePear)
  fmt.Printf("За %v грн. ми можемо купити %v груші\n\n", cashUA, countPear)

  fmt.Println("3. Скільки яблук ми можемо купити?")
  countApple = int32(cashUA/priceApple)
  fmt.Printf("За %v грн. ми можемо купити %v яблука\n\n", cashUA, countApple)

  fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?")
  countPear  = 2
  countApple = 2
  tempCashUA =  float64(countPear) * pricePear
  tempCashUA += float64(countApple) * priceApple
  fmt.Printf("Для купівлі %v груші та %v яблука потрібно %v грн.\n", countPear, countApple, tempCashUA)

  if cashUA > tempCashUA {
    fmt.Printf("Так можемо купити %v груші та %v яблука\n", countPear, countApple)
  } else {
    fmt.Printf("Ні не можемо купити %v груші та %v яблука\n", countPear, countApple)
  }
  
}
