package embed

import (
	"reflect"
	"testing"
)

type S1 struct {
	F1 int
	F2 string
	F3 *int
}

type S2 struct {
	F1 int
	F2 string
	F4 string
	F3 *int
}

type T1 struct {
	F1 S1
	F2 S2
	F3 int
}

type T2 struct {
	F3 int
	F4 string
	F2 S2
	F1 S2
}

func TestEmbed(t *testing.T) {
	f3 := 12345
	s1 := &S1{
		F1: 12345,
		F2: "12345",
		F3: &f3,
	}
	s2 := &S2{}
	embeddedS2 := &S2{
		F1: 12345,
		F2: "12345",
		F3: &f3,
	}

	if err := Embed(s2, s1); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s2, embeddedS2) {
		t.Fatal()
	}

	if err := Embed(s1, s2); err == nil {
		t.Fatal()
	}

	t1 := T1{
		F1: *s1,
		F2: *s2,
		F3: 12345,
	}
	t2 := T2{}

	if err := Embed(&t2, &t1); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(embeddedS2, &t2.F1) || !reflect.DeepEqual(t1.F2, t2.F2) || !reflect.DeepEqual(t1.F3, t2.F3) {
		t.Fatal()
	}

	if err := Embed(&t1, &t2); err == nil {
		t.Fatal()
	}
}
