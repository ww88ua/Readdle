package main

import (
  "time"
  "fmt"
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

type Error struct {
  CurrentNumber int
  Position int
}



func main() {
  count := 100
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
         error := Error{fibonacciArray[i], i}
         fmt.Println(error)
         errorsCount++

         continue OutLoop
     case <-number:
       if fibonacciArray[i] == currentNumber {
        fmt.Println("ok")
        answersCount++;
        continue OutLoop
       } else {
         error := Error{fibonacciArray[i], i}
         fmt.Println(error)
        errorsCount++;
        continue OutLoop
       }
       time.Sleep(1*time.Second)
     }
   }
 }




}




