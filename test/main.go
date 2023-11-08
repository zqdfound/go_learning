package main

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func main() {

	//key := "nJ8MFgKxWvvJNZ1HyPEJEVy679MJLPFL"
	//iv := "brXgwsG60EYQHlee"
	//keyBase64 := "bko4TUZnS3hXdnZKTloxSHlQRUpFVnk2NzlNSkxQRkw="
	//ivBase64 := "YnJYZ3dzRzYwRVlRSGxlZQ=="
	// 生成32字节key
	keyByte := make([]byte, 32)
	rand.Read(keyByte)
	// 生成16字节iv
	ivByte := make([]byte, 16)
	rand.Read(keyByte)
	// keyByte和ivByte进行base64编码
	key := base64.StdEncoding.EncodeToString(keyByte)
	iv := base64.StdEncoding.EncodeToString(ivByte)

	fmt.Println("===key====")
	fmt.Println(key)
	fmt.Println("===iv====")
	fmt.Println(iv)

	name := "庄庆典"
	idcard := "330327199606137719"
	//encKey := ""
	//encIV := ""
	//平台公钥
	publicKey := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzX2SwF2IILkYk9Hxyodn\nJTr7U4FeVNhXfbrL2XtYMuAuhdMmiTrQaI2DqvQHWeb+FcBa7NewiP8s+jqg9OOK\nEKhzQyvzIUJasuYPw8Kniu36k4KQpGzW6RlNAgulUstaZmy6oetZMFNx3MaMvIMZ\nQEmsrvJPmQQvk81vHebA9HtGv2SWvcV/+A9qq2F5/tDiU65OwM7PoLKuY2lKr30a\nbLMj98yoNHqeBc2WncR4VJr48KJzGvF4yvfaXsS8FcnSk1PW9Wb+X45kcxj0iHy8\nH6bGPClIbwOkgEol8b7iJmNgwkgmXpi7Pt3XnbVJokaJumRij/LxEZaL9TL+ALLa\neQIDAQAB\n-----END PUBLIC KEY-----"
	//应用公钥
	//publicKey := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5jSnlHWp00vrYesMlDxu\nJ9wcOkcPDoSB/4tN2o1rV/j6OYmN/bzsL+JGyzwYwoZAKN3TmpSwKukoAUWOngAa\nO+yFs/zisgFQmp4J7y+BoPxqumtJQc6PGSgiZeFHLuvxlOK3X9tw5/vGFwD4cH2E\n7GM5NgizpKcWQt/uoIdk5ZNuvHPO0cw8IGi9pe2tk7Qc2GlwkTIeyxPEY4LQ7huu\nInjA+wB12cqvBiNQyp7yn7wHqWCgLfs4Nx/Skrx5EmGzT1wrSXzcMHEBTDUL5IXp\nf2zur78Q/AhQpCNpdfQes4NqDihEpc2adW7UhmAnSFFir+0VrMK1O8aV70GywK/j\n0wIDAQAB\n-----END PUBLIC KEY-----\n"
	// rsa公钥匙加密示例
	block, _ := pem.Decode([]byte(publicKey))
	platformPublicKeyInstance, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Sprintf(err.Error())
	} else {
		publicKeyInstance := platformPublicKeyInstance.(*rsa.PublicKey)

		ciperDataBytesKey, err := rsa.EncryptPKCS1v15(rand.Reader, publicKeyInstance, []byte(key))
		if err != nil {
			fmt.Sprintf(err.Error())
		} else {
			ciperDataKEY := base64.StdEncoding.EncodeToString(ciperDataBytesKey)
			fmt.Println("===加密key==")
			fmt.Println(ciperDataKEY)
			//encKey = ciperDataKEY
		}

		ciperDataBytesIV, err := rsa.EncryptPKCS1v15(rand.Reader, publicKeyInstance, []byte(iv))
		if err != nil {
			fmt.Sprintf(err.Error())
		} else {
			ciperDataIV := base64.StdEncoding.EncodeToString(ciperDataBytesIV)
			fmt.Println("===加密IV==")
			fmt.Println(ciperDataIV)
			//encIV = ciperDataIV
		}
	}
	encName, err := AesCBCEncrypt(nil, name, key, iv)
	fmt.Println("===加密Name==")
	//fmt.Println(err.Error())
	fmt.Println(encName)

	//
	encIdCard, err := AesCBCEncrypt(nil, idcard, key, iv)
	fmt.Println("===加密Idcard==")
	//fmt.Println(err.Error())
	fmt.Println(encIdCard)

}

// AES-CBC加密明文数据
// key：       密钥
// iv：        初始向量
// originData：敏感数据

// 加密
func AesCBCEncrypt(ctx context.Context, originData, key, iv string) (string, error) {
	// base64解码key和iv
	keyByte, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	ivByte, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}

	// 加密流程
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	originDataByte := PKCS7Padding([]byte(originData), blockSize)
	blockMode := cipher.NewCBCEncrypter(block, ivByte)
	ciperData := make([]byte, len(originDataByte))
	blockMode.CryptBlocks(ciperData, originDataByte)
	return base64.StdEncoding.EncodeToString(ciperData), nil
}

// 解密
func AesCBCDecrypt(ctx context.Context, ciperData, key, iv string) (originText string, err error) {
	// base64解码密文ciperData、密钥key、初始向量iv
	ciperDataBytes, err := base64.StdEncoding.DecodeString(ciperData)
	if err != nil {
		return "", err
	}
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}

	// 解密流程
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, ivBytes)
	originDataByte := make([]byte, len(ciperDataBytes))
	blockMode.CryptBlocks(originDataByte, ciperDataBytes)
	// unpadding
	originDataByte = PKCSUnPadding(originDataByte)
	return string(originDataByte), nil
}

func PKCS7Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - len(plainText)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plainText, padtext...)
}
func PKCSUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
