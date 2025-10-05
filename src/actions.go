package main

type ActionInfo struct {
	// User Usage help/error message
	Usage string
	// Strategy Formating String
	Format string
	// Formating Hints for the Fields passed with the command
	FormatHint []FormatHint
	// Format order of Fields, use -1 for FormatHints that return empty
	FormatOrder []int
	// If it accepts extra Fields it will not error out
	AcceptsExtras bool
	// Adds to team Encounters
	IsEncounter bool
	// Doesn't take this Encounter into Team Damage Recommentation
	SkippableEncounter bool
	// Consumes team turn (0 = no, 1 = yes)
	MoveCost int
}

func init() {
	ValidActionsPlayer = make([]string, 0, len(ActionsPlayerMap))
	for k := range ActionsPlayerMap {
		ValidActionsPlayer = append(ValidActionsPlayer, k)
	}

	ValidActionsEnemy = make([]string, 0, len(ActionsEnemyMap))
	for k := range ActionsEnemyMap {
		ValidActionsEnemy = append(ValidActionsEnemy, k)
	}

	ValidExtraActions = make([]string, 0, len(ExtraActionsMap))
	for k := range ExtraActionsMap {
		ValidExtraActions = append(ValidExtraActions, k)
	}

	for k, v := range ActionsPlayerMap {
		if v.IsEncounter {
			ValidEncounterActionsPlayer = append(ValidEncounterActionsPlayer, k)
		}
	}
	
	for k, v := range ActionsEnemyMap {
		if v.IsEncounter {
			ValidEncounterActionsEnemy = append(ValidEncounterActionsEnemy, k)
		}
	}
}

var ValidActionsPlayer []string
var ValidEncounterActionsPlayer []string

var ActionsPlayerMap = map[string]ActionInfo{
	"move": {
		Usage: "move (Letter) (Direction) {Extras}",
		Format: "* Move %s %s",
		FormatHint: []FormatHint{FormatHintTeamLetter, FormatHintMoveDirection},
		AcceptsExtras: true,
		MoveCost: 1,
	},
	"swap": {
		Usage: "swap (Letter) (Letter) {Extras}",
		Format: "* Swap %s with %s",
		FormatHint: []FormatHint{FormatHintTeamLetter, FormatHintTeamLetter},
		AcceptsExtras: true,
	},
	"wait": {
		Usage: "wait (Letter)",
		Format: "* Make %s Wait",
		FormatHint: []FormatHint{FormatHintTeamLetter},
		MoveCost: 1,
	},
	"teleport": {
		Usage: "teleport (Letter)",
		Format: "* Teleport with %s",
		FormatHint: []FormatHint{FormatHintTeamLetter},
	},
	"teleportattack": {
		Format: "** Defeat %s Enemy on the other side%s",
		FormatHint: []FormatHint{FormatHintTeamLetter, FormatHintEnemyArmorType, FormatHintTeleportAttackReturn},
		FormatOrder: []int{-1, 1, 2},
		IsEncounter: true,
	},
	"startteleport": {
		Usage: "startteleport (Letter) (StartPosition)",
		Format: "* %s Teleports to %s Start Point",
		FormatHint: []FormatHint{FormatHintTeamLetter, FormatHintStartPosition},
	},
	"attack": {
		Usage: "attack (Letter) (Direction) (Armor Types) {Extras}",
		Format: "* Attack %s Enemy %s %s",
		FormatHint: []FormatHint{FormatHintTeamLetter, FormatHintAttackDirection, FormatHintEnemyArmorType},
		FormatOrder: []int{2, 1, 0},
		AcceptsExtras: true,
		IsEncounter: true,
		MoveCost: 1,
	},
	"boss": {
		Usage: "boss (Letter) (Armor Types) (Clear Type)",
		Format: "* Attack %s Boss with %s",
		FormatHint: []FormatHint{FormatHintTeamLetter, FormatHintBossArmorType, FormatHintRequiresExtraActionForBoss},
		FormatOrder: []int{1, 0, -1},
		// there is a chance that the boss moves and steps into a gift/drone etc now are they gonna make a map that uses that idk
		// also the program doesn't support that it thinks clear takes gift as input and compĺains
		//AcceptsExtras: true,
		IsEncounter: true,
		MoveCost: 1,
	},
	"ignore": {
		Usage: "ignore (Letter)",
		Format: "* Ignore %s",
		FormatHint: []FormatHint{FormatHintTeamLetter},
		MoveCost: 2,
	},
}

var ValidActionsEnemy []string
var ValidEncounterActionsEnemy []string

var ActionsEnemyMap = map[string]ActionInfo{
	"attack": {
		Usage: "attack (Letter) (Direction) (Armor Types)",
		// Format: "* %s will be attacked by the %s %s Enemy",
		Format: "* %s will be attacked by %s Enemy %s",
		FormatHint: []FormatHint{FormatHintTeamLetter, FormatHintAttackedDirection, FormatHintEnemyArmorType},
		FormatOrder: []int{0, 2, 1},
		IsEncounter: true,
	},
	"withdraw": {
		Usage: "withdraw (Letter) (Direction) (Armor Types)",
		Format: "* Withdraw with %s from %s %s Enemy attack",
		FormatHint: []FormatHint{FormatHintTeamLetter, FormatHintWithdrawDirection, FormatHintEnemyArmorType},
		IsEncounter: true,
		SkippableEncounter: true,
		MoveCost: 2,
	},
	"boss": {
		Usage: "boss (Letter) (Armor Types) (Clear Type)",
		Format: "* %s will be attacked by %s Boss",
		FormatHint: []FormatHint{FormatHintTeamLetter, FormatHintBossArmorType, FormatHintRequiresExtraActionForBoss},
		IsEncounter: true,
	},
}

var ValidExtraActions []string

var ExtraActionsMap = map[string]ActionInfo{
	"clear": {
		Usage: "Clear",
		Format: " for a %s-turn clear",
		FormatHint: []FormatHint{FormatHintTurns},
	},
	"timed": {
		Usage: "timed (Seconds)",
		Format: " to clear in %s seconds.\n** Try to clear each battle in under %s seconds on average",
		FormatHint: []FormatHint{FormatHintClearSeconds, FormatHintEncounterSeconds},
	},
	"gift": {
		Format: ".\n** Obtain Gift Box ({{ItemCard|Pyroxene|quantity=50}})",
	},
	"teleport": {
		Format: ", and Teleport",
	},
	"noteleport": {
		Format: ", but don't Teleport",
	},
	"drone": {
		Format: ", activating the Drone",
	},
	"heal": {
		Format: ", picking up HP Recovery",
	},
	"atkbuff": {
		Format: ", picking up ATK Buff",
	},
	"defbuff": {
		Format: ", picking up DEF Buff",
	},
	/*
	"reveal": {
		// Format: ", onto 2 Reveal Tile",
	},
	"lapsed": {
		// Format: ", into 2 Lapsed Tile",
	},
	"broken": {
		// Format: ", to Broken Tile",
	},
	"toggle": {
		// Format: ", onto 2 Toggle Tile",
	},
	*/
}
