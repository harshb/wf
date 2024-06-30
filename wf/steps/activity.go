package steps

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
)

func CreateActivity(activityName string) (string, error) {
	// Create a new session
	sess := session.Must(session.NewSession())

	// Create a new Step Functions service client
	svc := sfn.New(sess)

	// Create a new activity
	response, err := svc.CreateActivity(&sfn.CreateActivityInput{
		Name: &activityName,
	})
	if err != nil {
		return "", err
	}

	// Return the ARN of the newly created activity
	return *response.ActivityArn, nil
}
