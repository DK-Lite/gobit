package bybit

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	BaseURL = "https://api.bybit.com"
)

type APIWrapper struct {
	AccessKey        string
	SecretKey        string
	ServerTimeOffset int64
	DebugMode        bool
	Https            *http.Client
}

func NewWrapper(accessKey string, secretKey string) *APIWrapper {
	return &APIWrapper{
		AccessKey: accessKey,
		SecretKey: secretKey,
		Https: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (b *APIWrapper) signature(p HttpsParam) HttpsParam {
	p.Add("api_key", b.AccessKey)
	p.Add("timestamp", GetNowUTC())
	p.Add("sign", b.getSigned(p.Encode()))

	return p
}

func (b *APIWrapper) getSigned(param string) string {
	sig := hmac.New(sha256.New, []byte(b.SecretKey))
	sig.Write([]byte(param))
	signature := hex.EncodeToString(sig.Sum(nil))
	return signature
}

/* ======================== *
 * 		API					*
 * ========================	*/

func (b *APIWrapper) PostRequest(path string, param interface{}, v interface{}) error {
	p := NewHttpsParam()
	p.SetData(param)

	signedParam := b.signature(p)
	return b.Request(http.MethodPost, path, signedParam, &v)
}

func (b *APIWrapper) PrivateRequest(path string, param interface{}, v interface{}) error {
	p := NewHttpsParam()
	p.SetParam(param)

	signedParam := b.signature(p)
	return b.Request(http.MethodGet, path, signedParam, &v)
}

func (b *APIWrapper) PublicRequest(path string, param interface{}, v interface{}) error {
	p := NewHttpsParam()
	p.SetParam(param)

	return b.Request(http.MethodGet, path, p, &v)
}

func (b *APIWrapper) Request(method string, path string, param HttpsParam, v interface{}) error {
	param.SetUrl(BaseURL)
	param.SetPath(path)

	req, err := http.NewRequest(method, param.URL(), param.Body())
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	return doRequest(b.Https, req, &v)
}

func doRequest(client *http.Client, req *http.Request, v interface{}) error {
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("Client Request Error")
	}
	defer resp.Body.Close()

	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !statusOK {
		return errors.New("Status not OK")
	}

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return errors.New(fmt.Sprintf("Json Decoder Error: %s", err))
	}

	return nil
}
