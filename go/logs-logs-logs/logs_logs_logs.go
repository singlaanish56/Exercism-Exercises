package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	ans:="default"
	for _, char := range log{
		
		if(char=='\U00002757'){
			ans ="recommendation"
			break
		}else if(char=='\U0001F50D'){
			ans ="search"
			break
		}else if(char=='\U00002600'){
			ans="weather"
			break
		}
	}

	return ans
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	
	runes := []rune(log)
	for i, c := range runes{
		if(c==oldRune){
			runes[i] = newRune
		}
	}

	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	 x := utf8.RuneCountInString(log)

	 return x<=limit
}
