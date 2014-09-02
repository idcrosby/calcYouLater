package calcYouLater

import (
	"bytes"
	"fmt"
	"strconv"
)

var diffMap map[string]string

func main() {
		fmt.Println("testing...")
}

type Polynomial2 struct {
	//int x, y
	// [2]int pair
 // [][2]int Pairs
 Pairs []PolyPair
}

// This type should be 'private'?
type PolyPair struct {
	// Can 'Exp' be defaulted to 1?
  Coef,Exp string
}

func Differentiate(p Polynomial2) (result Polynomial2) {
 count := 0
 newPairs := make([]PolyPair, len(p.Pairs), cap(p.Pairs))
 for _,val := range p.Pairs {
 	 			// fmt.Println("here " + val.Coef + " " + val.Exp)
 		if coef, err := strconv.Atoi(val.Coef); err == nil {
 			// fmt.Println("why u no!?")
 			// coef is number
	 		if val.Exp == "0" {
 			 // x^0 differentiates away so skip
	 			// fmt.Println("skipping exponent of 0")
 			} else {
 				// fmt.Println("differentiate exponent of " + val.Exp)
 				// basic differentiatiom
 				// add function on Polynomial to incorporate string <-> int conversions
 				exp, _ := strconv.Atoi(val.Exp)
 				// fmt.Printf("result Pairs length is %d\n", len(newPairs))

 				//coef, _ := strconv.Atoi(val.Coef)
 				newPairs[count] = PolyPair{strconv.Itoa(exp * coef), strconv.Itoa(exp - 1)}
 				count++
 			}
 		} else if newCoef := diffMap[val.Coef]; len(newCoef) > 0 {
 			fmt.Println("Unable to convert " + val.Coef + " to number")
 			newPairs[count] = PolyPair{newCoef, val.Exp}
 			count++
 		} else {
 			fmt.Println("oops " + val.Coef)
 			// ...
 		}
 }
 	// only take the pairs which were added
	result.Pairs = newPairs[0:count]
 return 
}

func (p Polynomial2) ToString() string {
	var buffer bytes.Buffer
	for i := range p.Pairs {
		// if p.Pairs[i][0] < 0 {
			// buffer.WriteString("-")
		// } else if p.Pairs[i][0] > 0 {
			// buffer.WriteString("+")
		// }
		coef := p.Pairs[i].Coef
		if len(coef) > 1 && coef[1:len(coef)] == "-" {
			buffer.WriteString("-")
		}
		buffer.WriteString(p.Pairs[i].Coef + "x")
		if (p.Pairs[i].Exp != "0") {
			buffer.WriteString("^" + p.Pairs[i].Exp)
		}
	}
	return buffer.String()
}

func init() {

	// Build basic integration Map
	diffMap = map[string]string{
		"sin" : "cos",
		"cos" : "sin",
	}
}

func Integrate(x float64) float64 {

		// fmt.Printf("integrating %e \n", x)
		return x
}



type Polynomial struct {
	Pairs [][2]int
}

// functions on Polynomial

// basic constructor
func New(coef int, exp int) *Polynomial {
	p := new(Polynomial)
	p.Pairs = make([][2]int, 1)
	p.Pairs[0][0] = coef
	p.Pairs[0][1] = exp
	return p
}

func (p Polynomial) Add(coef int, exp int) *Polynomial {
	newPairs := make([][2]int, len(p.Pairs)+1, cap(p.Pairs)+1)

	copy(newPairs, p.Pairs)
	newPairs[len(p.Pairs)][0] = coef
	newPairs[len(p.Pairs)][1] = exp

	p1 := new(Polynomial)
	p1.Pairs = newPairs
	return p1
}

func (p Polynomial) ToString() string {
	var buffer bytes.Buffer
	for i := range p.Pairs {
		if p.Pairs[i][0] < 0 {
			buffer.WriteString("-")
		} else if p.Pairs[i][0] > 0 {
			buffer.WriteString("+")
		}
		buffer.WriteString(strconv.Itoa(p.Pairs[i][0]) + "x")
		if (p.Pairs[i][1] != 0) {
			buffer.WriteString("^" + strconv.Itoa(p.Pairs[i][1]))
		}
	}
	return buffer.String()
}
