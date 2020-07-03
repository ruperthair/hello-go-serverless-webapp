#!/usr/bin/env bash

set -e
GOOS=linux go build -o main
aws cloudformation package --template-file template.yaml --output-template-file serverless-output.yaml --s3-bucket gotestrh
rm main
aws cloudformation deploy --template-file serverless-output.yaml --stack-name gotestrh --capabilities CAPABILITY_IAM