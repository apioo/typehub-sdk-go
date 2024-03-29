// document automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package sdk

import "time"

type Document struct {
	Id          int       `json:"id"`
	User        *User     `json:"user"`
	Source      *Document `json:"source"`
	Status      int       `json:"status"`
	Stars       int       `json:"stars"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	BaseUrl     string    `json:"baseUrl"`
	Keywords    []string  `json:"keywords"`
	Homepage    string    `json:"homepage"`
	License     string    `json:"license"`
	Spec        any       `json:"spec"`
	UpdateDate  time.Time `json:"updateDate"`
	InsertDate  time.Time `json:"insertDate"`
}
