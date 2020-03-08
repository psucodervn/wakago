package model

import "time"

type SummaryData struct {
	Categories []struct {
		Digital      string  `json:"digital"`
		Hours        int     `json:"hours"`
		Minutes      int     `json:"minutes"`
		Name         string  `json:"name"`
		Percent      float64 `json:"percent"`
		Seconds      int     `json:"seconds"`
		Text         string  `json:"text"`
		TotalSeconds float64 `json:"total_seconds"`
	} `json:"categories"`
	Dependencies []interface{} `json:"dependencies"`
	Editors      []struct {
		Digital      string  `json:"digital"`
		Hours        int     `json:"hours"`
		Minutes      int     `json:"minutes"`
		Name         string  `json:"name"`
		Percent      float64 `json:"percent"`
		Seconds      int     `json:"seconds"`
		Text         string  `json:"text"`
		TotalSeconds float64 `json:"total_seconds"`
	} `json:"editors"`
	GrandTotal struct {
		Digital      string  `json:"digital"`
		Hours        int     `json:"hours"`
		Minutes      int     `json:"minutes"`
		Text         string  `json:"text"`
		TotalSeconds float64 `json:"total_seconds"`
	} `json:"grand_total"`
	Languages []struct {
		Digital      string  `json:"digital"`
		Hours        int     `json:"hours"`
		Minutes      int     `json:"minutes"`
		Name         string  `json:"name"`
		Percent      float64 `json:"percent"`
		Seconds      int     `json:"seconds"`
		Text         string  `json:"text"`
		TotalSeconds float64 `json:"total_seconds"`
	} `json:"languages"`
	Machines []struct {
		Digital       string  `json:"digital"`
		Hours         int     `json:"hours"`
		MachineNameID string  `json:"machine_name_id"`
		Minutes       int     `json:"minutes"`
		Name          string  `json:"name"`
		Percent       float64 `json:"percent"`
		Seconds       int     `json:"seconds"`
		Text          string  `json:"text"`
		TotalSeconds  float64 `json:"total_seconds"`
	} `json:"machines"`
	OperatingSystems []struct {
		Digital      string  `json:"digital"`
		Hours        int     `json:"hours"`
		Minutes      int     `json:"minutes"`
		Name         string  `json:"name"`
		Percent      float64 `json:"percent"`
		Seconds      int     `json:"seconds"`
		Text         string  `json:"text"`
		TotalSeconds float64 `json:"total_seconds"`
	} `json:"operating_systems"`
	Projects []struct {
		Digital      string  `json:"digital"`
		Hours        int     `json:"hours"`
		Minutes      int     `json:"minutes"`
		Name         string  `json:"name"`
		Percent      float64 `json:"percent"`
		Seconds      int     `json:"seconds"`
		Text         string  `json:"text"`
		TotalSeconds float64 `json:"total_seconds"`
	} `json:"projects"`
	Range struct {
		Date     string    `json:"date"`
		End      time.Time `json:"end"`
		Start    time.Time `json:"start"`
		Text     string    `json:"text"`
		Timezone string    `json:"timezone"`
	} `json:"range"`
}
