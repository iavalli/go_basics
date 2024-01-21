// valiable initialization

var num int = 11
var num = 11
num := 11

// strict variable types
// change  the variable
num = 22

camelCase := "short variable name"

// multiple variables initialization

var (
  a string // ""
  b bool // false
  c int // 0
)

package math

import "errors"

// static error

var errCannotSum = errors.New("cannot sum")

func sum(...)

// string concatenation + printing

package pkg

import "fmt"

func PrintName() {
	firstName := "John "
	 lastName := "Smith"
	fmt.Println(firstName + lastName)
}

// functions. Arguments (x int, y int) = (x, y int). functionName in camelCase (private) or BeginsWithBigLetter (exported func
// exported functions may be used in other packages:  package.Function()

func multiply(x int, y int) int {
    return x * y
}

// one  func may return multiple results (meanings). Is used for errors management.

package math

import "errors"

func divide(x, y int) (int, error) {
    if y == 0 {
        return 0, errors.New("cannot divide on zero")
    }

    return x / y, nil
}

// bad practice: return named results with empty return:

func multiply(x, y int) (res int) {
    res = x * y
    return
}

// using exported functions:

import "fmt"

func myPrint(msg string) {
    // пакет.функция
    fmt.Println(msg)
}

// strconv package turns int to string and vice versa:

package pkg

import "strconv"

func IntToString(x int) string {
  return strconv.Itoa(x)
}

// data types
int
int64 //  for bigint
float64 // for math operations, for example for math.Max()

// math operations

x := 10
y := 5

// addition
x + y // 15

// subtraction
x - y // 5

// division
x / y // 2

// multiplication
x * y // 50

// both numbers in operation should be of similar data type:

x := 5.05
y := 10

x + y  // invalid operation: x + y (mismatched types float64 and int)

// prohibitions:

// float64 number cannot be converted to an int if there are any numbers except 0 after point
x := int64(5.05) // compilation error: constant 5.05 truncated to integer

x := int64(5.00) // this will work

// uint cannot be < 0
x := uint(-5) // constant -5 overflows uint

// returns minimal of two int values:

package pkg

import "math"

// BEGIN (write your solution here)
func MinInt(x, y int) int {
  minValue := math.Min(float64(x), float64(y))
  result := int(minValue)
  return result
}

// boolean

var a bool = false // eqv. a := false
var b bool = true // eqv. b := true

// operators for booleans:
&& // AND
== // =, IS, EQUALS
|| // OR
! // NOT, for example a != b will return true

true == false // false
false == false // true

// boleans cannot be compared directly to other data types:

true == "hello" // invalid operation: false == "hello" (mismatched types untyped bool and untyped string)

// string cannot be converted to bool
// empty string equals false in boolean type

flag := true
text := "hello"

flag && bool(text) // cannot convert text (type string) to type bool

// checking if the string is empty(false)
flag && text != "" // true

// fuction checks that id is a positive number and text is not empty:

package pkg

func IsValid(id int, text string) bool {
  return id >= 1 && text != ""
}

// Strings

var s string = "hello"







 
