package iteration

func Repeat(ch string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += ch	
	}
	return repeated
}