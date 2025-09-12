package upload

import (
	"mime/multipart"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// OSS 对象存储接口
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

// NewOss OSS的实例化方法
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
func NewOss() OSS {
	switch global.GVA_CONFIG.System.OssType {
	case "local":
		return &Local{}
	case "qiniu":
		return &Qiniu{}
	case "tencent-cos":
		return &TencentCOS{}
	case "aliyun-oss":
		return &AliyunOSS{}
	case "huawei-obs":
		return HuaWeiObs
	case "aws-s3":
		return &AwsS3{}
	case "cloudflare-r2":
		return &CloudflareR2{}
	case "minio":
		minioClient, err := GetMinio(global.GVA_CONFIG.Minio.Endpoint, global.GVA_CONFIG.Minio.AccessKeyId, global.GVA_CONFIG.Minio.AccessKeySecret, global.GVA_CONFIG.Minio.BucketName, global.GVA_CONFIG.Minio.UseSSL)
		if err != nil {
			global.GVA_LOG.Warn("You configured to use minio, but initialization failed. Please check minio availability or security configuration.: " + err.Error())
			panic("Minio initialization failed") // 建议这样做，用户自己配置了minio，如果报错了还要把服务开起来，使用起来也很危险
		}
		return minioClient
	default:
		return &Local{}
	}
}
