package facade

import "testing"

var expect = "A module running\nB module running"

func TestFacadeAPI(t *testing.T) {
	api := NewAPI()
	actual := api.Test()
	if expect != actual {
		t.Fatalf("expect %s,actual %s", expect, actual)
	}
}
