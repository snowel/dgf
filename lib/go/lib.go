package dgf

const boardHeaderBytes byte = 4
const BOff int = 3 // move offset to treat a move as a board index

/* -- Editor funcitons -- */
// Modifying a board regardless of gameplay rules


/* -- Game play functions -- */
// manipulating a board under the rules of a game

// Checks if an intersection is free of stones
// as of of now: either an insersetion in empty or a ko liberty
func IsLib(board []byte, inter int) bool {
	if board[inter + BOff] == 0 { return true }
	if board[inter + BOff] == 3 { return true }
	return false
}


// Counts the lierties of a stone.
// -1 means there is no stone at the given inter.
func countLiberties(inter int, board []byte) int {
	if board[inter + BOff] == 0 {return -1}

	n := (int)(board[3])
	libs := 0
  if inter == 1 { // 1-1 conrner
		if IsLib(board, inter + 1) {libs++}
		if IsLib(board, inter + n) {libs++}
  		return libs
	}
  if inter == n { // 19-1 corner(or equiv)
		if IsLib(board, inter - 1) {libs++}
		if IsLib(board, inter + n) {libs++}
  		return libs
	}
  if inter == n * (n-1) + 1 { // 1-19 conrenr 
		if board[inter + 1 + BOff] == 0 {libs++}
		if IsLib(board, inter - n) {libs++}
  		return libs
	}
	if inter == intPow(n) {
		if IsLib(board, inter - 1) {libs++}
		if IsLib(board, inter - n) {libs++}
  		return libs
	}
  if (inter - 1) % n == 0 { // Left side, 1-x
		if IsLib(board, inter + 1) {libs++}
		if IsLib(board, inter - n) {libs++}
		if IsLib(board, inter + n) {libs++}
  		return libs
	}
  if inter % n == 0 { // rigth side n-x
		if IsLib(board, inter - 1) {libs++}
		if IsLib(board, inter - n) {libs++}
		if IsLib(board, inter + n) {libs++}
  		return libs
	}
  if inter < n { // top side x-1
		if IsLib(board, inter + 1) {libs++}
		if IsLib(board, inter - 1) {libs++}
		if IsLib(board, inter + n) {libs++}
  		return libs
	}
  if inter > n * (n-1) { // bottom side
		if IsLib(board, inter + 1) {libs++}
		if IsLib(board, inter - n) {libs++}
		if IsLib(board, inter - 1) {libs++}
  		return libs
	}

	 if IsLib(board, inter - 1) {libs++}
	 if IsLib(board, inter + 1) {libs++}
	 if IsLib(board, inter - n) {libs++}
	 if IsLib(board, inter + n) {libs++}
	 return libs
}

// TODO store conditional position behaviour in a function accepting a lambda
// COuld also count liberties at the same time...
// TODO if there is no need for seperate, I coudl absolutely combine them...

//For a given stone, list all stones that form a group with that stone as well as how many are in it
// TODO player grabbed from intersection value
func GetGroup(board []byte, inter int) []int {
	var group []int
	if IsLib(board, inter) { return group }
	player := board[inter + BOff]
	return getGroup(board, inter, group, player)
}

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

//TODO really beed to extract the neighboring intersectoin if checks with anonymous funcitons
// For a group gie the liberties of the group
func listGroupLiberties(board []byte, group []int) []int {
  var libs []int
	n := (int)(board[3])
  for _, v := range group {
  if v == 1 { // 1-1 conrner
		if IsLib(board, v + 1) { libs = ExclusiveAppend(libs, v + 1) }
		if IsLib(board, v + n) { libs = ExclusiveAppend(libs, v + n) }
		continue
	}
  if v == n { // 19-1 corner(or equiv)
		if IsLib(board, v - 1) { libs = ExclusiveAppend(libs, v - 1) }
		if IsLib(board, v + n) { libs = ExclusiveAppend(libs, v + n) }
		continue
	}
  if v == n * (n-1) + 1 { // 1-19 conrenr 
		if IsLib(board, v + 1) { libs = ExclusiveAppend(libs, v + 1) }
		if IsLib(board, v - n) { libs = ExclusiveAppend(libs, v - n) }
		continue
	}
	if v == intPow(n) {
		if IsLib(board, v - 1) { libs = ExclusiveAppend(libs, v - 1) }
		if IsLib(board, v - n) { libs = ExclusiveAppend(libs, v - n) }
		continue
	}
  if (v - 1) % n == 0 { // Left side, 1-x
		if IsLib(board, v + 1) { libs = ExclusiveAppend(libs, v + 1) }
		if IsLib(board, v - n) { libs = ExclusiveAppend(libs, v - n) }
		if IsLib(board, v + n) { libs = ExclusiveAppend(libs, v + n) }
		continue
	}
  if v % n == 0 { // rigth side n-x
		if IsLib(board, v - 1) { libs = ExclusiveAppend(libs, v - 1) }
		if IsLib(board, v - n) { libs = ExclusiveAppend(libs, v - n) }
		if IsLib(board, v + n) { libs = ExclusiveAppend(libs, v + n) }
		continue
	}
  if v < n { // top side x-1
		if IsLib(board, v + 1) { libs = ExclusiveAppend(libs, v + 1) }
		if IsLib(board, v - 1) { libs = ExclusiveAppend(libs, v - 1) }
		if IsLib(board, v + n) { libs = ExclusiveAppend(libs, v + n) }
		continue
	}
  if v > n * (n-1) { // bottom side
		if IsLib(board, v + 1) { libs = ExclusiveAppend(libs, v + 1) }
		if IsLib(board, v - n) { libs = ExclusiveAppend(libs, v - n) }
		if IsLib(board, v - 1) { libs = ExclusiveAppend(libs, v - 1) }
		continue
	}

	 if IsLib(board, v - 1) { libs = ExclusiveAppend(libs, v - 1) }
	 if IsLib(board, v + 1) { libs = ExclusiveAppend(libs, v + 1) }
	 if IsLib(board, v - n) { libs = ExclusiveAppend(libs, v - n) }
	 if IsLib(board, v + n) { libs = ExclusiveAppend(libs, v + n) }
  }
  return libs
}

// For a group with one liberty, return the liberty
//int -1 : group has more than 1 liberty or no libeties
func groupAtary(board []byte, group []int) int {
	libList := listGroupLiberties(board, group)
	if len(libList) == 1 { return libList[0] }
	return -1
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
			group := GetGroup(board, i-BOff)
			//fmt.Println(group, i, v)
			libs := countGroupLiberties(board, group)
			if libs <= 0 {captureGroup(board, group)}
		}
	}
}

// Premptive marking illigal ko moves
// Mut
func MarkKo(past []byte, present []byte, lastMove int) bool {
	// Check if the stone played is a single stone (Multi capture cannot be ko)
	group := GetGroup(present, lastMove)
	if len(group) != 1 { return false }
	// Check if the stone played has a single liberty
	if countGroupLiberties(present, group) != 1 { return false }
	// Check if capturing the stone would result in the past board position
	recap := groupAtary(present, group)
	if BoardPositionEqual(ApplyMove(present, recap), past) {
		// Mut the present board with ko mark
		present[recap + BOff] = 3
		return true
	}
	return false
}

// Clear any marking of ko on a board
// Mut
func UnmarkKo(b []byte) {
	for i, v := range b {
		if i < BOff { continue }
		if v == 3 { b[i] = 0 }
	}
}

func ApplyMove(board []byte, move int) []byte {
	player := GetPlayer(board)
	var newBoard []byte
	newBoard = append(newBoard, board...)
	// place stone of player x
	newBoard[move + BOff] = player
	// check player y liberties & remove stones
	RemoveCaptures(player ^ 0b00000011, newBoard) 
	// check player x liberties & remove stones
	RemoveCaptures(player, newBoard)
	//Pass the turn
	SwitchPlayer(newBoard)
	return newBoard
}

// Checks if 2 board positions, not states, are equal
func BoardPositionEqual(b1 []byte, b2 []byte) bool {
  length := len(b1)
  if length != len(b2) {return false}
  for i := BOff; i < length; i++ {
		if b1[i] != b2[i] {return false}	
  	}
	return true
}

// Valitate assuming marked kos
func ValidateMoveWko(board []byte, move int) (bool, []byte) {
	// Check the spot is open and not ko
	if board[move + BOff] != 0 { return false, board }
	//TODO instead of zero check, eventually a IsLegal could replace it, for anotations and other uses of the byte
	newBoard := ApplyMove(board, move)
	if newBoard[move + BOff] == 0 { return false, board }

	//Unmakr previous kos
	UnmarkKo(newBoard)
	// Mark the ko
	MarkKo(board, newBoard, move)

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

// Mut
func SwitchPlayer(board []byte) {
	board[0] = board[0] ^ 0b00000011
}

/*  Get data from Board  */

func GetPlayer(board []byte) byte {
	return board[0] & 0b00000011
}

func GetSize(board []byte) int {
	return (int)(board[3])
	// TODO bytes can be named i.e. enum for magic, size, etc
	// Esspetially goodif they change
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
