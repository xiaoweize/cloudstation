package aliyun

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/xiaoweize/cloudstation/store"
)

var (
	// 对象是否实现了接口的约束
	// a string = "abc"
	// _ store.Uploader 我不需要这个变量的值, 我只是做变量类型的判断
	// &AliOssStore{} 这个对象 必须满足 store.Uploader
	// _ store.Uploader = &AliOssStore{} 声明了一个空对象, 只是需要一个地址
	// nil 空指针, nil有没有类型？: 有类型
	// a *AliOssStore = nil   nil是一个AliOssStore 的指针
	// 如何把nil 转化成一个 指定类型的变量
	//    a int = 16
	//    b int64 = int64(a)
	//    (int64类型)(值)
	//	  (*Aliyunoss)(nil)
	_ store.Uploader = (*Aliyunoss)(nil)
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
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
}

//验证参数
func (o *Options) Validate() error {
	if o.Endpoint == "" || o.AccessKeyID == "" || o.AccessKeySecret == "" {
		return fmt.Errorf("Endpoint,AccessKeyID,AccessKeySecret has one empty!")
	}
	return nil
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
