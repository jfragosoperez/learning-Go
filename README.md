Learning the basics in Go progr. languageg. Some of the examples (all the ones located into samples folder) are taken from the reads in Go lang. official webpage -> https://golang.org/doc/code.html 


************* TYPES *******************

-- NUMBERS --

	-Integer:

		*uint8 -> same as byte

		*uint16
		
		*uint32
		
		*uint64
		
		*int8 
		
		*int16
		
		*int32 -> same as rune
		
		*int64
		
		*Machine dependent -> their size is dependent on the type of architecture of the machine -> uint, int, uintptr 

		*Go allows to increment/decrement by a unit using the operator ++/--. The language also enables to increment/decrement using the operator +=/-= .
	

	-Float:

		*float32 -> single precision
	
		*float64 -> double precision
	

	-Complex:
	
		*complex64
	
		*complex128	


-- STRINGS --

Some operations:

	*Length -> len("Hello world")

	*Char. accessing -> "Hello World"[1] -> returns 101 instead of e as a character is represented as a byte.  

	*Concatenation -> "Hello " + " world!"	


-- BOOLEANS --	

-1 bit integer representing true/false

-Operations:

	* &&

	* ||

	* !	


************* VARS ********************

-Declaration and assignation:

	* 	var $NAME $TYPE -> Examples: 

		var x string = "Hello"

		var y string

		y = "What's up?"
	
	* 	x := "Hello world"

		y := 5

		-No needed to add keyword var" and the type. Go compiler is able to infer the type based on the literal value you assign the variable. 

		-Use this form whenever possible


-Constants:

	* const x String = "Hello world"


-Definition of multiple var/consts -> Use keyword var/const and put the vars inside parentheses with each variable on its own line. Example:
	
	var (a = 5
		 b = 10
		 c = 15 )	


***** CONTROL STRUCTURES **************		 	

-For

	*Example:

	```python
	for i:=1; i <= 10; i++ {
		//stuff inside the loop
	}``
