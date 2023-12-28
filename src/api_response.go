package main

type LatestCommittee struct {
	StartingRound uint64                   `json:"starting_round"`
	Members       map[string][]interface{} `json:"members"`
	TotalStake    uint64
}
