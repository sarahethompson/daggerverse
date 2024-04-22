// Run serverless deploys

package main

import (
	"context"
)

type Serverless struct{}

// Runs a serverless deploy
func (m *Serverless) Deploy(
	ctx context.Context,
	configFile *File,
	awsAccessKeyID *Secret,
	awsSecretAccessKey *Secret,
	// +optional
	stage string,
	// +optional
	region string,
	// +optional
	force string,

) (string, error) {
	deployArgs := []string{"serverless", "deploy"}
	if stage != "" {
		deployArgs = append(deployArgs, "--stage", stage)
	}
	if region != "" {
		deployArgs = append(deployArgs, "--region", region)
	}
	if force != "" {
		deployArgs = append(deployArgs, "--force", force)
	}
	return dag.Container().From("alpine:latest").
		WithExec([]string{"apk", "add", "curl", "bash", "npm"}).
		WithExec([]string{"npm", "install", "-g", "serverless"}).
		WithFile("serverless.yml", configFile).
		WithSecretVariable("AWS_ACCESS_KEY_ID", awsAccessKeyID).
		WithSecretVariable("AWS_SECRET_ACCESS_KEY", awsSecretAccessKey).
		WithExec(deployArgs).
		Stderr(ctx)
}
