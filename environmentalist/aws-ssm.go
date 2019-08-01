package environmentalist

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// EnvSSM is a data structure that holds relevent data for accessing AWS SSM
type EnvSSM struct {
	Session    *session.Session
	SSMService *ssm.SSM
}

// NewSSM returns a new pointer to a new instance of SSM with default credential values (set via environmental variables or profile and lastly, manual)
//
// Arguments:
//     regionName (string): The AWS region name to connect to
//
// Returns:
//     (*SSM): A pointer to a new instance of SSM
func NewSSM(regionName string) *EnvSSM {
	var (
		accessKeyID, secretAccessKey, awsToken string
		sess                                   *session.Session
		err                                    error
	)
	var S EnvSSM
	sess, err = S.WithEnvVars(regionName)
	if err != nil {
		sess, err = S.WithProfile(regionName, "default")
		if err != nil {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Please enter your AWS Access Key ID: ")
			accessKeyID, err = reader.ReadString('\n')
			fmt.Println("Please enter your AWS Secret Access Key: ")
			secretAccessKey, err = reader.ReadString('\n')
			fmt.Println("Please enter your AWS Access Token: ")
			awsToken, err = reader.ReadString('\n')
			sess, err = S.WithManualCredentials(regionName, accessKeyID, secretAccessKey, awsToken)
			log.Fatal(err)
		}
	}
	S.Session = sess
	S.SSMService = ssm.New(S.Session)
	return &S
}

// WithManualCredentials returns a pointer to an aws.session.Session struct to use to connect to SSM with manaual credentials
//
// Arguments:
//     regionName (string):      the AWS region to connect to
//     accessKeyID (string):     your AWS access key ID to use to connect
//     secretAccessKey (string): your AWS secret access key to use to connect
//     awsToken (string):        your AWS connection token (if available)
//
// Returns:
//     (*session.Session): The pointer to the aws.session.Session struct
func (*EnvSSM) WithManualCredentials(regionName string, accessKeyID string, secretAccessKey string, awsToken string) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(regionName),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, awsToken),
	})
	if err != nil {
		return nil, err
	}
	return sess, nil
}

// WithEnvVars returns a pointer to an aws.session.Session struct to use to connect to SSM via environmental variables
//
// Arguments:
//     regionName (string): the AWS region to connect to
//
// Returns:
//     (*session.Session): The pointer to the aws.session.Session struct
func (*EnvSSM) WithEnvVars(regionName string) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(regionName),
	})
	if err != nil {
		return nil, err
	}
	return sess, nil
}

// WithProfile returns a pointer to an aws.session.Session struct to use to connect to SSM via a shared credentials file
//
// Arguments:
//     regionName (string):  the AWS region to connect to
//     profileName (string): the AWS profile to use
//
// Returns:
//     (*session.Session): the pointer to the aws.session.Session struct
func (*EnvSSM) WithProfile(regionName string, profileName string) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(regionName),
		Credentials: credentials.NewSharedCredentials("", profileName),
	})
	if err != nil {
		return nil, err
	}
	return sess, nil
}

// Add adds a parameter to the AWS SSM store
//
// Arguments:
//     parameterName (string): The name of the parameter to add (e.g. - /root/folder/folder/parameterName)
//     parameterValue (interface{}): the value to set, could be a string or a []string
//     parameterType (string): the type of parameter (either String or StringList)
//
// Returns:
//     (error): an error if one exists
func (S *EnvSSM) Add(parameterName string, parameterValue interface{}, parameterType string) error {
	var (
		err error
	)

	switch parameterType {
	case "String":
		_, err = S.SSMService.PutParameter(&ssm.PutParameterInput{
			Name:  aws.String(parameterName),
			Value: aws.String(parameterValue.(string)),
			Type:  aws.String(parameterType),
		})
	case "StringList":
		values := strings.Join(parameterValue.([]string), ",")
		_, err = S.SSMService.PutParameter(&ssm.PutParameterInput{
			Name:  aws.String(parameterName),
			Value: aws.String(values),
			Type:  aws.String(parameterType),
		})
	}

	if err != nil {
		return fmt.Errorf("unable to add the %s parameter '%s':'%s due to error, err: %v", parameterType, parameterName, parameterValue, err)
	}
	return nil
}

// Modify modifies a parameter in the AWS SSM store
//
// Arguments:
//     parameterName (string): The name of the parameter to modify (e.g. - /root/folder/folder/parameterName)
//     parameterValue (interface{}): the value to set, could be a string or a []string
//     parameterType (string): the type of parameter (either String or StringList)
//     overwrite (bool): overwrite the value in AWS SSM rather than create a new version
//
// Returns:
//     (error): an error if one exists
func (S *EnvSSM) Modify(parameterName string, parameterValue interface{}, parameterType string, overwrite bool) error {
	var (
		err error
	)

	switch parameterType {
	case "String":
		_, err = S.SSMService.PutParameter(&ssm.PutParameterInput{
			Name:      aws.String(parameterName),
			Value:     aws.String(parameterValue.(string)),
			Type:      aws.String(parameterType),
			Overwrite: aws.Bool(overwrite),
		})
	case "StringList":
		values := strings.Join(parameterValue.([]string), ",")
		_, err = S.SSMService.PutParameter(&ssm.PutParameterInput{
			Name:      aws.String(parameterName),
			Value:     aws.String(values),
			Type:      aws.String(parameterType),
			Overwrite: aws.Bool(overwrite),
		})
	}

	if err != nil {
		return fmt.Errorf("unable to add the %s parameter '%s':'%s due to error, err: %v", parameterType, parameterName, parameterValue, err)
	}
	return nil
}

// Delete removes a parameter from AWS SSM
//
// Arguments:
//     parameterName (string): The name of the parameter to remove
//
// Returns:
//     (error): an error if one exists
func (S *EnvSSM) Delete(parameterName string) error {
	_, err := S.SSMService.DeleteParameter(&ssm.DeleteParameterInput{
		Name: aws.String(parameterName),
	})
	if err != nil {
		return fmt.Errorf("unable to delete the parameter '%s' due to error, err: %v", parameterName, err)
	}
	return nil
}

// View returns the response from AWS SSM for a particular parameter
//
// Arguments:
//     parameterName (string): The name of the parameter to view
//
// Returns:
//     (*ssm.GetParameterOutput): The aws.ssm.GetParameter output value
//     (error): an error if one exists
func (S *EnvSSM) View(parameterName string) (*ssm.GetParameterOutput, error) {
	retVal, err := S.SSMService.GetParameter(&ssm.GetParameterInput{
		Name: aws.String(parameterName),
	})
	if err != nil {
		return nil, fmt.Errorf("unable to fetch the parameter '%s' due to error, err: %v", parameterName, err)
	}
	return retVal, nil
}
