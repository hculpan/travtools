package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hculpan/travtools/pkg/util"
)

type SystemSearchResponse struct {
	Results SystemSearchResults `json:"Results"`
}

type SystemSearchResults struct {
	Count int                `json:"Count"`
	Items []SystemSearchItem `json:"Items"`
}

type SystemSearchItem struct {
	Info SystemSearchInfo `json:"World"`
}

type SystemSearchInfo struct {
	HexX       int    `json:"HexX"`
	HexY       int    `json:"HexY"`
	Sector     string `json:"Sector"`
	Uwp        string `json:"Uwp"`
	Name       string `json:"Name"`
	SectorTags string `json:"SectorTags"`
}

func SearchSystemsHandler(w http.ResponseWriter, r *http.Request) {
	result := ""

	r.ParseForm()
	query := r.FormValue("systemName")
	if len(strings.Trim(query, " ")) > 0 {
		data, err := util.FetchData("https://travellermap.com/api/search?milieu=M1105&q=" + query)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		var response SystemSearchResponse
		err = json.Unmarshal(data, &response)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		for _, r := range response.Results.Items {
			if len(r.Info.Uwp) > 0 {
				result += fmt.Sprintf("<option class=\"list-group-item\" value=\"%s\">%s : %s : %s : %s</option>",
					r.Info.Name, r.Info.Name, r.Info.Sector, fmt.Sprintf("%02d%02d", r.Info.HexX, r.Info.HexY), r.Info.Uwp)
			}
		}
	}

	w.Write([]byte(result))
}
