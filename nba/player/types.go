package player

type Player struct {
	PlayerID                            int    `json:"PlayerID" gorm:"primaryKey"`
	SportsDataID                        string `json:"SportsDataID"`
	Status                              string `json:"Status"`
	TeamID                              int    `json:"TeamID"`
	Team                                string `json:"Team"`
	Jersey                              int    `json:"Jersey"`
	PositionCategory                    string `json:"PositionCategory"`
	Position                            string `json:"Position"`
	FirstName                           string `json:"FirstName"`
	LastName                            string `json:"LastName"`
	Height                              int    `json:"Height"`
	Weight                              int    `json:"Weight"`
	BirthDate                           string `json:"BirthDate"`
	BirthCity                           string `json:"BirthCity"`
	BirthState                          string `json:"BirthState"`
	BirthCountry                        string `json:"BirthCountry"`
	HighSchool                          string `json:"HighSchool"`
	College                             string `json:"College"`
	Salary                              int    `json:"Salary"`
	PhotoUrl                            string `json:"PhotoUrl"`
	Experience                          int    `json:"Experience"`
	SportRadarPlayerID                  string `json:"SportRadarPlayerID"`
	RotoworldPlayerID                   int    `json:"RotoworldPlayerID"`
	RotoWirePlayerID                    int    `json:"RotoWirePlayerID"`
	FantasyAlarmPlayerID                int    `json:"FantasyAlarmPlayerID"`
	StatsPlayerID                       int    `json:"StatsPlayerID"`
	SportsDirectPlayerID                int    `json:"SportsDirectPlayerID"`
	XmlTeamPlayerID                     int    `json:"XmlTeamPlayerID"`
	InjuryStatus                        string `json:"InjuryStatus"`
	InjuryBodyPart                      string `json:"InjuryBodyPart"`
	InjuryStartDate                     string `json:"InjuryStartDate"`
	InjuryNotes                         string `json:"InjuryNotes"`
	FanDuelPlayerID                     int    `json:"FanDuelPlayerID"`
	DraftKingsPlayerID                  int    `json:"DraftKingsPlayerID"`
	YahooPlayerID                       int    `json:"YahooPlayerID"`
	FanDuelName                         string `json:"FanDuelName"`
	DraftKingsName                      string `json:"DraftKingsName"`
	YahooName                           string `json:"YahooName"`
	DepthChartPosition                  string `json:"DepthChartPosition"`
	DepthChartOrder                     int    `json:"DepthChartOrder"`
	GlobalTeamID                        int    `json:"GlobalTeamID"`
	FantasyDraftName                    string `json:"FantasyDraftName"`
	FantasyDraftPlayerID                int    `json:"FantasyDraftPlayerID"`
	UsaTodayPlayerID                    int    `json:"UsaTodayPlayerID"`
	UsaTodayHeadshotUrl                 string `json:"UsaTodayHeadshotUrl"`
	UsaTodayHeadshotNoBackgroundUrl     string `json:"UsaTodayHeadshotNoBackgroundUrl"`
	UsaTodayHeadshotUpdated             string `json:"UsaTodayHeadshotUpdated"`
	UsaTodayHeadshotNoBackgroundUpdated string `json:"UsaTodayHeadshotNoBackgroundUpdated"`
	NbaDotComPlayerID                   int    `json:"NbaDotComPlayerID"`
}
