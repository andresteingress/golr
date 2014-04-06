package api

// Definition for the search parameters
type SearchParams struct {
	Q string `json:"q"`
	Sort string `json:"sort"`
	Start int `json:"start"`
	Rows int `json:"rows"`
	FQ string `json:"fq"`
	FL string `json:"fl"`
	DebugQuery string `json:"debugQuery"`
	Debug string `json:"debug"`
	ExplainOther string `json:"explainOther"`
	DefType string `json:"defType"`
	TimeAllowed int `json:"timeAllowed"`

	Wt string `json:"wt"`
}

// Definition for the search results
type SolrResponse struct {
	Results SolrResults `json:"response"`
}

type SolrResults struct {
	TotalResults int            `json:"numFound"`
	Start        int            `json:"start"`
	Documents    []map[string]interface{} `json:"docs"`
}
