package awsutil

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/sts"
)

func GetAWSAccountID() (string, error) {
	sess := session.Must(session.NewSession())
	svc := sts.New(sess)

	// Call GetCallerIdentity to get the account ID
	callerIdentity, err := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		return "", err
	}

	return *callerIdentity.Account, nil
}
func DeleteResource(resourceType, resourceName string) error {
	sess := session.Must(session.NewSession())

	accountID, err := GetAWSAccountID()
	if err != nil {
		return err
	}

	arn := fmt.Sprintf("arn:aws:%s::%s:%s", resourceType, accountID, resourceName)

	switch resourceType {
	case "lambda":
		svc := lambda.New(sess)
		_, err := svc.DeleteFunction(&lambda.DeleteFunctionInput{
			FunctionName: &arn,
		})
		return err
	case "step function":
		svc := sfn.New(sess)
		_, err := svc.DeleteStateMachine(&sfn.DeleteStateMachineInput{
			StateMachineArn: &arn,
		})
		return err
	case "iam":
		svc := iam.New(sess)
		_, err := svc.DeleteRole(&iam.DeleteRoleInput{
			RoleName: &resourceName, // IAM uses the role name, not the ARN
		})
		return err
	default:
		return fmt.Errorf("Invalid resource type: %s", resourceType)
	}
}

func DetachPolicy(resourceType, resourceName, policyArn string) error {
	if resourceType != "iam" {
		return fmt.Errorf("Can only detach policies from IAM roles")
	}

	sess := session.Must(session.NewSession())
	svc := iam.New(sess)

	_, err := svc.DetachRolePolicy(&iam.DetachRolePolicyInput{
		RoleName:  &resourceName,
		PolicyArn: &policyArn,
	})

	return err
}
