package aliyun_test 

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xiaoweize/cloudstation/store"
	"github.com/xiaoweize/cloudstation/store/aliyun"
)

var (
	uploader store.Uploader
)

var (
	endpoint        = os.Getenv("ALI_endpoint")
	accessKeyID     = os.Getenv("ALI_accessKeyID")
	accessKeySecret = os.Getenv("ALI_accessKeySecret")
)

var (
	ossbucket = os.Getenv("ALI_OssBucket")
)

//测试代码上传功能是否正常
func TestUpload(t *testing.T) {
	should := assert.New(t)
	err := uploader.Upload(ossbucket, "test.go", "store_test.go")
	if should.NoError(err) {
		t.Log("upload ok")
	}
}

//测试代码上传报错
func TestUploadError(t *testing.T) {
	should := assert.New(t)
	//制造错误，设置一个不存在的本地文件
	err := uploader.Upload(ossbucket, "test.go", "store_test.g")
	//断言错误是否为上面制造的预期错误
	should.Error(err, "open store_test.g: no such file or directory")
}

//初始化uploader
func init() {
	c, err := aliyun.NewAliyunOss(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		panic(err)
	}
	uploader = c
}
