package main

import (
	"fmt"
	"sort"
	"strings"
)

type teamLetter string
type teamUnit string

type teamUnitIndex int

const (
	tuName teamUnitIndex = iota
	tuDamage
	tuEnconters
	tuStartPosition
)

var teamUnitMap = map[teamUnitIndex]string {
	tuName: "| Team%s = ",
	tuDamage: "| DamageType%s = ",
	tuEnconters: "| Enemies%s = ",
	tuStartPosition: "| Start%s = ",
}
 
type Strategy struct {
	teams map[teamLetter]map[teamUnit]string
	canTeamsMove map[teamLetter]bool
	isTeamDead map[teamLetter]bool
	actions []string
	turns int
	maxNumOfTurns int
	// In Seconds
	timeToClear int
	encountersCount int
}

func NewStrategy() Strategy {
	return Strategy{
		teams: map[teamLetter]map[teamUnit]string{},
		canTeamsMove: map[teamLetter]bool{},
		isTeamDead: map[teamLetter]bool{},
		maxNumOfTurns: 10,
	}
}

func (strat *Strategy) AddCountEncounter() {
	strat.encountersCount += 1
}

func (strat *Strategy) GetNumOfEncounters() int {
	return strat.encountersCount
}

func (strat *Strategy) IsMaxTurns() bool {
	return strat.turns == strat.maxNumOfTurns
}

func (strat *Strategy) AdvanceTurn(playerTurn bool) {
	if strat.turns == 0 {
		strat.actions = []string{}
	}
	var turnText = ""
	if playerTurn {
		strat.turns += 1
		turnText = fmt.Sprintf("| Turn%d =", strat.turns)
		strat.ResetTeamMoves()
	} else {
		turnText = fmt.Sprintf("| Turn%dEnemy =", strat.turns)
	}
	strat.actions = append(strat.actions, turnText) 
}

func (strat *Strategy) ResetTeamMoves() {
	for k := range strat.canTeamsMove {
		if !strat.isTeamDead[k] {
			strat.canTeamsMove[k] = true
		}
	}
}

func (strat Strategy) String() string {
	var closure = func(a teamUnitIndex) string { 
		var strs []string
		for k, v := range strat.teams {
			l := toTeamUnit(k, a)
			r := v[l]
			strs = append(strs, fmt.Sprintf("%s%s\n", l, r))
		}
		// sort because map doesn't remember order
		sort.Strings(strs)
		var str string
		for _, v := range strs {
			str += v
		}
		str += "\n"
		return str
	}
	var str string
	str += "{{StrategyTable\n"
	str += closure(tuName)
	str += closure(tuDamage)
	str += closure(tuEnconters)
	str += closure(tuStartPosition)
	for _, v := range strat.actions {
		str += v+"\n"
	}
	str += "}}\n"
	return str
}

func ToTeamLetter(s string) teamLetter {
	return teamLetter(strings.ToUpper(s))
}

func ToTeamName(tl teamLetter) string {
	return "Team "+string(tl)
}

func toTeamUnit(tl teamLetter, ti teamUnitIndex) teamUnit {
	return teamUnit(fmt.Sprintf(teamUnitMap[ti], tl))
}
