package steps

import (
	"fmt"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

// PrintSteps defines Gherkin steps around printing things to the terminal.
func PrintSteps(suite *godog.Suite, fs *FeatureState) {
	suite.Step(`^it does not print "([^\"]*)"$`, func(text string) error {
		if strings.Contains(fs.activeScenarioState.lastRunResult.OutputSanitized(), text) {
			return fmt.Errorf("text found: %q", text)
		}
		return nil
	})

	suite.Step(`^it prints$`, func(expected *gherkin.DocString) error {
		if !strings.Contains(fs.activeScenarioState.lastRunResult.OutputSanitized(), expected.Content) {
			return fmt.Errorf("text not found:\n\nEXPECTED: %q\n\nACTUAL:\n\n%q", expected.Content, fs.activeScenarioState.lastRunResult.OutputSanitized())
		}
		return nil
	})

	suite.Step(`^it prints the error:$`, func(expected *gherkin.DocString) error {
		if !strings.Contains(fs.activeScenarioState.lastRunResult.OutputSanitized(), expected.Content) {
			return fmt.Errorf("text not found: %s\n\nactual text:\n%s", expected.Content, fs.activeScenarioState.lastRunResult.OutputSanitized())
		}
		if fs.activeScenarioState.lastRunErr == nil {
			return fmt.Errorf("expected error")
		}
		return nil
	})
}
