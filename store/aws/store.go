package aws

type Awsoss struct{}

func NewAwsoss() (*Awsoss, error) {
	return &Awsoss{}, nil
}

func (awsoss *Awsoss) Upload(ossbucket, objectKey, localfile string) error {
	return nil
}
