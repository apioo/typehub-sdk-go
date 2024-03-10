// TagTag automatically generated by SDKgen please do not edit this file manually
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

type TagTag struct {
	internal *sdkgen.TagAbstract
}

// Changelog Generates the changelog for the current release
func (client *TagTag) Changelog(user string, document string, payload Passthru) (TagChangelog, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document

	queryParams := make(map[string]interface{})

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/changelog", pathParams))
	if err != nil {
		return TagChangelog{}, err
	}

	u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return TagChangelog{}, err
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("POST", u.String(), reqBody)
	if err != nil {
		return TagChangelog{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return TagChangelog{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return TagChangelog{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response TagChangelog
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TagChangelog{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TagChangelog{}, err
		}

		return TagChangelog{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TagChangelog{}, err
		}

		return TagChangelog{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TagChangelog{}, err
		}

		return TagChangelog{}, &MessageException{
			Payload: response,
		}
	default:
		return TagChangelog{}, errors.New("the server returned an unknown status code")
	}
}

// Create Creates a new tag
func (client *TagTag) Create(user string, document string, payload Passthru) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document

	queryParams := make(map[string]interface{})

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/tag", pathParams))
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

// Get Returns a tag
func (client *TagTag) Get(user string, document string, id string) (Tag, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document
	pathParams["id"] = id

	queryParams := make(map[string]interface{})

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/tag/:id", pathParams))
	if err != nil {
		return Tag{}, err
	}

	u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Tag{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Tag{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Tag{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Tag
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Tag{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Tag{}, err
		}

		return Tag{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Tag{}, err
		}

		return Tag{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Tag{}, err
		}

		return Tag{}, &MessageException{
			Payload: response,
		}
	default:
		return Tag{}, errors.New("the server returned an unknown status code")
	}
}

// GetAll Returns all tags for a document
func (client *TagTag) GetAll(user string, document string, startIndex int, count int, search string) (TagCollection, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document

	queryParams := make(map[string]interface{})
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/tag", pathParams))
	if err != nil {
		return TagCollection{}, err
	}

	u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return TagCollection{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return TagCollection{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return TagCollection{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response TagCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TagCollection{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TagCollection{}, err
		}

		return TagCollection{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TagCollection{}, err
		}

		return TagCollection{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TagCollection{}, err
		}

		return TagCollection{}, &MessageException{
			Payload: response,
		}
	default:
		return TagCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewTagTag(httpClient *http.Client, parser *sdkgen.Parser) *TagTag {
	return &TagTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}