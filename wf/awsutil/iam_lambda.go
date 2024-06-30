package awsutil

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"wf/locals"
)

func GetBasicLambdaRole() (string, error) {
	sess := session.Must(session.NewSession())
	svc := iam.New(sess)

	// Check if the role already exists
	getRoleInput := &iam.GetRoleInput{
		RoleName: aws.String(locals.BasicLambdaExecutionRoleName),
	}
	if _, err := svc.GetRole(getRoleInput); err == nil {
		accountID, err := GetAWSAccountID()
		if err != nil {
			return "", err
		}
		// If the role exists, return its ARN
		return fmt.Sprintf("arn:aws:iam::%s:role/%s", accountID, locals.BasicLambdaExecutionRoleName), nil
	}

	// If the role does not exist, create it
	createRoleInput := &iam.CreateRoleInput{
		AssumeRolePolicyDocument: aws.String(`{
			"Version": "2012-10-17",
			"Statement": [
				{
					"Effect": "Allow",
					"Principal": {
						"Service": "lambda.amazonaws.com"
					},
					"Action": "sts:AssumeRole"
				}
			]
		}`),
		Path:     aws.String("/"),
		RoleName: aws.String(locals.BasicLambdaExecutionRoleName),
	}

	result, err := svc.CreateRole(createRoleInput)
	if err != nil {
		return "", err
	}

	// Attach the policy to the role
	attachRolePolicyInput := &iam.AttachRolePolicyInput{
		PolicyArn: aws.String("arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"),
		RoleName:  aws.String(locals.BasicLambdaExecutionRoleName),
	}

	if _, err := svc.AttachRolePolicy(attachRolePolicyInput); err != nil {
		return "", err
	}

	return *result.Role.Arn, nil
}
