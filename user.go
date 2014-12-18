package goaoj

import (
	"errors"
	"fmt"
	"net/url"
)

type User struct {
	ID             string    `xml:"id"`
	Name           string    `xml:"name"`
	RegisterDate   int       `xml:"registerdate"`
	LastSubmitDate int       `xml:"lastsubmitdate"`
	Status         *Status   `xml:"status"`
	SolvedList     []Problem `xml:"solved_list>problem"`
}

type Status struct {
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

type Problem struct {
	judgeID        int    `xml:"judge_id"`
	ID             int    `xml:"id"`
	SubmissionDate int    `xml:"submissiondate"`
	Language       string `xml:"language"`
	CPUTime        int    `xml:"cputime"`
	Memory         int    `xml:"memory"`
	CodeSize       int    `xml:"code_size"`
}

func (a *API) GetUser(id string) (*User, error) {
	v := url.Values{}
	v.Add("id", id)

	u := &User{}
	err := a.apiGet(BaseURL+"/user", v, u)
	if err != nil {
		return nil, err
	}
	if u.Status == nil {
		return nil, errors.New(fmt.Sprintf("%s does not exists.", id))
	}

	return u, nil
}
