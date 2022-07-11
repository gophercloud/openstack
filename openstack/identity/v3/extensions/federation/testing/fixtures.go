package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/identity/v3/extensions/federation"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

// ListOutput provides a single page of List results.
const ListOutput = `
{
    "links": {
        "next": null,
        "previous": null,
        "self": "http://example.com/identity/v3/OS-FEDERATION/mappings"
    },
    "mappings": [
        {
            "id": "ACME",
            "links": {
                "self": "http://example.com/identity/v3/OS-FEDERATION/mappings/ACME"
            },
            "rules": [
                {
                    "local": [
                        {
                            "user": {
                                "name": "{0}"
                            }
                        },
                        {
                            "group": {
                                "id": "0cd5e9"
                            }
                        }
                    ],
                    "remote": [
                        {
                            "type": "UserName"
                        },
                        {
                            "type": "orgPersonType",
                            "not_any_of": [
                                "Contractor",
                                "Guest"
                            ]
                        }
                    ]
                }
            ]
        }
    ]
}
`

const CreateRequest = `
  {
    "mapping": {
        "rules": [
            {
                "local": [
                    {
                        "user": {
                            "name": "{0}"
                        }
                    },
                    {
                        "group": {
                            "id": "0cd5e9"
                        }
                    }
                ],
                "remote": [
                    {
                        "type": "UserName"
                    },
                    {
                        "type": "orgPersonType",
                        "not_any_of": [
                            "Contractor",
                            "Guest"
                        ]
                    }
                ]
            }
        ]
    }
}
`

const CreateOutput = `
{
    "mapping": {
        "id": "ACME",
        "links": {
            "self": "http://example.com/identity/v3/OS-FEDERATION/mappings/ACME"
        },
        "rules": [
            {
                "local": [
                    {
                        "user": {
                            "name": "{0}"
                        }
                    },
                    {
                        "group": {
                            "id": "0cd5e9"
                        }
                    }
                ],
                "remote": [
                    {
                        "type": "UserName"
                    },
                    {
                        "type": "orgPersonType",
                        "not_any_of": [
                            "Contractor",
                            "Guest"
                        ]
                    }
                ]
            }
        ]
    }
}
`

const UpdateRequest = `
{
    "mapping": {
        "rules": [
            {
                "local": [
                    {
                        "user": {
                            "name": "{0}"
                        }
                    },
                    {
                        "group": {
                            "id": "0cd5e9"
                        }
                    }
                ],
                "remote": [
                    {
                        "type": "UserName"
                    },
                    {
                        "type": "orgPersonType",
                        "any_one_of": [
                            "Contractor",
                            "SubContractor"
                        ]
                    }
                ]
            }
        ]
    }
}
`

const UpdateOutput = `
{
    "mapping": {
        "id": "ACME",
        "links": {
            "self": "http://example.com/identity/v3/OS-FEDERATION/mappings/ACME"
        },
        "rules": [
            {
                "local": [
                    {
                        "user": {
                            "name": "{0}"
                        }
                    },
                    {
                        "group": {
                            "id": "0cd5e9"
                        }
                    }
                ],
                "remote": [
                    {
                        "type": "UserName"
                    },
                    {
                        "type": "orgPersonType",
                        "any_one_of": [
                            "Contractor",
                            "SubContractor"
                        ]
                    }
                ]
            }
        ]
    }
}
`

const GetOutput = `
{
    "mapping": {
        "id": "ACME",
        "links": {
            "self": "http://example.com/identity/v3/OS-FEDERATION/mappings/ACME"
        },
        "rules": [
            {
                "local": [
                    {
                        "user": {
                            "name": "{0}"
                        }
                    },
                    {
                        "group": {
                            "id": "0cd5e9"
                        }
                    }
                ],
                "remote": [
                    {
                        "type": "UserName"
                    },
                    {
                        "type": "orgPersonType",
                        "not_any_of": [
                            "Contractor",
                            "Guest"
                        ]
                    }
                ]
            }
        ]
    }
}
`

var TestMapping = federation.Mapping{
	ID: "ACME",
	Links: map[string]interface{}{
		"self": "http://example.com/identity/v3/OS-FEDERATION/mappings/ACME",
	},
	Rules: []federation.MappingRule{
		{
			Local: []federation.RuleLocal{
				{
					User: &federation.RuleUser{
						Name: "{0}",
					},
				},
				{
					Group: &federation.Group{
						ID: "0cd5e9",
					},
				},
			},
			Remote: []federation.RuleRemote{
				{
					Type: "UserName",
				},
				{
					Type: "orgPersonType",
					NotAnyOf: []string{
						"Contractor",
						"Guest",
					},
				},
			},
		},
	},
}

var MappingUpdated = federation.Mapping{
	ID: "ACME",
	Links: map[string]interface{}{
		"self": "http://example.com/identity/v3/OS-FEDERATION/mappings/ACME",
	},
	Rules: []federation.MappingRule{
		{
			Local: []federation.RuleLocal{
				{
					User: &federation.RuleUser{
						Name: "{0}",
					},
				},
				{
					Group: &federation.Group{
						ID: "0cd5e9",
					},
				},
			},
			Remote: []federation.RuleRemote{
				{
					Type: "UserName",
				},
				{
					Type: "orgPersonType",
					AnyOneOf: []string{
						"Contractor",
						"SubContractor",
					},
				},
			},
		},
	},
}

// ExpectedMappingsSlice is the slice of mappings expected to be returned from ListOutput.
var ExpectedMappingsSlice = []federation.Mapping{TestMapping}

// HandleListMappingsSuccessfully creates an HTTP handler at `/mappings` on the
// test handler mux that responds with a list of two mappings.
func HandleListMappingsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/OS-FEDERATION/mappings", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ListOutput)
	})
}

// HandleCreateMappingSuccessfully creates an HTTP handler at `/mappings` on the
// test handler mux that tests mapping creation.
func HandleCreateMappingSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/OS-FEDERATION/mappings/ACME", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateRequest)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, CreateOutput)
	})
}

// HandleUpdateMappingSuccessfully creates an HTTP handler at `/mappings` on the
// test handler mux that tests mapping update.
func HandleUpdateMappingSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/OS-FEDERATION/mappings/ACME", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, UpdateRequest)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, UpdateOutput)
	})
}

// HandleGetMappingSuccessfully creates an HTTP handler at `/mappings` on the
// test handler mux that responds with a single mapping.
func HandleGetMappingSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/OS-FEDERATION/mappings/ACME", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, GetOutput)
	})
}

// HandleDeleteMappingSuccessfully creates an HTTP handler at `/mappings` on the
// test handler mux that tests mapping deletion.
func HandleDeleteMappingSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/OS-FEDERATION/mappings/ACME", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}
