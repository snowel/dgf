package dgf
// Record to board + sequence
func RecordToBS(rec []int) ([]byte, []int) {
	size := rec[3]
	bStart := BOff + 1
	bEnd := BOff + intPow(size)

	sStart := bEnd + 2

	var b []byte
	var s []int

	for ; bStart <= bEnd; bStart++ {
		b = append(b, (byte)(rec[bStart]))
	}
	
	s = append(s, rec[sStart:]...)
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

// For a given record, validates all the moves are legal. Checks ko, self capture, play over stone
//func ValidateRecMoves(rec []int)
//From a recore, checkif a move is valid anr returns the new updated boardstate (without altering the record)
func RecValidateMove(rec []int, move int) (bool, []byte) {
	length := SeqLen(rec)
	if length == 0 {
		b, _ := RecordToBS(rec)//TODO this allows to play on top of a handycap stone. Another reason t default to ko markings
		newBoard := ApplyMove(b, move)
		return true, newBoard
	}
	current := RecordToPartial(rec, length)
	valid, next := ValidateMoveWko(current, move)

	return valid, next
}
