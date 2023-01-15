package game

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
