package cloud

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jphillips2121/games-api/models"
)

// AWS struct provides the configs required for AWS integration.
type AWS struct {
	FileName  string
	AwsRegion string
	AwsID     string
	AwsSecret string
	AwsToken  string
	AwsBucket string
}

// IsValidDeveloper returns if the developer is present in a list of approved developers.
func (a *AWS) IsValidDeveloper(game *models.Game) (bool, error) {

	// Download list of valid developers from S3
	err := a.downloadDeveloperFile()
	if err != nil {
		return false, err
	}

	return a.checkIfValidDeveloper(game)
}

// Downloads the valid developer list from S3
func (a *AWS) downloadDeveloperFile() error {

	// Create new file for valid developers
	file, err := os.Create(a.FileName)
	if err != nil {
		return err
	}

	defer file.Close()

	// Create session with AWS
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(a.AwsRegion),
		Credentials: credentials.NewStaticCredentials(a.AwsID, a.AwsSecret, a.AwsToken)},
	)
	if err != nil {
		return err
	}

	// Download file from S3 and save to file
	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(a.AwsBucket),
			Key:    aws.String(a.FileName),
		})

	return err
}

// Checks if the proposed developer is in the downloaded list
func (a *AWS) checkIfValidDeveloper(game *models.Game) (bool, error) {
	// Read the JSON file
	developerFile, err := ioutil.ReadFile(a.FileName)
	if err != nil {
		return false, err
	}

	// Unmarshal bytes into an array of Developers
	developers := &models.Developers{}
	err = json.Unmarshal(developerFile, developers)
	if err != nil {
		return false, err
	}

	// Loop through each developer to see if it matches the proposed developer
	for _, developer := range developers.Developers {
		if developer.Name == game.Developer {
			return true, nil
		}
	}

	return false, nil
}
