package webhook

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type WebHook struct {
	OrgName string `json:"orgName"`
	Event   string `json:"event"`
	Host    Host   `json:"host"`
	Alert   Alert  `json:"alert"`
}

type Host struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	IsRetired bool   `json:"isRetired"`
	ID        string `json:"id"`
	Status    string `json:"status"`
	Memo      string `json:"memo"`
	Roles     []Role `json:"roles"`
}

type Role struct {
	Fullname    string `json:"fullname"`
	ServiceName string `json:"serviceName"`
	RoleName    string `json:"roleName"`
	ServiceURL  string `json:"serviceUrl"`
	RoleURL     string `json:"roleUrl"`
}

type Alert struct {
	URL               string `json:"url"`
	CreatedAt         Time   `json:"createdAt"`
	Status            string `json:"status"`
	IsOpen            bool   `json:"isOpen"`
	Trigger           string `json:"trigger"`
	MonitorName       string `json:"monitorName"`
	MetricLabel       string `json:"metricLabel"`
	MetricValue       Number `json:"metricValue"`
	CriticalThreshold Number `json:"criticalThreshold"`
	WarningThreshold  Number `json:"warningThreshold"`
	MonitorOperator   string `json:"monitorOperator"`
	Duration          Number `json:"duration"`
}

type Number struct {
	asInt64   *int64
	asFloat64 *float64
}

func NewInt64AsNumber(i int64) *Number {
	return &Number{asInt64: &i, asFloat64: nil}
}

func NewFloat64AsNumber(f float64) *Number {
	return &Number{asInt64: nil, asFloat64: &f}
}

func (t *Number) AsInt64() *int64 {
	return t.asInt64
}

func (t *Number) AsFloat64() *float64 {
	return t.asFloat64
}

func (t *Number) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if t.asFloat64 != nil {
		fmt.Fprintf(&buf, `%f`, *t.asFloat64)
	} else if t.asInt64 != nil {
		fmt.Fprintf(&buf, `%d`, *t.asInt64)
	}

	return buf.Bytes(), nil
}

func (t *Number) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	rawNumber := string(data)
	if strings.Contains(rawNumber, ".") {
		f, err := strconv.ParseFloat(rawNumber, 64)
		if err != nil {
			return err
		}
		t.asFloat64 = &f
		t.asInt64 = nil
	} else {
		i, err := strconv.ParseInt(rawNumber, 10, 64)
		if err != nil {
			return err
		}
		t.asInt64 = &i
		t.asFloat64 = nil
	}
	return err
}

type Time struct {
	time.Time
}

func (t *Time) MarshalJSON() ([]byte, error) {
	u := make([]byte, 0, 8)
	binary.LittleEndian.PutUint64(u, uint64(t.Time.Unix()))
	return u, nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	msec, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	t.Time = time.Unix(msec/1000, 0)
	return err
}
