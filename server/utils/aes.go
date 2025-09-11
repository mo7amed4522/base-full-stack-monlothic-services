package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
)

// 密钥（可自定义，长度必须为 16、24 或 32 字节）
var key = []byte("pay451437841#xXy")

// EncryptAES 使用 AES 加密明文，返回加密后的 Base64 编码字符串
func EncryptAES(plaintext string) (string, error) {
	// 创建一个 cipher.Block 实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 将明文转换为字节数组
	plaintextBytes := []byte(plaintext)

	// 对数据进行填充
	paddedPlaintext := pad(plaintextBytes, block.BlockSize())

	// 初始化向量 IV
	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 创建一个 CBC 模式的加密器
	ciphertext := make([]byte, len(paddedPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	// 将 IV 附加到密文的开头
	ciphertext = append(iv, ciphertext...)

	// 返回加密后的 Base64 编码密文
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptAES 解密密文，返回明文字符串
func DecryptAES(ciphertextBase64 string) (string, error) {
	// 解码 Base64 编码的密文
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return "", err
	}

	// 创建一个 cipher.Block 实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 提取 IV（前 16 个字节）
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// 创建 CBC 解密器
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// 去除填充字符
	plaintext = pkcs7Unpad(plaintext, block.BlockSize())

	// 返回解密后的明文字符串
	return string(plaintext), nil
}

// pad 使用 PKCS#7 填充数据，确保长度是块大小的倍数
func pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padded := make([]byte, len(data)+padding)
	copy(padded, data)
	for i := len(data); i < len(padded); i++ {
		padded[i] = byte(padding)
	}
	return padded
}

// pkcs7Unpad 移除 PKCS#7 填充
func pkcs7Unpad(data []byte, blockSize int) []byte {
	padding := data[len(data)-1]
	if padding < 1 || padding > byte(blockSize) {
		return nil
	}
	for i := 0; i < int(padding); i++ {
		if data[len(data)-i-1] != padding {
			return nil
		}
	}
	return data[:len(data)-int(padding)]
}

func main() {
	// 需要加密的明文
	plaintext := "Hello, AES encryption in Go!"

	// 加密
	encrypted, err := EncryptAES(plaintext)
	if err != nil {
		log.Fatal("加密失败:", err)
	}
	fmt.Println("加密后的密文 (Base64 编码):", encrypted)

	// 解密
	decrypted, err := DecryptAES(encrypted)
	if err != nil {
		log.Fatal("解密失败:", err)
	}
	fmt.Println("解密后的明文:", decrypted)
}
