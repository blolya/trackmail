package main

import (
    "fmt"
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
  var text string
  fmt.Print("Enter path: ")
  fmt.Scanln(&text)
  dat, err := ioutil.ReadFile(text)
  check(err)
  strCounter := 1
  for i := 0; i < len(dat); i++ {
    if dat[i] == 13 {strCounter ++}
  }
  fmt.Println(strCounter)
}
