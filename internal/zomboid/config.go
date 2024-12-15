package zomboid

import "zomboid-server-tool/internal/utils"

type ZomboidConfig struct {
	PVP                       bool   `ini:"PVP"`
	PauseEmpty                bool   `ini:"PauseEmpty"`
	GlobalChat                bool   `ini:"GlobalChat"`
	ChatStreams               string `ini:"ChatStreams"`
	Open                      bool   `ini:"Open"`
	ServerWelcomeMessage      string `ini:"ServerWelcomeMessage"`
	AutoCreateUserInWhiteList bool   `ini:"AutoCreateUserInWhiteList"`

	DisplayUserName      bool `ini:"DisplayUserName"`
	ShowFirstAndLastName bool `ini:"ShowFirstAndLastName"`

	SpawnPoint          string `ini:"SpawnPoint"`
	SafetySystem        bool   `ini:"SafetySystem"`
	ShowSafety          bool   `ini:"ShowSafety"`
	SafetyToggleTimer   int    `ini:"SafetyToggleTimer"`
	SafetyCooldownTimer int    `ini:"SafetyCooldownTimer"`

	SpawnItems string `ini:"SpawnItems"`

	DefaultPort                 int    `ini:"DefaultPort"`
	UDPPort                     int    `ini:"UDPPort"`
	ResetID                     int    `ini:"ResetID"`
	Mods                        string `ini:"Mods"`
	Map                         string `ini:"Map"`
	DoLuaChecksum               bool   `ini:"DoLuaChecksum"`
	DenyLoginOnOverloadedServer bool   `ini:"DenyLoginOnOverloadedServer"`

	Public            bool   `ini:"Public"`
	PublicName        string `ini:"PublicName"`
	PublicDescription string `ini:"PublicDescription"`
	MaxPlayers        int    `ini:"MaxPlayers"`
	PingLimit         int    `ini:"PingLimit"`

	HoursForLootRespawn             int  `ini:"HoursForLootRespawn"`
	MaxItemsForLootRespawn          int  `ini:"MaxItemsForLootRespawn"`
	ConstructionPreventsLootRespawn bool `ini:"ConstructionPreventsLootRespawn"`

	DropOffWhiteListAfterDeath bool    `ini:"DropOffWhiteListAfterDeath"`
	NoFire                     bool    `ini:"NoFire"`
	AnnounceDeath              bool    `ini:"AnnounceDeath"`
	MinutesPerPage             float64 `ini:"MinutesPerPage"`
	SaveWorldEveryMinutes      int     `ini:"SaveWorldEveryMinutes"`

	PlayerSafehouse                bool `ini:"PlayerSafehouse"`
	AdminSafehouse                 bool `ini:"AdminSafehouse"`
	SafehouseAllowTrepass          bool `ini:"SafehouseAllowTrepass"`
	SafehouseAllowFire             bool `ini:"SafehouseAllowFire"`
	SafehouseAllowLoot             bool `ini:"SafehouseAllowLoot"`
	SafehouseAllowRespawn          bool `ini:"SafehouseAllowRespawn"`
	SafehouseDaySurvivedToClaim    int  `ini:"SafehouseDaySurvivedToClaim"`
	SafeHouseRemovalTime           int  `ini:"SafeHouseRemovalTime"`
	SafehouseAllowNonResidential   bool `ini:"SafehouseAllowNonResidential"`
	AllowDestructionBySledgehammer bool `ini:"AllowDestructionBySledgehammer"`
	SledgehammerOnlyInSafehouse    bool `ini:"SledgehammerOnlyInSafehouse"`

	KickFastPlayers bool `ini:"KickFastPlayers"`
	ServerPlayerID  int  `ini:"ServerPlayerID"`

	RCONPort     int    `ini:"RCONPort"`
	RCONPassword string `ini:"RCONPassword"`

	DiscordEnable    bool   `ini:"DiscordEnable"`
	DiscordToken     string `ini:"DiscordToken"`
	DiscordChannel   string `ini:"DiscordChannel"`
	DiscordChannelID string `ini:"DiscordChannelID"`

	Password                      string  `ini:"Password"`
	MaxAccountsPerUser            int     `ini:"MaxAccountsPerUser"`
	AllowCoop                     bool    `ini:"AllowCoop"`
	SleepAllowed                  bool    `ini:"SleepAllowed"`
	SleepNeeded                   bool    `ini:"SleepNeeded"`
	KnockedDownAllowed            bool    `ini:"KnockedDownAllowed"`
	SneakModeHideFromOtherPlayers bool    `ini:"SneakModeHideFromOtherPlayers"`
	WorkshopItems                 string  `ini:"WorkshopItems"`
	SteamScoreboard               bool    `ini:"SteamScoreboard"`
	SteamVAC                      bool    `ini:"SteamVAC"`
	UPnP                          bool    `ini:"UPnP"`
	VoiceEnable                   bool    `ini:"VoiceEnable"`
	VoiceMinDistance              float64 `ini:"VoiceMinDistance"`
	VoiceMaxDistance              float64 `ini:"VoiceMaxDistance"`
	Voice3D                       bool    `ini:"Voice3D"`
	SpeedLimit                    float64 `ini:"SpeedLimit"`

	LoginQueueEnabled        bool `ini:"LoginQueueEnabled"`
	LoginQueueConnectTimeout int  `ini:"LoginQueueConnectTimeout"`

	ServerBrowserAnnouncedIP            string  `ini:"server_browser_announced_ip"`
	PlayerRespawnWithSelf               bool    `ini:"PlayerRespawnWithSelf"`
	PlayerRespawnWithOther              bool    `ini:"PlayerRespawnWithOther"`
	FastForwardMultiplier               float64 `ini:"FastForwardMultiplier"`
	DisableSafehouseWhenPlayerConnected bool    `ini:"DisableSafehouseWhenPlayerConnected"`

	Faction                      bool `ini:"Faction"`
	FactionDaySurvivedToCreate   int  `ini:"FactionDaySurvivedToCreate"`
	FactionPlayersRequiredForTag int  `ini:"FactionPlayersRequiredForTag"`

	DisableRadioStaff     bool `ini:"DisableRadioStaff"`
	DisableRadioAdmin     bool `ini:"DisableRadioAdmin"`
	DisableRadioGM        bool `ini:"DisableRadioGM"`
	DisableRadioOverseer  bool `ini:"DisableRadioOverseer"`
	DisableRadioModerator bool `ini:"DisableRadioModerator"`
	DisableRadioInvisible bool `ini:"DisableRadioInvisible"`

	ClientCommandFilter                string  `ini:"ClientCommandFilter"`
	ClientActionLogs                   string  `ini:"ClientActionLogs"`
	PerkLogs                           bool    `ini:"PerkLogs"`
	ItemNumbersLimitPerContainer       int     `ini:"ItemNumbersLimitPerContainer"`
	BloodSplatLifespanDays             int     `ini:"BloodSplatLifespanDays"`
	AllowNonAsciiUsername              bool    `ini:"AllowNonAsciiUsername"`
	BanKickGlobalSound                 bool    `ini:"BanKickGlobalSound"`
	RemovePlayerCorpsesOnCorpseRemoval bool    `ini:"RemovePlayerCorpsesOnCorpseRemoval"`
	TrashDeleteAll                     bool    `ini:"TrashDeleteAll"`
	PVPMeleeWhileHitReaction           bool    `ini:"PVPMeleeWhileHitReaction"`
	MouseOverToSeeDisplayName          bool    `ini:"MouseOverToSeeDisplayName"`
	HidePlayersBehindYou               bool    `ini:"HidePlayersBehindYou"`
	PVPMeleeDamageModifier             float64 `ini:"PVPMeleeDamageModifier"`
	PVPFirearmDamageModifier           float64 `ini:"PVPFirearmDamageModifier"`
	CarEngineAttractionModifier        float64 `ini:"CarEngineAttractionModifier"`
	PlayerBumpPlayer                   bool    `ini:"PlayerBumpPlayer"`
	MapRemotePlayerVisibility          int     `ini:"MapRemotePlayerVisibility"`

	BackupsCount           int  `ini:"BackupsCount"`
	BackupsOnStart         bool `ini:"BackupsOnStart"`
	BackupsOnVersionChange bool `ini:"BackupsOnVersionChange"`
	BackupsPeriod          int  `ini:"BackupsPeriod"`

	AntiCheatProtectionType1  bool `ini:"AntiCheatProtectionType1"`
	AntiCheatProtectionType2  bool `ini:"AntiCheatProtectionType2"`
	AntiCheatProtectionType3  bool `ini:"AntiCheatProtectionType3"`
	AntiCheatProtectionType4  bool `ini:"AntiCheatProtectionType4"`
	AntiCheatProtectionType5  bool `ini:"AntiCheatProtectionType5"`
	AntiCheatProtectionType6  bool `ini:"AntiCheatProtectionType6"`
	AntiCheatProtectionType7  bool `ini:"AntiCheatProtectionType7"`
	AntiCheatProtectionType8  bool `ini:"AntiCheatProtectionType8"`
	AntiCheatProtectionType9  bool `ini:"AntiCheatProtectionType9"`
	AntiCheatProtectionType10 bool `ini:"AntiCheatProtectionType10"`
	AntiCheatProtectionType11 bool `ini:"AntiCheatProtectionType11"`
	AntiCheatProtectionType12 bool `ini:"AntiCheatProtectionType12"`
	AntiCheatProtectionType13 bool `ini:"AntiCheatProtectionType13"`
	AntiCheatProtectionType14 bool `ini:"AntiCheatProtectionType14"`
	AntiCheatProtectionType15 bool `ini:"AntiCheatProtectionType15"`
	AntiCheatProtectionType16 bool `ini:"AntiCheatProtectionType16"`
	AntiCheatProtectionType17 bool `ini:"AntiCheatProtectionType17"`
	AntiCheatProtectionType18 bool `ini:"AntiCheatProtectionType18"`
	AntiCheatProtectionType19 bool `ini:"AntiCheatProtectionType19"`
	AntiCheatProtectionType20 bool `ini:"AntiCheatProtectionType20"`
	AntiCheatProtectionType21 bool `ini:"AntiCheatProtectionType21"`
	AntiCheatProtectionType22 bool `ini:"AntiCheatProtectionType22"`
	AntiCheatProtectionType23 bool `ini:"AntiCheatProtectionType23"`
	AntiCheatProtectionType24 bool `ini:"AntiCheatProtectionType24"`

	AntiCheatProtectionType2ThresholdMultiplier  float64 `ini:"AntiCheatProtectionType2ThresholdMultiplier"`
	AntiCheatProtectionType3ThresholdMultiplier  float64 `ini:"AntiCheatProtectionType3ThresholdMultiplier"`
	AntiCheatProtectionType4ThresholdMultiplier  float64 `ini:"AntiCheatProtectionType4ThresholdMultiplier"`
	AntiCheatProtectionType9ThresholdMultiplier  float64 `ini:"AntiCheatProtectionType9ThresholdMultiplier"`
	AntiCheatProtectionType15ThresholdMultiplier float64 `ini:"AntiCheatProtectionType15ThresholdMultiplier"`
	AntiCheatProtectionType20ThresholdMultiplier float64 `ini:"AntiCheatProtectionType20ThresholdMultiplier"`
	AntiCheatProtectionType22ThresholdMultiplier float64 `ini:"AntiCheatProtectionType22ThresholdMultiplier"`
	AntiCheatProtectionType24ThresholdMultiplier float64 `ini:"AntiCheatProtectionType24ThresholdMultiplier"`
}

func LoadServerConfig(filename string) (*ZomboidConfig, error) {
	var zc ZomboidConfig

	err := utils.LoadIni(filename, &zc)
	if err != nil {
		return nil, err
	}

	return &zc, nil
}

func (zc *ZomboidConfig) SaveServerConfig(filename string) error {
	return utils.SaveIni(filename, zc)
}
