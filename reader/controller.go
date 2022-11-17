package reader

import (
	"bufio"
	"fmt"
	"os"
	"quake/utils"
	"strings"
)

// IReaderController: interface of Reader controller
type IReaderController interface {
	FindTheGames() []Game
	GetTheReports([]Game) map[string]Report
	GetTheDeathsCauses([]Game) map[string]DeathsCauses
}

// ReaderController: struct of Reader controller
type ReaderController struct {
	path string
}

// NewReaderController: create a new Reader controller
func NewReaderController(path string) IReaderController {
	return &ReaderController{path}
}

// FindTheGames: This function it is to find the Game with the properties: Name, Start (line) and End (line)
func (c *ReaderController) FindTheGames() []Game {
	var count_games, last_line int
	var games []Game

	file, _ := os.Open(c.path) // Open the file

	defer file.Close() // Close the file in the end of this scope

	scanner := bufio.NewScanner(file) // Start a new scanner of the file
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "InitGame") { // Verify if the line contains the InitGame string

			if len(games) < count_games+1 && len(games) >= 1 { // Check if the length of games is lower than count_games more one and the length of games is higher or equal one
				games[len(games)-1].End = last_line + 1 // If true, the previous game receive the last line more one, that this indicate the line of the next game
			}

			count_games++               // increase the count of games
			games = append(games, Game{ // append a new game in the games array
				Name:  fmt.Sprintf("game_%d", count_games), // Naming the game with the count_games variable
				Start: last_line + 1,                       // The next line after the start of the game
			})
		}

		last_line = last_line + 1 // increase the last line
	}

	games[len(games)-1].End = last_line // Put the last line value inside the last game

	return games // return the array of games
}

// GetTheReports: This function it is to find the Reports of the game
func (c *ReaderController) GetTheReports(games []Game) map[string]Report {
	reports := make(map[string]Report)

	for _, game := range games {
		var last_line int
		var total_kills int64
		var players []string
		kills := make(map[string]int64)

		file, _ := os.Open(c.path) // Open the file

		defer file.Close() // Close the file in the end of this scope

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if last_line > game.Start && last_line < game.End { // Check if the last line is between the game start line and the game end line
				if strings.Contains(scanner.Text(), "killed") { // Check if the line contains the "killed" word
					total_kills++ // increase the total kills because the line contains the "killed" word

					/* Killer Player */
					killed_index := strings.Index(scanner.Text(), "killed")              // Get the index of "killed" string
					dots_index := strings.LastIndex(scanner.Text(), ":")                 // Get the last index of ":" character
					words := strings.Fields(scanner.Text()[dots_index+2 : killed_index]) // Get the words between the last ":" character and the "killed" string (the sum + 2 it's because have the ": " space)
					username := strings.Join(words, " ")                                 // Convert the words array in a string

					if words[0] != "<world>" { // Check if the first index of array it's different of <world>
						kills[username] = kills[username] + 1   // The player will receive one more score
						if !utils.Contains(players, username) { // Check if the player already has in the players array
							players = append(players, username) // Append the new player in the players array
						}
					}

					/* Dead Player */
					by_index := strings.Index(scanner.Text(), "by")                               // Get the index of "by" string
					words = strings.Fields(scanner.Text()[killed_index+len("killed") : by_index]) // Get the words between the "killed" string more the length of "killed" string and "by" string
					username = strings.Join(words, " ")                                           // Convert the words array in a string
					kills[username] = kills[username] - 1                                         // Decrease one score of the player
					if !utils.Contains(players, username) {                                       // Check if the player already has in the players array
						players = append(players, username) // Append the new player in the players array
					}
				}
			}

			last_line++ // increase the last line
		}

		reports[game.Name] = Report{ // Create the report of the game
			TotalKills: total_kills,
			Players:    players,
			Kills:      kills,
		}
	}

	return reports // Return the reports
}

// GetTheDeathsCauses: This function it is to find the death causes inside the game
func (c *ReaderController) GetTheDeathsCauses(games []Game) map[string]DeathsCauses {
	deaths_causes := make(map[string]DeathsCauses)

	for _, game := range games {
		var last_line int
		killsByMeans := make(map[string]int64)

		file, _ := os.Open(c.path) // Open the file

		defer file.Close() // Close the file in the of this scope

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if last_line > game.Start && last_line < game.End { // Check if the last line is between the game start line and the game end line
				if strings.Contains(scanner.Text(), "killed") { // Check if the line contains the "killed" word

					/* Death Cause */
					by_index := strings.Index(scanner.Text(), "by")             // Get the index of "by" string
					words := strings.Fields(scanner.Text()[by_index+3:])        // Get the words after the "by" string
					deaths_cause := strings.Join(words, " ")                    // Convert the words array in a string
					killsByMeans[deaths_cause] = killsByMeans[deaths_cause] + 1 // Increase one score in the death cause
				}
			}

			last_line++ // increase the last line
		}

		deaths_causes[game.Name] = DeathsCauses{ // Create the death causes of the game
			KillsByMeans: killsByMeans,
		}
	}

	return deaths_causes // Return the death causes
}
