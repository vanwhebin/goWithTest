package iteration

func Repeat(msg string, times int) string {
	if times == 1 {
		return msg
	}
	var repeated string
	for i := 0; i < times; i++ {
		repeated += msg
	}
	return repeated
}
