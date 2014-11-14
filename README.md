Learning the basics in Go progr. languageg. Some of the examples (all the ones located into samples folder) are taken from the reads in Go lang. official webpage -> https://golang.org/doc/code.html 

The basic doc we can find in this README.md file, is based on the contents of the book:

"An introduction to programming in GO"
Caleb Doxsey
2012



## BASIC TYPES 

**NUMBERS**

	-Integer:

		uint8 -> same as byte

		uint16
		
		uint32
		
		uint64
		
		int8 
		
		int16
		
		int32 -> same as rune
		
		int64
		
		*Machine dependent -> their size is dependent on the type of architecture of the machine -> uint, int, uintptr 

		*Go allows to increment/decrement by a unit using the operator ++/--. The language also enables to increment/decrement using the operator +=/-= .
	

	-Float:

		float32 -> single precision
	
		float64 -> double precision
	

	-Complex:
	
		complex64
	
		complex128	


**STRINGS**

Some operations:

	Length -> len("Hello world")

	Char. accessing -> "Hello World"[1] -> returns 101 instead of e as a character is represented as a byte.  

	Concatenation -> "Hello " + " world!"	


**BOOLEANS**	

-1 bit integer representing true/false

-Operations:

	&&

	||

	!	



## OTHER TYPES 

**Arrays**

	var integerArray [10]int

	x := [5]float64{ 2, 5, 3, 1}


**Slices**	

	Are like arrays, but their size is allowed to change.

	Examples:

		//slice associated with an underlying
		//float64 array of length 5
		x := make([]float64, 5)

		//slice associated with an underlying
		//float64 array of length 5,
		//where 10 is the capacity of the underlying
		//array which the slice points to
		x := make([]float64, 5, 10)

		arr := []float64{1,2,3,4,5}
		x := arr[0:4] // this will assign to x values [1,2,3,4] because the high index is not
					  //included

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


**Maps**	

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
	
	- In go, if we try to search for a key that is not into the map, it returns the zero value 
	for its value type. For example, if value is a String it will return "", or if values are integers will return 0.  

	- Go can return several elements, if we do:

		elements := make(map[string]string)
		name, ok := elements["Hello"]	

	  then name contains the result of the lookup, and ok informs if the lookup was successful.

	  Thus, Go allows to do things like:

	  	if name, ok := elements["Un"]; ok {
	  		//if lookup was successful
	    	fmt.Println(name, ok)
		}


## VARS 

-Declaration and assignation:

	var $NAME $TYPE -> Examples: 

	var x string = "Hello"

	var y string

	y = "What's up?"
	
	x := "Hello world"

	y := 5

	-No needed to add keyword var" and the type. Go compiler is able to infer the type based on 
	the literal value you assign the variable. 

	-Use this form whenever possible


-Constants:

	const x String = "Hello world"


-Definition of multiple var/consts -> Use keyword var/const and put the vars inside parentheses with each variable on its own line. Example:
	
	var (a = 5
		 b = 10
		 c = 15 )	


## CONTROL STRUCTURES 

**For**

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
	

**If**
	
	*Example:

	if i % 2 == 0 {
		// even
	} else if i % 3 == 0 {
		// divisible by 3
	} else if i % 4 == 0 {
		// divisible by 4
	}


**Switch**

	*Example:

	switch i {
		case 0: fm.Println("Zero")
		case 1: fm.Println("One")
		default: fm.Println("Unknown")
	}	


## FUNCTIONS 

-Multiple returning values: Go allows to return multiple values from a function

	func performTransaction(...) (boolean, float64) {
		// all the stuff related to performing the transaction

		currentBalance := ... //incrementing or decrementing the current balance of the account

		//informing that the transaction has been performed successfully and returning 
		the current balance
		return true, currentBalance
	}

	func main() {
		transactionOk, accountBalance := performTransaction(...)
	}


-Variadic functions: by using "..." before the type name of the last parameter we are indicating that this functions takes 0 to more of these parameters

Example: fmt.Println function

	func Println(a ...interface{}) (n int, err error)


-Closure: we can create functions inside functions. Example:

	func main() {
		sum := func(x, y int) int {
			return x + y
		}
		fmt.Println(sum(3+4))
	}


-Recursion
Example:

	func factorial(x uint) uint {
		if x == 0 {
			return 1
		}

		return x * factorial(x-1)
	}	


**SPECIAL FUNCTIONS**

-defer -> schedules a function to be run after the function completes.
	
	Example:

		f, _ := os.Open(fileName)
		defer f.Close()
		//work with the file

	This will make to close the file at the very end of the function in which this code sequence is called. 

	By this way, we can have the opening of the file and its closing very closely, so we do not forget 

	to close the file, and we can work with the file without worring about closing the file at the end.
	

-Panic and Recover	

*panic -> Causes a run time error. Generally indicates a programmer error or an exceptional condition that there's no easy way to recover from.

*recover -> Handling a run-time panic.

IMPORTANT!!! -> the call to panic immediately stops execution of the functions, so if we want to recover from a panic, the recover must be in a defer function.

Example:

	func main() {
		defer func() {
			str := recover()
			fmt.Println(str)
		}()
		panic("PANIC")
	}	


## POINTERS 

And... Yes! Go allows to work with pointers! As all of you know pointers reference a location in memory where a values is stored rather than the value itself, by this way we are able to modify the original variable.

Example:

	func zero(xPtr *int) {
    	*xPtr = 0
	}
	func main() {
	    x := 5
	    zero(&x)
	    fmt.Println(x) // x is 0
	}

A pointer can also be created using the built-in new function. Example: xPtr := new(int)

Because of being a garbage collected programming language, everything will be cleaned up automatically when nothing refers anymore.

Pointers can be really helpful when paired with structs.



## STRUCTS AND INTERFACES 

**Structs**

As probably many of you already know, a struct is a type which contians named fields and introduces a new type.

Example:
	
	type Circle struct {
		x, y, r float64
	}

-Initialization -> when it is a created a new instance of the type, it initializes to 0 state, which means all its fields are by default set to their 0 state, which becomes into for example 0 for ints, 0.0 for floats, "" for strings, nil for pointers, ...

Examples:

	var c Circle    //creates a new instance and puts all its fields to 0.0 as all of them belong 
					//to the float type.

	c := Circle{x: 0, y: 0, r: 5}
	c := Circle{0, 0, 5}

	c := new(Circle)  //allocating memory for all the fields, sets each of them to its 0 state 
					  //and returning a pointer.


IMPORTANT!!! -> remember that in Go, arguments of a function are always copied, if we want to modify fields of the object inside a function, we need to work with pointers.

Example:

	func circleArea(c *Circle) float64 {
    	return math.Pi * c.r*c.r
	}


**Methods**	

If we add a "receiver" insted of passing the object as a parameter, we can call then the function using the '.' operator. 

Example: 

	type Rectangle struct {
	    x1, y1, x2, y2 float64
	}
	func (r *Rectangle) area() float64 {
	    l := distance(r.x1, r.y1, r.x1, r.y2)
	    w := distance(r.x1, r.y1, r.x2, r.y1)
	    return l * w
	}

	main() will call:
		r := Rectangle{0, 0, 10, 10}
		r.are()


**Embedded Types**

Struct fields usually represent the has-a relationship. Anonymous fields (embedded types) -> type has a field that depends on another type.

Example:

	type Person struct {
	    Name string
	}
	func (p *Person) Talk() {
	    fmt.Println("Hi, my name is", p.Name)
	}

	And we want an Android type:

	type Android struct {
	    Person
	    Model string
	}

	As we see we define the type and don't give it a name. Then we can call:

	a := new(Android)
	a.Person.Talk()

	or directly

	a.Talk()


**Interfaces**

Like a struct, an interface is created using the type keyworkd, bollowed by a name and the keyword interface. Instead of defining fields, we define a "method set" which is the list of methods or behaviour that a type must have in order to "implement" the interface.

Example:

	type Shape interface {
		area() float64
		perimeter() float64
	}	

	func totalArea(shapes ...Shape) float64 {
		var area float64
		for _, s := range shapes {
			area += s.area()
		}
		return area
	}

	Example of call -> totalArea(&c, &r) //calculating the total area of the sum of the area of a 
										 //circle and of a rectangle

Interfaces can also be used as fields. Example:

	type MultiShape struct{
		shapes []Shape
	}	


## CONCURRENCY

Go has rich support for concurrency using goroutines and channels.

**Goroutines**

A goroutine is a function that is capable of running concurrently with other functions. 

-To create a goroutines we use the keyword "go" followed by a function invocation.

-Goroutines are lightweight and we can easily create thousands of them.

Example:

	package main

	import (
	    "fmt"
		"time"
	    "math/rand"
	)

	func f(n int) {
	    for i := 0; i < 10; i++ {
	    	fmt.Println(n, ":", i)
			amt := time.Duration(rand.Intn(250))
			time.Sleep(time.Millisecond * amt)
		} 
	}

	func main() {
		// creating 10 subroutines that perform each one the f function
	    for i := 0; i < 10; i++ {
			go f(i) 
		}
    	var input string
    	fmt.Scanln(&input) //forcing the program to not exit, because we are calling subroutines. 
    					   //Else the program would exit without waiting for the subroutines to finish. 
    					   //We can stop if we hit although any subroutine is being executed at 
    					   //the time we hit the inter.
	}


**Channels**

Allow two goroutines to communicate and synchronize their execution.

-Representation -> keyword "chan" followed by the type of the things that are passed on the channel	

-The left arrow "<-" operator is used to send and receive messages on the channel.

Example:

	package main

	import (
	    "fmt"
		"time" 
	)

	func pinger(c chan string) {
	    for i := 0; ; i++ {
			c <- "ping" 
		}
	}
	func ponger(c chan string) {
    	for i := 0; ; i++ {
			c <- "pong" 
		}
	}
	func printer(c chan string) {
	     for {
	           msg := <- c
	           fmt.Println(msg)
	           time.Sleep(time.Second * 1)
	     }
	}
	func main() {
	     var c chan string = make(chan string)  //creating the channel declaring that the messages 
	     										//that will communicate it, will be strings

	     go pinger(c)  // first sender subroutine
	     go ponger(c)  // second sender subroutine
	     go printer(c) // receiver subroutine

	     var input string
	     fmt.Scanln(&input)
	}

	// this program prints "ping" and "pong" in turns till we hit the enter key.


-A channel can be bi-directional, or we can specify a direction on a channel type. If the channel can only be sent to, attempting to receive from the channel will result in a compiler error.

Example:

	func pinger(c chan<- string) //the channel c can only be sent to	



SPECIAL SWITCH STATEMENT FOR MESSAGING: select

Example: This program prints "from 1" every 2 seconds and "from 2" every 3 seconds. select picks first channel that is ready and receives from it (or sends to it). If more than one are ready, then it randomly picks which one to receive from. 

	func main() {
	    c1 := make(chan string)
	    c2 := make(chan string)
	    go func() {
	        for {
	            time.Sleep(time.Second * 2)
	        }
	    }()
	    go func() {
	        for {
	            c2 <- "from 2"
	            time.Sleep(time.Second * 3)
	        }
	    }()
	    go func() {
			for {
				select {
	                case msg1 := <- c1:
	                    fmt.Println(msg1)
	                case msg2 := <- c2:
	                    fmt.Println(msg2)
	                case <- time.After(time.Second): //time.After creates a channel and after the 
	                								 //given duration will send the current time on it.
     					fmt.Println("timeout") 
     				default: //happens immediately if none of the channels are ready
     					fmt.Println("nothing ready")	
				} 
			}
		}()
	    var input string
	    fmt.Scanln(&input)
	}	


**Buffered Channels**	

It's also possible to pass a second parameter to the make function when creating a channel:

	c := make(chan int, 1)  //creates a buffered channel with capacity 1. 

Normally channels are synchronous; both sizes of the channel will wait until the other side is ready. A buffered channel is asynchronous; sending or receiving a message will not wait unless the channel is already full.


## PACKAGES