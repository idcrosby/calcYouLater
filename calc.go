package main

import (
	"fmt"
	"strings"
)

var diffMap map[string]string

func main() {
		fmt.Println("testing...")
		
		var p = []Polynomial{
				Polynomial{
					[]int{2,3},
					[]int{4,1}
				}
		}
}

type Polynomial struct {
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

func Differentiate(p Polynomial) (result Polynomial) {
 count := 0
 for _,val := range p.Pairs {
 		if coef := strings.Atoi(val.Coef); coef != nil {
 			// coef is number
	 		if val[1]	 == 0 {
 			 // x^0 differentiates away so skip
 			} else {
 				// basic differentiatiom
 				// add function on Polynomial to incorporate string <-> int conversions
 				result.Pairs[count++] = PolyPair{strings.Itoa(strings.Atoi(val.Exp) * strings.Atoi(val.Coef)), strings.Itoa(strings.Atoi(val.Exp) - 1)}
 			}
 		} else if newCoef := diffMap[val.Coef]; newCoef != nil {
 			result.Pairs[count++] = PolyPair{newCoef, val.Exp}
 		} else {
 			// ...
 		}
 		return
 } 
}


func init() {
	// Build basic integration Map
	// This needs to be 'global'
	diffMap := make(map[string]string)
	diffMap["sin"] = "cos"
	diffMap["cos"] = "sin"
	

}

