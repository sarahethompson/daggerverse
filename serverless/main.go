// Run serverless deploys

package main

import (
	"context"
)

type Serverless struct{}

// Runs a serverless deploy
func (m *Serverless) Deploy(
	ctx context.Context,
	//configFile *File,
	configDir *Directory,
	awsAccessKeyID *Secret,
	awsSecretAccessKey *Secret,
	awsSessionToken *Secret,
	// +optional
	serverlessAccessKey *Secret,
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

	//create container
	container := dag.Container().From("alpine:latest")

	//set optional secret
	if serverlessAccessKey != nil {
		container = container.WithSecretVariable("SERVERLESS_ACCESS_KEY", serverlessAccessKey)
	}

	container = container.WithExec([]string{"apk", "add", "curl", "bash", "npm"}).
		WithExec([]string{"npm", "install", "-g", "serverless"}).
		WithDirectory("/host", configDir).
		WithWorkdir("/host").
		WithSecretVariable("AWS_ACCESS_KEY_ID", awsAccessKeyID).
		WithSecretVariable("AWS_SECRET_ACCESS_KEY", awsSecretAccessKey).
		WithSecretVariable("AWS_SESSION_TOKEN", awsSessionToken).
		WithExec(deployArgs)

	return container.Stdout(ctx)
}
