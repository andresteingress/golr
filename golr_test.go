package golr

import (
	"testing"
	"github.com/andresteingress/golr/api"
)

func TestStructToMap(t *testing.T)  {
	searchParams := api.SearchParams{
		Q: "this is a query",
		Start: 42,
	}

	result := structToMap(&searchParams)
	if result == nil {
		t.Error("Result must not be null!")
	}
}

func TestLA(t *testing.T)  {
	if lA("HALLO") != "hALLO"  {
		t.Fail()
	}
}

func TestSearchSolr(t *testing.T)  {
	if (testing.Short())  {
		t.SkipNow()
	}

	searchParams := api.SearchParams{
		Q: "solr",
	}

	result, err := Search(&searchParams)
	if err != nil {
		t.Error("Search hasn't been successful: " + err.Error())
	}

	if result == nil {
		t.Error("No result has been returned!")
	}


	if result.Results.Start != 0 {
		t.Error("Start needs to be equal to the given search parameters but is " + string(result.Results.Start))
	}

	if len(result.Results.Documents) == 0 {
		t.Error("At least one search result must have been found!")
	}

	firstDocument := result.Results.Documents[0]

	if firstDocument["id"] == nil  {
		t.Error("Document ID must not be nil!")
	}

	cats := firstDocument["cat"]
	if len(cats.([]interface{})) == 0 {
		t.Error("Inner collections must not be empty!")
	}
 }
