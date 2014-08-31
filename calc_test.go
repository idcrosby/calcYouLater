package calcYouLater

import (
	"testing"
	"fmt"
)

func TestIntegrate(t *testing.T) {
	const in, out = 5, 10
	if x := Integrate(in); x!= out {
		t.Errorf("Integrate(%v) = %v, want %v", in, x, out)
	}	
}

func TestConstructor(t *testing.T) {
	const coef, exp = 2, 3
	x := New(coef, exp)
	if x.Pairs[0][0] != coef || x.Pairs[0][1] != exp {
		t.Errorf("New(%v, %v) = %v and %v", coef, exp, x.Pairs[0][0], x.Pairs[0][1])
	}
}

func TestAdd(t *testing.T) {
	const coef1, exp1 = 1, 2
	const coef2, exp2 = 3, 4
	x := New(coef1, exp1)
	y := x.Add(coef2, exp2)
	if y.Pairs[0][0] != coef1 || y.Pairs[0][1] != exp1 {
		t.Errorf("New(%v, %v) = %v and %v", coef1, exp1, y.Pairs[0][0], y.Pairs[0][1])
	}
	fmt.Printf("Pairs length: %v", len(y.Pairs))
	if y.Pairs[1][0] != coef2 || y.Pairs[1][1] != exp2 {
		t.Errorf("Add(%v, %v) = %v and %v", coef2, exp2, y.Pairs[1][0], y.Pairs[1][1])
	} 
}

func TestPrint(t *testing.T) {
	const coef, exp = 3, 4
	val := "+3x^4"
	x := New(coef, exp)
	if str := x.ToString(); str != val {
		t.Errorf("ToString = " + str + ", want " + val)
	}
}