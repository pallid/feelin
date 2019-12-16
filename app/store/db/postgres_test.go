package db

import (
	"testing"

	"github.com/pallid/feelin/app/model"
)

func TestSetQueryTextForSelectData(t *testing.T) {
	expected := `SELECT * from  WHERE area = 100500 AND test_field_1 = ? AND test_field_2 = ? AND test_field_n = ?`
	q := &model.QueryResult{Area: 100500}
	q.SelectionFields = []string{"test_field_1", "test_field_2", "test_field_n"}
	r := &PostgresRepository{}
	r.SetQueryTextForSelectData(q)
	sd := q.SelectData
	if sd != expected {
		t.Error(
			"SetQueryTextForSelectData",
			"expected", expected,
			"got", sd,
		)
	}

}
