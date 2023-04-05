package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/tfe/provider"
	"github.com/sourcegraph/controller-cdktf/gen/tfe/workspace"
)

func newStack(scope constructs.Construct, name string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &name)

	provider.NewTfeProvider(stack, jsii.String("tfe"), &provider.TfeProviderConfig{})

	workspace.NewWorkspace(stack, jsii.String("workspace"), &workspace.WorkspaceConfig{
		Name:         jsii.String("test"),
		Organization: jsii.String("test"),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	newStack(app, "test")

	app.Synth()
}
