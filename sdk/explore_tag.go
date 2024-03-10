// ExploreTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package sdk

import (
	"encoding/json"
	"errors"
	"github.com/apioo/sdkgen-go"
	"io"
	"net/http"
	"net/url"
)

type ExploreTag struct {
	internal *sdkgen.TagAbstract
}

// GetAll Returns trending documents
func (client *ExploreTag) GetAll(startIndex int, count int, search string) (DocumentCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	var queryStructNames []string

	u, err := url.Parse(client.internal.Parser.Url("/explore", pathParams))
	if err != nil {
		return DocumentCollection{}, err
	}

	u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return DocumentCollection{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return DocumentCollection{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return DocumentCollection{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response DocumentCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return DocumentCollection{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return DocumentCollection{}, err
		}

		return DocumentCollection{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return DocumentCollection{}, err
		}

		return DocumentCollection{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return DocumentCollection{}, err
		}

		return DocumentCollection{}, &MessageException{
			Payload: response,
		}
	default:
		return DocumentCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewExploreTag(httpClient *http.Client, parser *sdkgen.Parser) *ExploreTag {
	return &ExploreTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}
