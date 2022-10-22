package main

import (
	"fmt"
)

type getAny struct{}
type AnyObj struct {
	getAny
	Opis    string
	Validan string
}

type AnyObj1 struct {
	getAny
	Opis    string
	Broj    int
	Iznos   float64
	Validan string
}

var Tabela string = "tabela"
var Tabela1 string = "tabela1"

// Genericki metod bez (de)generika
func (getAny) GetAny(a ...any) any {
	//Na primer dohvatanje argumenata
	stmSQL := fmt.Sprintf("select %s from %s a where a.id = %s", a[0], a[1], a[2])
	return stmSQL
}

func main() {
	anyObj := AnyObj{}
	anyObj1 := AnyObj1{}
	nesto := anyObj.GetAny("a.sifra, a.naziv", Tabela, "1", anyObj)
	fmt.Println(nesto)
	nesto = anyObj1.GetAny("a.sifra", Tabela1, "2", 22, 31.5)
	fmt.Println(nesto)

}

// Ne mogu se prosledjivati genericke promenljive razlicitog tipa, jer se kompajliranje odvija u dva koraka
// prvi korak se genericima dodeljuju pravi tip podataka, a onda se kompajlira sa striktnim promenljivama
func print[T any](output ...T) {
	fmt.Println(output)
}
