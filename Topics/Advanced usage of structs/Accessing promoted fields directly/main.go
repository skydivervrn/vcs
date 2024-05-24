package main

import "fmt"

type FootballPlayer struct {
	Name     string
	Age      int
	Position string
	Team     // add 'Team' as an anonymous field within the FootballPlayer struct
}

type Team struct {
	teamName  string
	coachName string
}

func main() {
	// do not delete the 'player' FootballPlayer declaration!
	var player FootballPlayer
	fmt.Scanln(&player.Name, &player.Age, &player.Position)

	// do not delete the code block below:
	var tn, cn string
	fmt.Scanln(&tn, &cn)

	// Assign the two variables 'tn' and 'cn' to the promoted fields:
	player.teamName = tn
	player.coachName = cn

	fmt.Println(player) // print the player struct here!
}
