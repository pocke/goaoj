package goaoj

import (
	"fmt"
	"net/url"
)

type problemList struct {
	List []Problem `xml:"problem"`
}

var _ respObj = &problemList{}

func (pl *problemList) getError() error {
	if len(pl.List) == 0 {
		return fmt.Errorf("Invalid volume")
	}
	return nil
}

func (a *API) GetProblemList(volume int) ([]Problem, error) {
	v := url.Values{}
	v.Add("volume", fmt.Sprintf("%d", volume))

	pl := &problemList{}
	err := a.apiGet(BaseURL+"/problem_list", v, pl)
	if err != nil {
		return nil, err
	}

	return pl.List, nil
}
