package block

type Player struct {
	Id       uint32
	Category uint16
	Points   uint32

	MatchPoints   uint32
	MatchesPlayed uint16
	Victories     uint16
	Defeats       uint16
	Draws         uint16

	WinningStreak  uint16
	BestStreak     uint16
	Disconnections uint16
	Division       byte

	Teams         [5]uint16
	GoalsScored   uint32
	GoalsReceived uint32
	TimePlayed    uint32
	LastLogin     uint32
	Position      uint32
	OldCategory   uint32
	OldPoints     uint32

	Name     string
	Comment  string
	Lang     uint16
	Settings []byte
	LoggedIn bool
	Admin    int

	RoomId     uint32
	GameStatus byte

	GroupId     uint32
	GroupName   string
	GroupStatus byte

	RoomData *PlayerRoomData

	Link *Link
}

type PlayerRoomData struct {
	Team          byte
	Spectator     byte
	Participation byte
}

type Link struct {
	Ip1   string
	Port1 uint16
	Ip2   string
	Port2 uint16
}

type PlayerInternal struct {
	Id            uint32
	Name          [48]byte
	GroupId       uint32
	GroupName     [48]byte
	GroupStatus   byte
	Division      byte
	RoomId        uint32 // Not clear whether this could be the room id, we will fill it with ff's
	Points        uint32
	Category      uint16
	MatchesPlayed uint16
	Victories     uint16
	Defeats       uint16
	Draws         uint16
	Lang          uint16
	GameStatus    byte // 0 = idle, 1 = competition
}

func (info Player) buildInternal() PieceInternal {
	var internal PlayerInternal
	internal.Id = info.Id
	copy(internal.Name[:], info.Name)
	internal.GroupId = info.GroupId
	copy(internal.GroupName[:], info.GroupName)
	internal.GroupStatus = info.GroupStatus
	internal.Division = info.Division
	internal.RoomId = info.RoomId
	internal.Points = info.Points
	internal.Category = info.Category
	internal.MatchesPlayed = info.MatchesPlayed
	internal.Victories = info.Victories
	internal.Defeats = info.Defeats
	internal.Draws = info.Draws
	internal.Lang = info.Lang
	internal.GameStatus = info.GameStatus
	return internal
}

func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
	}
}

func (info *Player) ResetRoomData() {
	info.RoomData = &PlayerRoomData{
		Team:          0xff,
		Spectator:     0,
		Participation: 0xff,
	}
}

func (info *Player) isParticipating() bool {
	return info.RoomData.Participation != 0xff
}
