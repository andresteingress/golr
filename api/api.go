package api

// SearchParams contains all parameters for a Solr "select" search request. Parameter
// Q is necessary.
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

// SolrResponse contains the search results from a "select" search request.
type SolrResponse struct {
	Results SolrResults `json:"response"`
}

// SolrResults is part of SolrResponse and contains the actual Solr documents and the number of total results.
type SolrResults struct {
	TotalResults int            `json:"numFound"`
	Start        int            `json:"start"`
	Documents    []map[string]interface{} `json:"docs"`
}
