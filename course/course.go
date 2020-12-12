package course

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

// A diferencia de otros lenguajes como JAVA, el método se declara fuera y se pone el tipo de dato y el valor
// entre paréntesis delante del nomrbe de la función.

func (c *Course) ChangePrice(price float64) {
	c.Price = price
}

func (c *Course) PrintClasses() {
	text := "La clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}
