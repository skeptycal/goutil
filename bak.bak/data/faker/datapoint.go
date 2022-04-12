package faker

import (
	"strings"
	"time"
)

type (
	KindMap map[string]int

	// dataPoint struct {
	// 	t0   time.Time
	// 	t1   time.Time
	// 	data []byte
	// }

	dataPoint struct {
		t0   time.Time
		t1   time.Time
		fn   func() Any
		data Any
	}
)

func (d *dataPoint) Start() {
	d.t0 = time.Now()
}

func (d *dataPoint) Stop() {
	d.t1 = time.Now()
}

func (d *dataPoint) Collect() {
	d.Start()
	d.data = d.fn()
	d.Stop()
}

func GetData() DataPoint {
	d := new(dataPoint)
	d.Collect()
	return d
}

func getEncodedString1(n int) string {
	b := []byte("")
	for i := 0; i < n; i++ {
		b = append(b, byte(RandomKind(true))+65)
	}
	return string(b)
}

func getEncodedString2(n int) string {
	sb := strings.Builder{}
	defer sb.Reset()
	for i := 0; i < n; i++ {
		sb.WriteByte(byte(RandomKind(true) + 65))
	}
	return sb.String()
}

func GetEncodedString(n int) string {
	return getEncodedString1(n)
}
