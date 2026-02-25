package dsco

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

var (
	// ErrMissingToken 表示调用需要鉴权的接口时，客户端未配置 token。
	ErrMissingToken = errors.New("缺少访问令牌（bearer token）")
)

func (c *Client) doJSON(ctx context.Context, method, path string, query interface{}, in interface{}, out interface{}) error {
	return c.do(ctx, requestSpec{
		Method:      method,
		Path:        path,
		Query:       query,
		NeedAuth:    true,
		ContentType: "application/json",
		Accept:      "application/json",
		Body:        in,
	}, out)
}

func (c *Client) doForm(ctx context.Context, method, path string, values url.Values, out interface{}) error {
	return c.do(ctx, requestSpec{
		Method:      method,
		Path:        path,
		NeedAuth:    false,
		ContentType: "application/x-www-form-urlencoded",
		Accept:      "application/json",
		RawBody:     strings.NewReader(values.Encode()),
	}, out)
}

type requestSpec struct {
	Method string
	Path   string

	Query interface{}

	NeedAuth bool

	ContentType string
	Accept      string

	Body    interface{} // JSON body
	RawBody io.Reader   // 非 JSON body（例如 form）
}

func (c *Client) do(ctx context.Context, spec requestSpec, out interface{}) error {
	fullURL, err := c.buildURL(spec.Path, spec.Query)
	if err != nil {
		return err
	}

	var bodyReader io.Reader
	if spec.RawBody != nil {
		bodyReader = spec.RawBody
	} else if spec.Body != nil {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(spec.Body); err != nil {
			return err
		}
		bodyReader = &buf
	}

	req, err := http.NewRequestWithContext(ctx, spec.Method, fullURL, bodyReader)
	if err != nil {
		return err
	}

	if spec.Accept != "" {
		req.Header.Set("Accept", spec.Accept)
	}
	if spec.ContentType != "" {
		req.Header.Set("Content-Type", spec.ContentType)
	}
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	if spec.NeedAuth {
		if c.token == "" {
			return ErrMissingToken
		}
		req.Header.Set("Authorization", "bearer "+c.token)
	}

	resp, err := c.httpCli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// fmt.Println(string(raw))

	if resp.StatusCode >= 400 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Method:     spec.Method,
			URL:        fullURL,
			Body:       raw,
		}
	}

	if out == nil {
		return nil
	}
	if len(raw) == 0 {
		return nil
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return err
	}
	return nil
}

func (c *Client) buildURL(apiPath string, query interface{}) (string, error) {
	if c.baseURL == nil {
		return "", errors.New("BaseURL 未初始化")
	}

	clean := strings.TrimSpace(apiPath)
	clean = strings.TrimPrefix(clean, "/")

	u := *c.baseURL
	basePath := strings.TrimSuffix(u.Path, "/")
	if clean == "" {
		u.Path = basePath
	} else {
		// 说明：
		// - 部分 API 会把 “position/eventId” 等参数放在 path 中，该值可能包含已编码的 "%2F"。
		// - 直接把带 '%' 的字符串赋值给 u.Path，然后调用 u.String() 会触发二次编码（"%2F" -> "%252F"），导致服务端认为参数非法。
		// - 这里通过同时设置 Path（解码后）与 RawPath（编码后）来避免二次编码。
		rawPath := basePath + "/" + clean
		if strings.Contains(rawPath, "%") {
			if decoded, err := url.PathUnescape(rawPath); err == nil {
				u.Path = decoded
				u.RawPath = rawPath
			} else {
				u.Path = rawPath
			}
		} else {
			u.Path = rawPath
		}
	}

	values := encodeQuery(query)
	if len(values) > 0 {
		u.RawQuery = values.Encode()
	}
	return u.String(), nil
}

func encodeQuery(q interface{}) url.Values {
	values := url.Values{}
	if q == nil {
		return values
	}

	switch v := q.(type) {
	case url.Values:
		return v
	case map[string]string:
		for k, val := range v {
			if val == "" {
				continue
			}
			values.Add(k, val)
		}
		return values
	case map[string][]string:
		for k, arr := range v {
			for _, val := range arr {
				if val == "" {
					continue
				}
				values.Add(k, val)
			}
		}
		return values
	}

	rv := reflect.ValueOf(q)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return values
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return values
	}

	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)
		ft := rt.Field(i)

		tag := ft.Tag.Get("url")
		if tag == "" {
			continue
		}
		name, omitEmpty := splitQueryTag(tag)

		if fv.Kind() == reflect.Ptr {
			if fv.IsNil() {
				continue
			}
			fv = fv.Elem()
		}

		if omitEmpty && isZeroValue(fv) {
			continue
		}

		switch fv.Kind() {
		case reflect.String:
			if fv.String() == "" && omitEmpty {
				continue
			}
			values.Add(name, fv.String())
		case reflect.Bool:
			values.Add(name, strconv.FormatBool(fv.Bool()))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			values.Add(name, strconv.FormatInt(fv.Int(), 10))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			values.Add(name, strconv.FormatUint(fv.Uint(), 10))
		case reflect.Slice, reflect.Array:
			for j := 0; j < fv.Len(); j++ {
				ev := fv.Index(j)
				if ev.Kind() == reflect.String {
					if ev.String() == "" && omitEmpty {
						continue
					}
					values.Add(name, ev.String())
				}
			}
		}
	}
	return values
}

func splitQueryTag(tag string) (name string, omitEmpty bool) {
	parts := strings.Split(tag, ",")
	name = strings.TrimSpace(parts[0])
	for _, p := range parts[1:] {
		if strings.TrimSpace(p) == "omitempty" {
			omitEmpty = true
			break
		}
	}
	return name, omitEmpty
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Slice, reflect.Array, reflect.Map:
		return v.Len() == 0
	case reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}

func (c *Client) doJSONWithRawBody(ctx context.Context, method, path string, query interface{}, in interface{}, out interface{}) (string, error) {
	return c.doWithRawBody(ctx, requestSpec{
		Method:      method,
		Path:        path,
		Query:       query,
		NeedAuth:    true,
		ContentType: "application/json",
		Accept:      "application/json",
		Body:        in,
	}, out)
}

func (c *Client) doFormWithRawBody(ctx context.Context, method, path string, values url.Values, out interface{}) (string, error) {
	return c.doWithRawBody(ctx, requestSpec{
		Method:      method,
		Path:        path,
		NeedAuth:    false,
		ContentType: "application/x-www-form-urlencoded",
		Accept:      "application/json",
		RawBody:     strings.NewReader(values.Encode()),
	}, out)
}

func (c *Client) doWithRawBody(ctx context.Context, spec requestSpec, out interface{}) (string, error) {
	fullURL, err := c.buildURL(spec.Path, spec.Query)
	if err != nil {
		return "", err
	}

	var bodyReader io.Reader
	if spec.RawBody != nil {
		bodyReader = spec.RawBody
	} else if spec.Body != nil {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(spec.Body); err != nil {
			return "", err
		}
		bodyReader = &buf
	}

	req, err := http.NewRequestWithContext(ctx, spec.Method, fullURL, bodyReader)
	if err != nil {
		return "", err
	}

	if spec.Accept != "" {
		req.Header.Set("Accept", spec.Accept)
	}
	if spec.ContentType != "" {
		req.Header.Set("Content-Type", spec.ContentType)
	}
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	if spec.NeedAuth {
		if c.token == "" {
			return "", ErrMissingToken
		}
		req.Header.Set("Authorization", "bearer "+c.token)
	}

	resp, err := c.httpCli.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	rawStr := string(raw)

	if resp.StatusCode >= 400 {
		return rawStr, &APIError{
			StatusCode: resp.StatusCode,
			Method:     spec.Method,
			URL:        fullURL,
			Body:       raw,
		}
	}

	if out == nil {
		return rawStr, nil
	}
	if len(raw) == 0 {
		return rawStr, nil
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return rawStr, err
	}
	return rawStr, nil
}
