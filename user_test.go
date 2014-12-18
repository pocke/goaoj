package goaoj_test

import (
	"testing"

	"github.com/pocke/goaoj"
)

var api = goaoj.NewAPI()

func TestGetUser(t *testing.T) {
	u, err := api.GetUser("pck")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%d", len(u.SolvedList))
	// Too many log
	u.SolvedList = make([]goaoj.Problem, 0)
	t.Logf("%+v", u)
	t.Logf("%+v", u.Status)
}
