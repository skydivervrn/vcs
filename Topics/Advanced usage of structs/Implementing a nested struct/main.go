package main

import "fmt"

type FootballPlayer struct {
	Name     string
	Age      int
	Position string
	Team     Team // What additional field should the FootballPlayer struct have?
}

type Team struct {
	teamName  string
	coachName string
}

// DO NOT delete or modify the code within the main function!
func main() {
	var player FootballPlayer
	fmt.Scanln(&player.Name, &player.Age, &player.Position)
	fmt.Scanln(&player.Team.teamName, &player.Team.coachName)

	fmt.Println(player)
}
