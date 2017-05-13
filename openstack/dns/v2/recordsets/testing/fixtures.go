package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/recordsets"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

// ListByZoneOutput is a sample response to a ListByZone call.
const ListByZoneOutput = `
{
    "recordsets": [
        {
            "description": "This is an example record set.",
            "links": {
                "self": "https://127.0.0.1:9001/v2/zones/2150b1bf-dee2-4221-9d85-11f7886fb15f/recordsets/f7b10e9b-0cae-4a91-b162-562bc6096648"
            },
            "updated_at": null,
            "records": [
                "10.1.0.2"
            ],
            "ttl": 3600,
            "id": "f7b10e9b-0cae-4a91-b162-562bc6096648",
            "name": "example.org.",
            "project_id": "4335d1f0-f793-11e2-b778-0800200c9a66",
            "zone_id": "2150b1bf-dee2-4221-9d85-11f7886fb15f",
            "zone_name": "example.com.",
            "created_at": "2014-10-24T19:59:44.000000",
            "version": 1,
            "type": "A",
            "status": "PENDING",
            "action": "CREATE"
        },
        {
            "description": "This is another example record set.",
            "links": {
                "self": "https://127.0.0.1:9001/v2/zones/2150b1bf-dee2-4221-9d85-11f7886fb15f/recordsets/f7b10e9b-0cae-4a91-b162-562bc6096648"
            },
            "updated_at": "2017-03-04T14:29:07.000000",
            "records": [
                "10.1.0.3",
                "10.1.0.4"
            ],
            "ttl": 3600,
            "id": "7423aeaf-b354-4bd7-8aba-2e831567b478",
            "name": "foo.example.org.",
            "project_id": "4335d1f0-f793-11e2-b778-0800200c9a66",
            "zone_id": "2150b1bf-dee2-4221-9d85-11f7886fb15f",
            "zone_name": "example.com.",
            "created_at": "2014-10-24T19:59:44.000000",
            "version": 1,
            "type": "A",
            "status": "PENDING",
            "action": "CREATE"
        }
    ],
    "links": {
        "self": "http://127.0.0.1:9001/v2/zones/2150b1bf-dee2-4221-9d85-11f7886fb15f/recordsets?limit=1"
    },
    "metadata": {
        "total_count": 2
    }
}
`

// GetOutput is a sample response to a Get call.
const GetOutput = `
{
		"description": "This is an example record set.",
		"links": {
				"self": "https://127.0.0.1:9001/v2/zones/2150b1bf-dee2-4221-9d85-11f7886fb15f/recordsets/f7b10e9b-0cae-4a91-b162-562bc6096648"
		},
		"updated_at": null,
		"records": [
				"10.1.0.2"
		],
		"ttl": 3600,
		"id": "f7b10e9b-0cae-4a91-b162-562bc6096648",
		"name": "example.org.",
		"project_id": "4335d1f0-f793-11e2-b778-0800200c9a66",
		"zone_id": "2150b1bf-dee2-4221-9d85-11f7886fb15f",
		"zone_name": "example.com.",
		"created_at": "2014-10-24T19:59:44.000000",
		"version": 1,
		"type": "A",
		"status": "PENDING",
		"action": "CREATE"
}
`

// FirstRecordSet is the first result in ListByZoneOutput
var FirstRecordSetCreatedAt, _ = time.Parse(gophercloud.RFC3339MilliNoZ, "2014-10-24T19:59:44.000000")
var FirstRecordSet = recordsets.RecordSet{
	ID:          "f7b10e9b-0cae-4a91-b162-562bc6096648",
	Description: "This is an example record set.",
	UpdatedAt:   time.Time{},
	Records:     []string{"10.1.0.2"},
	TTL:         3600,
	Name:        "example.org.",
	ProjectID:   "4335d1f0-f793-11e2-b778-0800200c9a66",
	ZoneID:      "2150b1bf-dee2-4221-9d85-11f7886fb15f",
	ZoneName:    "example.com.",
	CreatedAt:   FirstRecordSetCreatedAt,
	Version:     1,
	Type:        "A",
	Status:      "PENDING",
	Action:      "CREATE",
	Links: map[string]interface{}{
		"self": "https://127.0.0.1:9001/v2/zones/2150b1bf-dee2-4221-9d85-11f7886fb15f/recordsets/f7b10e9b-0cae-4a91-b162-562bc6096648",
	},
}

// SecondRecordSet is the first result in ListByZoneOutput
var SecondRecordSetCreatedAt, _ = time.Parse(gophercloud.RFC3339MilliNoZ, "2014-10-24T19:59:44.000000")
var SecondRecordSetUpdatedAt, _ = time.Parse(gophercloud.RFC3339MilliNoZ, "2017-03-04T14:29:07.000000")
var SecondRecordSet = recordsets.RecordSet{
	ID:          "7423aeaf-b354-4bd7-8aba-2e831567b478",
	Description: "This is another example record set.",
	UpdatedAt:   SecondRecordSetUpdatedAt,
	Records:     []string{"10.1.0.3", "10.1.0.4"},
	TTL:         3600,
	Name:        "foo.example.org.",
	ProjectID:   "4335d1f0-f793-11e2-b778-0800200c9a66",
	ZoneID:      "2150b1bf-dee2-4221-9d85-11f7886fb15f",
	ZoneName:    "example.com.",
	CreatedAt:   SecondRecordSetCreatedAt,
	Version:     1,
	Type:        "A",
	Status:      "PENDING",
	Action:      "CREATE",
	Links: map[string]interface{}{
		"self": "https://127.0.0.1:9001/v2/zones/2150b1bf-dee2-4221-9d85-11f7886fb15f/recordsets/f7b10e9b-0cae-4a91-b162-562bc6096648",
	},
}

// ExpectedRecordSetSlice is the slice of results that should be parsed
// from ListByZoneOutput, in the expected order.
var ExpectedRecordSetSlice = []recordsets.RecordSet{FirstRecordSet, SecondRecordSet}

// HandleListByZoneSuccessfully configures the test server to respond to a ListByZone request.
func HandleListByZoneSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/zones/2150b1bf-dee2-4221-9d85-11f7886fb15f/recordsets",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, ListByZoneOutput)
		})
}

// HandleGetSuccessfully configures the test server to respond to a Get request.
func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/zones/2150b1bf-dee2-4221-9d85-11f7886fb15f/recordsets/f7b10e9b-0cae-4a91-b162-562bc6096648",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, GetOutput)
		})
}
