package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	f, exists := cb[file]
    if exists {
        count := 0
        for _, i := range f{
            if i == true {
                count++
            }
        }
        return count
    } else {
        return 0
    }
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank >= 1 && rank <= 8 {
        count := 0
        for _, v := range cb {
        for i, f := range v {
            if i == (rank-1) && f == true {
                count++
            }
        }
    }
        return count
    } else {
        return 0
    }
    
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    count := 0
	for _, i := range cb {
        count += len(i)
    }
    return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    count := 0
	for k := range cb {
        count += CountInFile(cb, k)
    }
    return count
}
