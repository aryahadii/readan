package simplifier

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

type MercurySimplifier struct {
	parserURL string
	APIKey    string
}

var (
	DefaultMercurySimplifier = &MercurySimplifier{
		parserURL: "https://mercury.postlight.com/parser?url=%s",
	}
)

func (mc *MercurySimplifier) SimplifyHTML(url *url.URL) (string, error) {
	mercuryRes, err := mc.requestMercury(url)
	if err != nil {
		return "", err
	}
	return mercuryRes.Content, nil
}

func (mc *MercurySimplifier) requestMercury(url *url.URL) (*mercuryResponse, error) {
	req, err := mc.newMercuryRequest(url)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	mercuryRes, err := mc.readMercuryResponse(res)
	if err != nil {
		return nil, err
	}
	return mercuryRes, nil
}

func (mc *MercurySimplifier) newMercuryRequest(url *url.URL) (*http.Request, error) {
	mercuryURL := fmt.Sprintf(mc.parserURL, url.String())
	req, err := http.NewRequest("GET", mercuryURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-api-key", mc.APIKey)
	return req, nil
}

type mercuryResponse struct {
	Title        string `json:"title"`
	Author       string `json:"author"`
	Content      string `json:"content"`
	Direction    string `json:"direction"`
	LeadImageURL string `json:"lead_image_url"`
	WordCount    int    `json:"word_count"`
}

func (mc *MercurySimplifier) readMercuryResponse(res *http.Response) (*mercuryResponse, error) {
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, errors.Wrap(err, "cannot read Mercury's response")
	}
	logrus.Debugf("Mercury's response: %v", string(body))

	mercuryResp := &mercuryResponse{}
	err = json.Unmarshal(body, mercuryResp)
	if err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal Mercury's response")
	}
	return mercuryResp, nil
}
