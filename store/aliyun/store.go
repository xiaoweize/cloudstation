package aliyun

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/xiaoweize/cloudstation/store"
)

var (
	//对象是否实现接口的约束
	_ store.Uploader = &Aliyunoss{}
)

type Aliyunoss struct {
	client *oss.Client
}

//阿里云实现上传操作
func (alioss *Aliyunoss) Upload(ossbucket, objectKey, localfile string) error {
	bucket, err := alioss.client.Bucket(ossbucket)
	if err != nil {
		return err
	}
	//后续优化点：不允许上传目录，去除localfile中的路径仅获取文件名
	err = bucket.PutObjectFromFile(objectKey, localfile)
	if err != nil {
		return err
	}
	downloadUrl, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载地址:%s\n,有效期1天", downloadUrl)
	return nil
}

//Aliyunoss对象构造函数
func NewAliyunOss(endpoint, accessKeyID, accessKeySecret string) (*Aliyunoss, error) {
	c, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return nil, err
	} else {
		return &Aliyunoss{
			client: c,
		}, nil
	}
}
