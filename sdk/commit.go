// commit automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package sdk

import "time"

type Commit struct {
	Id         int            `json:"id"`
	Document   Document       `json:"document"`
	User       User           `json:"user"`
	Previous   CommitPrevious `json:"previous"`
	Message    string         `json:"message"`
	CommitHash string         `json:"commitHash"`
	Spec       any            `json:"spec"`
	InsertDate time.Time      `json:"insertDate"`
}