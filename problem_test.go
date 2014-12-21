package goaoj_test

import "testing"

func TestGetProblem(t *testing.T) {
	p, err := api.GetProblem("0002", false)
	if err != nil {
		t.Error(err)
	}
	if len(p.SolvedList) != 0 {
		t.Error("Problem.SolvedList should be empty when status false. But not empty!")
	}
	t.Logf("Problem: %+v", p)

	p, err = api.GetProblem("0002", true)
	if err != nil {
		t.Error(err)
	}
	if len(p.SolvedList) == 0 {
		t.Error("Problem.SolvedList should not be empty when status true. But empty!")
		t.Logf("Problem: %+v", p)
	}

	p, err = api.GetProblem("InvalidProblemID", false)
	if err == nil {
		t.Logf("%+v", p)
		t.Error("Should return error when problem does not exists. But not return error.")
	}
}
