package dsco

import (
	"context"
	"errors"
	"net/http"
)

func (s *ReturnService) CreateWithRawBody(ctx context.Context, req *ReturnCreateRequest) (*ReturnResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out ReturnResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPost, "/return/", nil, req, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *ReturnService) CompleteWithRawBody(ctx context.Context, req *ReturnCompleteRequest) (*ReturnResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out ReturnResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPut, "/return/", nil, req, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *ReturnService) GetChangeLogWithRawBody(ctx context.Context, q ReturnChangeLogQuery) (*ReturnChangeLogResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out ReturnChangeLogResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodGet, "/return/log", q, nil, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}
