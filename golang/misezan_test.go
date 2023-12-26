package misezan_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/BonyChops/misezan/golang"
)

func internalTestMisezanWithCases[T misezan.Number](t *testing.T, cases []struct {
	A, B T
	Ans  float64
}) {
	for _, c := range cases {
		got := misezan.Calc(c.A, c.B)
		if got != c.Ans {
			t.Errorf("Calc(%v, %v) == %v, want %v", c.A, c.B, got, c.Ans)
		}

		got = misezan.Calc(c.B, c.A)
		if got != c.Ans {
			t.Errorf("Calc(%v, %v) == %v, want %v", c.B, c.A, got, c.Ans)
		}
	}
}

func internalTestMisezanBasic[T misezan.Number](t *testing.T) {
	rawJson, err := os.ReadFile("../test/test_cases.json")
	if err != nil {
		t.Errorf("Failed to read test_cases.json: %v", err)
	}

	var cases struct {
		Basic []struct {
			A, B T
			Ans  float64
		}
	}

	if err = json.Unmarshal(rawJson, &cases); err != nil {
		t.Errorf("Failed to unmarshal test_cases.json: %v", err)
		return
	}

	internalTestMisezanWithCases[T](t, cases.Basic)
}

func internalTestMisezanExtra[T misezan.Number](t *testing.T) {
	rawJson, err := os.ReadFile("../test/test_cases.json")
	if err != nil {
		t.Errorf("Failed to read test_cases.json: %v", err)
	}

	var cases struct {
		Extra []struct {
			A, B T
			Ans  float64
		}
	}

	if err = json.Unmarshal(rawJson, &cases); err != nil {
		t.Errorf("Failed to unmarshal test_cases.json: %v", err)
		return
	}

	internalTestMisezanWithCases[T](t, cases.Extra)
}

func TestMisezanInt(t *testing.T) {
	internalTestMisezanBasic[int](t)
}

func TestMisezanIntExtra(t *testing.T) {
	internalTestMisezanExtra[int](t)
}

func TestMisezanInt8(t *testing.T) {
	internalTestMisezanBasic[int8](t)
}

func TestMisezanInt16(t *testing.T) {
	internalTestMisezanBasic[int16](t)
}

func TestMisezanInt32(t *testing.T) {
	internalTestMisezanBasic[int32](t)
}

func TestMisezanInt32Extra(t *testing.T) {
	internalTestMisezanExtra[int32](t)
}

func TestMisezanInt64(t *testing.T) {
	internalTestMisezanBasic[int64](t)
}

func TestMisezanInt64Extra(t *testing.T) {
	internalTestMisezanExtra[int64](t)
}
