package RSA

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

var Bits int

// 生成密钥对并保存到文件
func GenerateRSAKey(uid string) error {
	// 1、RSA生成私钥文件的核心步骤：
// 	flag.IntVar(&RSA.Bits, "key flag", 1024, "密钥长度，默认值为1024位")
	// 1) 生成RSA密钥对
	privateKer, err := rsa.GenerateKey(rand.Reader, Bits)
	if err != nil {
		return err
	}
	// 2) 将私钥对象转换成DER编码形式
	derPrivateKer := x509.MarshalPKCS1PrivateKey(privateKer)
	// 3) 创建私钥pem文件
	path := "./files/private/" + uid + ".pem"
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	// 4) 对密钥信息进行编码，写入到私钥文件中
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derPrivateKer,
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	// 2、RSA生成公钥文件的核心步骤：
	// 1) 生成公钥对象
	publicKey := &privateKer.PublicKey
	// 2) 将公钥对象序列化为DER编码格式
	derPublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	// 3) 创建公钥pem文件
	path2 := "./files/public/" + uid + ".pem"
	file, err = os.Create(path2)
	// file, err = os.Create("./files/public.pem")
	if err != nil {
		return err
	}
	// 4) 对公钥信息进行编码，写入到公钥文件中
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPublicKey,
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	return nil
}
