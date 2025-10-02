package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

const DEBUG_DONT_CHECK_MOVE_COST = false

var tokens TokenList = NewTokenList()
var strategy Strategy = NewStrategy()

func main()  {
	flag.BoolFunc("actions", "shows the actions for writing a strategy", func(s string) error {
		printActionHelp := func(m map[string]ActionInfo) {
			keys := make([]string, 0, len(m))
			for k := range m {
				keys = append(keys, k)
			}
			slices.Sort(keys)
			for _, k := range keys {
				v := m[k]
				print("  ", k)
				for _, fh := range v.FormatHint {
					_, ok := FormatHintMap[fh]
					if ok {
						print(" (")
						switch fh {
						case FormatHintTeamLetter:
							print(fh.String())
						default:
							first := true
							keys := make([]string, 0, len(FormatHintMap[fh]))
							for k := range FormatHintMap[fh] {
								keys = append(keys, k)
							}
							slices.Sort(keys)
							for _, k := range keys {
								if first {
									print(k)
								} else {
									print("|",k)
								}
								first = false
							}
						}
						print(")")
					} else {
						switch fh {
						case FormatHintClearSeconds:
							print(" (Seconds)")
						}
					}
				}	
				if v.AcceptsExtras {
					print(" {Extras}")
				}
				println()
			}
		}

		println("Special Actions:")
		println("  team (TeamLetter) (bottom|bottomleft|bottommiddle|bottomright|left|middle|right|top|topleft|topmiddle|topright)")
		println("  turn")
		println("  turnenemy")
		println("Player Turn Actions:")
		printActionHelp(ActionsPlayerMap)
		println("Enemy Turn Actions:")
		printActionHelp(ActionsEnemyMap)
		println("Extras:")
		printActionHelp(ExtraActionsMap)

		os.Exit(1)
		return nil
	})
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)

	// tokenization
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		line = strings.ToLower(line)
		if line == "" {
			continue
		}

		fields := strings.Fields(line)

		tokens.Append(fields...)
		tokens.Append("\n")
	}
	tokens.Append("EOF")
	// fmt.Printf("%#v", tokens.list)

	// pre-pass one
	fmt.Printf("%v",parse(&tokens, &strategy))
}
