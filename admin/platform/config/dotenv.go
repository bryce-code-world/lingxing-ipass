package config

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// LoadDotEnv 从指定文件加载环境变量（不覆盖进程已存在的 env）。
// 约定：
// - 仅支持 KEY=VALUE
// - 允许以 # 开头的注释行
// - 不做复杂转义（一期保持简单直白）
func LoadDotEnv(path string) error {
	path = strings.TrimSpace(path)
	if path == "" {
		return errors.New("path 不能为空")
	}
	f, err := os.Open(path)
	if err != nil {
		// .env 文件不存在时不算错误，方便线上用系统环境变量注入
		return nil
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		k, v, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		k = strings.TrimSpace(k)
		v = strings.TrimSpace(v)
		if k == "" {
			continue
		}
		if _, exists := os.LookupEnv(k); exists {
			continue
		}
		v = strings.Trim(v, `"'`)
		_ = os.Setenv(k, v)
	}
	return sc.Err()
}

