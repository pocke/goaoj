package goaoj

import (
	"fmt"
	"net/url"
)

type Problem struct {
	ID                 sstring        `xml:"id"`
	Name               sstring        `xml:"name"`
	Available          iint           `xml:"available"`
	ProblemTimeLimit   iint           `xml:"problemtimelimit"`
	ProblemMemoryLimit iint           `xml:"problemmemorylimit"`
	Status             *ProblemStatus `xml:"status"`
	SolvedList         []ProblemUser  `xml:"solved_list>user"`
}

var _ respObj = &Problem{}

type ProblemStatus struct {
	Submission   iint `xml:"submission"`
	Accepted     iint `xml:"accepted"`
	WrongAnswer  iint `xml:"wronganswer"`
	TimeLimit    iint `xml:"timelimit"`
	MemoryLimit  iint `xml:"memorylimit"`
	OutputLimit  iint `xml:"outputlimit"`
	RuntimeError iint `xml:"runtimeerror"`
}

type ProblemUser struct {
	ID             sstring `xml:"id"`
	SubmissionDate iint    `xml:"submissiondate"`
	Language       sstring `xml:"language"`
	CPUTime        iint    `xml:"cputime"`
	Memory         iint    `xml:"memory"`
	CodeSize       iint    `xml:"code_size"`
}

func (p *Problem) getError() error {
	if p.ID == "" {
		// TODO: return ewAPIError("Invalid ID")
		return fmt.Errorf("Invalid ID")
	}
	return nil
}

func (a *API) GetProblem(id string, status bool) (*Problem, error) {
	v := url.Values{}
	v.Add("id", id)
	if !status {
		v.Add("status", "false")
	}

	p := &Problem{}
	err := a.apiGet(BaseURL+"/problem", v, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
