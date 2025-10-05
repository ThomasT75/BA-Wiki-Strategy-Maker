package main

type FormatHint int

const (
	FormatHintAny FormatHint = iota
	FormatHintTeamLetter 
	FormatHintMoveDirection
	FormatHintAttackDirection
	FormatHintAttackedDirection
	FormatHintStartPosition
	FormatHintDamageType
	FormatHintEnemyArmorType
	FormatHintBossArmorType
	FormatHintEnemyEncounterType
	FormatHintBossEncounterType
	FormatHintTurns
	FormatHintMaskToTeamDamageType
	FormatHintConvertBetweenDamageArmorTypes
	FormatHintRequiresExtraActionForBoss
	FormatHintWithdrawDirection
	FormatHintClearSeconds
	FormatHintEncounterSeconds
	FormatHintTeleportAttackReturn

	// keep it at the bottom
	MaxFormatHint
)

func (f FormatHint) String() string {
	switch f {
	case FormatHintAny:
		return "Any"
	case FormatHintTeamLetter:
		return "TeamLetter"
	case FormatHintMoveDirection:
		return "MoveDirection"
	case FormatHintAttackDirection:
		return "AttackDirection"
	case FormatHintAttackedDirection:
		return "AttackedDirection"
	case FormatHintStartPosition:
		return "StartPosition"
	case FormatHintDamageType:
		return "DamageType"
	case FormatHintEnemyArmorType:
		return "EnemyArmorType"
	case FormatHintBossArmorType:
		return "BossArmorType"
	case FormatHintEnemyEncounterType:
		return "EnemyEncounterType"
	case FormatHintBossEncounterType:
		return "BossEncounterType"
	case FormatHintTurns:
		return "Turns"
	case FormatHintMaskToTeamDamageType:
		return "MaskToTeamDamageType"
	case FormatHintConvertBetweenDamageArmorTypes:
		return "ConvertBetweenDamageArmorTypes"
	case FormatHintRequiresExtraActionForBoss:
		return "RequiresExtraActionForBoss"
	case FormatHintWithdrawDirection:
		return "WithdrawDirection"
	case FormatHintClearSeconds:
		return "ClearSeconds"
	case FormatHintEncounterSeconds:
		return "EncounterSeconds"
	case FormatHintTeleportAttackReturn:
		return "TeleportAttackReturn"
	}

	return ""
}

var FormatHintMap = map[FormatHint]map[string]string {
	FormatHintTeamLetter: {
		// Will be propulated at runtime.
	},
	FormatHintStartPosition: {
		"top": "Top",
		"bottom": "Bottom",
		"middle": "Middle",
		"left": "Left",
		"right": "Right",
		"topleft": "Top Left",
		"topmiddle": "Top Middle",
		"topright": "Top Right",
		"bottomleft": "Bottom Left",
		"bottommiddle": "Bottom Middle",
		"bottomright": "Bottom Right",
	},
	FormatHintMoveDirection: {
		"up": "Up",
		"down": "Down",
		"left": "Left",
		"right": "Right",
		"upleft": "Up-Left",
		"upright": "Up-Right",
		"downleft": "Down-Left",
		"downright": "Down-Right",
	},
	FormatHintAttackDirection: {
		"up": "Above of",
		"down": "Below of",
		"left": "Left of",
		"right": "Right of",
		"upleft": "Up-Left of",
		"upright": "Up-Right of",
		"downleft": "Down-Left of",
		"downright": "Down-Right of",
	},
	FormatHintAttackedDirection: {
		"up": "from Above",
		"down": "from Below",
		"left": "from the Left",
		"right": "from the Right",
		"upleft": "from the Up-Left",
		"upright": "from the Up-Right",
		"downleft": "from the Down-Left",
		"downright": "from the Down-Right",
	},
	FormatHintWithdrawDirection: {
		"up": "Above",
		"down": "Below",
		"left": "Left",
		"right": "Right",
		"upleft": "Up-Left",
		"upright": "Up-Right",
		"downleft": "Down-Left",
		"downright": "Down-Right",
	},
	FormatHintTeleportAttackReturn: {
		"back": ", and Teleport Back",
		"stay": ", but don't Teleport Back",
	},
	FormatHintBossArmorType: {
		"normal": "{{TypeIcon|Normalarmor|enemy=boss}}",
		"light": "{{LightB}}",
		"heavy": "{{HeavyB}}",
		"special": "{{SpecialB}}",
		"elastic": "{{ElasticB}}",
	},
	FormatHintEnemyArmorType: {
		"normal": "{{TypeIcon|Normalarmor|enemy=}}",
		"light": "{{LightE}}",
		"heavy": "{{HeavyE}}",
		"special": "{{SpecialE}}",
		"elastic": "{{ElasticE}}",
	},
	FormatHintDamageType: {
		"any": "{{TypeIcon|Normal}}",
		"explosive": "{{Explosive}}",
		"piercing": "{{Piercing}}",
		"mystic": "{{Mystic}}",
		"sonic": "{{Sonic}}",
	},
	FormatHintEnemyEncounterType: {
		"normal": "{{TypeIcon|Normalarmor|enemy=|size=32}}",
		"light": "{{LightE|32}}",
		"heavy": "{{HeavyE|32}}",
		"special": "{{SpecialE|32}}",
		"elastic": "{{ElasticE|32}}",
	},
	FormatHintBossEncounterType: {
		"normal": "{{TypeIcon|Normalarmor|enemy=boss|size=32}}",
		"light": "{{LightB|32}}",
		"heavy": "{{HeavyB|32}}",
		"special": "{{SpecialB|32}}",
		"elastic": "{{ElasticB|32}}",
	},
	FormatHintConvertBetweenDamageArmorTypes: {
		"any": "normal",
		"explosive": "light",
		"piercing": "heavy",
		"mystic": "special",
		"sonic": "elastic",
		"normal": "any",
		"light": "explosive",
		"heavy": "piercing",
		"special": "mystic",
		"elastic": "sonic",
	},
	FormatHintMaskToTeamDamageType: {
		"0000": "Any",
		"0001": "Explosive",
		"0010": "Piercing",
		"0011": "{{TypeList|Damage|Explosive,Piercing}}",
		"0100": "Mystic",
		"0101": "{{TypeList|Damage|Explosive,Mystic}}",
		"0110": "{{TypeList|Damage|Piercing,Mystic}}",
		"0111": "{{TypeList|Damage|Explosive,Piercing,Mystic}}",
		"1000": "Sonic",
		"1001": "{{TypeList|Damage|Explosive,Sonic}}",
		"1010": "{{TypeList|Damage|Piercing,Sonic}}",
		"1011": "{{TypeList|Damage|Explosive,Piercing,Sonic}}",
		"1100": "{{TypeList|Damage|Mystic,Sonic}}",
		"1101": "{{TypeList|Damage|Explosive,Mystic,Sonic}}",
		"1110": "{{TypeList|Damage|Piercing,Mystic,Sonic}}",
		"1111": "{{TypeList|Damage|Explosive,Piercing,Mystic,Sonic}}",
	},
	FormatHintRequiresExtraActionForBoss: {
		"clear": "",
		"timed": "",
	},
}

