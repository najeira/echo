package metrics

import (
	"sync"
	"time"
	"net/http"

	"github.com/labstack/echo"
	mt "github.com/rcrowley/go-metrics"
)

var (
	metricsHttp *metrics
)

func init() {
	metricsHttp = &metrics{timers: make(map[string]metric)}
}

func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer Measure(c, time.Now())
			return next(c)
		}
	}
}

func GetStats() []Stats {
	return metricsHttp.GetStats()
}

func Handler(c echo.Context) error {
	return c.JSON(http.StatusOK, GetStats())
}

func Measure(c echo.Context, start time.Time) {
	key := c.Path()
	res := c.Response()
	metricsHttp.Measure(key, start, res.Size, res.Status)
}

type metric struct {
	timer     mt.Timer
	size      mt.Meter
	status200 mt.Meter
	status300 mt.Meter
	status400 mt.Meter
	status500 mt.Meter
}

type metrics struct {
	mu     sync.RWMutex
	timers map[string]metric
}

func (m *metrics) Measure(key string, start time.Time, size int64, code int) {
	elapsed := time.Now().Sub(start)

	m.mu.RLock()
	t, ok := m.timers[key]
	m.mu.RUnlock()

	if !ok {
		m.mu.Lock()
		t, ok = m.timers[key]
		if !ok {
			t = metric{
				timer:     mt.NewTimer(),
				size:      mt.NewMeter(),
				status200: mt.NewMeter(),
				status300: mt.NewMeter(),
				status400: mt.NewMeter(),
				status500: mt.NewMeter(),
			}
			m.timers[key] = t
		}
		m.mu.Unlock()
	}

	t.timer.Update(elapsed)
	t.size.Mark(size)
	if 200 <= code && code <= 299 {
		t.status200.Mark(1)
	} else if 300 <= code && code <= 399 {
		t.status300.Mark(1)
	} else if 400 <= code && code <= 499 {
		t.status400.Mark(1)
	} else {
		t.status500.Mark(1)
	}
}

func (m *metrics) GetStats() []Stats {
	m.mu.RLock()
	defer m.mu.RUnlock()
	res := make([]Stats, 0, len(m.timers))
	for query, t := range m.timers {
		res = append(res, Stats{
			Query:         query,
			Count:         t.timer.Count(),
			Sum:           float64(t.timer.Sum()) / float64(time.Millisecond),
			Min:           float64(t.timer.Min()) / float64(time.Millisecond),
			Max:           float64(t.timer.Max()) / float64(time.Millisecond),
			Avg:           t.timer.Mean() / float64(time.Millisecond),
			Rate:          t.timer.Rate15(),
			P50:           t.timer.Percentile(0.5) / float64(time.Millisecond),
			P95:           t.timer.Percentile(0.95) / float64(time.Millisecond),
			Size:          t.size.Count(),
			SizeRate:      t.size.Rate15(),
			Status200:     t.status200.Count(),
			Status200Rate: t.status200.Rate15(),
			Status300:     t.status300.Count(),
			Status300Rate: t.status300.Rate15(),
			Status400:     t.status400.Count(),
			Status400Rate: t.status400.Rate15(),
			Status500:     t.status500.Count(),
			Status500Rate: t.status500.Rate15(),
		})
	}
	return res
}

type Stats struct {
	Query         string  `csv:"query" json:"query"`
	Count         int64   `csv:"count" json:"count"`
	Sum           float64 `csv:"sum" json:"sum"`
	Min           float64 `csv:"min" json:"min"`
	Max           float64 `csv:"max" json:"max"`
	Avg           float64 `csv:"avg" json:"avg"`
	Rate          float64 `csv:"rate" json:"rate"`
	P50           float64 `csv:"p50" json:"p50"`
	P95           float64 `csv:"p95" json:"p95"`
	Size          int64   `csv:"size" json:"size"`
	SizeRate      float64 `csv:"size_rate" json:"size_rate"`
	Status200     int64   `csv:"status_200" json:"status_200"`
	Status200Rate float64 `csv:"status_200_rate" json:"status_200_rate"`
	Status300     int64   `csv:"status_300" json:"status_300"`
	Status300Rate float64 `csv:"status_300_rate" json:"status_300_rate"`
	Status400     int64   `csv:"status_400" json:"status_400"`
	Status400Rate float64 `csv:"status_400_rate" json:"status_400_rate"`
	Status500     int64   `csv:"status_500" json:"status_500"`
	Status500Rate float64 `csv:"status_500_rate" json:"status_500_rate"`
}
