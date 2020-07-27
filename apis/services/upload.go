package services

import (
	"context"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"io"
	"smiling-blog/config"
)

type UploadServices interface {
	Qiniu(localFile io.Reader, size int64, filename string) (string, error)
}

type Upload struct {
}

func NewUploadServices() UploadServices {
	return &Upload{}
}

// 上传七牛云
func (slf *Upload) Qiniu(localFile io.Reader, size int64, filename string) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: config.Configs.QiNiu.Bucket,
	}
	mac := qbox.NewMac(config.Configs.QiNiu.Ak, config.Configs.QiNiu.Sk)
	upToken := putPolicy.UploadToken(mac)
	cig := slf.qiniuStorage()
	formUploader := storage.NewFormUploader(&cig)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err := formUploader.Put(context.Background(), &ret, upToken, filename, localFile, size, &putExtra)
	if err != nil {
		return "", err
	}
	return ret.Key, nil
}

func (slf *Upload) qiniuStorage() storage.Config {
	c := storage.Config{}
	c.Zone = &storage.ZoneHuanan
	c.UseHTTPS = false
	c.UseCdnDomains = false
	return c
}
