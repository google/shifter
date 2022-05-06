package openshift

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	// Bearer Token Authentication
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AuthOptions.BearerToken))
	return req, nil
}
func (c *Client) do(req *http.Request, target interface{}) (*http.Response, error) {

	resp, err := c.httpClient.Do(req)
	if err != nil {
		// Handle Errors Here..
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Non-OK HTTP status:", resp.StatusCode)
		fmt.Println(resp)
		// You may read / inspect response body
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	return resp, err
}
