package terraform

import (
	"context"

	"get.porter.sh/porter/pkg/exec/builder"
)

type InvokeOptions struct {
	Action string
}

// Invoke runs a custom terraform action
func (m *Mixin) Invoke(ctx context.Context, opts InvokeOptions) error {
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
	commands := []string{opts.Action}
	if len(step.Arguments) > 0 {
		commands = step.GetArguments()
	}
	step.Arguments = commands

	applyVarsToStepFlags(&step)

	action.Steps[0] = step
	_, err = builder.ExecuteSingleStepAction(ctx, m.RuntimeConfig, action)
	if err != nil {
		return err
	}

	return m.handleOutputs(ctx, step.Outputs)
}
