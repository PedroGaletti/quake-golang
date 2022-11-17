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
	var last_line int
	count_games := 0
	var games []Game

	file, _ := os.Open(c.path)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "InitGame") {

			if len(games) < count_games+1 && len(games) >= 1 {
				games[len(games)-1].End = last_line + 1
			}

			count_games++
			games = append(games, Game{
				Name:  fmt.Sprintf("game_%d", count_games),
				Start: last_line + 1,
			})
		}

		last_line = last_line + 1
	}

	games[len(games)-1].End = last_line

	return games
}

// GetTheReports: This function it is to find the Reports of the game
func (c *ReaderController) GetTheReports(games []Game) map[string]Report {
	reports := make(map[string]Report)

	for _, game := range games {
		var last_line int
		var total_kills int64
		var players []string
		kills := make(map[string]int64)

		file, _ := os.Open(c.path)

		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if last_line > game.Start && last_line < game.End {
				if strings.Contains(scanner.Text(), "killed") {
					total_kills++

					/* Killer Player */
					killed_index := strings.Index(scanner.Text(), "killed")
					dots_index := strings.LastIndex(scanner.Text(), ":")
					words := strings.Fields(scanner.Text()[dots_index+2 : killed_index])
					username := strings.Join(words, " ")

					if words[0] != "<world>" {
						kills[username] = kills[username] + 1
						if !utils.Contains(players, username) {
							players = append(players, username)
						}
					}

					/* Dead Player */
					by_index := strings.Index(scanner.Text(), "by")
					words = strings.Fields(scanner.Text()[killed_index+len("killed") : by_index])
					username = strings.Join(words, " ")
					kills[username] = kills[username] - 1
					if !utils.Contains(players, username) {
						players = append(players, username)
					}
				}
			}

			last_line++
		}

		reports[game.Name] = Report{
			TotalKills: total_kills,
			Players:    players,
			Kills:      kills,
		}
	}

	return reports
}

// GetTheDeathsCauses: This function it is to find the death causes inside the game
func (c *ReaderController) GetTheDeathsCauses(games []Game) map[string]DeathsCauses {
	deaths_causes := make(map[string]DeathsCauses)

	for _, game := range games {
		var last_line int
		killsByMeans := make(map[string]int64)

		file, _ := os.Open(c.path)

		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if last_line > game.Start && last_line < game.End {
				if strings.Contains(scanner.Text(), "killed") {

					/* Death Cause */
					by_index := strings.Index(scanner.Text(), "by")
					words := strings.Fields(scanner.Text()[by_index+3:])
					deaths_cause := strings.Join(words, " ")
					killsByMeans[deaths_cause] = killsByMeans[deaths_cause] + 1
				}
			}

			last_line++
		}

		deaths_causes[game.Name] = DeathsCauses{
			KillsByMeans: killsByMeans,
		}
	}

	return deaths_causes
}
