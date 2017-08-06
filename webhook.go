package webhook

import (
	"encoding/binary"
	"strconv"
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
	URL               string  `json:"url"`
	CreatedAt         Time    `json:"createdAt"`
	Status            string  `json:"status"`
	IsOpen            bool    `json:"isOpen"`
	Trigger           string  `json:"trigger"`
	MonitorName       string  `json:"monitorName"`
	MetricLabel       string  `json:"metricLabel"`
	MetricValue       float64 `json:"metricValue"`
	CriticalThreshold int     `json:"criticalThreshold"`
	WarningThreshold  int     `json:"warningThreshold"`
	MonitorOperator   string  `json:"monitorOperator"`
	Duration          int     `json:"duration"`
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
