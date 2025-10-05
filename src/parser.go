package main

import (
	"fmt"
	"strconv"
)


func parse(tl *TokenList, strat *Strategy) Strategy {
	team(tl, strat)

	for tl.IsCurrent("team") {
		team(tl, strat)
	}

	turn(tl, strat)

	var validActions = ValidActionsPlayer
	var validExtras = ValidExtraActions
	var actionsToUse = ActionsPlayerMap
	for !tl.IsCurrent("EOF") {
		if tl.IsCurrentAny(validActions...) {
			Action(tl, strat, actionsToUse, false)
			if !tl.Accept("\n") {
				for tl.IsCurrentAny(validExtras...) {
					Action(tl, strat, ExtraActionsMap, true)
					tl.Accept("\n")
				}
			}
		} else if tl.IsCurrentAny("turn", "turnenemy") {
			switch tl.GetCurrent() {
			case "turn":
				turn(tl, strat)
				validActions = ValidActionsPlayer
				actionsToUse = ActionsPlayerMap
			case "turnenemy":
				turnEnemy(tl, strat)
				validActions = ValidActionsEnemy
				actionsToUse = ActionsEnemyMap
			}
		} else {
			tl.Err()
			break
		}
	}

	return *strat
}

func turn(tl *TokenList, strat *Strategy) {
	tl.Expect("turn")
	if strat.IsMaxTurns() {
		tl.ErrTxt("Maximun Number of Turns Reached: %d\n", strat.turns)
	}
	strat.AdvanceTurn(true)
	tl.Expect("\n")
}

func turnEnemy(tl *TokenList, strat *Strategy) {
	tl.Expect("turnenemy")
	strat.AdvanceTurn(false)
	tl.Expect("\n")
}

func Action(tl *TokenList, strat *Strategy, actionMap map[string]ActionInfo, appendToLast bool) {
	action := tl.GetCurrent()
	act, ok := actionMap[action]
	if !ok {
		return
	}
	tl.Next()

	if action == "or" {
		strat.ResetTeamMoves()
	}

	var letter teamLetter = ""
	var values []string
	for _, v := range act.FormatHint {
		var s string
		switch v {
		case FormatHintTeamLetter:
			l := ToTeamLetter(tl.GetCurrent())
			if letter == "" {
				letter = l
			}
			s = string(l)
			f, ok := FormatHintMap[v][s]
			if ok {
				s = f
			} else {
				tl.ErrTxt("No Team With Letter: %s\n", s)
			}
			tl.Next()
		case FormatHintTurns:
			s = strconv.Itoa(strat.turns)
		case FormatHintAny:
			if tl.IsCurrent("\n") {
				tl.ErrTxt("Not A Field Of Type %s: %s\n", v, tl.GetCurrent())
			}
			s = tl.GetCurrent()
			tl.Next()
		case FormatHintClearSeconds:
			s = tl.GetCurrent()
			sec, err := strconv.Atoi(s)
			if err != nil {
				tl.ErrTxt("Field of Type %s is not a number: %s", v, err)
			}
			strat.timeToClear = sec
			tl.Next()
		case FormatHintEncounterSeconds:
			s = strconv.Itoa(strat.timeToClear/strat.GetNumOfEncounters())
		default:
			s = tl.GetCurrent()
			f, ok := FormatHintMap[v][s]
			if ok {
				s = f
			} else {
				tl.ErrTxt("Not A Field Of Type %s: %s\n", v, s)
			}
			// s should only be empty if ok=true and f=""
			if s != "" {
				tl.Next()
			} else {
				// soft enable extra action
				act.AcceptsExtras = true
			}
		}
		if s != "" {
			values = append(values, s)
		}
	}

	var orderedValues []any
	if len(act.FormatHint) == len(act.FormatOrder) && len(act.FormatOrder) > 0 {
		for _, v := range act.FormatOrder {
			if v == -1 {
				continue
			}
			orderedValues = append(orderedValues, values[v])
		}
	} else {
		for _, v := range values {
			orderedValues = append(orderedValues, v)
		}
	}

	// If extra action is detected add to the same strat.action
	if appendToLast {
		strat.actions[len(strat.actions)-1] += fmt.Sprintf(act.Format, orderedValues...)
	} else {
		strat.actions = append(strat.actions, fmt.Sprintf(act.Format, orderedValues...))
	}

	// Check if this action accepts Extras
	if !tl.IsCurrent("\n") && !act.AcceptsExtras {
		tl.ErrTxt("This Action: %s Doesn't Accept Extra Actions\n", action)
	}

	if tl.IsCurrent("\n") {
		strat.actions[len(strat.actions)-1] += "."
	}

	// Check If Team can do anything
	if strat.isTeamDead[letter] {
		tl.ErrTxt("This team can't take anymore actions (dead): %s\n", letter)
	}

	// Calculate Movement
	if letter != "" && act.MoveCost > 0 && !DEBUG_DONT_CHECK_MOVE_COST {
		switch act.MoveCost {
		case 2:
			strat.isTeamDead[letter] = true
			strat.canTeamsMove[letter] = false
		case 1:
			if strat.canTeamsMove[letter] {
				strat.canTeamsMove[letter] = false
			} else {
				tl.ErrTxt("This team already used his turn: %s\n", letter)
			}
		}
	}
}

