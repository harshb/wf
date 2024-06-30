package awsutil

import (
	"fmt"
	"wf/locals"
)

func CleanUp() error {
	// Detach the policy before deleting the role
	//err := DetachPolicy("iam", locals.BasicLambdaExecutionRoleName, "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole")
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	//
	//err = DeleteResource("iam", locals.BasicLambdaExecutionRoleName)
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}

	// Detach the policy before deleting the StepFunctionExecutionRole
	err := DetachPolicy("iam", locals.StepFunctionExecutionRoleName, "arn:aws:iam::aws:policy/service-role/AWSStepFunctionsFullAccess")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = DeleteResource("iam", locals.StepFunctionExecutionRoleName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successful cleanup!")
	return nil
}
