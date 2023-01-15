package nba

type GamesByDateRequest struct {
	Date string `uri:"date" binding:"required"`
}

type PlayersByTeamRequest struct {
	TeamName string `uri:"team" binding:"required"`
}

type PlayerDetailsByNameRequest struct {
	FirstName string `uri:"firstname" binding:"required"`
	LastName  string `uri:"lastname" binding:"required"`
}

type Game struct {
	GameID                       int     `json:"GameID"`
	Season                       int     `json:"Season"`
	SeasonType                   int     `json:"SeasonType"`
	Status                       string  `json:"Status"`
	Day                          string  `json:"Day"`
	DateTime                     string  `json:"DateTime"`
	AwayTeam                     string  `json:"AwayTeam"`
	HomeTeam                     string  `json:"HomeTeam"`
	AwayTeamID                   int     `json:"AwayTeamID"`
	HomeTeamID                   int     `json:"HomeTeamID"`
	StadiumID                    int     `json:"StadiumID"`
	Channel                      string  `json:"Channel"`
	Attendance                   int     `json:"Attendance"`
	AwayTeamScore                int     `json:"AwayTeamScore"`
	HomeTeamScore                int     `json:"HomeTeamScore"`
	Updated                      string  `json:"Updated"`
	Quarter                      string  `json:"Quarter"`
	TimeRemainingMinutes         int     `json:"TimeRemainingMinutes"`
	TimeRemainingSeconds         int     `json:"TimeRemainingSeconds"`
	PointSpread                  float32 `json:"PointSpread"`
	OverUnder                    float32 `json:"OverUnder"`
	AwayTeamMoneyLine            int     `json:"AwayTeamMoneyLine"`
	HomeTeamMoneyLine            int     `json:"HomeTeamMoneyLine"`
	GlobalGameID                 int     `json:"GlobalGameID"`
	GlobalAwayTeamID             int     `json:"GlobalAwayTeamID"`
	GlobalHomeTeamID             int     `json:"GlobalHomeTeamID"`
	PointSpreadAwayTeamMoneyLine int     `json:"PointSpreadAwayTeamMoneyLine"`
	PointSpreadHomeTeamMoneyLine int     `json:"PointSpreadHomeTeamMoneyLine"`
	LastPlay                     string  `json:"LastPlay"`
	IsClosed                     bool    `json:"IsClosed"`
	GameEndDateTime              string  `json:"GameEndDateTime"`
	HomeRotationNumber           int     `json:"HomeRotationNumber"`
	AwayRotationNumber           int     `json:"AwayRotationNumber"`
	NeutralVenue                 bool    `json:"NeutralVenue"`
	OverPayout                   float32 `json:"OverPayout"`
	UnderPayout                  float32 `json:"UnderPayout"`
	CrewChiefID                  int     `json:"CrewChiefID"`
	UmpireID                     int     `json:"UmpireID"`
	RefereeID                    int     `json:"RefereeID"`
	AlternateID                  int     `json:"AlternateID"`
	DateTimeUTC                  string  `json:"DateTimeUTC"`
	SeriesInfo                   string  `json:"SeriesInfo"`
}

type Team struct {
	TeamID               int    `json:"TeamID"`
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

type Player struct {
	PlayerID                            int    `json:"PlayerID"`
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
