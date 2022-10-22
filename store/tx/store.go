package tx

type Txoss struct{}

func NewTxoss() (*Txoss, error) {
	return &Txoss{}, nil
}

func (txoss *Txoss) Upload(ossbucket, objectKey, localfile string) error {
	return nil
}
