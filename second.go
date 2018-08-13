package main

import (
	"strings"
	"sort"
	"fmt"
)

var names string = `AMELIA
OLIVIA
EMILY
ISLA
JESSICA
POPPY
ISABELLA
SOPHIE
MIA
RUBY
LILY
GRACE
EVIE
SOPHIA
ELLA
SCARLETT
CHLOE
ISABELLE
FREYA
CHARLOTTE
SIENNA
DAISY
PHOEBE
MILLIE
EVA
ALICE
LUCY
FLORENCE
SOFIA
LAYLA
LOLA
HOLLY
IMOGEN
MOLLY
MATILDA
LILLY
ROSIE
ELIZABETH
ERIN
MAISIE
LEXI
ELLIE
HANNAH
EVELYN
ABIGAIL
ELSIE
SUMMER
MEGAN
JASMINE
MAYA
AMELIE
LACEY
WILLOW
EMMA
BELLA
ELEANOR
ESME
ELIZA
GEORGIA
HARRIET
GRACIE
ANNABELLE
EMILIA
AMBER
IVY
BROOKE
ROSE
ANNA
ZARA
LEAH
MOLLIE
MARTHA
FAITH
HOLLIE
AMY
BETHANY
VIOLET
KATIE
MARYAM
FRANCESCA
JULIA
MARIA
DARCEY
ISABEL
TILLY
MADDISON
VICTORIA
ISOBEL
NIAMH
SKYE
MADISON
DARCY
AISHA
BEATRICE
SARAH
ZOE
PAIGE
HEIDI
LYDIA
SARA
OLIVER
JACK
HARRY
JACOB
CHARLIE
THOMAS
OSCAR
WILLIAM
JAMES
GEORGE
ALFIE
JOSHUA
NOAH
ETHAN
MUHAMMAD
ARCHIE
LEO
HENRY
JOSEPH
SAMUEL
RILEY
DANIEL
MOHAMMED
ALEXANDER
MAX
LUCAS
MASON
LOGAN
ISAAC
BENJAMIN
DYLAN
JAKE
EDWARD
FINLEY
FREDDIE
HARRISON
TYLER
SEBASTIAN
ZACHARY
ADAM
THEO
JAYDEN
ARTHUR
TOBY
LUKE
LEWIS
MATTHEW
HARVEY
HARLEY
DAVID
RYAN
TOMMY
MICHAEL
REUBEN
NATHAN
BLAKE
MOHAMMAD
JENSON
BOBBY
LUCA
CHARLES
FRANKIE
DEXTER
KAI
ALEX
CONNOR
LIAM
JAMIE
ELIJAH
STANLEY
LOUIE
JUDE
CALLUM
HUGO
LEON
ELLIOT
LOUIS
THEODORE
GABRIEL
OLLIE
AARON
FREDERICK
EVAN
ELLIOTT
OWEN
TEDDY
FINLAY
CALEB
IBRAHIM
RONNIE
FELIX
AIDEN
CAMERON
AUSTIN
KIAN
RORY
SETH
ROBERT
ALBERT
SONNY`

// First Version
//func main()  {
//	names := strings.Split(names, "\n")
//	var small_score = 0
//	var winner string
//	for _, name := range names{
//		count := getNameScore(name)
//		//fmt.Println(i, name, count)
//		if count > small_score {
//			small_score = count
//			winner = name
//		}
//	}
//	fmt.Println(winner)
//}

func getNameScore(s string) int{
	score := 0
	for _, v := range s {
		score += int(v)
	}
	return score
}

type Person struct {
	name string
	score int
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByScore []Person

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].score > a[j].score }

// Second Version
func main(){
	names := strings.Split(names, "\n")
	people := make([]Person, 200)
	for i, name := range names {
		score := getNameScore(name)
		people[i] = Person{name, score}
	}
	sort.Sort(ByScore(people))
	fmt.Println(people)
}

type Odam struct {
	name string
	age int
}

func (p Odam) getAge() int  {
	return p.age
}
