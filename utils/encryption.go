package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	lua "github.com/yuin/gopher-lua"

	"go-jocy/config"
	"go-jocy/internal/model"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const publicKeyPEM = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1eiQTbqff456g/MlFTOi
cxAsw5kMac3ZGmjm+xCCggkuITKHx0Ae3B9EdoKDrYfZXSrC2Llty89RRGasZ36S
zlZem+s8c4A/OIcXzbbGteCIk/ItlPMZZzZlVnIWC1OtLFuisadzbZjOLpxmAl+C
cgUP2gcsaY8gQvXnkzJch8LoIcR+9orCW8zamPJOwoFq/sTDq0xP+TvUtGt2pijp
ed0uv6fAE6rOqZoRjheFAEJMLQyNeZxeQfN7OqFnJGNq1MhwIwZ2BP78TEf+zqsh
YdZPkXIlzISzoEi8P+HCtTc1veC9pJrDh0s7HoWmAFE8tFTj65gLWNS+0PvcLCOW
DQIDAQAB
-----END PUBLIC KEY-----`

const privateKeyPEM = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAxm2Kzu9L/FNX42em9Xo73JXtJCtrhleKN9jqclpK6/Iyah/T
UjH5RCNItWiTKHg6LcsGxZs9+4fP6uU8oO5Qp1akaOrJTg3QTQQFyRxDrv+LN/nL
6/MpSf3SnihyVPQWwlkj3yHZWrVC9HI3q2JmzGV/kzwnpVIj2as8zl4cO7OZr0F9
bR+G4jblqLPmB/x/BBOGrWWCxn+YI2RVHw23dev9jql284eN/KV4tlDlbtoJGy4+
Cb7nEV/THvRVZYbHAp+fMY0+NyqyslLX/btJqT8eSH6Hb8c+BSC77Dry4G8/m/wU
YPdvXiL3cVhZmEaqjs8rUafGyQmW3mrflAbIJwIDAQABAoIBAQCbRhUdIbzAUyev
V+kapvA5CUlsyF133wDV+vRbT7TZNcmlqgnfhCOe4k1/R7oALTS5qOo/r9+s+PYG
xiPPey26BN7bCv9ECSM7YS511ZxRUL9MqjidBscEk49BHD17pRY6Ny8O6JoBlV4z
kz1k67etsq9GNAiCIejT6F/IzXQicmO5MaJWCjBNSP+IPTvd5NW3DUNlt2NcCBvO
2sCgSq2Z4B0IdNWeSvd4ZmA2qSkqk60A8glNR4HdRTG9VWR0fUOd/qzpN1vjBUKM
aIHeUX50NCRdK8EGqrVOCq4uUgBRj7bjt0DOb9ck4vYxgBdkyK4HMYsAGdirYKxd
DkseicqhAoGBAOZ750Vky38kq3MucAE/uaFpaDUSeOKDy03fumM6TLlkeAxnTQZW
NBDzlrqNgQPMLu+tmm+ZsEN5buF8C2oKc97+Rz21rvrOufr4sX2PqHfj/kerYGq/
NzX1jpRbqsmcs+3JxveeozHuXBbOFpd8teCGZPFPHREFDe4sZFtFmwX1AoGBANxl
JNlKIz0TVmPgGCUZ4j8BRBiLPHMeFkUoam6Djou2iJLYey3ZNHhyMiRER7Smia0f
Y/QjqJIeSRWQlZExu6s9ijl4VSmoh4hLanOxxAE+gFnuhgK4XwMV0hvHbqSaupQd
fkULZ+t3rGKzt+t+0ob7xx+LjYWEwpLsKCQKRKgrAoGADLPvfyea/5rpyCNbEPaO
KJNCpwopl3JkFhqqjyV7bQxYgXaADEVcAUMrn4SFA8yRGaybwmLaEB31OoA3sNR6
pmOlUYVd63zRSz/BqIXuZw0tyo1rdvaq+FJcVVjoBMyaLhTc3nDj1bCpaqhZHmhF
Lea6UYJmu7VnmyTfMxiW/rECgYAh4MJLTGQiTUioTZgoi9QFT1KCW1TNdUCDHPVP
S5Wr0EEqIXC92XeBVDx06rIDCN586ChbLOgKnfEqCXGUQgrRBcKrlt2wa6F5x+3z
Hs48Srk8Gbgrzt97/+yuLHfLgaVQg0AXqOsufNTYzztkTbha23T+WltEvOWT5A0/
jPyExQKBgCGbq62piyIEeMNoP/SoLvh4hTq/eeNw5yCcLEsLrgt45Xb/2YgeyXWv
xTXl4c8bPdZTFYQ9A7IUYvhizpH032tDouqCsvgu3KtDO/pW6IteL17YBco7fRMQ
JhBuQjGDCMEGEJW76GwlXj/xUW32TN/5KeQXtHHZ4z2lZlJLU81B
-----END RSA PRIVATE KEY-----`

// pkcs7Padding 填充
func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// pkcs7Unpad 填充移除
func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("data is empty")
	}
	padding := int(data[length-1])
	if padding > length || padding > aes.BlockSize {
		return nil, fmt.Errorf("invalid padding size")
	}
	return data[:length-padding], nil
}

// RsaEncryption RSA加密
func RsaEncryption(plaintext string) (string, error) {
	// 解码公钥
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return "", fmt.Errorf("failed to decode PEM block")
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("not a valid RSA public key")
	}

	// 加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plaintext))
	if err != nil {
		return "", fmt.Errorf("failed to encrypt: %v", err)
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// RsaDecryption RSA解密
func RsaDecryption(encryptedText string) (string, error) {
	// 解码私钥
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return "", fmt.Errorf("failed to decode PEM block containing private key")
	}

	// 解析PKCS1格式私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}

	// 解码密文
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 ciphertext: %v", err)
	}

	// 解密
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt: %v", err)
	}

	return string(plaintext), nil
}

// AesEncryption AES加密
func AesEncryption(plaintext, key, iv string) (string, error) {
	// Convert key and iv to byte slices
	keyBytes := []byte(key)
	ivBytes := []byte(iv)
	plaintextBytes := []byte(plaintext)

	// Create AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Pad the plaintext to a multiple of the block size
	blockSize := block.BlockSize()
	plaintextBytes = pkcs7Padding(plaintextBytes, blockSize)

	// Create a CBC encrypter and encrypt the data
	ciphertext := make([]byte, len(plaintextBytes))
	mode := cipher.NewCBCEncrypter(block, ivBytes)
	mode.CryptBlocks(ciphertext, plaintextBytes)

	// Encode the ciphertext in Base64
	ciphertextBase64 := base64.StdEncoding.EncodeToString(ciphertext)
	return ciphertextBase64, nil
}

// AesEncryptionBase64 AES加密(Base64)
func AesEncryptionBase64(plaintext, key, iv string) (string, error) {
	// Convert key and iv to byte slices
	keyBytes := []byte(key)
	ivBytes := []byte(iv)
	plaintextBytes := []byte(plaintext)

	// 将文本 Base64 编码
	data2 := []byte(base64.StdEncoding.EncodeToString(plaintextBytes))

	// 创建 AES Cipher
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// 使用 CBC 模式加密
	ciphertext := make([]byte, len(pkcs7Padding(data2, aes.BlockSize)))
	cipher.NewCBCEncrypter(block, ivBytes).CryptBlocks(ciphertext, pkcs7Padding(data2, aes.BlockSize))

	// 返回加密后的 Base64 字符串
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AesDecryption AES解密
func AesDecryption(encryptedText, key, iv string) (string, error) {
	// 将密文解码为字节数组
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 ciphertext: %v", err)
	}

	// 将密钥和IV转换为字节数组
	keyBytes := []byte(key)
	ivBytes := []byte(iv)

	// 检查密钥和IV长度是否匹配AES要求
	if len(keyBytes) != 16 && len(keyBytes) != 24 && len(keyBytes) != 32 {
		return "", fmt.Errorf("invalid key length: %d bytes", len(keyBytes))
	}
	if len(ivBytes) != aes.BlockSize {
		return "", fmt.Errorf("invalid IV length: %d bytes (must be %d)", len(ivBytes), aes.BlockSize)
	}

	// 创建AES块
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %v", err)
	}

	// 检查密文长度是否为块大小的倍数
	if len(ciphertext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	// 使用 CBC 模式解密
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// 移除填充
	plaintext, err = pkcs7Unpad(plaintext)
	if err != nil {
		return "", fmt.Errorf("failed to unpad plaintext: %v", err)
	}

	return string(plaintext), nil
}

// MD5Encryption MD5加密
func MD5Encryption(plaintext string) string {
	hash := md5.Sum([]byte(plaintext))
	return hex.EncodeToString(hash[:])
}

// ResponseDecryption 响应解密
func ResponseDecryption(encryptedText string) (string, error) {
	// 根据[.]分割字符串
	parts := strings.Split(encryptedText, ".")
	rsaText := parts[0]
	aesText := parts[1]

	// RSA解密
	rsaKey, err := RsaDecryption(rsaText)
	if err != nil {
		return "", err
	}
	// 反转rsaKey得到rsaIV
	rsaIV := ReverseString(rsaKey)

	// AES解密
	return AesDecryption(aesText, rsaKey, rsaIV)
}

// MD5PlayUrlSign 生成播放地址签名
func MD5PlayUrlSign(jmStr, salt, ts string) string {
	combined := jmStr + salt + ts
	hash := md5.Sum([]byte(combined))
	return hex.EncodeToString(hash[:])
}

// DecryptPlayUrl 解密播放地址
func DecryptPlayUrl(source string) (any, error) {
	// 模拟设备信息
	platform := "Android"
	appVersion := config.GinConfig.App.AppVersion

	// 加密盐
	salt := "v50gjcy"

	// 时间戳
	ts := strconv.FormatInt(time.Now().Unix(), 10)

	modifiedSource := fmt.Sprintf("%s&t=%s", source, ts)

	// 创建请求
	client := resty.New()

	client.SetHeaderVerbatim("User-Agent", "Dart/2.17 (dart:io)")
	client.SetHeaderVerbatim("x-time", ts)
	client.SetHeaderVerbatim("x-form", platform)
	client.SetHeaderVerbatim("x-sign1", MD5PlayUrlSign(appVersion, salt, ts))
	client.SetHeaderVerbatim("x-sign2", MD5PlayUrlSign(source, salt, ts))

	// 退避重试
	client.SetRetryCount(3)
	client.SetRetryWaitTime(time.Second / 2)

	resp, err := client.R().Get("http://yhhy.xj.6b7.xyz/vo1v03.php?url=" + modifiedSource)
	if err != nil {
		return nil, err
	}
	config.GinLOG.Debug(resp.String())

	// AES解密
	key := "wcyjmnnnawozmydn"
	iv := "wcivwyjmlnzbhlmq"

	decrypted, err := AesDecryption(resp.String(), key, iv)
	if err != nil {
		return nil, err
	}
	config.GinLOG.Debug(decrypted)

	var pu model.PlayURL
	err = json.Unmarshal([]byte(decrypted), &pu)
	if err != nil {
		return nil, err
	}
	return pu, nil
}

// DecryptPlayUrlLUA 解密播放地址LUA版本
func DecryptPlayUrlLUA(luaScript, source string) (any, error) {
	// 创建一个新的Lua状态
	L := lua.NewState()
	defer L.Close()

	// 注册工具函数到Lua环境
	L.SetGlobal("utils", L.NewTable())
	L.SetField(L.GetGlobal("utils"), "md5", L.NewFunction(func(L *lua.LState) int {
		input := L.ToString(1)
		result := MD5Encryption(input)
		L.Push(lua.LString(result))
		return 1
	}))

	L.SetField(L.GetGlobal("utils"), "timestamp", L.NewFunction(func(L *lua.LState) int {
		ts := strconv.FormatInt(time.Now().Unix(), 10)
		L.Push(lua.LString(ts))
		return 1
	}))

	L.SetField(L.GetGlobal("utils"), "aes128cbc_decrypt", L.NewFunction(func(L *lua.LState) int {
		key := L.ToString(1)
		iv := L.ToString(2)
		encryptedText := L.ToString(3)

		decrypted, err := AesDecryption(encryptedText, key, iv)
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}

		L.Push(lua.LString(decrypted))
		return 1
	}))

	// 设置设备信息
	L.SetGlobal("device_info", L.NewTable())
	L.SetField(L.GetGlobal("device_info"), "platform", lua.LString("Android"))
	L.SetField(L.GetGlobal("device_info"), "app_version", lua.LString(config.GinConfig.App.AppVersion))

	// 注册HTTP请求函数
	L.SetGlobal("httpGet", L.NewFunction(func(L *lua.LState) int {
		url := L.ToString(1)
		options := L.ToTable(2)

		// 创建HTTP客户端
		client := resty.New()
		client.SetRetryCount(3)
		client.SetRetryWaitTime(time.Second / 2)

		// 设置请求头
		req := client.R()
		if options != nil {
			headerTable := options.RawGetString("header")
			if headerTable, ok := headerTable.(*lua.LTable); ok {
				headerTable.ForEach(func(k, v lua.LValue) {
					req.SetHeaderVerbatim(k.String(), v.String())
				})
			}
		}

		// 发送请求
		resp, err := req.Get(url)
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}

		L.Push(lua.LString(resp.String()))
		return 1
	}))

	// 注册JSON解析函数
	jsonMod := L.NewTable()
	L.SetGlobal("json", jsonMod)

	// 添加decode函数
	L.SetField(jsonMod, "decode", L.NewFunction(func(L *lua.LState) int {
		jsonStr := L.ToString(1)

		var result interface{}
		err := json.Unmarshal([]byte(jsonStr), &result)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LNumber(0))
			L.Push(lua.LString(err.Error()))
			return 3
		}

		// 将Go对象转换为Lua对象
		luaObj := valueToLua(L, result)
		L.Push(luaObj)
		L.Push(lua.LNumber(len(jsonStr) + 1))
		L.Push(lua.LNil)
		return 3
	}))

	// 添加encode函数
	L.SetField(jsonMod, "encode", L.NewFunction(func(L *lua.LState) int {
		value := L.CheckAny(1)

		goValue := luaToValue(L, value)
		jsonBytes, err := json.Marshal(goValue)
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}

		L.Push(lua.LString(string(jsonBytes)))
		return 1
	}))

	// 执行Lua脚本
	err := L.DoString(luaScript)
	if err != nil {
		return nil, fmt.Errorf("lua script execution error: %v", err)
	}

	// 调用parser函数
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("parser"),
		NRet:    4,
		Protect: true,
	}, lua.LString(source)); err != nil {
		return nil, fmt.Errorf("failed to call parser function: %v", err)
	}

	// 获取返回值
	status := L.Get(-4).String()
	url := L.Get(-3).String()
	headerStr := L.Get(-2).String()
	mediaType := L.Get(-1).String()

	// 清理栈
	L.Pop(4)

	if status != "OK" {
		return nil, fmt.Errorf("parser returned error status: %s", status)
	}

	// 构建返回结果
	var urlField model.URLField
	if mediaType == "multi" {
		// 解析多URL格式
		var multiUrls []model.URLDetail
		err := json.Unmarshal([]byte(url), &multiUrls)
		if err != nil {
			return nil, fmt.Errorf("failed to parse multi URLs: %v", err)
		}
		urlField.Multi = multiUrls
	} else {
		// 单URL格式
		urlField.Single = url
	}

	// 如果有自定义请求头，可以在这里处理
	_ = headerStr

	result := model.PlayURL{
		Code:    200,
		Success: 1,
		Type:    mediaType,
		Url:     urlField,
		Msg:     "success",
	}
	return result, nil
}

// valueToLua 将Go值转换为Lua值
func valueToLua(L *lua.LState, v interface{}) lua.LValue {
	switch val := v.(type) {
	case nil:
		return lua.LNil
	case bool:
		return lua.LBool(val)
	case float64:
		return lua.LNumber(val)
	case int:
		return lua.LNumber(val)
	case int64:
		return lua.LNumber(val)
	case string:
		return lua.LString(val)
	case []interface{}:
		table := L.NewTable()
		for i, item := range val {
			L.RawSetInt(table, i+1, valueToLua(L, item))
		}
		return table
	case map[string]interface{}:
		table := L.NewTable()
		for k, item := range val {
			L.SetField(table, k, valueToLua(L, item))
		}
		return table
	default:
		return lua.LNil
	}
}

// luaToValue 将Lua值转换为Go值
func luaToValue(L *lua.LState, v lua.LValue) interface{} {
	switch v.Type() {
	case lua.LTNil:
		return nil
	case lua.LTBool:
		return lua.LVAsBool(v)
	case lua.LTNumber:
		return float64(v.(lua.LNumber))
	case lua.LTString:
		return v.String()
	case lua.LTTable:
		table := v.(*lua.LTable)

		// 检查是否为数组
		maxn := table.MaxN()
		if maxn > 0 {
			// 处理为数组
			array := make([]interface{}, 0, maxn)
			for i := 1; i <= maxn; i++ {
				item := table.RawGetInt(i)
				if item != lua.LNil {
					array = append(array, luaToValue(L, item))
				}
			}
			return array
		}

		// 处理为对象
		obj := make(map[string]interface{})
		table.ForEach(func(key, value lua.LValue) {
			if key.Type() == lua.LTString {
				obj[key.String()] = luaToValue(L, value)
			}
		})
		return obj
	default:
		return nil
	}
}

// EncryptRequests 加密请求数据
func EncryptRequests(data string) (string, error) {
	config.GinLOG.Debug(fmt.Sprintf("加密请求数据: %s", data))
	rsaKey := RandomString(16)
	rsaIV := ReverseString(rsaKey)

	// RSA加密
	encryptedRSA, err := RsaEncryption(rsaKey)
	if err != nil {
		return "", err
	}

	// AES加密
	encryptedAES, err := AesEncryption(data, rsaKey, rsaIV)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s.%s", encryptedRSA, encryptedAES), nil
}
