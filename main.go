package main

import (
	"fmt"
)

type Fields struct {
	A1 string
	B1 string
	C1 string
	A2 string
	B2 string
	C2 string
	A3 string
	B3 string
	C3 string
}

func Field(A1 string, B1 string, C1 string, A2 string, B2 string, C2 string, A3 string, B3 string, C3 string) {
	fmt.Printf(`   			  A   B   C
        		1  %s |  %s |  %s
        		2  %s |  %s |  %s
        		3  %s |  %s |  %s`, A1, B1, C1, A2, B2, C2, A3, B3, C3)
	fmt.Print("\n\t\t\t------------\n")
}

func setField(pos string, player uint16, currentFields Fields) (Fields, error) {
	var targetField *string

	switch pos {
	case "A1":
		targetField = &currentFields.A1
	case "A2":
		targetField = &currentFields.A2
	case "A3":
		targetField = &currentFields.A3
	case "B1":
		targetField = &currentFields.B1
	case "B2":
		targetField = &currentFields.B2
	case "B3":
		targetField = &currentFields.B3
	case "C1":
		targetField = &currentFields.C1
	case "C2":
		targetField = &currentFields.C2
	case "C3":
		targetField = &currentFields.C3
	default:
		return currentFields, fmt.Errorf("Invalid field: %s", pos)
	}

	if *targetField != "" {
		return currentFields, fmt.Errorf("Field %s is already occupied", pos)
	}

	*targetField = getPlayerSymbol(player)
	return currentFields, nil
}

func getPlayerSymbol(player uint16) string {
	if player == 1 {
		return "X"
	}
	return "O"
}

func getFieldContent(fields Fields, field string) string {
	switch field {
	case "A1":
		return fields.A1
	case "A2":
		return fields.A2
	case "A3":
		return fields.A3
	case "B1":
		return fields.B1
	case "B2":
		return fields.B2
	case "B3":
		return fields.B3
	case "C1":
		return fields.C1
	case "C2":
		return fields.C2
	case "C3":
		return fields.C3
	default:
		return ""
	}
}

func checkWinningCombination(fields Fields, winningCombinations [][]string) bool {
	for _, combination := range winningCombinations {
		field1 := getFieldContent(fields, combination[0])
		field2 := getFieldContent(fields, combination[1])
		field3 := getFieldContent(fields, combination[2])

		if field1 == field2 && field2 == field3 && field1 != "" {
			return true
		}
	}
	return false
}

func Move() string {
	var f string
	fmt.Printf("Choose your field: ")
	fmt.Scan(&f)
	return f
}

func main() {
	var player uint16
	winningCombinations := [][]string{
		{"A1", "B2", "C3"},
		{"C1", "B2", "A3"},
		{"A1", "B1", "C1"},
		{"A2", "B2", "C2"},
		{"A3", "B3", "C3"},
		{"A1", "A2", "A3"},
		{"B1", "B2", "B3"},
		{"C1", "C2", "C3"},
	}
	fmt.Println("Hey, Welcome to the Game, GoTacToe, X is Starting")
	currentFields := Fields{}
	Field(currentFields.A1, currentFields.B1, currentFields.C1, currentFields.A2, currentFields.B2, currentFields.C2, currentFields.A3, currentFields.B3, currentFields.C3)

	for {
		for player = 1; player <= 2; player++ {
			result := Move()
			if result == "Wrong Field" {
				Field(currentFields.A1, currentFields.B1, currentFields.C1, currentFields.A2, currentFields.B2, currentFields.C2, currentFields.A3, currentFields.B3, currentFields.C3)
				player--
				fmt.Println("Wrong Field, please choose a valid field")
			} else {
				updatedFields, err := setField(result, player, currentFields)
				if err != nil {
					fmt.Printf("Error: %s\n", err)
					player--
				} else {
					currentFields = updatedFields
					Field(currentFields.A1, currentFields.B1, currentFields.C1, currentFields.A2, currentFields.B2, currentFields.C2, currentFields.A3, currentFields.B3, currentFields.C3)

					if checkWinningCombination(currentFields, winningCombinations) {
						fmt.Printf("Player %d wins\n", player)
						return
					}

					if currentFields.A1 != "" && currentFields.A2 != "" && currentFields.A3 != "" &&
						currentFields.B1 != "" && currentFields.B2 != "" && currentFields.B3 != "" &&
						currentFields.C1 != "" && currentFields.C2 != "" && currentFields.C3 != "" {
						fmt.Println("It's a draw")
						return
					}
				}
			}
		}
	}
}
