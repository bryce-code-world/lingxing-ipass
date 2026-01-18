package config

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// LoadDotEnv 加载 .env 文件到进程环境变量中。
//
// 约定：
// - 文件不存在时返回 nil（可选配置）。
// - 不覆盖进程中已存在的同名环境变量（避免线上配置被本地文件覆盖）。
// - 只支持最常见的 KEY=VALUE 形式；VALUE 允许使用单/双引号包裹。
func LoadDotEnv(path string) error {
	path = strings.TrimSpace(path)
	if path == "" {
		return errors.New("path 不能为空")
	}

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// 支持行尾注释：KEY=VAL # comment（简化实现：仅在不在引号内时处理 #）。
		line = stripTrailingComment(line)
		if line == "" {
			continue
		}

		key, val, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		if key == "" {
			continue
		}
		if os.Getenv(key) != "" {
			continue
		}
		val = unquote(val)
		_ = os.Setenv(key, val)
	}
	if err := sc.Err(); err != nil {
		return err
	}
	return nil
}

func unquote(v string) string {
	v = strings.TrimSpace(v)
	if len(v) < 2 {
		return v
	}
	if (v[0] == '"' && v[len(v)-1] == '"') || (v[0] == '\'' && v[len(v)-1] == '\'') {
		return v[1 : len(v)-1]
	}
	return v
}

func stripTrailingComment(line string) string {
	inSingle := false
	inDouble := false
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '\'':
			if !inDouble {
				inSingle = !inSingle
			}
		case '"':
			if !inSingle {
				inDouble = !inDouble
			}
		case '#':
			if !inSingle && !inDouble {
				return strings.TrimSpace(line[:i])
			}
		}
	}
	return strings.TrimSpace(line)
}
