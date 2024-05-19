package iteration

func Repeat(repeat string, times int) string {
	var repeated string
	if times < 0 {
		return "Invalid number, Overflow expected"
	}
	for i := 0; i < times; i++ {
		repeated += repeat
	}
	return repeated
}

func Repeat2(repeat string, times int) {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += repeat
		println(repeated)
	}
}

func main() {
	Repeat2("a", -2)
}
