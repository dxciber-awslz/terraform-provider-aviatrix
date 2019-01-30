package goaviatrix

import (
	"encoding/json"
	"errors"
	"log"
)

// Controller Http Access enabled get result struct
type ControllerHttpAccessResp struct {
	Return bool   `json:"return"`
	Result string `json:"result"`
}

func (c *Client) EnableHttpAccess() error {
    url := "?CID=%s&action=config_http_access&operation=enable"
	path := c.baseURL + fmt.Sprintf(url, c.CID)
	resp, err := c.Get(c.baseURL, nil)
	if err != nil {
		return err
	}
	var data APIResp
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}
	if !data.Return {
		return errors.New(data.Reason)
	}
	return nil
}

func (c *Client) DisableHttpAccess() error {
    url := "?CID=%s&action=config_http_access&operation=disable"
	path := c.baseURL + fmt.Sprintf(url, c.CID)
	resp, err := c.Get(c.baseURL, nil)
	if err != nil {
		return err
	}
	var data APIResp
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}
	if !data.Return {
		return errors.New(data.Reason)
	}
	return nil
}

func (c *Client) GetHttpAccessEnabled() (string,error) {
    url := "?CID=%s&action=config_http_access&operation=get"
	path := c.baseURL + fmt.Sprintf(url, c.CID)
	resp, err := c.Get(c.baseURL, nil)
	if err != nil {
		return err
	}
	var data ControllerHttpAccessResp
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}
	if !data.Return {
		return errors.New(data.Reason)
	}
	result := data.Result
	return result
}

