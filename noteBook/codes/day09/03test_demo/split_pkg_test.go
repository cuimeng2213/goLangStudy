package split_pkg

import (
	"reflect"
	"testing"
)

func TestMySplit(t *testing.T) {
	s := "abcbdbcbf"
	r := [...]string{"a", "c", "d", "c", "f"}
	ret := MySplit(s, "b")

	for i, v := range ret {
		if r[i] != v {
			t.Error("切割结果不对")
		}

	}
	// reflect.DeepEqual(ret, r)
	t.Logf("AAAAAAAAAAAAAAAAAA %v ", ret)
}
func TestMySplit2(t *testing.T) {
	s := "abcbdbcbf"
	r := []string{"a", "bd", "bf"}
	ret := MySplit(s, "bc")

	// for i, v := range ret {
	// 	if r[i] != v {
	// 		t.Error("切割结果不对")
	// 	}

	// }
	if !reflect.DeepEqual(ret, r) {
		t.Errorf("failed %#v %#v \n", r, ret)
	}
	t.Logf("AAAAAAAAAAAAAAAAAA %v ", ret)
}
