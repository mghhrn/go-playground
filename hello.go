package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"

	MyStringUtil "github.com/mghhrn/go-stringutil"
)

var c, python, java bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Printf("number of cpu cores = ")
	fmt.Printf("%d\n", runtime.NumCPU())

	fmt.Println(MyStringUtil.Reverse("!oG ,olleH"))

	fmt.Printf("Sum is = %d\n" , sum(3, 5))

	sum := func(a, b int) int { return a+b } (2, 4)
	fmt.Printf("sum using closures = %d \n", sum)
	//prime_sieve.Sieve();

	// the ':=' used instead of 'var', you can only use ':=' inside a function body.
	x, y := swap("hello", "meti")
	fmt.Println(x, y)

	var sack = 99
	myShare, yourShare := splitPotatos(sack)
	fmt.Printf("My share of potatos is %d, and your share is %d!\n", myShare, yourShare)

	var gamma int
	fmt.Println(gamma, c, python, java)

	var second, minute int = 12 , 9
	var age, sex = 27 , "male"
	fmt.Println(second, minute, age, sex)

	var anInt int
	var aFloat float64
	var aBoolean bool
	var aString string
	fmt.Printf("%v %v %v %q\n", anInt, aFloat, aBoolean, aString)

	i, j := 42, 32
	f := math.Sqrt(float64(i*i + j*j))
	u := uint(f)
	fmt.Println(i, f, u)

	v := 0.83 + 0.5i // change me!
	fmt.Printf("v is of type %T\n", v)

	const finalValue = 'a'
	fmt.Println("a =", finalValue)

	const Big = 1 << 100
	const Small = Big >> 99
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	sum = 0
	//the init and post statements are optional
	for variable := 0; variable < 10; variable++ {
		sum += variable
	}
	fmt.Println(sum)

	// this a WHILE loop
	for sum < 52 {
		sum += 1
	}
	fmt.Println(sum)

	if val := "really?"; sum == 52 {
		fmt.Println("Yaaay!", val)
	}
	// val is are in access only in its declared IF scope
	//fmt.Println(val)



}

func sum(a int , b int) int {
	return a + b
}

func sumV2(a , b int) int {
	return a + b
}

func splitPotatos(sackOfPotatos int) (myPart, yourPart int) {
	yourPart = sackOfPotatos / 2
	myPart = sackOfPotatos - yourPart
	return
}

func swap(a, b string) (string, string) {
	return b, a
}

/*****  different types in Go *******
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

--------
The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
 *************************************/

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}