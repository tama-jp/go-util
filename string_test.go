package go_util

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func TestRandString(t *testing.T) {

	for i := 10; i <= 100; i = i + 20 {
		rnd := RandString(i)

		t.Log("len(rnd)", len(rnd))
		t.Log("rnd", rnd)

	}

}
