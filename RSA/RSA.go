package RSA

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

// 公钥加密
func RsaEncrypt(origData []byte, uid string) ([]byte, error) {
	//1.将文件当中的加密私钥读出存入缓冲区
	publicKeyFileName := "./files/public/" + uid + ".pem"
	privateFile, err := os.Open(publicKeyFileName)
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	fileInfo, err := privateFile.Stat()
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, fileInfo.Size())
	// fmt.Println("-----", buffer)

	privateFile.Read(buffer)

	block, _ := pem.Decode(buffer)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 私钥解密
func RsaDecrypt(ciphertext []byte, uid string) ([]byte, error) {
	//解密
	path := "./files/private/" + uid + ".pem"
	content, err := os.ReadFile(path)
	// fmt.Println("-----", content)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
	privateKey := content
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// 签名
func RSAUsePrivateKeySign(plainText []byte, uid string) []byte {
	// GenerateRSAKey(uid)
	//1.将文件当中的加密私钥读出存入缓冲区
	privateKeyFileName := "./files/private/" + uid + ".pem"
	privateFile, err := os.Open(privateKeyFileName)
	if err != nil {
		fmt.Println(uid)
		panic(err)
	}
	defer privateFile.Close()
	fileInfo, err := privateFile.Stat()
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, fileInfo.Size())
	privateFile.Read(buffer)
	//2.pem解码，将数据流转换成一个存有DER字符串的pem.Block块
	block, _ := pem.Decode(buffer)
	//3.使用x509方法，对DER字符串解析，获取到私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//4.生成明文消息对应的散列值
	hash256 := sha256.New()
	hash256.Write(plainText)
	hashValue := hash256.Sum(nil)
	//5.使用[私钥]和[散列值]完成签名（RSA包内的方法）
	codeSign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashValue)
	if err != nil {
		panic(err)
	}
	return codeSign
}

//使用RSA公钥校验函数
func RSAUsePublicKeyVerify(plainText, codeSign []byte, uid string) (flag bool) {
	// GenerateRSAKey(uid)
	//1.将文件当中的加密公钥读出存入缓冲区
	publicKeyFileName := "./files/public/" + uid + ".pem"
	publicFile, err := os.Open(publicKeyFileName)
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	fileInfo, err := publicFile.Stat()
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, fileInfo.Size())
	publicFile.Read(buffer)
	//2.pem解码，将数据流转换成一个存有DER字符串的pem.Block块
	block, _ := pem.Decode(buffer)
	//3.使用x509方法，对DER字符串解析，获取到公钥（记得断言）
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey := publicInterface.(*rsa.PublicKey)
	//4.生成明文消息对应的散列值
	hash256 := sha256.New()
	hash256.Write(plainText)
	hashValue := hash256.Sum(nil)
	//5.使用[公钥]和[散列值]完成签名校验（RSA包内的方法）
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashValue, codeSign)
	// fmt.Println("===")
	if err != nil {
		flag = false
	} else {
		flag = true
	}
	return
}
