package lingxing

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// 保留 doAndDecode 以便按需复用；通过显式引用避免被 lint 判定为 unused。
var _ = (*Client).doAndDecode

var (
	// ErrMissingAccessToken 表示调用需要鉴权的接口时，未配置 access_token。
	ErrMissingAccessToken = errors.New("缺少 access_token")
	// ErrMissingAppID 表示调用需要签名的接口时，未配置 app_key（AppID）。
	ErrMissingAppID = errors.New("缺少 app_key（AppID）")
	// ErrMissingAppSecret 表示调用获取/刷新 token 接口时，未配置 appSecret。
	ErrMissingAppSecret = errors.New("缺少 appSecret")
)

func (c *Client) doSignedJSON(ctx context.Context, method, path string, query any, body any, out any) error {
	_, err := c.doSignedJSONWithEnvelope(ctx, method, path, query, body, out)
	return err
}

func (c *Client) doSignedJSONWithEnvelope(ctx context.Context, method, path string, query any, body any, out any) (ResponseEnvelope, error) {
	env, _, err := c.doSignedJSONWithEnvelopeWithRawBody(ctx, method, path, query, body, out)
	return env, err
}

func (c *Client) doSignedJSONWithEnvelopeWithRawBody(ctx context.Context, method, path string, query any, body any, out any) (ResponseEnvelope, string, error) {
	if strings.TrimSpace(c.appID) == "" {
		return ResponseEnvelope{}, "", ErrMissingAppID
	}
	if c.autoToken {
		if err := c.ensureAccessToken(ctx); err != nil {
			return ResponseEnvelope{}, "", err
		}
	}
	if strings.TrimSpace(c.accessToken) == "" {
		return ResponseEnvelope{}, "", ErrMissingAccessToken
	}

	ts := c.now().Unix()
	common := url.Values{}
	common.Set("access_token", c.accessToken)
	common.Set("app_key", c.appID)
	common.Set("timestamp", strconv.FormatInt(ts, 10))

	signParams, err := c.buildSignParams(ts, query, body)
	if err != nil {
		return ResponseEnvelope{}, "", err
	}
	plain, err := buildSignPlain(signParams)
	if err != nil {
		return ResponseEnvelope{}, "", err
	}
	sign, err := signFromPlain(plain, c.appID)
	if err != nil {
		return ResponseEnvelope{}, "", err
	}
	common.Set("sign", sign)

	fullURL, err := c.buildURL(path, common)
	if err != nil {
		return ResponseEnvelope{}, "", err
	}

	var bodyReader io.Reader
	if body != nil {
		b, err := jsonMarshalNoEscape(body)
		if err != nil {
			return ResponseEnvelope{}, "", err
		}
		fmt.Printf("Request body: %s\n", string(b))
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	fmt.Printf("Request URL: %s\n", fullURL)
	fmt.Printf("Request Method: %s\n", method)
	if err != nil {
		return ResponseEnvelope{}, "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	return c.doAndDecodeWithRawBody(req, out)
}

func (c *Client) doSignedGET(ctx context.Context, path string, query any, out any) error {
	_, err := c.doSignedGETWithEnvelope(ctx, path, query, out)
	return err
}

func (c *Client) doSignedGETWithEnvelope(ctx context.Context, path string, query any, out any) (ResponseEnvelope, error) {
	env, _, err := c.doSignedGETWithEnvelopeWithRawBody(ctx, path, query, out)
	return env, err
}

func (c *Client) doSignedGETWithEnvelopeWithRawBody(ctx context.Context, path string, query any, out any) (ResponseEnvelope, string, error) {
	if strings.TrimSpace(c.appID) == "" {
		return ResponseEnvelope{}, "", ErrMissingAppID
	}
	if c.autoToken {
		if err := c.ensureAccessToken(ctx); err != nil {
			return ResponseEnvelope{}, "", err
		}
	}
	if strings.TrimSpace(c.accessToken) == "" {
		return ResponseEnvelope{}, "", ErrMissingAccessToken
	}

	ts := c.now().Unix()
	values := encodeQuery(query)
	values.Set("access_token", c.accessToken)
	values.Set("app_key", c.appID)
	values.Set("timestamp", strconv.FormatInt(ts, 10))

	signParams, err := c.buildSignParams(ts, query, nil)
	if err != nil {
		return ResponseEnvelope{}, "", err
	}
	plain, err := buildSignPlain(signParams)
	if err != nil {
		return ResponseEnvelope{}, "", err
	}
	sign, err := signFromPlain(plain, c.appID)
	if err != nil {
		return ResponseEnvelope{}, "", err
	}
	values.Set("sign", sign)

	fullURL, err := c.buildURL(path, values)
	if err != nil {
		return ResponseEnvelope{}, "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return ResponseEnvelope{}, "", err
	}
	req.Header.Set("Accept", "application/json")
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	return c.doAndDecodeWithRawBody(req, out)
}

func (c *Client) doMultipartForm(ctx context.Context, path string, form url.Values, out any) error {
	_, err := c.doMultipartFormWithEnvelope(ctx, path, form, out)
	return err
}

func (c *Client) doMultipartFormWithEnvelope(ctx context.Context, path string, form url.Values, out any) (ResponseEnvelope, error) {
	env, _, err := c.doMultipartFormWithEnvelopeWithRawBody(ctx, path, form, out)
	return env, err
}

func (c *Client) doMultipartFormWithEnvelopeWithRawBody(ctx context.Context, path string, form url.Values, out any) (ResponseEnvelope, string, error) {
	fullURL, err := c.buildURL(path, nil)
	if err != nil {
		return ResponseEnvelope{}, "", err
	}

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	for k, arr := range form {
		for _, v := range arr {
			if err := w.WriteField(k, v); err != nil {
				_ = w.Close()
				return ResponseEnvelope{}, "", err
			}
		}
	}
	if err := w.Close(); err != nil {
		return ResponseEnvelope{}, "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, &buf)
	if err != nil {
		return ResponseEnvelope{}, "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", w.FormDataContentType())
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	return c.doAndDecodeWithRawBody(req, out)
}

func (c *Client) doAndDecode(req *http.Request, out any) (ResponseEnvelope, error) {
	resp, err := c.httpCli.Do(req)
	if err != nil {
		return ResponseEnvelope{}, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseEnvelope{}, err
	}
	// fmt.Printf("Response body: %s\n", string(raw))

	if resp.StatusCode >= 400 {
		return ResponseEnvelope{}, &APIError{
			StatusCode: resp.StatusCode,
			Method:     req.Method,
			URL:        req.URL.String(),
			Body:       raw,
		}
	}

	var env ResponseEnvelope
	if err := json.Unmarshal(raw, &env); err != nil {
		// 返回体不是预期的“统一响应壳”时，直接把原始 body 带上，方便排查。
		return ResponseEnvelope{}, &APIError{
			StatusCode: resp.StatusCode,
			Method:     req.Method,
			URL:        req.URL.String(),
			Body:       raw,
		}
	}

	if !env.IsSuccess() {
		return env, &APIError{
			StatusCode: resp.StatusCode,
			Method:     req.Method,
			URL:        req.URL.String(),
			Code:       env.Code,
			Message:    env.Message(),
			RequestID:  env.RequestID(),
			Body:       raw,
		}
	}

	if out == nil {
		return env, nil
	}
	if len(env.Data) == 0 || string(env.Data) == "null" {
		return env, nil
	}
	if err := json.Unmarshal(env.Data, out); err != nil {
		return env, err
	}
	return env, nil
}

func (c *Client) doAndDecodeWithRawBody(req *http.Request, out any) (ResponseEnvelope, string, error) {
	resp, err := c.httpCli.Do(req)
	if err != nil {
		return ResponseEnvelope{}, "", err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseEnvelope{}, "", err
	}
	// fmt.Printf("Response body: %s\n", string(raw))
	rawStr := string(raw)

	if resp.StatusCode >= 400 {
		return ResponseEnvelope{}, rawStr, &APIError{
			StatusCode: resp.StatusCode,
			Method:     req.Method,
			URL:        req.URL.String(),
			Body:       raw,
		}
	}

	var env ResponseEnvelope
	if err := json.Unmarshal(raw, &env); err != nil {
		// 返回体不是预期的“统一响应壳”时，直接把原始 body 带上，方便排查。
		return ResponseEnvelope{}, rawStr, &APIError{
			StatusCode: resp.StatusCode,
			Method:     req.Method,
			URL:        req.URL.String(),
			Body:       raw,
		}
	}

	if !env.IsSuccess() {
		return env, rawStr, &APIError{
			StatusCode: resp.StatusCode,
			Method:     req.Method,
			URL:        req.URL.String(),
			Code:       env.Code,
			Message:    env.Message(),
			RequestID:  env.RequestID(),
			Body:       raw,
		}
	}

	if out == nil {
		return env, rawStr, nil
	}
	if len(env.Data) == 0 || string(env.Data) == "null" {
		return env, rawStr, nil
	}
	if err := json.Unmarshal(env.Data, out); err != nil {
		return env, rawStr, err
	}
	return env, rawStr, nil
}

func (c *Client) buildURL(apiPath string, values url.Values) (string, error) {
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
		u.Path = basePath + "/" + clean
	}

	if len(values) > 0 {
		u.RawQuery = values.Encode()
	}
	return u.String(), nil
}

func (c *Client) buildSignParams(ts int64, query any, body any) (map[string]any, error) {
	params := map[string]any{
		"access_token": c.accessToken,
		"app_key":      c.appID,
		"timestamp":    ts,
	}

	// query 参与签名（GET：业务参数在 query 上；POST：业务参数通常不在 URL，但也需要参与签名）。
	qv := encodeQuery(query)
	for k, arr := range qv {
		if len(arr) == 0 {
			continue
		}
		if len(arr) == 1 {
			params[k] = arr[0]
			continue
		}
		// 多值场景在文档中未明确，这里使用逗号拼接，至少保证稳定性。
		params[k] = strings.Join(arr, ",")
	}

	if body == nil {
		return params, nil
	}

	m, err := bodyToSignParams(body)
	if err != nil {
		return nil, err
	}
	for k, v := range m {
		params[k] = v
	}
	return params, nil
}

// bodyToSignParams 将 body 转为“参与签名”的参数集合。
//
// 关键点：
//   - 对于嵌套对象/数组：按文档要求“转成 string 参与签名”。
//   - 为了保证签名与实际请求 body 一致，这里基于“即将发送的 JSON”提取顶层字段的原始 JSON 文本，
//     避免二次反序列化再序列化导致字段顺序/格式变化（从而触发 api sign not correct）。
func bodyToSignParams(body any) (map[string]any, error) {
	switch v := body.(type) {
	case map[string]any:
		return v, nil
	case map[string]string:
		out := make(map[string]any, len(v))
		for k, val := range v {
			out[k] = val
		}
		return out, nil
	}

	b, err := jsonMarshalNoEscape(body)
	if err != nil {
		return nil, err
	}

	var rawMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &rawMap); err != nil {
		return nil, err
	}

	out := make(map[string]any, len(rawMap))
	for k, raw := range rawMap {
		rb := bytes.TrimSpace(raw)
		if len(rb) == 0 || bytes.Equal(rb, []byte("null")) {
			out[k] = nil
			continue
		}
		if rb[0] == '"' {
			var s string
			if err := json.Unmarshal(rb, &s); err != nil {
				return nil, err
			}
			if s == "" {
				// 空字符串不参与签名
				continue
			}
			out[k] = s
			continue
		}
		// number/bool/object/array：使用原始 JSON 文本作为签名入参（不再二次序列化）。
		out[k] = string(rb)
	}
	return out, nil
}

func encodeQuery(q any) url.Values {
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
