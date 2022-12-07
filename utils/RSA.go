package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var (
	privateKeyStr string
	publicKeyStr  string
)

func init() {
	////GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	////Reader是一个全局、共享的密码用强随机数生成器
	//privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	//if err != nil {
	//	panic(err)
	//}
	//
	////通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	//X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	////使用pem格式对x509输出的内容进行编码
	////构建一个pem.Block结构体对象
	//privateBlock := &pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//
	//privateKeyStr = string(pem.EncodeToMemory(privateBlock))
	//
	////获取公钥的数据
	//publicKey := privateKey.PublicKey
	//
	////X509对公钥编码
	//X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	//if err != nil {
	//	panic(err)
	//}
	////pem格式编码
	////创建一个pem.Block结构体对象
	//publicBlock := &pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//
	//publicKeyStr = string(pem.EncodeToMemory(publicBlock))

	privateKeyStr = `-----BEGIN RSA Private Key-----
MIICXgIBAAKBgQC3WTCs9B1Q7iSEDgcK1IDOvPdcLv5IbyUQFBKwLArjctUMqwky
HaBmr4MfY2s3wn7bUslP8p1VzAW1Y+z4xThZzJbw0hhR6BYHwbX5u65RQwHv2R4i
VeA8byGVdsNEBlt9r0KINZrMSa+U/n9uut20zwx3taSt2uXnT8+Cf2hhSwIDAQAB
AoGBAInoxWsiRzbLmZ3Wq0djJevbSTgGhO/Y4gjhAFmRFzOT+VqI/+a5UCBM3hTr
BAex/RddgtzmwZ96UcTpf6JzPflOXaP9CvxAsRTXTL775fTqFLD/+AGD2mnkEh4t
Gr0yIdWwgi/0Z2T+b4J4tKLPO8FlbCWyqkZp2zCSfoqsRH9hAkEAxdn0sBGP3P+T
Xq5bIrV/IqfCWZ/LRbtujfcmKwJNKSjy5B7oqkfzk9UM2CyqjFubReSMOIoLe3eh
w8OaedeYGwJBAO08DQXeblnaA1wJYOi5gDP8FM5WRuxaORn+w34U7UMv3690aBfV
urMZPuXtQG949itCdjoF9v7lXZvhS4kJTpECQB0Liy1R1rHV3zeWFxD4XgqjYBey
KrA4/NjggHzt4I/7T/UHSJa+61Y3f6Q4omzqgdf33lHqihb9EAViMMpUSSMCQQC3
dis0ELISNnFGqdIR5/LnQNjuQPaULanfGpAgXxqlTM6Vp1YgqJ67hpiHw5SDIG5v
QfSgJqC+uq5LsE1z6oihAkEAmCgsupSXMljYUyOfto6eJLbkc0kkettSTwT7lJP2
ir7gHx6sgLdTXV2opZ4s1DSjGe9tGGSV5yZZU0LJ1pNUGA==
-----END RSA Private Key-----`
	publicKeyStr = `-----BEGIN RSA Public Key-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC3WTCs9B1Q7iSEDgcK1IDOvPdc
Lv5IbyUQFBKwLArjctUMqwkyHaBmr4MfY2s3wn7bUslP8p1VzAW1Y+z4xThZzJbw
0hhR6BYHwbX5u65RQwHv2R4iVeA8byGVdsNEBlt9r0KINZrMSa+U/n9uut20zwx3
taSt2uXnT8+Cf2hhSwIDAQAB
-----END RSA Public Key-----`
}

//
// RSAEncrypt
//  @Description: 加密
//  @param plainText
//  @param path
//  @return []byte
//
func RSAEncrypt(plainText []byte) ([]byte, error) {
	//pem解码
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		return nil, errors.New("public key error")
	}

	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	//对明文进行加密 主要是计算 c = m**e mod |n|
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
}

//
// RSADecrypt
//  @Description: 解密
//  @param cipherText
//  @param path
//  @return []byte
//
func RSADecrypt(cipherText []byte) ([]byte, error) {
	//pem解码
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return nil, errors.New("private key error")
	}

	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	//对密文进行解密 主要执行 m = c**d mod |n|
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		return nil, err
	}

	//返回明文
	return plainText, nil
}
