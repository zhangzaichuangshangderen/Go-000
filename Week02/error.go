package Week02

import (
	"github.com/pkg/errors"
)

func Wym() error {
	err := errors.New("你搞错啦，哈哈哈\r\n")
	return errors.Wrap(err,"Wym")
}

func Wym2() error {
	err := Wym()
	if err != nil {
		err = errors.WithMessage(err,"Wym2")
	}
	return err
}