package aws

import "fmt"

type Awsoss struct{}

func NewAwsoss() (*Awsoss, error) {
	return &Awsoss{}, fmt.Errorf("aws oss not impl")
}

func (awsoss *Awsoss) Upload(ossbucket, objectKey, localfile string) error {
	return nil
}
