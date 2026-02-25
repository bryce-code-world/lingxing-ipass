package lingxing

import (
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"strconv"
)

var (
	// ErrInvalidAESKeyLength 表示签名所用的 AES 密钥长度不合法（只能是 16/24/32 字节）。
	//
	// 领星文档写明“密钥为 appId”，但未说明当 appId 长度不满足 AES 要求时如何处理，
	// SDK 这里选择直接报错，避免“自作主张”导致签名与官方不一致。
	ErrInvalidAESKeyLength = errors.New("签名密钥长度不合法（仅支持16/24/32字节）")
)

func buildSignPlain(params map[string]any) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	first := true
	for _, k := range keys {
		val, ok, err := normalizeSignValue(params[k])
		if err != nil {
			return "", fmt.Errorf("参数 %s 规范化失败: %w", k, err)
		}
		if !ok {
			continue
		}
		if !first {
			buf.WriteByte('&')
		}
		first = false
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(val)
	}
	return buf.String(), nil
}

func normalizeSignValue(v any) (string, bool, error) {
	if v == nil {
		return "null", true, nil
	}

	switch x := v.(type) {
	case string:
		if x == "" {
			return "", false, nil
		}
		return x, true, nil
	case []byte:
		if len(x) == 0 {
			return "", false, nil
		}
		return string(x), true, nil
	case bool:
		return strconv.FormatBool(x), true, nil
	case int:
		return strconv.FormatInt(int64(x), 10), true, nil
	case int8:
		return strconv.FormatInt(int64(x), 10), true, nil
	case int16:
		return strconv.FormatInt(int64(x), 10), true, nil
	case int32:
		return strconv.FormatInt(int64(x), 10), true, nil
	case int64:
		return strconv.FormatInt(x, 10), true, nil
	case uint:
		return strconv.FormatUint(uint64(x), 10), true, nil
	case uint8:
		return strconv.FormatUint(uint64(x), 10), true, nil
	case uint16:
		return strconv.FormatUint(uint64(x), 10), true, nil
	case uint32:
		return strconv.FormatUint(uint64(x), 10), true, nil
	case uint64:
		return strconv.FormatUint(x, 10), true, nil
	case float32:
		return strconv.FormatFloat(float64(x), 'f', -1, 32), true, nil
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64), true, nil
	}

	// 对于集合/对象类型，按文档建议参与签名时转成 string。
	// 这里直接使用 JSON 字符串（稳定、可读、跨语言容易复现）。
	b, err := jsonMarshalNoEscape(v)
	if err != nil {
		return "", false, err
	}
	// JSON 结果为 "" 时，按“空值不参与”处理。
	if len(b) == 0 || string(b) == `""` {
		return "", false, nil
	}
	return string(b), true, nil
}

func signFromPlain(plain string, appID string) (string, error) {
	key := []byte(appID)
	if l := len(key); l != 16 && l != 24 && l != 32 {
		return "", ErrInvalidAESKeyLength
	}

	// c) MD5(32位)并转大写。
	md5Sum := md5.Sum([]byte(plain))
	hexed := hex.EncodeToString(md5Sum[:])
	upperMD5 := toUpperASCII(hexed)

	// d) AES/ECB/PKCS5PADDING 加密（密钥为 appId）。
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	padded := pkcs5Pad([]byte(upperMD5), bs)

	out := make([]byte, len(padded))
	for off := 0; off < len(padded); off += bs {
		block.Encrypt(out[off:off+bs], padded[off:off+bs])
	}
	return base64.StdEncoding.EncodeToString(out), nil
}

func pkcs5Pad(in []byte, blockSize int) []byte {
	pad := blockSize - (len(in) % blockSize)
	out := make([]byte, 0, len(in)+pad)
	out = append(out, in...)
	for i := 0; i < pad; i++ {
		out = append(out, byte(pad))
	}
	return out
}

func toUpperASCII(s string) string {
	b := []byte(s)
	for i := range b {
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = b[i] - 'a' + 'A'
		}
	}
	return string(b)
}
