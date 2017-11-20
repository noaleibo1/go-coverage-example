package go_coverage_example

import "testing"

func TestLightAWandHappyPath(t *testing.T) {
	wizard := Wizard{
		IsTrainedWizard: true,
	}
	err := wizard.LightWand()
	if err != nil {
		t.Errorf("Wizard returned an error: %s", err.Error())
	}
}