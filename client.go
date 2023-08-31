package goperset

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/opentracing/opentracing-go"
	"io"
	"net/http"
	"strings"
)

const (
	urlSeparator string = "/"
)

func makeURL(path string) string {
	if strings.HasSuffix(path, urlSeparator) {
		return strings.TrimSuffix(path, urlSeparator)
	} else {
		return path
	}
}

func NewContext() context.Context {
	return context.WithValue(context.Background(), tracerContextKey, opentracing.GlobalTracer())
}

func NewClient(basePath string) *Goperset {
	return &Goperset{
		BasePath: basePath,
		Client:   &http.Client{},
	}
}

func getCsrftoken(ctx context.Context, client *Goperset, token string) (string, error) {
	body, err := ClientResty(ctx, client, token, "", "application/json", "GET", csrfEndpoint, nil)
	if err != nil {
		return "", err
	}
	var csrfResp CsrfResponse
	if err = json.Unmarshal(body, &csrfResp); err != nil {
		return "", err
	}

	return csrfResp.AccessToken, nil
}

func GetAccessTokens(ctx context.Context, client *Goperset, username string, password string) (string, string, error) {
	payload := LoginPayload{
		Password: password,
		Provider: "db",
		Refresh:  true,
		Username: username,
	}
	body, err := ClientResty(ctx, client, "", "", "application/json", "POST", loginEndpoint, payload)
	if err != nil {
		return "", "", err
	}
	bodyReader := bytes.NewReader(body)
	ioBody := io.NopCloser(bodyReader)
	var loginResp LoginResponse
	if err = json.NewDecoder(ioBody).Decode(&loginResp); err != nil {
		return "", "", err
	}
	csrfToken, err := getCsrftoken(ctx, client, loginResp.AccessToken)
	if err != nil {
		return "", "", err
	}
	return loginResp.AccessToken, csrfToken, nil
}

func ClientResty(ctx context.Context, client *Goperset, token string, csrfToken string, contentType string, method string, endpoint string, payload interface{}) ([]byte, error) {
	basePath := makeURL(client.BasePath)
	var payloadBuffer io.Reader
	// Check if payload is of type *bytes.Buffer
	if buffer, _type := payload.(*bytes.Buffer); _type {
		payloadBuffer = buffer
	} else {
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		payloadBuffer = bytes.NewBuffer(payloadBytes)
	}

	req, err := http.NewRequest(method, basePath+endpoint, payloadBuffer)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", contentType)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	if csrfToken != "" {
		req.Header.Add("X-CSRFToken", csrfToken)
	}

	resp, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
