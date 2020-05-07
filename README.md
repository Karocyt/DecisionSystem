# Expert System
Lexer and backward chaining inference engine written as two independant Go Modules in Golang.

## Input format
(see included valid and invalid examples for the full range of possibilities, as some operators are not accepted on the right side of the equation in my implementation)  
One instruction by line, using Capitalized keys, in 3 sections:
### Operations / Rules
	Forty + Two  = Fortytwo
	Rain | Umbrella = !Sun
The following symbols are defined, in order of decreasing priority:  
• ( and ) which are fairly obvious. Example : A + (B | C) => D  
• ! which means NOT. Example : !B  
• + which means AND. Example : A + B  
• | which means OR. Example : A | B  
• ˆ which means XOR. Example : A ˆ B  
• => which means "implies". Example : A + B => C
### Facts / Statements
	=AB CUmbrella
A single line starting by an equal sign.  
State which keys/symbols are defined as True.  
As keys are Capitalized, spacing doesn't matter.  
The above example implies that both A, B, C and Umbrella are True.  
### Queries
	?SunA
A single line starting by a question mark.  
The engine should be able to output the booleans corresponding to these Keys once rules have been applied.

## Lexer
State functions based lexer following [Rob Pike presentation on Lexical Scanning in Go](https://www.youtube.com/watch?v=HxaD_trXwRE)
Lexical errors shown using gcc's syntax.

## Parser / Engine
Builds a tree of Nodes using Go interfaces.
Each Node is an operation involving other Nodes and has an Eval method trickling down the tree to return a boolean.
Keys Nodes have only one child (their defining operation or True/False leaves Nodes) and default to False if childless.
Hence, True and False are childless operations, evaluating to their respective booleans.
When "trickling down", a list of used nodes is built and an error is raised if we enter a recursive loop (if at any point, evaluating a key involves... evaluating this same key).

# Compilation

	go build -o expert_system
Compile expert_system app

# Usage

./expert_system testdata/input1.txt

### First Golang project caveats
As this was my first project ever using this wonderfull language, there might be a few non-idiomatic things.
For instance, all methods and variables are exported (Capitalized). This was very usefull for debugging as my code moved around a lot but most struct members and methods could/should/must be set to private if these modules were to be included in any production-grade project.
Also, coming from a C background, I might have done manually things that would be more efficiently done by the wonderfull standard library.

## Author
* **Kevin Azoulay** @ 42 Lyon
