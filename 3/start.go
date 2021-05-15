package main

import (
    "fmt"
    "math/rand"
    "time"
)

func GenerateSlice(size int) []int {
    slice := make([]int, size, size)
    for i := 0; i < size; i++ {
        slice[i] = rand.Intn(999) - rand.Intn(999)
    }
    return slice
}

func RevertSlice(size int) []int {
    slice := make([]int, size, size)
    for i := 0; i < size; i++ {
        slice[i] = size - i
    }
    return slice
}


func NotRevertSlice(size int) []int {
    slice := make([]int, size, size)
    for i := 0; i < size; i++ {
        slice[i] = i
    }
    return slice
}

func BubbleSort(numbers []int) []int{
  //defer duration(track("BubbleSort"))
  for i:= len(numbers); i > 0; i--{
      for j := 1; j < i; j++{
          if numbers[j-1] > numbers[j] {
              t := numbers[j]
              numbers[j] = numbers[j-1]
              numbers[j-1] = t
          }
      }
  }
  return numbers
}

func InsertionSort(arr []int) []int {
    //defer duration(track("BubbleSort"))
    l := len(arr)
    for i := 1; i < l; i++ {
        for j := 0; j < i; j++ {
            if arr[j] > arr[i] {
                arr[j], arr[i] = arr[i], arr[j]
            }
        }
    }
    return arr
}

func QuickSort(a []int) []int {
    if len(a) < 2 {
        return a
    }

    left, right := 0, len(a)-1

    pivot := rand.Int() % len(a)

    a[pivot], a[right] = a[right], a[pivot]

    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }

    a[left], a[right] = a[right], a[left]

    QuickSort(a[:left])
    QuickSort(a[left+1:])

    return a
}

func track(msg string) (string, time.Time) {
    return msg, time.Now()
}

func duration(msg string, start time.Time) {
    fmt.Printf("%v: %v\n", msg, time.Since(start))
}

func main(){
/*
  slice1 := NotRevertSlice(100000)
  slice2 := RevertSlice(100000)
  InsertionSort(slice1)
  InsertionSort(slice2)
    */
    slice := GenerateSlice(11)
  fmt.Println("Standart slice: ", slice)
  fmt.Println("Sorted with insertion sort: ", InsertionSort(slice))
  fmt.Println("Sorted with bubble sort: ", BubbleSort(slice))
  fmt.Println("Sorted with bubble sort: ", QuickSort(slice))

}
