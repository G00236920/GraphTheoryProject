package main

func main() {

}

//Shunting Yard algorithm
//Change the string from infix to postfix notation
func intoPost(infix string) string {

	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	postfix, s := []rune{}, []rune{}

	for _, r := range infix {

	}

	return string(postfix)

}
