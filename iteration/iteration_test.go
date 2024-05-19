package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("test repeating 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q, but got %q", expected, repeated)
		}
	})
	t.Run("test repeating 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("expected %q, but got %q", expected, repeated)
		}
	})

	t.Run("Test if repeat < 0 times should return invalid", func(t *testing.T) {
		repeated := Repeat("a", -10)
		expected := "Invalid number, Overflow expected"

		if repeated != expected {
			t.Errorf("expected %q, but got %q", expected, repeated)
		}

	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
