package main

import (
	"fmt"
	"math/big"
	"testing"
)

var (
	A048881List = []int{
		0, 0, 1, 0, 1, 1, 2, 0, 1, 1, 2, 1, 2, 2, 3, 0, 1, 1, 2, 1, 2, 2, 3, 1,
		2, 2, 3, 2, 3, 3, 4, 0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2,
		2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2,
		3, 2, 3, 3, 4, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 1, 2, 2, 3,
		2, 3, 3, 4, 2, 3,
	}
	A000108List = []string{
		"1", "1", "2", "5", "14", "42", "132", "429", "1430", "4862", "16796", "58786",
		"208012", "742900", "2674440", "9694845", "35357670", "129644790",
		"477638700", "1767263190", "6564120420", "24466267020",
		"91482563640", "343059613650", "1289904147324",
		"4861946401452", "18367353072152", "69533550916004",
		"263747951750360", "1002242216651368", "3814986502092304",
	}
	A002596List = []string{
		"1", "1", "-1", "1", "-5", "7", "-21", "33", "-429", "715", "-2431", "4199", "-29393",
		"52003", "-185725", "334305", "-9694845", "17678835", "-64822395",
		"119409675", "-883631595", "1641030105", "-6116566755",
		"11435320455", "-171529806825", "322476036831",
		"-1215486600363", "2295919134019",
	}
	A056981List = []string{
		"1", "1", "1", "1", "25", "49", "441", "1089", "184041", "511225", "5909761",
		"17631601", "863948449", "2704312009", "34493775625",
		"111759833025", "93990019574025", "312541206957225",
		"4201942893536025", "14258670483605625",
		"780804795682244025",
	}
	A005187List = []int{
		0, 1, 3, 4, 7, 8, 10, 11, 15, 16, 18, 19, 22, 23, 25, 26, 31, 32,
		34, 35, 38, 39, 41, 42, 46, 47, 49, 50, 53, 54, 56, 57, 63, 64,
		66, 67, 70, 71, 73, 74, 78, 79, 81, 82, 85, 86, 88, 89, 94, 95,
		97, 98, 101, 102, 104, 105, 109, 110, 112, 113, 116, 117, 119,
		120, 127, 128,
	}
	A056982List = []string{
		"1", "4", "64", "256", "16384", "65536", "1048576", "4194304",
		"1073741824", "4294967296", "68719476736", "274877906944",
		"17592186044416", "70368744177664", "1125899906842624",
		"4503599627370496", "4611686018427387904",
		"18446744073709551616", "295147905179352825856",
		"1180591620717411303424",
	}
)

func TestA048881(t *testing.T) {
	for n, entry := range A048881List {
		guess := A048881(int64(n))
		if guess != entry {
			t.Errorf("Incorrect result. For Value %d, expected %d. Got %d\n",
				n, entry, guess)
		}
	}
	fmt.Printf("Tested first %d entries of A048881.\n", len(A048881List))
}

func TestCatalan(t *testing.T) {
	for n, entry := range A000108List {
		var cat big.Int
		Catalan(&cat, n)
		guess := cat.Text(10)
		if guess != entry {
			t.Errorf("Incorrect result. For N=%d, expected %s. Got %s.\n",
				n, entry, guess)
		}
	}
	fmt.Printf("Tested first %d entries of A000108.\n", len(A000108List))
}

func TestA002596(t *testing.T) {
	for n, entry := range A002596List {
		var cat big.Int
		A002596(&cat, n)
		guess := cat.Text(10)
		if guess != entry {
			t.Errorf("Incorrect result. For N=%d, expected %s. Got %s.\n",
				n, entry, guess)
		}
	}
	fmt.Printf("Tested first %d entries of A002596.\n", len(A002596List))
}

func TestA056981(t *testing.T) {
	for n, entry := range A056981List {
		var cat big.Int
		A056981(&cat, n)
		guess := cat.Text(10)
		if guess != entry {
			t.Errorf("Incorrect result. For N=%d, expected %s. Got %s.\n",
				n, entry, guess)
		}
	}
	fmt.Printf("Tested first %d entries of A056981.\n", len(A056981List))
}

func TestA005187(t *testing.T) {
	for n, entry := range A005187List {
		guess := A005187(n)
		if guess != entry {
			t.Errorf("Incorrect result. For N=%d, expected %d. Got %d.\n",
				n, entry, guess)
		}
	}
	fmt.Printf("Tested first %d entries of A005187.\n", len(A005187List))
}

func TestA056982(t *testing.T) {
	for n, entry := range A056982List {
		var cat big.Int
		A056982(&cat, n)
		guess := cat.Text(10)
		if guess != entry {
			t.Errorf("Incorrect result. For N=%d, expected %s. Got %s.\n",
				n, entry, guess)
		}
	}
	fmt.Printf("Tested first %d entries of A056982.\n", len(A056982List))
}

//KummerGaussTerm(r *big.Rat, n int) *big.Rat
func TestKummerGaussTerm(t *testing.T) {
	var kg big.Rat
	fmt.Printf("------------ KummerGauss terms ------------\n")
	for n := 0; n < 100; n++ {
		KummerGaussTerm(&kg, n)
		fmt.Printf("%3d: %s\n", n, kg.FloatString(100))
		if n < len(KGTerm) {
			guess, _ := kg.Float64()
			if KGTerm[n] != guess {
				t.Errorf("Incorrect result. For N=%d.\n", n)
			}
		}
	}
}
