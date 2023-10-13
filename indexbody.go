package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	student1 := student{
		height: 1.7,
		weight: 40,
	}

	fmt.Println("Индекс массы тела =", student1.calcIWB())
	fmt.Println("Женится или нет", student1.married())
}

type student struct {
	height float64
	weight float64
}

// дописать расчет индекса массы тела вес (кг)/рост^2 (м)
// создать новую функцию которая будет вычисляется женится студент или нет
// берем рандоиное число если четное то да если нет то нет
func (s student) calcIWB() float32 {
	return float32(s.weight / math.Pow(s.height, 2.0))
}

func (s student) married() string {
	randInt := rand.Intn(1000)
	return getMarriedByNumber(randInt)
}

func getMarriedByNumber(number int) string {
	switch number {
	case 0:
		return "да"
	default:
		return "нет"
	}
}


package main

import (
"fmt"
"math"
"math/rand"
)

func main() {
	student1 := student{
		height: 1.7,
		weight: 40,
	}

	fmt.Println("Индекс массы тела ="student1.calcIWB())
	fmt.Println("Женится или нет"student1.married())
}

type student struct {
	height float64
	weight float64
}

// дописать расчет индекса массы тела вес (кг)/рост^2 (м)
// создать новую функцию которая будет вычисляется женится студент или нет
// берем рандоиное число если четное то да если нет то нет
func (s student) calcIWB() float32 {
	return float32(s.weight / math.Pow(s.height, 2.0))
}

func (s student) married() string {
	randInt := rand.Intn(1000)
	return getMarriedByNumber(randInt)
}

func getMarriedByNumber(number int)  string {
	switch number {
	case 0:
		return "да"
	default:
		return "нет"
	}
}


package main

import (
"fmt"
"math"
"math/rand"
)

func main() {
	student1 := student{
		name:    "Alina",
		surname: "Arakchaa",
		age:     18,
		height:  170,
		weight:  40,
	}

	fmt.Println(student1.calcIWB())
	fmt.Println(student1.wedding())
}

type student struct {
	name    string
	surname string
	age     int
	height  float64
	weight  float64
}

// дописать расчет индекса массы тела вес (кг)/рост^2 (м)
// создать новую функцию которая будет вычисляется женится студент или нет
// берем рандоиное число если четное то да если нет то нет
func (s student) calcIWB() float64 {
	return s.weight / math.Pow(s.height, 2)
}

func (s student) wedding() string {
	random := rand.Intn(1000)

	if random%2 == 0 {
		return "Женится"
	} else {
		return "Не женится"
	}

}
