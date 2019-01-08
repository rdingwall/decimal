package dectest

import (
	"os"
	"testing"
)

func TestParser2(t *testing.T) {
	s := &scanner{}
	if err := s.parse([]byte("precision: 4")); err != nil {
		t.Error(err)
	}
	if s.precision != 4 {
		t.Errorf("precison was %d", s.precision)
	}

	if err := s.parse([]byte("absx001 abs '1'      -> '1'")); err != nil {
		t.Error(err)
	}

	if s.c.ID != "absx001" {
		t.Errorf("ID was %s", s.c.ID)
	}

	if s.c.Op != Abs {
		t.Errorf("Op was %s", s.c.Op.String())
	}

}

func TestParser3(t *testing.T) {
	b, err := os.Open("../_dectest/abs.decTest")
	if err != nil {
		t.Error(err)
	}
	s := NewScanner(b)
	for s.Scan() {
		c := s.Case()
		t.Logf("%s %s %s %s %s (precision=%d, minExponent=%d, maxExponent=%d, rounding=%v)", c.ID, c.Op, c.Inputs, c.Output, c.Conditions, s.precision, s.minExponent, s.maxExponent, s.rounding)
	}
	if s.Err() != nil {
		t.Error(err)
	}
}
