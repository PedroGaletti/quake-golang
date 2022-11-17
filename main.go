package main

import (
	"encoding/json"
	"fmt"

	Reader "quake/reader"
)

func main() {
	reader := Reader.NewReaderController("logs/qgames.log") // Create a new reader controller

	games := reader.FindTheGames() // Find the games inside the log file

	reports := reader.GetTheReports(games) // Get the report about each game
	for i := 1; i < len(reports)+1; i++ {
		game_name := fmt.Sprintf("game_%d", i)
		marshalled, _ := json.Marshal(reports[game_name])
		fmt.Println(game_name + ": " + string(marshalled))
	}

	fmt.Println()

	deaths_causes := reader.GetTheDeathsCauses(games) // Get the death causes about each game
	for i := 1; i < len(deaths_causes)+1; i++ {
		game_name := fmt.Sprintf("game_%d", i)
		marshalled, _ := json.Marshal(deaths_causes[game_name])
		fmt.Println(game_name + ": " + string(marshalled))
	}
}
