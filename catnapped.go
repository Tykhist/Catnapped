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
  /*
  Currently, the code only chooses a random guest and place from the arrays.
  However, I want to add inputs to make this an actual game.
  The player should get 3 guesses and an option to quit, with unique dialogue.
  There should also be a section that restarts the game in the player chooses to, if possible.
  */
  rand.Seed(time.Now().UnixNano())
  
  guestList := [8]string{"Viridian", "Cerulean", "Umber", "Carmine", "Chartreuse", "Zaffre", "Ochre", "Vermillion"}
  fmt.Println("Guests:", guestList)

  catStorage := [...]string{"garage", "crate", "closet", "bathtub", "box", "cabinet"}
  fmt.Println("Hiding Places:", catStorage)

  guestAnswer := getRandomElement(guestList[:])
  catAnswer := getRandomElement(catStorage[:])
  fmt.Printf("My word! It seems that %s has hidden your cat in a %s!\n", guestAnswer, catAnswer)
}
