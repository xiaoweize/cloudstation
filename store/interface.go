package store

//定义上传接口 具体上传操作由各云厂商实现
type Uploader interface {
	Upload(ossbucket, objectKey, localfile string) error
}
