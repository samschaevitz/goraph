package spbf

import (
	"testing"

	"github.com/gyuho/goraph/graph/gs"
)

func TestSPBF(t *testing.T) {
	g12 := gs.FromJSON("../../files/testgraph.json", "testgraph.012")
	s12, ok12 := SPBF(g12, g12.FindVertexByID("S"), g12.FindVertexByID("T"))
	s12c := "S(=0) → A(=7) → C(=4) → B(=2) → T(=-2)"
	ok12c := true
	if s12 != s12c || ok12 != ok12c {
		t.Errorf("Should be same but %v", s12)
	}

	g13 := gs.FromJSON("../../files/testgraph.json", "testgraph.013")
	s13, ok13 := SPBF(g13, g13.FindVertexByID("S"), g13.FindVertexByID("T"))
	s13c := "There is negative weighted cycle (No Shortest Path)"
	ok13c := false
	if s13 != s13c || ok13 != ok13c {
		t.Errorf("Should be same but %v", s13)
	}
}
