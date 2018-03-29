# Graph Theory Project for year 3 GMIT

This is my project for Graph theory for year 3 in software development in GMIT.
The project is being assigned by Dr. Ian Mcloughlin.


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



## Built With

* [Golang](https://golang.org/) - The coding Language used
* [Visual Studio Code](https://code.visualstudio.com/) - The Development text editor
* [Cmder](http://cmder.net/) - Used in place of a command line or terminal

## Research and Theory

[Shunting Yard Algorithm](https://brilliant.org/wiki/shunting-yard-algorithm/) For converting from infix to postfix notation.
[Deterministic Finite Automata](https://www.tutorialspoint.com/automata_theory/deterministic_finite_automaton.htm) DFA Example.
[Non-Deterministic Finite Automata](https://www.tutorialspoint.com/automata_theory/non_deterministic_finite_automaton.htm) NFA Example.
[Thompson's Regular Expression Explaination](https://swtch.com/~rsc/regexp/regexp1.html) To explain how to implement regular Expressions in GoLang.



## Versioning

- Version 1.0


## Authors

- Name: Michael Kidd
- Student Number: G00236920


## License

- No License Needed


## Acknowledgments

Dr. Ian Mcloughlin for the introduction and videos in creating the project.
