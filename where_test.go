package where

import (
	"testing"
)

func TestSQLString(t *testing.T) {

	cb := ConditionBuilder{}

	s1 := Condition{
		Target:    "c1",
		Condition: "=",
		Value:     "c1.foo.com",
	}

	cb.Where(&s1)

	str, val := cb.Build(1)
	t.Logf(str)
	for _, v := range val {
		t.Logf("> %v", v)
	}

	s2 := Condition{
		Target:    "c2",
		Condition: "=",
		Value:     "c2.foo.com",
	}

	s3 := Condition{
		Target:    "c3",
		Condition: ">",
		Value:     "c3.foo.com",
	}

	cb = ConditionBuilder{}
	cb.Where(And(&s2, &s1))

	str, val = cb.Build(1)
	t.Logf(str)
	for _, v := range val {
		t.Logf("> %v", v)
	}

	cb = ConditionBuilder{}
	cb.Where(And(&s2, Or(&s3, &s2, &s1)))

	str, val = cb.Build(1)
	t.Logf(str)
	for _, v := range val {
		t.Logf("> %v", v)
	}

}
