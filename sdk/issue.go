
// issue automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

type Issue struct {
    Id int `json:"id"`
    User *User `json:"user"`
    Document *Document `json:"document"`
    Status int `json:"status"`
    Title string `json:"title"`
    UpdateDate string `json:"updateDate"`
    InsertDate string `json:"insertDate"`
}

