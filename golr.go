package golr

import (
	log "github.com/cihub/seelog"

	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/andresteingress/golr/api"
)

var URL string = "http://localhost:8983/solr/collection1"

func Search(sp *api.SearchParams) (*api.SolrResponse, error) {
	if sp == nil {
		panic("'sp' Parameter must not be nil!")
	}

	// enable JSON response
	sp.Wt = "json"

	values := structToMap(sp)
	response, err := http.PostForm(URL+"/select", values);

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	log.Debug("JSON received: " + string(data))

	var result api.SolrResponse

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}


