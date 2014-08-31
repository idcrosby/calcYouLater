package calcYouLater

import (
	"bytes"
	"fmt"
	"strconv"
)

var diffMap map[string]string

func main() {
		fmt.Println("testing...")
		
		// var p = []Polynomial{
		// 		Polynomial{
		// 			[]int{2,3},
		// 			[]int{4,1}
		// 		}
		// }
}

type Polynomial2 struct {
	//int x, y
	// [2]int pair
 // [][2]int Pairs
 Pairs []PolyPair
}

// This type should be 'private'
type PolyPair struct {
	// Can 'Exp' be defaulted to 1?
  Coef,Exp string
}

func Differentiate(p Polynomial2) (result Polynomial2) {
 count := 0
 for _,val := range p.Pairs {
 		if coef, err := strconv.Atoi(val.Coef); err != nil {
 			// coef is number
	 		if val.Exp == "0" {
 			 // x^0 differentiates away so skip
 			} else {
 				// basic differentiatiom
 				// add function on Polynomial to incorporate string <-> int conversions
 				exp, _ := strconv.Atoi(val.Exp)
 				//coef, _ := strconv.Atoi(val.Coef)
 				result.Pairs[count] = PolyPair{strconv.Itoa(exp * coef), strconv.Itoa(exp - 1)}
 				count++
 			}
 		} else if newCoef := diffMap[val.Coef]; len(newCoef) > 0 {
 			result.Pairs[count] = PolyPair{newCoef, val.Exp}
 			count++
 		} else {
 			// ...
 		}
 }
 return 
}


func init() {
	// Build basic integration Map
	// This needs to be 'global'
	diffMap := make(map[string]string)
	diffMap["sin"] = "cos"
	diffMap["cos"] = "sin"
	

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

