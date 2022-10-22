package aliyun

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-playground/validator"
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
	fmt.Printf("文件下载地址:%s\n,有效期1天\n", downloadUrl)
	return nil
}

type Options struct {
	Endpoint        string `validate:"required"`
	AccessKeyID     string `validate:"required"`
	AccessKeySecret string `validate:"required"`
}

//验证参数
func (o *Options) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

//Aliyunoss对象构造函数，使用构造函数生成示例，再调用Upload方法
func NewAliyunOss(opts *Options) (*Aliyunoss, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}
	c, err := oss.New(opts.Endpoint, opts.AccessKeyID, opts.AccessKeySecret)
	if err != nil {
		return nil, err
	} else {
		return &Aliyunoss{
			client: c,
		}, nil
	}
}
