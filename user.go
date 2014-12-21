package goaoj

import (
	"fmt"
	"net/url"
)

type respObj interface {
	getError() error
}

type User struct {
	ID             string        `xml:"id"`
	Name           string        `xml:"name"`
	RegisterDate   int           `xml:"registerdate"`
	LastSubmitDate int           `xml:"lastsubmitdate"`
	Status         *UserStatus   `xml:"status"`
	SolvedList     []UserProblem `xml:"solved_list>problem"`
}

var _ respObj = &User{}

type UserStatus struct {
	Submission   int `xml:"submission"`
	Solved       int `xml:"solved"`
	Accepted     int `xml:"accepted"`
	WrongAnswer  int `xml:"wronganswer"`
	TimeLimit    int `xml:"timelimit"`
	MemoryLimit  int `xml:"memorylimit"`
	OutputLimit  int `xml:"outputlimit"`
	RuntimeError int `xml:"runtimeerror"`
	CompileError int `xml:"compileerror"`
}

type UserProblem struct {
	judgeID        int    `xml:"judge_id"`
	ID             int    `xml:"id"`
	SubmissionDate int    `xml:"submissiondate"`
	Language       string `xml:"language"`
	CPUTime        int    `xml:"cputime"`
	Memory         int    `xml:"memory"`
	CodeSize       int    `xml:"code_size"`
}

func (u *User) getError() error {
	if u.ID == "" {
		// TODO: return ewAPIError("Invalid ID")
		return fmt.Errorf("Invalid ID")
	}
	return nil
}

func (a *API) GetUser(id string) (*User, error) {
	v := url.Values{}
	v.Add("id", id)

	u := &User{}
	err := a.apiGet(BaseURL+"/user", v, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
