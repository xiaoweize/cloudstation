package tx

import "fmt"

type Txoss struct{}

func NewTxoss() (*Txoss, error) {
	return &Txoss{}, fmt.Errorf("tx oss not impl!")
}

func (txoss *Txoss) Upload(ossbucket, objectKey, localfile string) error {
	return nil
}
