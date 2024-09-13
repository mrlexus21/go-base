package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("setup")
	res := m.Run()
	fmt.Println("tear-down")

	os.Exit(res)
}

func TestMultiple(t *testing.T) {
	t.Run("groupA", func(t *testing.T) {
		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("parallel simple")
			var x, y, result = 2, 2, 4

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)

				if r != 1 {
					t.Errorf("failed")
				}
			})
		})

		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("parallel medium")
			var x, y, result = 222, 222, 49284

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})

		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("parallel negative")
			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})
	})
}

/*func TestAdd(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		t.Parallel()
		t.Log("parallel")
		var x, y, result = 2, 2, 4

		realResult := Add(x, y)

		if realResult != result {
			t.Errorf("%d != %d", result, realResult)
		}
		t.Run("1", func(t *testing.T) {
			r := Add(1, 1)

			if r != 2 {
				t.Errorf("failed")
			}
		})
	})
}*/
