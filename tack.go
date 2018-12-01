package main

import (
	"net/http"
	"bytes"
	"github.com/kataras/iris/core/errors"
)

func (s *Ticker) RPC(schedulerType string) (int, error) {

	req, errReq := http.NewRequest(http.MethodPost, s.APIGateWay, bytes.NewReader([]byte(schedulerType)))
	if errReq != nil {
		return http.StatusBadGateway, errReq
	}

	resp, errResp := s.HTTPClient.Do(req)
	if errResp != nil {
		return http.StatusOK, errResp
	}

	if resp.StatusCode != 200 {
		return resp.StatusCode, errors.New("")
	}

	return http.StatusOK, nil
}
