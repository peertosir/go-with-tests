package iteration

func Repeat(str string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += str
	}
	return
}
