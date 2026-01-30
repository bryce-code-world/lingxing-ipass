package boarding

import (
	"context"
	"testing"
)

func TestGetToken(t *testing.T) {
	cli := newClient(t)
	token, res, err := cli.Authorization.GetTokenWithRawBody(context.Background())
	if err != nil {
		t.Fatalf("GetToken() err=%v", err)
	}
	t.Logf("GetToken() token=%+v", token)
	t.Logf("GetToken() raw_response=%+v", res)
}
