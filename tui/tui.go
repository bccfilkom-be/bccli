package tui

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		return err
	}

	model := m.(model)
	if model.step != stepFinish {
		return fmt.Errorf("cancelled")
	}

	pkg := model.packageName.Value()
	framework := model.frameworkChoices[model.selectedFramework]
	driver := model.driverChoices[model.selectedDriver]

	fmt.Printf(
		"\nSummary\n-----------\nPackage: %s\nFramework: %s\nDriver: %s\n\n",
		pkg, framework, driver,
	)

	if err := runStep("Initializing project structure", "bccli", "init", pkg, "--framework", framework); err != nil {
		return fmt.Errorf("project initialization failed: %w", err)
	}

	if driver != "" {
		if err := runStep(fmt.Sprintf("Generating %s infrastructure", driver), "bccli", "infra", "generate", driver); err != nil {
			return fmt.Errorf("infra generation failed: %w", err)
		}
	}

	fmt.Println("All done! Project and infrastructure successfully initialized.")

	return nil
}

func runStep(title string, name string, args ...string) error {
	fmt.Printf("🔧 %s...\n", title)

	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ %s failed: %v\n\n", title, err)
		return err
	}

	fmt.Printf("\n%s completed successfully.\n\n", title)
	return nil
}
