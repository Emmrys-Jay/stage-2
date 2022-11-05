package parser

var Operations = []string{"addition", "subtraction", "multiplication"}

func ParseRequest(op string, x int, y int) (operation string, result int) {

	switch op {
	case Operations[0]:
		operation = Operations[0]
		result = x + y
	case Operations[1]:
		operation = Operations[1]
		result = x - y
	case Operations[2]:
		operation = Operations[2]
		result = x * y
	default:
		operation = "You passed in an invalid operation"
	}

	return
}

func ParseRandomString() {

}
