# Команды  

**Запуск скрипта go**  
`go run main.go`

# Типы  

**Преоброзование типов** 
```
float64(valueInt) // тип(value)
```

**«signed integer» (знаковое целое). Только положительное значение и 0**
```
int8  // -128 through 127
int16 // -32768 through 32767
int32 // -2147483648 through 2147483647 // псевдоним rune  
int64 // -9223372036854775808 through 9223372036854775807
```

**«unsigned integer» (беззнаковое целое)**  
```
uint8  // 0 through 255 // псевдоним byte 
uint16 // 0 through 65535
uint32 // 0 through 4294967295
uint64 // 0 through 18446744073709551615
```

**Машиннозависимые типы**  
`uint, int и uintptr`  

**Числа с плавающей точкой**  
```
float32
float64
```

**Комплексные числа**  
```
complex64
complex128
```

**Прочие типы**  
```
string
boolean
NaN
+∞
−∞
```
# Выжимка

## Array  
len(array) // Получить длину массива  
newArray := [5]float64{ 98, 93, 77, 82, 83 } // Создание массива, сокращённая запись, тип не обязателен  
newArray := make([]float64, 5, 10) // Срез массива по индексу, где 10 длина массива на который указывает срез  
newArray := array[0:5] // Срез массива по индексу, сокращённая запись  
newArray := append(array, newValue) // Добавление в массив  
copy(newArray, copyArray) // Копирование массива  

## Map  
newMap = make(map[string]int) // Мапа key:string value:int  
newMap := map[string]string { "key1": "value1", "key2": "value2" } // сокращённая запись мапы  
myMap["myKey"] // Получить значение мапы по ключу  
delete(x, "myKey") // Удалить значение по ключу  
value, ok := x["key"] // Получить значение мапы по ключу и статус наличия этого значения  

## Потоки, условия  

```go
for i := 1; i <= 10; i++ {
	// ...
} 
```
```go
if isOk {
	// ...
} else if {
	// ...
} else {
	// ...
} 
```
```go
switch value {
	case itValue: // ...
	default: // ...
}
```

## Функции  

defer myFunc() // отложить выполнение функции, типо setTimeout 0 в js  
func(){ /* ... */}() // анонимная функция  


```go
func myFunc(value []float64) float64 {
	// ...
	return result
}
```
```go
func myFunc() (int, int) {
    return value, ok // использовать можно так x, y := myFunc()
}
```
```go
func myFunc() (value, err := f())  {
    return value, err // возвращает значение и ошибку
}
```
```go
func myFunc(args ...int) int {
	// ...
	return result
}
```

## Указатели (на ячейки памяти)  
```go
func main() {
	value := 5
	myFunc(&value) // & передаёт адрес переменной value
	fmt.Println(value)

	value := new(int) // создаём указатель
	myFunc(value) //передаём указатель
	fmt.Println(*value)

	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c.area()) // вызываем, как будто у объкта есть метод area
}
func myFunc(value *int) { // * говорит что мы работаем с ячейкой памяти на которую ссылается value
	*value = 0
}
func (c *Circle) area() float64 { // функция особого типа
	return math.Pi * c.r * c.r
}
```
```go
type Person struct {
	name string
}
type Android struct {
	person Person
	model string
}

func main() {
	p := new(Android) // создаём экземпляр типа с имплементацией его методов
    p.person.Talk()
}

func (p *Person) Talk() { // функция особого типа
	fmt.Println("Hi, my name is", p.name)
}
```

## Типизация, интерфесы/структуры  
c := new(Circle) // создаст экземпляр, присвоит значения по умолчанию и вернёт указатель
```go
type Circle struct {
	x float64
	y float64
	r float64
}
```
**Вложенные структуры**  
```go
type Person struct {
	name string
}
type Android struct {
	person Person
	model string
	abilities []Abilities // где Abilities это interface с описанием методов
}
```
**Интерфейсы (в них описываются методы)**  
```go
type Shape interface {
    area() float64
}
```

## Пакеты
```go
time.Now() // возвращает текущее время /* 2021-09-01 01:47:17.689248556 +0400 +04 m=+0.000055961 */  
rand.Intn(10) // возвращает случайное число для последовательных итераций  
```

