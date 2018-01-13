package constants

type Editors struct {
	Digital      string  `json:"digital"`
	Hours        uint    `json:"hours"`
	Minutes      uint    `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Seconds      uint    `json:"seconds"`
	Text         string  `json:"text"`
	TotalSeconds uint    `json:"total_seconds"`
}

type Entities struct {
	Digital      string  `json:"digital"`
	Hours        uint    `json:"hours"`
	Minutes      uint    `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Seconds      uint    `json:"seconds"`
	Text         string  `json:"text"`
	TotalSeconds uint    `json:"total_seconds"`
	Type         string  `json:"type"`
}

type GrandTotal struct {
	Digital      string `json:"digital"`
	Hours        uint   `json:"hours"`
	Minutes      uint   `json:"minutes"`
	Text         string `json:"text"`
	TotalSeconds uint   `json:"total_seconds"`
}

type Languages struct {
	Digital      string  `json:"digital"`
	Hours        uint    `json:"hours"`
	Minutes      uint    `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Seconds      uint    `json:"seconds"`
	Text         string  `json:"text"`
	TotalSeconds uint    `json:"total_seconds"`
}

type OperatingSystems struct {
	Digital      string  `json:"digital"`
	Hours        uint    `json:"hours"`
	Minutes      uint    `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Seconds      uint    `json:"seconds"`
	Text         string  `json:"text"`
	TotalSeconds uint    `json:"total_seconds"`
}

type Projects struct {
	Digital      string  `json:"digital"`
	Hours        uint    `json:"hours"`
	Minutes      uint    `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Seconds      uint    `json:"seconds"`
	Text         string  `json:"text"`
	TotalSeconds uint    `json:"total_seconds"`
}

type Range struct {
	Date     string `json:"date"`
	End      string `json:"end"`
	Start    string `json:"start"`
	Text     string `json:"text"`
	Timezone string `json:"timezone"`
}

type SummaryItem struct {
	Editors          []Editors          `json:"editors"`
	Entities         []Entities         `json:"entities"`
	GrandTotal       GrandTotal         `json:"grand_total"`
	Languages        []Languages        `json:"languages"`
	OperatingSystems []OperatingSystems `json:"operate_system"`
	Projects         []Projects         `json:"projects"`
	Range            Range              `json:"range"`
}

type Summaries struct {
	Data  []SummaryItem `json:"data"`
	End   string        `json:"end"`
	Start string        `json:"start"`
}
