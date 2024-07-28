package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File
// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	ans:=0


		for _, f := range cb[file]{
			if(f){
				ans++
			}
		}
	

	return ans
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if(rank<1 || rank>8){
		return 0
	}

	ans:=0

	for _, c := range cb{
		for i, j := range c{
			if(i+1==rank && j){
				ans++
			}
		}
	}

	return ans
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	ans:=0

	for _,v := range cb{
		ans+=len(v)
	}

	return ans
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	ans:=0

	for _, v := range cb{
		for _, f := range v{
			if(f){
				ans++
			}
		}
	}

	return ans
}
