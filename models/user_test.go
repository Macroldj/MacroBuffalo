package models

import (
	"encoding/json"
	"testing"
)

func Test_User(t *testing.T) {
	data := []byte(`
    {
        "id": 5,
				"first_name": "eddie",
				"last_name": "chen",
				"email": "eddie.chen@test.com"
    }
  `)

	var u User
	err := json.Unmarshal(data, &u)

	if err != nil {
		t.Error(err)
	}

	if u.ID != 5 || u.FirstName != "eddie" || u.LastName != "chen" || u.Email != "eddie.chen@test.com" {
		t.Error("invalid user json schema")
	}
}
