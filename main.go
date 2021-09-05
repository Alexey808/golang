package main // определение файла

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var globalValue1 bool = 1 == 1
var globalValue2 string = "Hello World"

const globalValue3 string = "HELLO_WORLD"
const (
	a = "abc"
	b = 123
)

type Circle struct {
	x float64
	y float64
	r float64
	// x, y, r float64
}

func main() {
	// testVariables()
	// testFtmScanf()
	// testFor()
	// testSwitch()
	// testArray()
	// testMap()
	// testCondition()
	// testFuncPanic()
	// testPointers()
	// testStruct()
	// testStreams()
	// testPackages()
	// testChannel()
}

// --- testVariables
func testVariables() {
	localValue := "test"
	fmt.Println(localValue)
	fmt.Println(globalValue1, globalValue2, globalValue3)
	fmt.Println(a, b)
}

// --- testFtmScanf
func testFtmScanf() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input) // заполнить переменную input числом ввода
	output := input * 2
	fmt.Println(output)
}

// --- testFor
func testFor() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	for i := 1; i <= 6; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

// --- testSwitch
func testSwitch() {
	i := 1

	switch i {
	case 0:
		fmt.Println("Zerro")
	case 1:
		fmt.Println("One")
	default:
		fmt.Println("Unknown number")
	}
}

// --- testArray
func testArray() {
	x := [5]float64{98, 93, 77, 82, 83} // короткая запись создания массива, тип можн не указ.

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	// ---

	var arr [5]int
	arr[4] = 100
	arrLength := len(arr) // длина массива

	var float float64 = 0
	float = float64(len(arr)) // конвертация int в float64

	fmt.Println(arr, arrLength, float)

	// --- Срезы
	a := make([]float64, 5)     // тип float64 длина 5
	b := make([]float64, 5, 10) // 10 длина массива, на который указывает срез
	c := [5]float64{1, 2, 3, 4, 5}
	cc := c[1:3] // arr[:] эквивалентно arr[0:len(arr)]
	fmt.Println(a, b, cc)
}

// --- testMap
func testMap() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
	delete(x, "key")
	fmt.Println(x["key"])
	value, status := x["key"]
	fmt.Println(value, status)
}

// --- testCondition
func testCondition() {
	myMap := make(map[string]string)     // создаём мапу
	myMap["myKey"] = "MYKEY"             // записываем ключ:значение
	if value, ok := myMap["myKey"]; ok { // если есть value и значение сущесвтует с таким ключом
		fmt.Println(value, ok)
	}
}

// --- testFuncPanic
func testFuncPanic() {
	defer func() { // defer отложить выполнение функции
		fmt.Println("test2")
	}()
	func() {
		fmt.Println("test1")
	}()
	panic("PANIC") // останавливает выполнение
	fmt.Println("0")
}

// --- testPointers
func testPointers() {
	a := 5
	pointerOne(&a)
	fmt.Println(a)

	b := new(int)
	pointerTwo(b)
	fmt.Println(*b)

	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c.area())
}

func pointerOne(a *int) {
	*a = 1
}

func pointerTwo(b *int) {
	*b = 2
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// --- testStruct
func testStruct() {
	c1 := new(Circle)              // создаст экземпляр, присвоит значения по умолчанию и вернёт указатель
	c2 := Circle{x: 0, y: 0, r: 5} // создаст экземпляр типа, присвоит значения
	c3 := Circle{0, 0, 5}          // создаст экземпляр типа, присвоит значения
	fmt.Println(*c1)
	fmt.Println(c2.r)
	fmt.Println(c3)

	p := new(Android) // создаём экземпляр типа с имплементацией его методов
	p.person.Talk()
}

type Person struct {
	name string
}

type Android struct {
	person Person
	model  string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.name)
}

// --- testStreams
func testStreams() {
	for i := 0; i < 10; i++ {
		go f(i) // горутина
	}

	var input string
	fmt.Scanln(&input) // хак для отладки чтобы видеть лог
}

func f(n int) {
	for i := 0; i < 10; i++ {
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
		fmt.Println(n, ":", i, "", rand.Intn(250))
	}
}

// --- testChannel
func testChannel() {
	var c chan string = make(chan string) // создание канала

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input) // хак для отладки
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" // отправка сообщения по каналу
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong" // отправка сообщения по каналу
	}
}

func printer(c chan string) {
	for {
		m := <-c // получение сообщения с канала
		fmt.Println(m)
		time.Sleep(time.Second * 1)
	}
}

// ---
// func testPackages() {
// t := time.Now()
// r := rand.Intn(10)
// r := rand.New(123)
// amt := time.Duration(rand.Intn(250))
// fmt.Println(r)
// }
