package dgf

import "fmt"

const boardHeaderBytes byte = 4
const BOff int = 3// move offset to treat a move as a board index
var Starpoints19 = []int{61, 175, 289, 295, 181, 67, 73, 187, 301}

// Counts the lierties of a stone.
// -1 means there is no stone at the given inter.
func countLiberties(inter int, board []byte) int {
	if board[inter + BOff] == 0 {return -1}

	n := (int)(board[3])
	libs := 0
  if inter == 1 { // 1-1 conrner
		if board[inter + 1 + BOff] == 0 {libs++}
		if board[inter + n + BOff] == 0 {libs++}
  		return libs
	}
  if inter == n { // 19-1 corner(or equiv)
		if board[inter - 1 + BOff] == 0 {libs++}
		if board[inter + n + BOff] == 0 {libs++}
  		return libs
	}
  if inter == n * (n-1) + 1 { // 1-19 conrenr 
		if board[inter + 1 + BOff] == 0 {libs++}
		if board[inter - n + BOff] == 0 {libs++}
  		return libs
	}
	if inter == intPow(n) {
		if board[inter - 1 + BOff] == 0 {libs++}
		if board[inter - n + BOff] == 0 {libs++}
  		return libs
	}
  if (inter - 1) % n == 0 { // Left side, 1-x
		if board[inter + 1 + BOff] == 0 {libs++}
		if board[inter - n + BOff] == 0 {libs++}
		if board[inter + n + BOff] == 0 {libs++}
  		return libs
	}
  if inter % n == 0 { // rigth side n-x
		if board[inter - 1 + BOff] == 0 {libs++}
		if board[inter - n + BOff] == 0 {libs++}
		if board[inter + n + BOff] == 0 {libs++}
  		return libs
	}
  if inter < n { // top side x-1
		if board[inter + 1 + BOff] == 0 {libs++}
		if board[inter - 1 + BOff] == 0 {libs++}
		if board[inter + n + BOff] == 0 {libs++}
  		return libs
	}
  if inter > n * (n-1) { // bottom side
		if board[inter + 1 + BOff] == 0 {libs++}
		if board[inter - n + BOff] == 0 {libs++}
		if board[inter - 1 + BOff] == 0 {libs++}
  		return libs
	}

	 if board[inter - 1 + BOff] == 0 {libs++}
	 if board[inter + 1 + BOff] == 0 {libs++}
	 if board[inter - n + BOff] == 0 {libs++}
	 if board[inter + n + BOff] == 0 {libs++}
	 return libs
}

// TODO store conditional position behaviour in a function accepting a lambda
// COuld also count liberties at the same time...
// TODO if there is no need for seperate, I coudl absolutely combine them...

//For a given stone, list all stones that form a group with that stone as well as how many are in it
func getGroup(board []byte, inter int, group []int, player byte) []int {
	if Contains(inter, group){
		return group
	} else {
		group = ExclusiveAppend(group, inter)
	 }
	
	n := (int)(board[3])
  if inter == 1 { // 1-1 conrnerg
		if board[inter + 1 + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter + 1, group, player)...)}
		if board[inter + n + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter + n, group, player)...)} 
  		return group
	}
  if inter == n { // 19-1 corner(or equiv)
		if board[inter - 1 + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter - 1, group, player)...)} 
		if board[inter + n + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter + n, group, player)...)} 
  		return group
	}
  if inter == n * (n-1) + 1 { // Bottom Left corner 
		if board[inter + 1 + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter + 1, group, player)...)} 
		if board[inter - n + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter - n, group, player)...)} 
  		return group
	}
	if inter == intPow(n) { // Bottom right corner
		if board[inter - 1 + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter - 1, group, player)...)} 
		if board[inter - n + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter - n, group, player)...)} 
  		return group
	}
  if (inter - 1) % n == 0 { // Left side
		if board[inter + 1 + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter + 1, group, player)...)} 
		if board[inter - n + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter - n, group, player)...)} 
		if board[inter + n + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter + n, group, player)...)} 
  		return group
	}
  if inter % n == 0 { // rigth side n-x
		if board[inter - 1 + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter - 1, group, player)...)} 
		if board[inter - n + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter - n, group, player)...)} 
		if board[inter + n + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter + n, group, player)...)} 
  		return group
	}
  if inter < n { // top side x-1
		if board[inter + 1 + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter + 1, group, player)...)} 
		if board[inter - 1 + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter - 1, group, player)...)} 
		if board[inter + n + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter + n, group, player)...)} 
  		return group
	}
  if inter > n * (n-1) { // bottom side
		if board[inter + 1 + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter + 1, group, player)...)} 
		if board[inter - n + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter - n, group, player)...)} 
		if board[inter - 1 + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter - 1, group, player)...)} 
  		return group
	}

	 if board[inter - 1 + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter - 1, group, player)...)} 
	 if board[inter + 1 + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter + 1, group, player)...)} 
	 if board[inter - n + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter - n, group, player)...)} 
	 if board[inter + n + BOff] == player {group = ExclusiveAppend(group, getGroup(board, inter + n, group, player)...)} 
	 return group
}

func countGroupLiberties(board []byte, group []int) int {
  libs := 0
  for _, v := range group {
		libs += countLiberties(v, board)
  }
  //fmt.Printf("In countGoup libs, the group %v has %d liberties.\n", group, libs)
  return libs
}

//removes stoens and adds them to the appropriate cpature score
func captureGroup(board []byte, group []int) {
  player := 0b00000011 ^ board[group[0] + BOff] //we are assuming a non 0 array. reminder vlaues at intersect represented the color of that stone

  for _, v := range group {
		board[v + BOff] = 0
		board[player]++
  }
}

// Checks all groups on the board and removes that player's captures stones
func RemoveCaptures(player byte, board []byte) {
	for i, v := range board {
		if i <= BOff {continue}
		if v == player {
		  var group []int
			group = getGroup(board, i-BOff, group, player)
			//fmt.Println(group, i, v)
			libs := countGroupLiberties(board, group)
			if libs <= 0 {captureGroup(board, group)}
		}
	}
 }

// Applies a move without validating it.
func ApplyMove(board []byte, move int) []byte {
	player := GetPlayer(board)
  	// place stone of player x
	board[move + BOff] = player
	// check player y liberties & remove stones
	RemoveCaptures(player ^ 0b00000011, board) 
	// check player x liberties & remove stones
	RemoveCaptures(player, board)
	//Pass the turn
	SwitchPlayer(board)
	return board
}

// Checks if 2 board positions, not states, are equal
func BoardPositionEqual(b1 []byte, b2 []byte) bool {
  length := len(b1)
  if length != len(b2) {return false}
  for i := 3; i < length; i++ {
		if b1[i] != b2[i] {return false}	
  	}
	return true
}

// Verifies weather or not the move is valid.
// Simple KO validations requires computing the resulting move, so it can optionally output the resulting board.
func ValidateMove(oldBoard []byte, currentBoard []byte, move int) (bool, []byte) {
	// Check for availability of placing stone (there isn't a stone there now)
	if currentBoard[move + 2] != 0 {return false, currentBoard} 
	// Check for KO
	newBoard := ApplyMove(currentBoard, move)
	
	if BoardPositionEqual(oldBoard, newBoard) {return false, currentBoard}
	// Check for Self capture
	// Because we currently do lazy ko check, and ApplyMove will capture what needs be captured, so if the stone we just placed is capped in the new board we know we plaeyed an illigal move
	if newBoard[move + 2] == 0 {return false, currentBoard}

	return true, newBoard
}

//Moves
// Get a move from Japanese style coordinates
func JCoordToMove(col int, row int, size int) int {
	if col <= 0 || row <= 0 || col > size || row > size {
		return -1
	}

	return col + (row-1) * size
}

/*  Set data in Board  */

func SwitchPlayer(board []byte) {
	board[0] = board[0] ^ 0b00000011
}

/*  Get data from Board  */

func GetPlayer(board []byte) byte {
	return board[0] & 0b00000011
}

func GetCaps(board []byte) (float32, float32) {
	komi := board[0] & 0b00001100
	bCaps := (float32)(board[1])
	wCaps := (float32)(board[2])
	switch komi {
	case 4:
		bCaps += 0.5
	case 8:
	 	wCaps += 0.5
	case 12:
		bCaps += 0.5
		wCaps += 0.5
	}
	return bCaps, wCaps
}

// Record to board + sequence
func RecordToBS(rec []int) ([]byte, []int) {
	size := rec[3]
	bStart := BOff + 1
	bEnd := BOff + intPow(size)

	sStart := bEnd + 2
	sEnd := len(rec)

	var b []byte
	var s []int

	for ; bStart <= bEnd; bStart++ {
		b = append(b, (byte)(rec[bStart]))
	}
	
	for ; sStart <= sEnd;sStart++ {
		s = append(s, (rec[sStart]))
	}
	return b, s
}

//How many moves are in the sequence of moves?
func SeqLen(rec []int) int {
	length := len(rec)
	size := rec[3]

	return length - intPow(size) - 5
}

// Record to nth move board
// For n = SeqLen -> current/most recent board
func RecordToPartial(rec []int, n int) []byte {
	board, seq := RecordToBS(rec)
	for i, _ := range seq {
		if i == n {break}
		ApplyMove(board, seq[i])
	}

	return board
}

/*---  Utills  ---*/
/*--- --- Graphical  ---*/

// Print using Nearfont format
// Currnelty only supports 19x19
func NfPrintBoard(board []byte) {
  size := (int)(board[3])
  if size != 19 {
		fmt.Println("Sorry, only 19x19 boards work for now...")
		return
  }
  s := "This is you current board.\n"
  player := GetPlayer(board)
  var pstone string
 	switch player {
	case 1:
	  pstone = ""
	case 2:
	  pstone = ""
	}
	bcap, wcap := GetCaps(board)
	s += fmt.Sprintf("Current Move: %s : %f : %f\n", pstone, bcap, wcap)
  for index, v := range board {
		if index <= BOff {continue}
	 i := index - BOff
	 if v == 0 {
		 if Contains(i, Starpoints19){
			s += fmt.Sprintf("󰧞 ")
		 } else {
			s += fmt.Sprintf("󰧟 ")
		 }
	 }
	 if v == 1 {
			s += fmt.Sprintf(" ")
	 }
	 if v == 2 {
			s += fmt.Sprintf(" ")
	 }
	if i % size == 0 {
		s += fmt.Sprintf("\n")
	}
}
	fmt.Println(s)
}


// INfamous contains
func Contains[E comparable](val E, slice []E) bool {
  for _, v := range slice {
		if v == val {return true}
  }
  return false
}

func intPow(x int) int {
	return x * x
}

func ExclusiveAppend[E comparable](slice []E, val ...E) []E {
	for _, v := range val {
		if !Contains(v, slice) {slice = append(slice, v)}
	}

	return slice
}
