package pai

import (
	filho "github.com/PedroHPuhl/Filho"
)

func Pai() string {
	a := 0

	if a == 0 {
		return filho.Filho()
	} else {
		return filho.Filha()
	}
}
