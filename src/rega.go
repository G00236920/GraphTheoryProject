package main

import (
	"fmt"
)

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

type nfa struct {
	initial *state
	accept  *state
}

func poregtonfa(postfix string) *nfa {

	//Array of pointers to nfa's
	nfastack := []*nfa{}

	for _, r := range postfix {

		switch r {
		case '.':

			//Pop 2 elements off the nfa stack
			//pop the last nfa off the stack
			frag2 := nfastack[len(nfastack)-1]
			//set the nfastack to everything except the last element
			nfastack = nfastack[:len(nfastack)-1]
			//pop the last nfa off the stack
			frag1 := nfastack[len(nfastack)-1]
			//set the nfastack to everything except the last element
			nfastack = nfastack[:len(nfastack)-1]

			//Frag1's 1st edge should point to edge2's initial state
			frag1.accept.edge1 = frag2.initial

			//push new fragment to the nfa stack,
			//set its initial state to the initial state of frag1
			//set its accept state to the accept state of frag2
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})

		case '|':

			//Pop 2 elements off the nfa stack
			//pop the last nfa off the stack
			frag2 := nfastack[len(nfastack)-1]
			//set the nfastack to everything except the last element
			nfastack = nfastack[:len(nfastack)-1]
			//pop the last nfa off the stack
			frag1 := nfastack[len(nfastack)-1]
			//set the nfastack to everything except the last element
			nfastack = nfastack[:len(nfastack)-1]

			//Frag1's 1st edge should point to edge2's initial state
			frag1.accept.edge1 = frag2.initial

			//initial state points to frag1 and frag2 initial states
			initial := state{edge1: frag1.initial, edge2: frag2.initial}

			//Accept becomes an empty state
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			//push new fragment to the nfa stack,
			//set its initial state to the new initial state
			//set its accept state to the new accept state
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case '*':
			//pop the last nfa off the stack
			frag := nfastack[len(nfastack)-1]
			//set the nfastack to everything except the last element
			nfastack = nfastack[:len(nfastack)-1]

			//Accept becomes an empty state
			accept := state{}

			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			//Push to the Nfa stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		default:
			//empty state for accept
			accept := state{}

			initial := state{symbol: r, edge1: &accept}

			//Push to the Nfa stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		}

	}

	return nfastack[0]
}

func addState(listState []*state, singleState *state, acceptState *state) []*state {

	listState = append(listState, singleState)

	if singleState != acceptState && singleState.symbol == 0 {

		//Recusive Function of the addState
		listState = addState(listState, singleState.edge1, acceptState)

		//edge2 is not equal the nil value run the method again
		//If the state has a second edge
		if singleState.edge2 != nil {

			//Recusive Function of the addState
			listState = addState(listState, singleState.edge1, acceptState)

		}

	}

	return listState

}

func pomatch(po string, s string) bool {

	//our bool value to be returned
	ismatch := false

	//Postfix to NFA regular expression
	ponfa := poregtonfa(po)

	current := []*state{}
	next := []*state{}

	//Add all the initial states to the current array
	current = addState(current[:], ponfa.initial, ponfa.accept)

	//Loop through the s String one char at a time
	for _, r := range s {

		for _, c := range current {

			if c.symbol == r {

				next = addState(next[:], c.edge1, ponfa.accept)

			}

		}

		//Replace current array with everything from the next arary
		//Then empty the next array and make each entry blank
		current, next = next, []*state{}

	}

	//Loop through the current state
	for _, c := range current {

		//if the current state matches the accept state in ponfa then return true
		//this means the current condition of the string is within an accept state in our Automataton
		if c == ponfa.accept {
			ismatch = true
			break
		}

	}

	return ismatch
}

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

func main() {

	//Our test String
	fmt.Println(pomatch(intoPost("ab.c*|"), "cccc"))

}
