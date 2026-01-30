package store

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

// PGTextArray 用于兼容 PostgreSQL text[] 字段的读写。
//
// 说明：
//   - 当前项目使用的 PostgreSQL driver 在读取 text[] 时，可能返回 string（例如：{"a","b"}），
//     GORM 无法直接 Scan 到 []string，因此需要自定义 Scanner。
//   - 该类型同时实现 driver.Valuer，便于未来若改为走 GORM 写入也能正常落库。
type PGTextArray []string

func (a *PGTextArray) Scan(src any) error {
	switch v := src.(type) {
	case nil:
		*a = nil
		return nil
	case []byte:
		items, err := parsePGTextArrayLiteral(string(v))
		if err != nil {
			return err
		}
		*a = PGTextArray(items)
		return nil
	case string:
		items, err := parsePGTextArrayLiteral(v)
		if err != nil {
			return err
		}
		*a = PGTextArray(items)
		return nil
	case []string:
		*a = PGTextArray(v)
		return nil
	default:
		return fmt.Errorf("pg text[] scan: unsupported src type %T", src)
	}
}

func (a PGTextArray) Value() (driver.Value, error) {
	return toPGTextArrayLiteral([]string(a)), nil
}

func parsePGTextArrayLiteral(s string) ([]string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}
	if len(s) < 2 || s[0] != '{' || s[len(s)-1] != '}' {
		return nil, fmt.Errorf("pg text[] parse: invalid array literal %q", s)
	}

	inner := s[1 : len(s)-1]
	if strings.TrimSpace(inner) == "" {
		return []string{}, nil
	}

	var items []string
	i := 0
	for i < len(inner) {
		// skip separators / whitespace
		for i < len(inner) && (inner[i] == ',' || inner[i] == ' ' || inner[i] == '\t' || inner[i] == '\n' || inner[i] == '\r') {
			i++
		}
		if i >= len(inner) {
			break
		}

		if inner[i] == '"' {
			// quoted element
			i++
			var sb strings.Builder
			for i < len(inner) {
				c := inner[i]
				if c == '\\' {
					if i+1 >= len(inner) {
						return nil, fmt.Errorf("pg text[] parse: invalid escape in %q", s)
					}
					sb.WriteByte(inner[i+1])
					i += 2
					continue
				}
				if c == '"' {
					i++
					break
				}
				sb.WriteByte(c)
				i++
			}
			items = append(items, sb.String())
			continue
		}

		// unquoted element
		start := i
		for i < len(inner) && inner[i] != ',' {
			i++
		}
		token := strings.TrimSpace(inner[start:i])
		if token == "NULL" {
			items = append(items, "")
		} else {
			items = append(items, unescapeBackslash(token))
		}
	}

	return items, nil
}

func unescapeBackslash(s string) string {
	if !strings.Contains(s, `\`) {
		return s
	}
	var sb strings.Builder
	sb.Grow(len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && i+1 < len(s) {
			sb.WriteByte(s[i+1])
			i++
			continue
		}
		sb.WriteByte(s[i])
	}
	return sb.String()
}
