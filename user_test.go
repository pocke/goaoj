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
	// Too many log
	u.SolvedList = make([]goaoj.UserProblem, 0)
	t.Logf("User: %+v", u)
	t.Logf("User status: %+v", u.Status)

	u, err = api.GetUser("InvalidUserNameあ")
	if err == nil {
		t.Logf("%+v", u)
		t.Error("Should return error when user does not exists. But not return error.")
	}
}
