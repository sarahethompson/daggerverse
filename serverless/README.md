# Serverless

Functions for using the [Serverless](https://www.serverless.com/framework/docs/providers/aws/cli-reference) framework.

Currently, only deploying to AWS is supported.

## Deploying

Deploys a serverless stack to AWS.

From the dagger cli:

```sh
dagger call -m github.com/sarahethompson/daggerverse/serverless deploy --aws-secret-access-key=env:AWS_SECRET_ACCESS_KEY --aws-access-key-id=env:AWS_ACCESS_KEY_ID --aws-session-token=env:AWS_SESSION_TOKEN --config-dir="example" --region=us-east-1
```

Required inputs:
- aws-secret-access-key
- aws-access-key-id
- aws-session-token
- config-dir

Optional inputs:
- region
- stage
- force
- serverlessAccessKey
