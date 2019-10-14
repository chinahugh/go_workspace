package mypag

type Person struct {
	Name string
	Age  int
}

func Baba(a *Person, b *Person) int {

	return a.Age + b.Age
}
