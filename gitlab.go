package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func NewGitlabClient(url string) GitlabClient {
	return GitlabClient{
		URL:        url,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(ioutil.Discard, "", log.Lshortfile|log.Ldate|log.Ltime),
	}
}

type GitlabClient struct {
	URL string

	HTTPClient *http.Client
	Logger     *log.Logger
}

func (gitlab *GitlabClient) CILint(fileContents []byte) (CILintResponse, error) {
	var (
		err error

		urlStr = gitlab.URL + "/api/v4/ci/lint"

		lintRequest = CILintRequest{
			Content: string(fileContents),
		}

		requestBody []byte
		request     *http.Request

		responseBody []byte
		response     *http.Response

		ciLintResponse CILintResponse
	)

	gitlab.Logger.Println("gitlab url:", urlStr)

	if requestBody, err = json.Marshal(lintRequest); err != nil {
		return CILintResponse{}, err
	}

	request, err = http.NewRequest("POST", urlStr, bytes.NewBuffer(requestBody))

	if err != nil {
		return CILintResponse{}, err
	}

	request.Header.Set("Content-Type", "application/json")

	if response, err = gitlab.HTTPClient.Do(request); err != nil {
		return CILintResponse{}, err
	}

	defer response.Body.Close()

	if responseBody, err = ioutil.ReadAll(response.Body); err != nil {
		return CILintResponse{}, err
	}

	gitlab.Logger.Println("response body:", string(responseBody))

	if err = json.Unmarshal(responseBody, &ciLintResponse); err != nil {
		return CILintResponse{}, err
	}

	return ciLintResponse, nil
}

type CILintResponse struct {
	Status string   `status`
	Errors []string `errors`
}

type CILintRequest struct {
	Content string `json:"content"`
}
