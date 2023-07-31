package utils

// Enum to get all move types
const (
	IntersectionDiagonalMove = 2
	FirstDiagonalMove        = 0
	SecondDiagonalMove       = 1
	NoDiagonalMove           = -1
)

// Enum for Result of the Tic Tac Toe
const (
	GameWinningString = "Congratulations %s, you have won the game :-)\n"
	Draw              = "Game is Draw"
)

// Enum for Players who will start
const (
	HashTag                      = "#"
	X                            = "X"
	O                            = "O"
	AutoAssigningChoicesFallback = "Auto-assigning %s to %s"
	TossWinningString            = "Hey %s, you have won the toss what would you like to choose between %s %s?"
	AssignOtherChoiceString      = "Assigning %s to %s\n"
	PlayerTurnToPick             = "Hey %s, please select {rowIndex, " +
		"colIndex} with which you want to mark board with %s: "
	InvalidMove = "INVALID_MOVE %s as it is pre-filled, you need to retry again :-(\n"
)
