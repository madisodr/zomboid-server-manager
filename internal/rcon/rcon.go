package rcon

import (
	"fmt"
	"log"
	"time"

	"github.com/gorcon/rcon"

	"zomboid-server-tool/internal/utils"
)

type rconConnection struct {
	conn *rcon.Conn
}

func NewRconConnection(host string, port int, password string) (*rconConnection, error) {
	conn, err := rcon.Dial(fmt.Sprintf("%s:%d", host, port), password)
	if err != nil {
		return nil, fmt.Errorf("failed to create rcon connection: %w", err)
	}

	return &rconConnection{conn: conn}, nil
}

func (rc *rconConnection) ServerShutdown(rconConnection *rconConnection, shutdownDuration time.Duration) {
	log.Println("Update required. Beginning shutdown process.")

	shutdownTimer := utils.NewSecondsTimer(shutdownDuration)
	ticker := time.NewTicker(15 * time.Second)
	defer shutdownTimer.Stop()
	defer ticker.Stop()

	rc.SendMessage(fmt.Sprintf("Server will shut down in %v minutes", shutdownDuration.Round(time.Minute)))

	for {
		select {
		case <-shutdownTimer.Timer.C:
			err := rconConnection.saveAndShutdown()
			if err != nil {
				log.Fatal(err)
			}

			return
		case <-ticker.C:
			remaining := shutdownTimer.TimeRemaining()
			if remaining <= 1*time.Minute {
				ticker.Reset(5 * time.Second)
			}
			if remaining <= 10*time.Second {
				ticker.Reset(1 * time.Second)
			}

			err := rc.SendMessage(fmt.Sprintf("Server will shut down in %v", remaining.Round(time.Second)))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (rc *rconConnection) saveAndShutdown() error {
	_, err := rc.conn.Execute("save")
	if err != nil {
		return err
	}

	log.Println("Saved server state")
	_, err = rc.conn.Execute("quit")
	if err != nil {
		return err
	}

	log.Println("Exiting...")
	return nil
}

func (rc *rconConnection) Close() {
	rc.conn.Close()
}

func (rc *rconConnection) SendMessage(message string) error {
	// any message sent to the server is logged to the console as well
	log.Printf("Sent to Server: %s", message)

	resp, err := rc.conn.Execute(fmt.Sprintf("servermsg \"%s\"", message))
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	log.Println(resp)

	return nil
}
