package const_test

import (
	"testing"
)

const a int = 0

const(
	Monday = iota + 1
	Tuesday
	Wednesday
)

const x = iota // x=0
const (
	y = iota     //y=0
	z            //z=1
)

func TestConst(t *testing.T) {
	t.Log(Monday,Tuesday,Wednesday)
}
