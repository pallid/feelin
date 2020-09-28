package metrics

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
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

func (ms *MetricsService) insertMetrics(arrayMetrics []*QuantityMetric) (err error) {
	valueStrings := []string{}
	valueArgs := []interface{}{}
	for _, m := range arrayMetrics {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, m.DateMetric)
		valueArgs = append(valueArgs, m.Area)
		valueArgs = append(valueArgs, m.TableName)
		valueArgs = append(valueArgs, m.Value)
		valueArgs = append(valueArgs, m.Hash)
	}
	smt := `INSERT INTO quantity_metrics (date_metric, area, table_name, value, hash) VALUES %s`
	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))
	tx, err := ms.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()
	_, err = tx.Exec(smt, valueArgs...)
	return
}
