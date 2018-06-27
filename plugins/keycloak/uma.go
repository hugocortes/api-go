package keycloak

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (client *keycloak) HandlePasswordGrant(username string, password string, scope string) (*TransformedUMAToken, error) {
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Add("username", username)
	data.Add("password", password)
	data.Add("client_id", client.clientName)
	data.Add("client_secret", client.clientSecret)

	if scope != "" {
		data.Add("scope", scope)
	}

	req, _ := http.NewRequest("POST", client.umaToken, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	decoder, err := doHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	values := &TransformedUMAToken{}
	err = decoder.Decode(values)

	if err != nil {
		return nil, err
	}

	return values, nil
}

func doHTTPRequest(req *http.Request) (*json.Decoder, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	status := resp.StatusCode
	if status >= 400 {
		err = fmt.Errorf("%v response", status)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	byteReader := bytes.NewReader(body)
	decoder := json.NewDecoder(byteReader)

	return decoder, nil
}
