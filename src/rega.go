package main

type state struct{

	symbol rune
	edge1 *state
	edge2 *state

}


type nfa struct{

	initial *state
	accept *state

}


func poregtonfa(postfix string) *nfa {

	//Array of pointers to nfa's
	nfastack := []*nfa{}

	for _, r :=  postfix {

		switch r
		{
			case '.':

				//Pop 2 elements off the nfa stack
				//pop the last nfa off the stack
				frag2 := nfastack[len(nfastack)-1]
				//set the nfastack to everything except the last element
				nfastack = nfastack[:nfastack-1]
				//pop the last nfa off the stack
				frag1 := nfastack[len(nfastack)-1]
				//set the nfastack to everything except the last element
				nfastack = nfastack[:nfastack-1]

				//Frag1's 1st edge should point to edge2's initial state
				frag1.accept.edge1 = frag2.initial

				//push new fragment to the nfa stack,
				//set its initial state to the initial state of frag1
				//set its accept state to the accept state of frag2
				nfastack = append(nfastack, 
					&nfa(
						initial: frag1.initial, 
						accept:  frag2.accept
					)
				)

			case '|':

				//Pop 2 elements off the nfa stack
				//pop the last nfa off the stack
				frag2 := nfastack[len(nfastack)-1]
				//set the nfastack to everything except the last element
				nfastack = nfastack[:nfastack-1]
				//pop the last nfa off the stack
				frag1 := nfastack[len(nfastack)-1]
				//set the nfastack to everything except the last element
				nfastack = nfastack[:nfastack-1]

				//Frag1's 1st edge should point to edge2's initial state
				frag1.accept.edge1 = frag2.initial

				//initial state points to frag1 and frag2 initial states
				initial := state{
					edge1: frag1.initial, 
					edge2: frag2.initial
				}

				
				//Accept becomes an empty state
				accept := state{}
				frag1.accept.edge1 = &accept
				frag2.accept.edge1 = &accept

				//push new fragment to the nfa stack,
				//set its initial state to the new initial state
				//set its accept state to the new accept state
				nfastack = append(nfastack, 
					&nfa(
						initial: &initial, 
						accept:  &accept
					)
				)

			case '*':
				//pop the last nfa off the stack
				frag := nfastack[len(nfastack)-1]
				//set the nfastack to everything except the last element
				nfastack = nfastack[:nfastack-1]

				//push new fragment to the nfa stack,
				//set its initial state to the new initial state
				//set its accept state to the new accept state
				nfastack = append(nfastack, 
					&nfa(
						initial: &initial, 
						accept:  &accept
					)
				)

				//Accept becomes an empty state
				accept := state{}

				initial := state{
					edge1: frag.initial,
					edge2: &accept
				}

				frag.accept.edge1 = frag.initial
				frag.accept.edge2 = &accept

				//Push to the Nfa stack
				nfastack = append(nfastack, 
					&nfa(
						initial: &initial, 
						accept:  &accept
					)
				)

			default:
				//empty state for accept
				accept := state{}
				
				initial := state{
					symbol: r,
					edge1: &accept
				}

				//Push to the Nfa stack
				nfastack = append(nfastack, 
					&nfa(
						initial: &initial, 
						accept:  &accept
					)
				)	

		}

	}


	func pomatch(po string,s string) bool{

		//our bool value to be returned
		ismatch := false

		//Postfix to NFA regular expression
		ponfa := poregtonfa(po)

		current := []*state{}
		next := []*state{}

		//Loop through the s String one char at a time
		for _, r := range s {

			for _, c := range current {
				
				if c.symbol == r {
					
				}

			}

			//Replace current array with everything from the next arary
			//Then empty the next array and make each entry blank
			current, next = next, []*state{}

		}

		return ismatch
	}
	

	func main() {

		//Our test String
		fmt.Println(pomatch("ab.c*|", "cccc"))

	}

}