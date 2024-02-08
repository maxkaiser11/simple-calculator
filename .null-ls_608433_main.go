package main

import "fmt"

func main() {
  fmt.Println(("Username: \n")
  var username string
  fmt.Scan(&username)

  fmt.Prinprintln("Password: \n")

  var password string
  fmt.Scan(&password)

  if username == "maxkaiser" && password == "max98" {
    fmt.Println("Succsefully Logged In!")
  } else {
    fmt.Println("Wrong Username or Password!, Please try again.")
  }
}
