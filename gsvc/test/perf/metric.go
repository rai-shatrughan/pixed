package perf

import "time"

type Metric struct {
	// maxDuration time.Duration
	// minDuration time.Duration
	// avgDuration time.Duration
	failed    int64
	count     int64
	duration  time.Duration
	startTime time.Time
	endTime   time.Time
}

func (m *Metric) SetStartTime(t time.Time) {
	m.startTime = t
}

func (m *Metric) StartTime() time.Time {
	return m.startTime.UTC()
}

func (m *Metric) SetEndTime(t time.Time) {
	m.endTime = t
}

func (m *Metric) EndTime() time.Time {
	return m.endTime.UTC()
}

func (m *Metric) IncrDuration(d time.Duration) {
	m.duration = m.duration + d
}

func (m *Metric) Duration() time.Duration {
	return m.duration
}

func (m *Metric) IncrCount(c int64) {
	m.count = m.count + c
}

func (m *Metric) Count() int64 {
	return m.count
}

func (m *Metric) IncrFailed(c int64) {
	m.failed = m.failed + c
}

func (m *Metric) Failed() int64 {
	return m.failed
}

func (m *Metric) AvgDuration() time.Duration {
	return m.duration / time.Duration(m.count)
}
