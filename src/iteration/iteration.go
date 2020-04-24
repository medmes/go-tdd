package iteration

func Repeat(s string, numbersOfTimes int) string {

	var repeated string
	for i := 0; i < numbersOfTimes; i++ {
		repeated += s
	}
	return repeated
}
