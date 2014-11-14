Learning the basics in Go progr. languageg. Some of the examples (all the ones located into samples folder) are taken from the reads in Go lang. official webpage -> https://golang.org/doc/code.html 

Other from the book:

"An introduction to programming in GO"
Caleb Doxsey
2012



********BASIC TYPES *******************

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



********OTHER TYPES *******************	

-- Arrays --

	* var integerArray [10]int

	* x := [5]float64{ 2, 5, 3, 1}


-- Slices --	

	Are like arrays, but their size is allowed to change.

	Examples:

		//slice associated with an underlying
		//float64 array of length 5
		* x := make([]float64, 5)

		//slice associated with an underlying
		//float64 array of length 5,
		//where 10 is the capacity of the underlying
		//array which the slice points to
		* x := make([]float64, 5, 10)

		* arr := []float64{1,2,3,4,5}
		  x := arr[0:4] // this will assign to x values [1,2,3,4] because the high index is not included

	Built-in functions:

		-APPEND -> creates new slice by taking an existing slice

		slice1 := []int{1,2,3}
     	slice2 := append(slice1, 4, 5)	

     	//RESULT:
     	// slice 1 value is [1,2,3]
     	// slice 2 value is [1,2,3,4,5]


     	-COPY 

     	slice1 := []int{1,2,3}
	    slice2 := make([]int, 2)
	    copy(slice2, slice1)

	    //RESULT:
	    // slice 2 now will have values [1,2] because slice2 has room for only two elements


-- Maps --	

	Examples:

		*Assigning:

			x := make(map[string]int)
			x["key"] = 10

			elements := map[string]string{
			    "H": "Hydrogen",
			    "He": "Helium",
			    "Li": "Lithium",
			    "Be": "Beryllium",
			    "B": "Boron",
			    "C": "Carbon",
			    "N": "Nitrogen",
			    "O": "Oxygen",
			    "F": "Fluorine",
			    "Ne": "Neon",
			}


		*Accessing:

			x["key"]

	Delete op. ->	delete(x, 1)	

	
	IMPORTANT!!!
	
	- In go, if we try to search for a key that is not into the map, it returns the zero value for its value type. For example, if value is a String it will return "", or if values are integers will return 0.  

	- Go can return several elements, if we do:

		elements := make(map[string]string)
		name, ok := elements["Hello"]	

	  then name contains the result of the lookup, and ok informs if the lookup was successful.

	  Thus, Go allows to do things like:

	  	if name, ok := elements["Un"]; ok {
	  		//if lookup was successful
	    	fmt.Println(name, ok)
		}


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

-- For --

	*Example:
	
	for i:=1; i <= 10; i++ {
		//stuff inside the loop
	}


	*Example 2:

	x := [5]float64{
	    98,
	    93,
	    77,
	    82,
	    83,
	}

	var total float64 = 0
	for _, value := range x {
    	total += value
	}
	fmt.Println(total / float64(len(x)))
	

-- If --
	
	*Example:

	if i % 2 == 0 {
		// even
	} else if i % 3 == 0 {
		// divisible by 3
	} else if i % 4 == 0 {
		// divisible by 4
	}


-- Switch --

	*Example:

	switch i {
		case 0: fm.Println("Zero")
		case 1: fm.Println("One")
		default: fm.Println("Unknown")
	}	


