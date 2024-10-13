package calculator

// simple calculator project
func Calc(operation string, n1 int, n2 int) int {
	switch operation {
	case "+":
		return n1 + n2
	case "-":
		return n1 - n2
	case "*":
		return n1 * n2
	case "/":
		if n2 == 0 {
			panic("you can not divide by zero.")
		}
		return int(n1 / n2)
	default:
		panic("invalid operation")
	}
}
