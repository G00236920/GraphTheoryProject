package main

//Shunting Yard algorithm
//Change the string from infix to postfix notation
func intoPost(infix string) string {

	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	postfix, s := []rune{}, []rune{}

	for _, r := range infix {

		switch {

		case r == '(':

			//Add r to the stack
			s = append(s, r)

		case r == ')':

			//pop off the stack until the opening bracket is located
			//Loop while the last character is not the opening bracket
			for s[len(s)-1] != '(' {

				//Add the top of the stack to postfix
				postfix = append(postfix, s[len(s)-1])

				//add everything except the last element to the stack
				//Removing the last element as we have appended it to the postfix rune
				s = s[:len(s)-1]

			}

		case specials[r] > 0:

			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {

				//Add the top of the stack to postfix
				postfix = append(postfix, s[len(s)-1])

				//add everything except the last element to the stack
				//Removing the last element as we have appended it to the postfix rune
				s = s[:len(s)-1]

			}

			s = append(s, r)

		default:

			//Append any other character to the postfix rune
			postfix = append(postfix, r)

		}

	}

	if len(s) > 0 {

		//Add the top of the stack to postfix
		postfix = append(postfix, s[len(s)-1])

		//add everything except the last element to the stack
		//Removing the last element as we have appended it to the postfix rune
		s = s[:len(s)-1]

	}

	return string(postfix)

}
