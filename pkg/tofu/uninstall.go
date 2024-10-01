package tofu

import (
	"context"

	"get.porter.sh/porter/pkg/exec/builder"
)

// Uninstall runs a OpenTofu destroy
func (m *Mixin) Uninstall(ctx context.Context) error {
	action, err := m.loadAction(ctx)
	if err != nil {
		return err
	}
	step := action.Steps[0]

	err = m.commandPreRun(ctx, &step)
	if err != nil {
		return err
	}

	// Update step fields that exec/builder works with
	step.Arguments = []string{"destroy"}
	// Always run in non-interactive mode
	step.Flags = append(step.Flags, builder.NewFlag("auto-approve"))
	step.Flags = append(step.Flags, builder.NewFlag("input=false"))

	applyVarsToStepFlags(&step)

	action.Steps[0] = step
	_, err = builder.ExecuteSingleStepAction(ctx, m.RuntimeConfig, action)
	if err != nil {
		return err
	}

	return m.handleOutputs(ctx, step.Outputs)
}
