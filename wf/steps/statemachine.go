package steps

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
)

func CreateStateMachine(stateMachineName string, definition string, roleArn string) (string, error) {
	// Create a new session
	sess := session.Must(session.NewSession())

	// Create a new Step Functions service client
	svc := sfn.New(sess)

	// Create a new state machine
	response, err := svc.CreateStateMachine(&sfn.CreateStateMachineInput{
		Name:       &stateMachineName,
		Definition: &definition,
		RoleArn:    &roleArn,
	})
	if err != nil {
		return "", err
	}

	// Return the ARN of the newly created state machine
	return *response.StateMachineArn, nil
}
