package metrics

import (
	_ "database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestMetrics_GoodFake(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ms, err := NewMetricsService(db, 1)

	if err != nil {
		t.Fatalf("ошибка при создании экземпляра MetricsService %s", err)
	}

	qm := &QuantityMetric{
		Area:      "100500",
		TableName: "fake_table",
		Value:     5000,
		Hash:      "15135345",
	}
	dm := time.Now()
	qm.DateMetric = dm
	ms.Submit(qm)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO quantity_metrics").
		WithArgs(qm.DateMetric, qm.Area, qm.TableName, qm.Value, qm.Hash).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	time.Sleep(time.Duration(3) * time.Second)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestMetrics_StopStart(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	ms, err := NewMetricsService(db, 1)
	if err != nil {
		t.Fatalf("ошибка при создании экземпляра MetricsService %s", err)
	}
	ms.Stop()
	err = ms.Start()
	if err != nil {
		t.Fatalf("ошибка запуска MetricsService %s", err)
	}
}
