package calcYouLater

import (
	"testing"
	"fmt"
)

func TestIntegrate(t *testing.T) {
	// const in, out = 5, 10
	// if x := Integrate(in); x!= out {
	// 	t.Errorf("Integrate(%v) = %v, want %v", in, x, out)
	// }	
}

func TestConstructor(t *testing.T) {
	const coef, exp = 2, 3
	x := New(coef, exp)
	if x.Pairs[0][0] != coef || x.Pairs[0][1] != exp {
		t.Errorf("New(%v, %v) = %v and %v", coef, exp, x.Pairs[0][0], x.Pairs[0][1])
	}
}

func TestDifferentiate(t *testing.T) {
	pair := PolyPair{"4","3"}
	p1 := Polynomial2{[]PolyPair{pair}}
	if p2 := Differentiate(p1); len(p2.Pairs) < 1 || p2.Pairs[0].Coef != "12" || p2.Pairs[0].Exp != "2" {
		t.Errorf("Differentiate(" + p1.ToString() + ") = " + p2.ToString() + ", want 12x^2")
	}
}

func TestDifferentiate2(t *testing.T) {
	pair1 := PolyPair{"4","3"}
	pair2 := PolyPair{"-2","1"}
	pair3 := PolyPair{"4","0"}
	
	p1 := Polynomial2{[]PolyPair{pair1, pair2, pair3}}
	if p2 := Differentiate(p1); len(p2.Pairs) != 2 || p2.Pairs[0].Coef != "12" || p2.Pairs[0].Exp != "2" || p2.Pairs[1].Coef != "-2" || p2.Pairs[1].Exp != "0" {
		t.Errorf("Differentiate(" + p1.ToString() + ") = " + p2.ToString() + ", want 12x^2-2x")
	}
}

func TestDifferentiate3(t *testing.T) {
	pair1 := PolyPair{"-2","1"}
	pair2 := PolyPair{"sin","1"}
	
	p1 := Polynomial2{[]PolyPair{pair1, pair2}}
	if p2 := Differentiate(p1); len(p2.Pairs) != 2 || p2.Pairs[0].Coef != "-2" || p2.Pairs[0].Exp != "0" || p2.Pairs[1].Coef != "cos" || p2.Pairs[1].Exp != "1" {
		t.Errorf("Differentiate(" + p1.ToString() + ") = " + p2.ToString() + ", want -2+cosx")
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