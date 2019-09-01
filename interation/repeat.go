package interation


func Repeat(character string, repeater int) string {
	var repeated string
	for i := 0; i < repeater; i++ {
		repeated += character
	}
	return repeated
}
