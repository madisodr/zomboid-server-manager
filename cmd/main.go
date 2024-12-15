package main

import (
	"log"
	"strings"
	"time"

	"github.com/ilyakaznacheev/cleanenv"

	"zomboid-server-tool/internal/rcon"
	"zomboid-server-tool/internal/steam"
	"zomboid-server-tool/internal/zomboid"
)

type config struct {
	SteamAPIKey   string `env:"STEAM_API_KEY"`
	ACFFile       string `env:"WORKSHOP_ACF_FILE"`
	ServerIniFile string `env:"SERVER_INI"`

	ModNames      []string `env:"MOD_NAMES" env-separator:";"`
	WorkshopItems []string `env:"WORKSHOP_ITEMS" env-separator:";"`
	Maps          []string `env:"MAPS" env-separator:";"`

	CheckWorkshopInterval time.Duration `env:"CHECK_WORKSHOP_INTERVAL"`
	ShutdownDuration      time.Duration `env:"SHUTDOWN_DURATION"`

	RconHost     string `env:"RCON_HOST" env-default:"localhost"`
	RconPassword string `env:"RCON_PASSWORD"`
	RconPort     int    `env:"RCON_PORT" env-default:"27015"`
}

var cfg config

func main() {
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err)
	}

	log.Println("Parsed Configuration")

	zomboidCfg, err := zomboid.LoadServerConfig(cfg.ServerIniFile)
	if err != nil {
		log.Fatal(err)
	}

	// turn this on when im ready for the values in .env to be the source of truth
	zomboidCfg.Mods = strings.Join(cfg.ModNames, ";")
	zomboidCfg.WorkshopItems = strings.Join(cfg.WorkshopItems, ";")

	//log.Printf("Parsed Zomboid Ini: %s\n", zomboidCfg.Mods)
	//log.Printf("Mods from config.env %s\n", cfg.ModNames)

	zomboidCfg.SaveServerConfig(cfg.ServerIniFile)

	// connect to the RCON server
	rconConnection, err := rcon.NewRconConnection(cfg.RconHost, cfg.RconPort, cfg.RconPassword)
	if err != nil {
		log.Fatal(err)
	}
	defer rconConnection.Close()

	// every CheckWorkshopInterval, check for updates
	ticker := time.NewTicker(cfg.CheckWorkshopInterval)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		log.Println("Checking for mod updates.")
		modUpdateCount := steam.CheckForUpdates(cfg.SteamAPIKey, cfg.WorkshopItems, cfg.ACFFile)
		if modUpdateCount != 0 {
			log.Printf("%d mods need updating.\n", modUpdateCount)
			// rconConnection.ServerShutdown(cfg.ShutdownDuration)
		} else {
			log.Println("No updates found.")
		}
	}
}
