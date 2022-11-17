package reader

type Report struct {
	TotalKills int64            `json:"total_kills"`
	Players    []string         `json:"players"`
	Kills      map[string]int64 `json:"kills"`
}

type Game struct {
	Name  string `json:"name"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

type DeathsCauses struct {
	KillsByMeans map[string]int64
}
