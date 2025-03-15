
// IssueTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "github.com/apioo/sdkgen-go/v2"
    "io"
    "net/http"
    "net/url"
    
)

type IssueTag struct {
    internal *sdkgen.TagAbstract
}



// Create Creates a new issue
func (client *IssueTag) Create(user string, document string, payload IssueCreate) (*Message, error) {
    pathParams := make(map[string]interface{})
    pathParams["user"] = user
    pathParams["document"] = document

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/issue", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("POST", u.String(), reqBody)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 400 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 404 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// CreateComment Creates a new issue comment
func (client *IssueTag) CreateComment(user string, document string, id string, payload CommentCreate) (*Message, error) {
    pathParams := make(map[string]interface{})
    pathParams["user"] = user
    pathParams["document"] = document
    pathParams["id"] = id

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/issue/:id/comment", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("POST", u.String(), reqBody)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 400 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 404 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Delete Removes an issue
func (client *IssueTag) Delete(user string, document string, id string) (*Message, error) {
    pathParams := make(map[string]interface{})
    pathParams["user"] = user
    pathParams["document"] = document
    pathParams["id"] = id

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/issue/:id", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("DELETE", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 400 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 404 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Get Returns an issue
func (client *IssueTag) Get(user string, document string, id string) (*Issue, error) {
    pathParams := make(map[string]interface{})
    pathParams["user"] = user
    pathParams["document"] = document
    pathParams["id"] = id

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/issue/:id", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data Issue
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 400 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 404 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetAll Returns all issues
func (client *IssueTag) GetAll(user string, document string, status int, startIndex int, count int, search string) (*IssueCollection, error) {
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
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data IssueCollection
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 400 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 404 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetAllComments Shows all issue comments
func (client *IssueTag) GetAllComments(user string, document string, id string, startIndex int, count int, search string) (*CommentCollection, error) {
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
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data CommentCollection
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 400 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 404 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// ReactComment Reacts to a comment
func (client *IssueTag) ReactComment(user string, document string, id string, comment string, reaction string, payload Passthru) (*Message, error) {
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
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("POST", u.String(), reqBody)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 400 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 404 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data Message
        err := json.Unmarshal(respBody, &data)

        return nil, &MessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}




func NewIssueTag(httpClient *http.Client, parser *sdkgen.Parser) *IssueTag {
	return &IssueTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
