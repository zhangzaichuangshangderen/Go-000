package Week02

import (
	"errors"
	"testing"
)

func TestWymError(t *testing.T) {
	test := Wym()
	t.Logf("%T",test)
	t.Log(test)

	errors.Unwrap()
	 Wym2()
	//t.Logf("%T",test2)
	//t.Log(test2)

	//io.Reader.Read()


}