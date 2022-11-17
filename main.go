package main

import (
	"encoding/json"
	"fmt"

	Reader "quake/reader"
)

func main() {
	reader := Reader.NewReaderController("logs/qgames.log")

	games := reader.FindTheGames()

	reports := reader.GetTheReports(games)
	for i := 1; i < len(reports)+1; i++ {
		game_name := fmt.Sprintf("game_%d", i)
		marshalled, _ := json.Marshal(reports[game_name])
		fmt.Println(game_name + ": " + string(marshalled))
	}

	fmt.Println()

	deaths_causes := reader.GetTheDeathsCauses(games)
	for i := 1; i < len(deaths_causes)+1; i++ {
		game_name := fmt.Sprintf("game_%d", i)
		marshalled, _ := json.Marshal(deaths_causes[game_name])
		fmt.Println(game_name + ": " + string(marshalled))
	}
}
