package goaoj_test

import "testing"

func TestGetProblemList(t *testing.T) {
	pl, err := api.GetProblemList(100)
	if err != nil {
		t.Errorf("ProblemList: %+v\n", pl)
		t.Error(err)
	}
	if len(pl) == 0 {
		t.Error("ProblemList should not be empty. But got empty.")
	}

	pl, err = api.GetProblemList(83715892)
	if err == nil {
		t.Errorf("ProblemList: %+v\n", pl)
		t.Error("Should return error when does not exists volume. But not return error.")
	}
}
