package go_coverage_example

import (
	"fmt"
	"errors"
)

type Wizard struct {
	IsTrainedWizard bool
}

func (w *Wizard) LightWand() error {
	if w.IsTrainedWizard {
		fmt.Println("Lumos!")
		return nil
	}

	fmt.Println("You still have a lot to learn!")
	return errors.New("Wizard is not skilled enough")
}