package team

type Team struct {
	TeamID               int    `json:"TeamID" gorm:"primaryKey"`
	Key                  string `json:"Key"`
	Active               bool   `json:"Active"`
	City                 string `json:"City"`
	Name                 string `json:"Name"`
	LeagueID             int    `json:"LeagueID"`
	StadiumID            int    `json:"StadiumID"`
	Conference           string `json:"Conference"`
	Division             string `json:"Division"`
	PrimaryColor         string `json:"PrimaryColor"`
	SecondaryColor       string `json:"SecondaryColor"`
	TertiaryColor        string `json:"TertiaryColor"`
	QuaternaryColor      string `json:"QuaternaryColor"`
	WikipediaLogoUrl     string `json:"WikipediaLogoUrl"`
	WikipediaWordMarkUrl string `json:"WikipediaWordMarkUrl"`
	GlobalTeamID         int    `json:"GlobalTeamID"`
	NbaDotComTeamID      int    `json:"NbaDotComTeamID"`
}
