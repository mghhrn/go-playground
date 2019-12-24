package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
	"strings"
	"time"

	MyStringUtil "github.com/mghhrn/go-stringutil"
)

var c, python, java bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

type Vertex struct {
	X int
	Y int
}

type Student struct {
	name string
	nationalCode int
	averageScore float32
}

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


	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}


	// this type of switch case is useful for long if-then-else chains
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}


	// the excecution of below line will be deferred until the surrounding function returns
	// Deferred function calls are pushed onto a stack and executed in last-in-first-out order.
	defer fmt.Println("Step 2")
	fmt.Println("Step 1")


	var intPointer *int
	someInt := 12
	intPointer = &someInt
	fmt.Println("someInt POINTER value = " , intPointer)
	fmt.Println("someInt BEFORE = " , *intPointer)
	*intPointer = *intPointer + 5
	fmt.Println("someInt AFTER = " , *intPointer)

	vert := Vertex{12, 4}
	vert.X = 44
	fmt.Println(vert)
	poinerToStruct := &vert
	poinerToStruct.Y = 72
	fmt.Println(*poinerToStruct)

	studentPointer := &Student{
		name:         "Jaber Khalil",
		nationalCode: 1113334445,
		averageScore: 19.4,
	}
	student2 := Student{}

	if student2.name == "" {
		fmt.Println("Student seems like a ghost!")
	}

	fmt.Println(studentPointer.name)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	s := primes[1:4]
	fmt.Println(s)

	//slices are like references to arrays
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	par1 := names[0:2]
	par2 := names[1:]
	fmt.Println(par1, par2)

	par2[0] = "XXX"
	fmt.Println(par1, par2)
	fmt.Println(names)

	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s2)

	s3 := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("len=%d cap=%d %v\n", len(s3), cap(s3), s3)

	// Slice the slice to give it zero length.
	s3 = s3[:0]
	fmt.Printf("len=%d cap=%d %v\n", len(s3), cap(s3), s3)

	// Extend its length.
	s3 = s3[:4]
	fmt.Printf("len=%d cap=%d %v\n", len(s3), cap(s3), s3)

	// Drop its first two values.
	// peculiar case !!!
	s3 = s3[2:]
	fmt.Printf("len=%d cap=%d %v\n", len(s3), cap(s3), s3)

	// peculiar case !!!
	s3 = s3[1:]
	fmt.Printf("len=%d cap=%d %v\n", len(s3), cap(s3), s3)

	// peculiar case !!!
	s3 = s3[1:2]
	fmt.Printf("len=%d cap=%d %v\n", len(s3), cap(s3), s3)

	var nilSlice []int
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice))
	if nilSlice == nil {
		fmt.Println("nilSlice is nil!")
	}

	sl1 := make([]int, 5)  // len(a)=5
	sl2 := make([]int, 0, 5)   // len(b)=0, cap(b)=5
	fmt.Println(sl1, sl2)

	////// slice of slices
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}


	/// appending
	var slc []int
	printSlice(slc)

	// append works on nil slices.
	slc = append(slc, 0)
	printSlice(slc)

	// The slice grows as needed.
	slc = append(slc, 1)
	printSlice(slc)

	// We can add more than one element at a time.
	slc = append(slc, 2, 3, 4)
	printSlice(slc)



	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {		// range returns two values, the index and the value at that index
		fmt.Printf("2**%d = %d\n", i, v)
	}
	for i := range pow {  // for i, _ := range pow
		fmt.Printf("2**%d\n", i)
	}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}



}


func Pic(dx, dy int) [][]uint8 {
	var finalPic  [][]uint8
	for y:=0; y < dy; y+=1 {
		var row []uint8
		row = make([]uint8, dx)
		for i := range row {
			row[i] = uint8(i * y)
		}
		finalPic = append(finalPic, row)
	}
	return finalPic
}



func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
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