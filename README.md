# Graph Theory Project for year 3 GMIT

This is my project for Graph theory for year 3 in software development in GMIT.
The project is being assigned by Dr. Ian Mcloughlin.

### Project Overview
You must write a program in the Go programming language [2] that can
build a non-deterministic finite automaton (NFA) from a regular expression,
and can use the NFA to check if the regular expression matches any given
string of text. You must write the program from scratch and cannot use the
regexp package from the Go standard library nor any other external library.
A regular expression is a string containing a series of characters, some
of which may have a special meaning. For example, the three characters
“.”, “|”, and “∗” have the special meanings “concatenate”, “or”, and “Kleene
star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1,
and 1∗ means any number of 1’s. These special characters must be used in
your submission.

Graph Theory Project 2018
Other special characters you might consider allowing as input are brackets
“()” which can be used for grouping, “+” which means “at least one of”, and
“?” which means “zero or one of”. You might also decide to remove the
concatenation character, so that 1.0 becomes 10, with the concatenation
implicit.
You may initially restrict the non-special characters your program works
with to 0 and 1, if you wish. However, you should at least attempt to expand
these to all of the digits, and the characters a to z, and A to Z.
You are expected to be able to break this project into a number of smaller
tasks that are easier to solve, and to plug these together after they have been
completed. You might do that for this project as follows:
1. Parse the regular expression from infix to postfix notation.
2. Build a series of small NFA’s for parts of the regular expression.
3. Use the smaller NFA’s to create the overall NFA.
4. Implement the matching algorithm using the NFA.


### Prerequisites
- Git installed on the local machine
- Golang installed on the local machine
- An understanding of NFA conditions in infix or postfic notation

### Installing and starting the program

- install git on local machine
- Open a command prompt or terminal
- cd to the location where you wish to store the program
- type "git clone https://github.com/G00236920/GraphTheoryProject"
- Each file for the project will then download to your system
- type "cd GraphTheoryProject/src"
- Run command "Go build rega.go"
- Run command "rega"
- The program will run in the command line

## Running the tests

To test the program do the following

- Install git on local machine
- Open a command prompt or terminal
- cd to the location where you wish to store the program
- type "git clone https://github.com/G00236920/GraphTheoryProject"
- Each file for the project will then download to your system
- type "cd GraphTheoryProject/src"
- Run command "Go build rega.go"
- Run command "rega"
- The program will run in the command line

You will be asked to enter an Option either to run a test or to exit the program.
- Choose option 1 to run the test
- Now you will be asked to enter a condition, you can enter a condition either is postfix or infix notation.
	
	EXAMPLES:
	- Postfix = "ab.c*|"
	- Infix = "(a.b)c*|"
	
- Once you enter a condition you will be asked to enter a string to test against  
	
	EXAMPLES That prove true with the previous Conditions
	- ab
	- c
	- ccc
	- cccccccccccccccccccccccccccccccccccccccccccccccccccccc
	- (no entry)
	
	EXAMPLES That prove false
	- b
	- a
	- ac
	- bc
	- abc

## Deployment

- The program can be run by using an Executable file in Windows or Mac OS.

## How the Program works
When the user starts the program they are presented with an option to run a test or end the program, if the user chooses to the end the program then the programn exits.

If the user chooses to run the program then it will start the runTest method,
- Here the program requests the user to enter a Condition for the NFA.
- The program then determines if the condition is in postfix notation or infix notation based on the string containing brackets or notation
	- If the string contains a Bracket the program will Convert the string to postfix notation by calling the intoPost Method.
	- In the intoPost Function the infix string that the user entered as testString will be passed, 
	- A map of runes is created, This contains the order of precedence each Special character is taken into.
		- My Specials for the moment contains the Kleane star (*), the Concatination notation (.) and the OR operator (|).
	- Two runes are Created a postfix rune and a rune labeled s, These will be used as stacks to add and remove from.
	- We loop through the infix string one character at a time.
		- If the current charcater is an opening bracket, 
			- The character is added to the stack s.
		- If the Current character is a closing bracket, 
			- Start popping from the s stack until the opening bracket is located.
			- Append the top of the stack to the postfix stack until the opening bracket is found.
			- The s stack is then stated to be the previous s stack minus the last item on the stack.
		- If the current character is equal to one of the characters in our special map,
			- Add the top of the s Stack to the postFix Stack.
			- Add all elements from the s Stack to s Stack except the last item on the stack. Thereby removing the last item.
			- then add the current character to the s Stack.
		- If the current character does not fit one of the previous conditions,
			- Append the character to the postfix stack.
	- Now if the length of the stack s is larger than 0 (Meaning it still has characters to deal with)
		- Add the last item of the stack s to the postfix stack
		- Add all characters from the s stack to itself except the char that has been added to the postfix stack (thereby removing it from this stack)
		- loop continues until the s stack is empty.
	- Now returning the Postfix rune cast to a string. (This string is now the infix notation converted to postfix).
- The user is not requested to enter a text string that will be tested against the condition.
- The string is tested against the condition in the pomatch function, which is passed the Condition and the testString.
	- A boolean variable is created that will later be returned to tell the user if the string was accepted or rejected.
	- A variable called ponfa is created which calls the poregnfa function (Which will check our reg expression), this is passed the condition.
		- here we create a nfastack which is an array of pointers to structs, these structs contain a pointer to an initial state and a pointer to an accept state.
		- each state is also a struct which contains a rune and pointers to two edges, (Similar to a binary tree).
		- We loop through the postfix string that we took in as testString to the function.
			- If the current character is "."
				- We then take the last twp elements from the nfastack - frag1 and frag2.
					- frag2 will be the last element from the stack
					- nfastack then becomes all elements from the stack except frag2
					- frag1 will be the last element from the stack
					- nfastack then becomes all elements from the stack except frag1
				- The second last element from this (frag1) will be taken and its edge1's accept pointer will be changed to be the pointer to the accept state of the first element removed (frag2) initial state.
				- Then we add a struct to the nfastack which consists of the initial state being equal to frag1's initial state and the accept state will now be frag2's accept state.
			- If the current character is "|"
				- We then take the last twp elements from the nfastack - frag1 and frag2.
					- frag2 will be the last element from the stack
					- nfastack then becomes all elements from the stack except frag2
					- frag1 will be the last element from the stack
					- nfastack then becomes all elements from the stack except frag1
				- frag1's  edge1's accept pointer will be changed to be the pointer to the accept state of frag2's initial state.
				- A new state is created labelled initial, this state will have edge1 pointer to frag1's initial state and the edge2 will point to frag2's initial state.
				- A new state labelled accept is created and left blank.
				- both frag1 and frag2's accept state are changed to that of the new blank accept state.
				- we then add to the nfastack, a new nfa struct which contains the initial state we created and the accept state.
			- If the current character is "*"
				- we pop the last element from the nfastack
					- create a element and make it equal to the last element on the stack
					- make the nfastack equal eeverything except the last element
				- create a new blank accept state as before
				- create a new initial state, which has the initial pointer of the last element that we popped off and the blank accept state we created.
				- set the edge1 of our popped element to the elements initial state
				- set its edge2 to the blank accept state we created.
				- add to the nfastack the new nfa consisting of the initial and accept states we created.
			- if the current character is anything except the characters mentioned above.
				- 
		


## Built With

* [Golang](https://golang.org/) - The coding Language used
* [Visual Studio Code](https://code.visualstudio.com/) - The Development text editor
* [Cmder](http://cmder.net/) - Used in place of a command line or terminal

## Research and Theory

* [Shunting Yard Algorithm](https://brilliant.org/wiki/shunting-yard-algorithm/) For converting from infix to postfix notation.
* [Deterministic Finite Automata](https://www.tutorialspoint.com/automata_theory/deterministic_finite_automaton.htm) DFA Example.
* [Non-Deterministic Finite Automata](https://www.tutorialspoint.com/automata_theory/non_deterministic_finite_automaton.htm) NFA Example.
* [Thompson's Regular Expression Explaination](https://swtch.com/~rsc/regexp/regexp1.html) To explain how to implement regular Expressions in GoLang.



## Versioning

- Version 1.0


## Authors

- Name: Michael Kidd
- Student Number: G00236920


## License

- No License Needed


## Acknowledgments

Dr. Ian Mcloughlin for the introduction and videos in creating the project.
