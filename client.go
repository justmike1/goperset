package goperset

import (
	"bytes"
	"context"
	"encoding/json"
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

func NewClient(ctx context.Context, basePath string) *Goperset {
	return &Goperset{
		BasePath: basePath,
		Client:   &http.Client{},
		Context:  ctx,
	}
}

func CreateDatabase(client *Goperset, tokens ClientToken, payload DatabasePayload) ([]byte, error) {
	body, err := ClientResty(client, tokens, "application/json", "POST", DatabaseController, payload)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func CreateDashboard(client *Goperset, tokens ClientToken, payload DashboardPayload) ([]byte, error) {
	body, err := ClientResty(client, tokens, "application/json", "POST", DashboardController, payload)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetDashboard(client *Goperset, tokens ClientToken, params GetDashboardParams) ([]byte, error) {
	body, err := ClientResty(client, tokens, "application/json", "GET", DashboardController, params)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetDashboardInfo(client *Goperset, tokens ClientToken, params DashboardInfoParams) ([]byte, error) {
	body, err := ClientResty(client, tokens, "application/json", "GET", DashboardInfoEndpoint, params)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func CreateDataset(client *Goperset, tokens ClientToken, payload DatasetPayload) ([]byte, error) {
	body, err := ClientResty(client, tokens, "application/json", "POST", DatasetController, payload)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func CreateChart(client *Goperset, tokens ClientToken, payload DatasetPayload) ([]byte, error) {
	body, err := ClientResty(client, tokens, "application/json", "POST", ChartController, payload)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func getCsrftoken(client *Goperset, tokens ClientToken) (string, error) {
	body, err := ClientResty(client, tokens, "application/json", "GET", csrfEndpoint, nil)
	if err != nil {
		return "", err
	}
	var csrfResp CsrfResponse
	if err = json.Unmarshal(body, &csrfResp); err != nil {
		return "", err
	}

	return csrfResp.AccessToken, nil
}

func GetAccessTokens(client *Goperset, username string, password string) (ClientToken, error) {
	var t ClientToken
	t = ClientToken{
		AccessToken: "",
		CsrfToken:   "",
	}
	payload := LoginPayload{
		Password: password,
		Provider: "db",
		Refresh:  true,
		Username: username,
	}
	body, err := ClientResty(client, t, "application/json", "POST", loginEndpoint, payload)
	if err != nil {
		return t, err
	}
	bodyReader := bytes.NewReader(body)
	ioBody := io.NopCloser(bodyReader)
	var loginResp LoginResponse
	if err = json.NewDecoder(ioBody).Decode(&loginResp); err != nil {
		return t, err
	}
	t = ClientToken{
		AccessToken: loginResp.AccessToken,
		CsrfToken:   "",
	}
	csrfToken, err := getCsrftoken(client, t)
	if err != nil {
		return t, err
	}
	t = ClientToken{
		AccessToken: loginResp.AccessToken,
		CsrfToken:   csrfToken,
	}
	return t, nil
}

func ClientResty(client *Goperset, tokens ClientToken, contentType string, method string, endpoint string, payload interface{}) ([]byte, error) {
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

	req = req.WithContext(client.Context)
	req.Header.Set("Content-Type", contentType)
	if tokens.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+tokens.AccessToken)
	}
	if tokens.CsrfToken != "" {
		req.Header.Add("X-CSRFToken", tokens.CsrfToken)
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
