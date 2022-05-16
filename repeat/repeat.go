package iteration


func Repeat(character string, repeatedCount int) string {

	var repeated string

	if repeatedCount == 0 {
		return character
	} else {

		for i := 0; i < repeatedCount; i++ {
			repeated += character
		}
	
		return repeated

	}

	
	
}