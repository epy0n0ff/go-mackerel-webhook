package webhook

import (
	"encoding/json"
	"testing"
	"time"
)

func TestUnmarshal(t *testing.T) {
	rawJson := `
{
	"orgName": "TestOrg",
	"event": "alert",
	"host": {
		"name": "epy0n0ff-TestPC",
		"url": "https://mackerel.io/orgs/TestOrg/hosts/32subXxxxxx",
		"isRetired": false,
		"id": "32subXxxxxx",
		"status": "working",
		"memo": "",
		"roles": [{
				"fullname": "TestSub: TestSubHDD",
				"serviceName": "TestSub",
				"roleName": "TestSubHDD",
				"serviceUrl": "https://mackerel.io/orgs/Test/services/TestSub",
				"roleUrl": "https://mackerel.io/orgs/Test/services/TestSub#role=TestSubHDD"
			},
			{
				"fullname": "TestNetwork: TestNetwork",
				"serviceName": "TestNetwork",
				"roleName": "TestNetwork",
				"serviceUrl": "https://mackerel.io/orgs/TestOrg/services/TestNetwork",
				"roleUrl": "https://mackerel.io/orgs/TestOrg/services/TestNetwork#role=TestNetwork"
			}
		]
	},
	"alert": {
		"url": "https://mackerel.io/orgs/TestOrg/alerts/359Mwoxxxxx",
		"createdAt": 1501610221657,
		"status": "ok",
		"isOpen": false,
		"trigger": "delete monitor",
		"monitorName": "Memory %テスト監視",
		"metricLabel": "Memory %",
		"criticalThreshold": 20,
		"warningThreshold": 19,
		"monitorOperator": ">",
		"duration": 1
	}
}`
	var hook WebHook
	err := json.Unmarshal([]byte(rawJson), &hook)
	if err != nil {
		t.Fatalf("unexpected error:%v", err)
	}

	ctime := hook.Alert.CreatedAt.Format(time.RFC3339)
	if "2017-08-02T02:57:01+09:00" != ctime {
		t.Fatalf("unexpected error: 2017-08-02T02:57:01+09:00 != %s", ctime)
	}
}

func TestUnmarshalNumber(t *testing.T) {
	rawJson := `
{
	"orgName": "TestOrg",
	"event": "alert",
	"host": {
		"name": "epy0n0ff-TestPC",
		"url": "https://mackerel.io/orgs/TestOrg/hosts/32subXxxxxx",
		"isRetired": false,
		"id": "32subXxxxxx",
		"status": "working",
		"memo": "",
		"roles": [{
				"fullname": "TestSub: TestSubHDD",
				"serviceName": "TestSub",
				"roleName": "TestSubHDD",
				"serviceUrl": "https://mackerel.io/orgs/Test/services/TestSub",
				"roleUrl": "https://mackerel.io/orgs/Test/services/TestSub#role=TestSubHDD"
			},
			{
				"fullname": "TestNetwork: TestNetwork",
				"serviceName": "TestNetwork",
				"roleName": "TestNetwork",
				"serviceUrl": "https://mackerel.io/orgs/TestOrg/services/TestNetwork",
				"roleUrl": "https://mackerel.io/orgs/TestOrg/services/TestNetwork#role=TestNetwork"
			}
		]
	},
	"alert": {
		"url": "https://mackerel.io/orgs/TestOrg/alerts/359Mwoxxxxx",
		"createdAt": 1501610221657,
		"status": "ok",
		"isOpen": false,
		"trigger": "delete monitor",
		"monitorName": "Memory %テスト監視",
		"metricLabel": "Memory %",
		"criticalThreshold": 20,
		"warningThreshold": 1.4665636369580741,
		"monitorOperator": ">",
		"duration": 1
	}
}`

	var hook WebHook
	err := json.Unmarshal([]byte(rawJson), &hook)
	if err != nil {
		t.Fatalf("unexpected error:%v", err)
	}

	if 20 != *hook.Alert.CriticalThreshold.AsInt64 {
		t.Fatalf("unexpected error: criticalThreshold: 20 != %d", *hook.Alert.CriticalThreshold.AsInt64)
	}
	if hook.Alert.CriticalThreshold.AsFloat64 != nil {
		t.Fatalf("unexpected error: CriticalThreshold.AsFloat64 must be null")
	}

	if 1.4665636369580741 != *hook.Alert.WarningThreshold.AsFloat64 {
		t.Fatalf("unexpected error: warningThreshold: 1.4665636369580741 != %f", *hook.Alert.WarningThreshold.AsFloat64)
	}
	if hook.Alert.WarningThreshold.AsInt64 != nil {
		t.Fatalf("unexpected error: WarningThreshold.AsInt64 must be null")
	}

}
