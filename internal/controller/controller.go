package ctx

import (
	"gopress/pkg"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RsaKeyItem struct {
	PublicKey  string `json:"public_key"`  // 公钥
	PrivateKey string `json:"private_key"` // 私钥
	Value      string `json:"value"`       // 加密值
	Expire     int64  `json:"expire"`      // 过期时间戳
}

var RsaMap = make(map[string]RsaKeyItem)

// 获取RSA密钥
func GetRsaKey() (string, string, error) {
	publicKey, privateKey, err := pkg.RsaGenerate()
	if err != nil {
		return "", "", err
	}

	rasItem := RsaKeyItem{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		Value:      "",
		Expire:     time.Now().Add(time.Minute * 10).UnixMilli(),
	}

	rsaKey := uuid.New().String()
	RsaMap[rsaKey] = rasItem

	// 创建定时器 删除过期的密钥
	time.AfterFunc(time.Minute*10, func() {
		delete(RsaMap, rsaKey)
	})

	list := strings.Split(rasItem.PublicKey, "\n")
	publicKey = strings.Join(list[1:len(list)-2], "\n")

	return rsaKey, publicKey, nil
}

// 成功数据
func CtxSuccess(data interface{}) interface{} {
	return fiber.Map{
		"code": 0,
		"msg":  "success",
		"data": data,
	}
}

// 失败数据
func CtxError(code int, msg string, data ...interface{}) interface{} {
	return fiber.Map{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}
