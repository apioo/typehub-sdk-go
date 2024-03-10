// PullRequestTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/apioo/sdkgen-go"
	"io"
	"net/http"
	"net/url"
)

type PullRequestTag struct {
	internal *sdkgen.TagAbstract
}

// ReactComment Reacts to a comment
func (client *PullRequestTag) ReactComment(user string, document string, id string, comment string, reaction string, payload Passthru) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document
	pathParams["id"] = id
	pathParams["comment"] = comment
	pathParams["reaction"] = reaction

	queryParams := make(map[string]interface{})

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/pull_request/:id/comment/:comment/:reaction", pathParams))
	if err != nil {
		return Message{}, err
	}

	u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return Message{}, err
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("POST", u.String(), reqBody)
	if err != nil {
		return Message{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Message{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// CreateComment Creates a new pull request comment
func (client *PullRequestTag) CreateComment(user string, document string, id string, payload CommentCreate) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document
	pathParams["id"] = id

	queryParams := make(map[string]interface{})

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/pull_request/:id/comment", pathParams))
	if err != nil {
		return Message{}, err
	}

	u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return Message{}, err
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("POST", u.String(), reqBody)
	if err != nil {
		return Message{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Message{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// GetAllComments Shows all pull request comments
func (client *PullRequestTag) GetAllComments(user string, document string, id string, startIndex int, count int, search string) (CommentCollection, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document
	pathParams["id"] = id

	queryParams := make(map[string]interface{})
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/pull_request/:id/comment", pathParams))
	if err != nil {
		return CommentCollection{}, err
	}

	u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return CommentCollection{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return CommentCollection{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return CommentCollection{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response CommentCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommentCollection{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommentCollection{}, err
		}

		return CommentCollection{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommentCollection{}, err
		}

		return CommentCollection{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommentCollection{}, err
		}

		return CommentCollection{}, &MessageException{
			Payload: response,
		}
	default:
		return CommentCollection{}, errors.New("the server returned an unknown status code")
	}
}

// Delete Removes a pull request
func (client *PullRequestTag) Delete(user string, document string, id string) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document
	pathParams["id"] = id

	queryParams := make(map[string]interface{})

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/pull_request/:id", pathParams))
	if err != nil {
		return Message{}, err
	}

	u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return Message{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Message{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// Update Merges a pull request
func (client *PullRequestTag) Update(user string, document string, id string, payload Passthru) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document
	pathParams["id"] = id

	queryParams := make(map[string]interface{})

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/pull_request/:id", pathParams))
	if err != nil {
		return Message{}, err
	}

	u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return Message{}, err
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("PUT", u.String(), reqBody)
	if err != nil {
		return Message{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Message{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// Create Creates a new pull request
func (client *PullRequestTag) Create(user string, document string, payload PullRequestCreate) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document

	queryParams := make(map[string]interface{})

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/pull_request", pathParams))
	if err != nil {
		return Message{}, err
	}

	u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return Message{}, err
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("POST", u.String(), reqBody)
	if err != nil {
		return Message{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Message{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, err
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// Get Returns a pull request
func (client *PullRequestTag) Get(user string, document string, id string) (PullRequest, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document
	pathParams["id"] = id

	queryParams := make(map[string]interface{})

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/pull_request/:id", pathParams))
	if err != nil {
		return PullRequest{}, err
	}

	u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return PullRequest{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return PullRequest{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return PullRequest{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response PullRequest
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return PullRequest{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return PullRequest{}, err
		}

		return PullRequest{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return PullRequest{}, err
		}

		return PullRequest{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return PullRequest{}, err
		}

		return PullRequest{}, &MessageException{
			Payload: response,
		}
	default:
		return PullRequest{}, errors.New("the server returned an unknown status code")
	}
}

// GetAll Returns all pull requests
func (client *PullRequestTag) GetAll(user string, document string, status int, startIndex int, count int, search string) (PullRequestCollection, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document

	queryParams := make(map[string]interface{})
	queryParams["status"] = status
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/pull_request", pathParams))
	if err != nil {
		return PullRequestCollection{}, err
	}

	u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return PullRequestCollection{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return PullRequestCollection{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return PullRequestCollection{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response PullRequestCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return PullRequestCollection{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return PullRequestCollection{}, err
		}

		return PullRequestCollection{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return PullRequestCollection{}, err
		}

		return PullRequestCollection{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return PullRequestCollection{}, err
		}

		return PullRequestCollection{}, &MessageException{
			Payload: response,
		}
	default:
		return PullRequestCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewPullRequestTag(httpClient *http.Client, parser *sdkgen.Parser) *PullRequestTag {
	return &PullRequestTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}