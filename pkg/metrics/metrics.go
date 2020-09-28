package metrics

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
)

const (
	// Время простоя таймера
	// перед записью накопившихся данных
	idleTimeoutSec = 10
)

// MetricsService ...
type MetricsService struct {
	timeout    time.Duration
	qmQueue    chan *QuantityMetric
	db         *sql.DB
	queueClose bool
}

// NewMetricsService ...
func NewMetricsService(database *sql.DB, timeoutSec time.Duration) (*MetricsService, error) {

	if err := database.Ping(); err != nil {
		return nil, err
	}

	if timeoutSec == 0 {
		timeoutSec = idleTimeoutSec
	}

	ms := &MetricsService{
		timeout:    time.Second * timeoutSec,
		qmQueue:    make(chan *QuantityMetric),
		db:         database,
		queueClose: false,
	}

	go ms.dispatch()
	return ms, nil
}

// Submit ...
func (ms *MetricsService) Submit(n *QuantityMetric) {
	if n != nil {
		ms.qmQueue <- n
	}
}

// Start ...
func (ms *MetricsService) Start() error {
	if ms.queueClose {
		ms.qmQueue = make(chan *QuantityMetric)
		ms.queueClose = false
	}
	go ms.dispatch()
	return nil
}

// Stop ...
func (ms *MetricsService) Stop() {
	close(ms.qmQueue)
	ms.queueClose = true
}

func (ms *MetricsService) dispatch() {
	timeout := time.NewTimer(ms.timeout)
	var (
		arr  []*QuantityMetric
		task *QuantityMetric
		ok   bool
	)
Loop:
	for {
		timeout.Reset(ms.timeout)
		select {
		case task, ok = <-ms.qmQueue:
			if !ok {
				break Loop
			}
			arr = append(arr, task)
		case <-timeout.C:
			if len(arr) != 0 {
				if err := ms.insertMetrics(arr); err != nil {
					fmt.Println("[metrics] error: ", err)
					break Loop
				}
				arr = make([]*QuantityMetric, 0)
			}
		}
	}
}

func (ms *MetricsService) insertMetrics(arrayMetrics []*QuantityMetric) error {
	tx, err := ms.db.Begin()
	if err != nil {
		return err
	}
	fNames := []string{"date_metric", "area", "table_name", "value", "hash"}
	stmt, err := tx.Prepare(pq.CopyIn("quantity_metrics", fNames...))
	if err != nil {
		return err
	}
	for _, m := range arrayMetrics {
		values := make([]interface{}, 0, len(fNames))
		values = append(values, m.DateMetric)
		values = append(values, m.Area)
		values = append(values, m.TableName)
		values = append(values, m.Value)
		values = append(values, m.Hash)
		_, err = stmt.Exec(values...)
		if err != nil {
			return err
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
