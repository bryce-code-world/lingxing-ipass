package dsco

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (s *StreamService) CreateStreamWithRawBody(ctx context.Context, stream Stream) (*Stream, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}
	var out Stream
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPost, "/stream", nil, stream, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *StreamService) DeleteStreamWithRawBody(ctx context.Context, id string) (*Stream, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}
	id = strings.TrimSpace(id)
	if id == "" {
		return nil, "", errors.New("stream id 不能为空")
	}
	var out Stream
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodDelete, "/stream/"+escapePath(id), nil, nil, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *StreamService) CreateStreamOperationWithRawBody(ctx context.Context, id string, op any) (*StreamOperationResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}
	id = strings.TrimSpace(id)
	if id == "" {
		return nil, "", errors.New("stream id 不能为空")
	}
	var out StreamOperationResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPost, "/stream/"+escapePath(id), nil, op, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *StreamService) GetStreamEventsFromPositionWithRawBody(ctx context.Context, id string, partitionID int, position string) (*StreamEventWrapper, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}
	id = strings.TrimSpace(id)
	position = strings.TrimSpace(position)
	if id == "" {
		return nil, "", errors.New("stream id 不能为空")
	}
	if partitionID < 0 {
		return nil, "", errors.New("partitionId 不允许为负数")
	}
	if position == "" {
		return nil, "", errors.New("position 不能为空")
	}

	var out StreamEventWrapper
	path := fmt.Sprintf("/stream/%s/%d/%s", escapePath(id), partitionID, escapeStreamPosition(position))
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodGet, path, nil, nil, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *StreamService) GetStreamEventsInRangeWithRawBody(ctx context.Context, id string, partitionID int, startPosition string, endPosition string) (*StreamEventWrapper, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}
	id = strings.TrimSpace(id)
	startPosition = strings.TrimSpace(startPosition)
	endPosition = strings.TrimSpace(endPosition)
	if id == "" {
		return nil, "", errors.New("stream id 不能为空")
	}
	if partitionID < 0 {
		return nil, "", errors.New("partitionId 不允许为负数")
	}
	if startPosition == "" || endPosition == "" {
		return nil, "", errors.New("start/end position 不能为空")
	}

	var out StreamEventWrapper
	path := fmt.Sprintf("/stream/%s/%d/%s/%s", escapePath(id), partitionID, escapeStreamPosition(startPosition), escapeStreamPosition(endPosition))
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodGet, path, nil, nil, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *StreamService) UpdateStreamPositionWithRawBody(ctx context.Context, id string, partitionID int, position string) (*SuccessFailResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}
	id = strings.TrimSpace(id)
	position = strings.TrimSpace(position)
	if id == "" {
		return nil, "", errors.New("stream id 不能为空")
	}
	if partitionID < 0 {
		return nil, "", errors.New("partitionId 不允许为负数")
	}
	if position == "" {
		return nil, "", errors.New("position 不能为空")
	}

	var out SuccessFailResponse
	path := fmt.Sprintf("/stream/%s/%d/%s", escapePath(id), partitionID, escapeStreamPosition(position))
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPut, path, nil, nil, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}
