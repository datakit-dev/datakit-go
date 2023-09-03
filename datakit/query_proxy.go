package datakit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const QUERY_PROXY_PATH = "/queryProxy"

type QueryProxyClient struct {
	Client *http.Client
	apiUrl string
}

func NewQueryProxyClient(apiUrl string, apiToken string) *QueryProxyClient {
	client := &http.Client{
		Transport: NewTransport(apiToken),
	}
	return &QueryProxyClient{
		Client: client,
		apiUrl: apiUrl,
	}
}

type QueryJSON struct {
	Select  []string `json:"select"`
	From    []string `json:"from"`
	Where   string   `json:"where,omitempty"`
	GroupBy string   `json:"groupby,omitempty"`
	Having  string   `json:"having,omitempty"`
	Qualify string   `json:"qualify,omitempty"`
	Window  string   `json:"window,omitempty"`
}

type QueryProxyRequestBody struct {
	ClientType  string     `json:"clientType"`
	WorkspaceId string     `json:"workspaceId"`
	SQL         string     `json:"sql,omitempty"`
	JSON        *QueryJSON `json:"json,omitempty"`
}

type Row map[string]interface{}

type QueryProxyResponseBody struct {
	Meta struct {
		JobId     string `json:"jobId"`
		NextQuery string `json:"nextQuery"`
	}
	Rows []Row `json:"rows"`
}

type QueryProxyResponseError struct {
	Message string `json:"message"`
}

func (b *QueryProxyResponseBody) Columns() []string {
	var columns []string
	for k := range b.Rows[0] {
		columns = append(columns, k)
	}
	return columns
}

func (qp *QueryProxyClient) query(body *QueryProxyRequestBody) (resp *QueryProxyResponseBody, err error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(reqBody)
	res, err := qp.Client.Post(qp.apiUrl+QUERY_PROXY_PATH, "application/json", reader)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var errBody QueryProxyResponseError
		if err := json.NewDecoder(res.Body).Decode(&errBody); err != nil {
			return nil, fmt.Errorf("query failed (%d): %s", res.StatusCode, "unknown error")
		}
		return nil, fmt.Errorf("query failed (%d): %s", res.StatusCode, errBody.Message)
	}

	err = json.NewDecoder(res.Body).Decode(&resp)
	return resp, err
}

func (qp *QueryProxyClient) QuerySQL(workspaceId, sql string) (resp *QueryProxyResponseBody, err error) {
	body := &QueryProxyRequestBody{
		ClientType:  "datakit:cli",
		WorkspaceId: workspaceId,
		SQL:         sql,
	}
	return qp.query(body)
}

func (qp *QueryProxyClient) QueryJSON(workspaceId, obj string) (resp *QueryProxyResponseBody, err error) {
	var jsonObj QueryJSON
	if err := json.Unmarshal([]byte(obj), &jsonObj); err != nil {
		return nil, fmt.Errorf("invalid JSON query: %s", err)
	}

	body := &QueryProxyRequestBody{
		ClientType:  "datakit:cli",
		WorkspaceId: workspaceId,
		JSON:        &jsonObj,
	}
	return qp.query(body)
}
