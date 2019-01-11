package main

import (
	"net/http"
	"errors"
	"fmt"
)

func (s *Ticker) RPC(schedulerType string) (int, error) {

	req, errReq := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", s.APIGateWay , schedulerType), nil)
	if errReq != nil {
		return http.StatusBadGateway, errReq
	}

	resp, errResp := s.HTTPClient.Do(req)
	if errResp != nil {
		return http.StatusOK, errResp
	}

	if resp.StatusCode != 200 {
		return resp.StatusCode, errors.New(resp.Status)
	}

	return http.StatusOK, nil
}
