package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  stuffToDo := []string{
    "Comment on an issue or JIRA.",
    "Create a pull request for an Open Source project.",
    "Work on your interpreter.",
    "Operationalize something.",
    "Write a script to make your life easier.",
    "Think of a catchy comeback for el heffe's mockery of machine learning.",
    "Think of a way to feign that a local MongoDB is better than RDS, and ask scriptory for an opinion.",
    "Think of a new Blue Ocean Strategy.",
    "Text wife.",
    "Drink a glass of water.",
    "Eat a snack.",
    "Get up and stretch.",
    "Watch OfficeSpace.",
    "Watch an Episode of Silicon Valley.",
    "Get up and take a walk.",
    "Charge the computer.",
    "Charge iPhone.",
    "Charge Apple Watch.",
    "Listen to Podcast.",
    "Play a game of StarCraft 2.",
    "Review Cobbler configuration for home lab.",
    "Review network configuration for home lab.",
    "Schedule a SOC audit for home lab.",
    "Find a source for inexpensive Broadwell servers.",
    "Work on a DevOps request.",
  }
  rand.Seed(time.Now().UnixNano())

  choosen := stuffToDo[rand.Intn(len(stuffToDo)-1)]
  fmt.Println(choosen)
}
