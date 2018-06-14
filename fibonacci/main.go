package main

import (
  "time"
  "fmt"
  "encoding/json"
)

var n1,n2,n3 int = 0,1,0
var answersCount,errorsCount = 0,0
var currentNumber = 0

func getFibonacci(count int) []int  {
  var fibonacciArray = make([]int, count)

  fibonacciArray[0] = 0
  fibonacciArray[1] = 1
  for i := 2; i < count; i++ {
    fibonacciArray[i] = fibonacciArray[i-1] + fibonacciArray[i-2]
  }

  return fibonacciArray

}


func main() {
  count :=  50
  var fibonacciArray = getFibonacci(count)


  OutLoop:
  for i := 0; i < len(fibonacciArray); i++ {
  if answersCount == 10 || errorsCount == 3 {
    break
  }
  timeout := time.After(10 * time.Second)

   for   {
     number := make(chan int)
     var currentNumber = 0
     go func() {
       var n int
       fmt.Println("Enter a number: ")

       _, err := fmt.Scanf("%d", &n)

       if err != nil {
         fmt.Println(err)
       }
       currentNumber = n
        number <- n
     }()
     select {
       case <-timeout:
         mapD := map[string]int{"currentNumber": fibonacciArray[i], "position": i}
         json, _ := json.Marshal(mapD)
         fmt.Println(string(json))
         errorsCount++

         continue OutLoop
     case <-number:
       if fibonacciArray[i] == currentNumber {
        fmt.Println("ok")
        answersCount++;
        continue OutLoop
       } else {
         mapD := map[string]int{"currentNumber": fibonacciArray[i], "position": i}
         json, _ := json.Marshal(mapD)
         fmt.Println(string(json))
         errorsCount++;
         continue OutLoop
       }
       time.Sleep(1*time.Second)
     }
   }
 }




}




