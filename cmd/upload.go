/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xiaoweize/cloudstation/store"
	"github.com/xiaoweize/cloudstation/store/aliyun"
	"github.com/xiaoweize/cloudstation/store/aws"
	"github.com/xiaoweize/cloudstation/store/tx"
)

var (
	provider     string
	ossBucket    string
	objectKey    string
	localFile    string
	endpoint     string
	accessKey    string
	accessSecret string
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload file",
	Long:  "upload 文件上传命令",
	Example: `upload -k accessKey  -s accessSecret -f localFile (default upload aliyunoss)
upload -p provider -b ossBucket -o objectKey -e endpoint -k accessKey  -s accessSecret -f localFile
	`,
	//执行upload命令时的处理逻辑
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			uploader store.Uploader
			err      error
		)
		switch provider {
		case "aliyun":
			uploader, err = aliyun.NewAliyunOss(&aliyun.Options{
				Endpoint:        endpoint,
				AccessKeyID:     accessKey,
				AccessKeySecret: accessSecret,
			})
		case "aws":
			uploader, err = aws.NewAwsoss()
		case "tx":
			uploader, err = tx.NewTxoss()
		default:
			return fmt.Errorf("no such a oss provide!")
		}
		if err != nil {
			return err
		}
		//使用对应云厂商已经实现的Upload方法上传，这里将objectKey设置成了localFile
		return uploader.Upload(ossBucket, localFile, localFile)
	},
}

func init() {
	//添加根命令root下的命令
	rootCmd.AddCommand(uploadCmd)
	f := rootCmd.PersistentFlags()
	f.StringVarP(&provider, "provider", "p", "aliyun", "oss provider[aliyun(default)/tx/aws]")
	f.StringVarP(&endpoint, "endpoint", "e", "oss-cn-shanghai.aliyuncs.com", "oss endpoint")
	f.StringVarP(&ossBucket, "ossBucket", "b", "test-delicloud", "oss bucket")
	f.StringVarP(&accessKey, "accessKey", "k", "", "oss accessKey")
	f.StringVarP(&accessSecret, "accessSecret", "s", "", "oss accessSecret")
	f.StringVarP(&localFile, "localFile", "f", "", "local file path name")
}
