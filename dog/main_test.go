package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	a := Years(2)
	if a != 14 {
		t.Errorf("Years was incorrect. got: %d., want: %d.", a, 14)
	}

}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(2)
	}
}

func ExampleYears() {
	fmt.Println(Years(3))
	// Output:
	// 21
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(3)
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(3))
	// Output:
	// 21
}
func TestYearsTwo(t *testing.T) {
	c := YearsTwo(3)
	if c != 21 {
		t.Errorf("YearsTwo was incorrect. got: %d., want: %d.", c, 21)
	}
}
