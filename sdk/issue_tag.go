// IssueTag automatically generated by SDKgen please do not edit this file manually
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

type IssueTag struct {
	internal *sdkgen.TagAbstract
}

// ReactComment Reacts to a comment
func (client *IssueTag) ReactComment(user string, document string, id string, comment string, reaction string, payload Passthru) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document
	pathParams["id"] = id
	pathParams["comment"] = comment
	pathParams["reaction"] = reaction

	queryParams := make(map[string]interface{})

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/issue/:id/comment/:comment/:reaction", pathParams))
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

// CreateComment Creates a new issue comment
func (client *IssueTag) CreateComment(user string, document string, id string, payload CommentCreate) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document
	pathParams["id"] = id

	queryParams := make(map[string]interface{})

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/issue/:id/comment", pathParams))
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

// GetAllComments Shows all issue comments
func (client *IssueTag) GetAllComments(user string, document string, id string, startIndex int, count int, search string) (CommentCollection, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document
	pathParams["id"] = id

	queryParams := make(map[string]interface{})
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/issue/:id/comment", pathParams))
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

// Delete Removes an issue
func (client *IssueTag) Delete(user string, document string, id string) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document
	pathParams["id"] = id

	queryParams := make(map[string]interface{})

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/issue/:id", pathParams))
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

// Create Creates a new issue
func (client *IssueTag) Create(user string, document string, payload IssueCreate) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document

	queryParams := make(map[string]interface{})

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/issue", pathParams))
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

// Get Returns an issue
func (client *IssueTag) Get(user string, document string, id string) (Issue, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document
	pathParams["id"] = id

	queryParams := make(map[string]interface{})

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/issue/:id", pathParams))
	if err != nil {
		return Issue{}, err
	}

	u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Issue{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Issue{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Issue{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Issue
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Issue{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Issue{}, err
		}

		return Issue{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Issue{}, err
		}

		return Issue{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Issue{}, err
		}

		return Issue{}, &MessageException{
			Payload: response,
		}
	default:
		return Issue{}, errors.New("the server returned an unknown status code")
	}
}

// GetAll Returns all issues
func (client *IssueTag) GetAll(user string, document string, status int, startIndex int, count int, search string) (IssueCollection, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document

	queryParams := make(map[string]interface{})
	queryParams["status"] = status
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/issue", pathParams))
	if err != nil {
		return IssueCollection{}, err
	}

	u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return IssueCollection{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return IssueCollection{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return IssueCollection{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response IssueCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return IssueCollection{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return IssueCollection{}, err
		}

		return IssueCollection{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return IssueCollection{}, err
		}

		return IssueCollection{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return IssueCollection{}, err
		}

		return IssueCollection{}, &MessageException{
			Payload: response,
		}
	default:
		return IssueCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewIssueTag(httpClient *http.Client, parser *sdkgen.Parser) *IssueTag {
	return &IssueTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}
