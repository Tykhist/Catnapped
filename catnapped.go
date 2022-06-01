package main

import (
  "fmt"
  "math/rand"
  "time"
)

func getRandomElement(slice []string) string {
  return slice[rand.Intn(len(slice))]
}

func main() {
  rand.Seed(time.Now().UnixNano())
  
  guestList := [8]string{"Viridian", "Cerulean", "Umber", "Cinnabar", "Chartreuse", "Periwinkle", "Ochre", "Vermillion"}
  fmt.Println("Guests:", guestList)

  catStorage := [...]string{"garage", "crate", "closet", "bathtub", "box", "cabinet"}
  fmt.Println("Hiding Places:", catStorage)

  guestAnswer := getRandomElement(guestList[:])
  catAnswer := getRandomElement(catStorage[:])
  fmt.Printf("My word! It seems that %s has hidden your cat in a %s!\n", guestAnswer, catAnswer)
}
