package adminweb

import "encoding/json"

func jsonValid(b []byte) bool {
	return json.Valid(b)
}
