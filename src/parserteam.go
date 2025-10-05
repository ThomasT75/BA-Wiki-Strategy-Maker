package main

import "fmt"

var StrToMaskTeamDamage = map[string]int {
	"explosive": 1,
	"piercing": 2,
	"mystic": 4,
	"sonic": 8,
}

func team(tl *TokenList, strat *Strategy) {
	tl.Expect("team")
	// get team letter
	letter :=	ToTeamLetter(tl.GetCurrent())
	strat.teams[letter] = make(map[teamUnit]string)
	tl.Next()
	// let strategy know the team can move
	strat.canTeamsMove[letter] = true
	// set team name with letter
	var unit teamUnit
	unit = toTeamUnit(letter, tuName)
	strat.teams[letter][unit] = ToTeamName(letter)
	// get start position
	unit = toTeamUnit(letter, tuStartPosition)
	strat.teams[letter][unit] = FormatHintMap[FormatHintStartPosition][tl.GetCurrent()]
	tl.Next()

	tl.Expect("\n")

	// create temporary snapshots
	tltmp := *tl
	strattmp := *strat
	var armors = []string{}
	var armorsSkip = []bool{}
	var validEnconter = ValidEncounterActionsPlayer
	var actionsToUse = ActionsPlayerMap
	for !tltmp.IsCurrent("EOF") {
		if tltmp.IsCurrentAny(validEnconter...) {
			action, ok := actionsToUse[tltmp.GetCurrent()]
			if !ok {
				return
			}
			tltmp.Next()
			// extract armor and letter info
			var armor = ""
			var actionLetter teamLetter = ""
			var fHint FormatHint
			for _, v := range action.FormatHint {
				switch v {
				case FormatHintEnemyArmorType:
					armor = tltmp.GetCurrent()
					fHint = FormatHintEnemyEncounterType
				case FormatHintBossArmorType:
					armor = tltmp.GetCurrent()
					fHint = FormatHintBossEncounterType
				case FormatHintTeamLetter:
					actionLetter = ToTeamLetter(tltmp.GetCurrent())
				}
				// format hint can be longer than number of fields
				if tltmp.IsCurrent("\n") {
					break
				}
				tltmp.Next()
			}
			if actionLetter == letter {
				armors = append(armors, armor)
				armorsSkip = append(armorsSkip, action.SkippableEncounter)
				unit = toTeamUnit(letter, tuEnconters)
				s := armor
				f, ok := FormatHintMap[fHint][armor] 
				if ok {
					s = f
				}
				if strat.teams[letter][unit] != "" {
					s = " "+s
				}
				strat.teams[letter][unit] += s
				strat.AddCountEncounter()
			}
			// eat extras
			for !tltmp.IsCurrent("\n") { 
				tltmp.Next()
			}
			tltmp.Next()
		} else if tltmp.IsCurrentAny("turn", "turnenemy") {
			switch tltmp.GetCurrent() {
			case "turn":
				turn(&tltmp, &strattmp)
				validEnconter = ValidEncounterActionsPlayer
				actionsToUse = ActionsPlayerMap
			case "turnenemy":
				turnEnemy(&tltmp, &strattmp)
				validEnconter = ValidEncounterActionsEnemy
				actionsToUse = ActionsEnemyMap
			}
		} else {
			for !tltmp.IsCurrent("\n") {
				tltmp.Next()
			}
			tltmp.Next()	
		}
	}

	// do damage type
	var at = ""
	var atNum = 0
	for i, v := range armors {
		if armorsSkip[i] {
			continue
		}
		f, ok := FormatHintMap[FormatHintConvertBetweenDamageArmorTypes][v]
		if !ok {
			continue
		}
		i, ok := StrToMaskTeamDamage[f]
		if !ok {
			continue
		}
		atNum |= i
	}
	at = fmt.Sprintf("%04b", atNum)
	unit = toTeamUnit(letter, tuDamage)
	strat.teams[letter][unit] += FormatHintMap[FormatHintMaskToTeamDamageType][at]

	if atNum & 1 > 0 {
		FormatHintMap[FormatHintTeamLetter][string(letter)] += 
			FormatHintMap[FormatHintDamageType]["explosive"]
	}
	if atNum & 2 > 0 {
		FormatHintMap[FormatHintTeamLetter][string(letter)] += 
			FormatHintMap[FormatHintDamageType]["piercing"]
	}
	if atNum & 4 > 0 {
		FormatHintMap[FormatHintTeamLetter][string(letter)] += 
			FormatHintMap[FormatHintDamageType]["mystic"]
	}
	if atNum & 8 > 0 {
		FormatHintMap[FormatHintTeamLetter][string(letter)] += 
			FormatHintMap[FormatHintDamageType]["sonic"]
	}
	if atNum == 0 {
		FormatHintMap[FormatHintTeamLetter][string(letter)] += 
			FormatHintMap[FormatHintDamageType]["any"]
	}
	unit = toTeamUnit(letter, tuName)
	FormatHintMap[FormatHintTeamLetter][string(letter)] += 
		" "+strat.teams[letter][unit]
} 
