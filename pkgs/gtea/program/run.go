package program

import (
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/moqsien/goutils/pkgs/gtea/gprint"
)

const (
	BubbleteaProgramSleepEnv string = "BUBBLETEA_PROGRAM_SLEEP"
)

func Run(app *tea.Program) {
	if _, err := app.Run(); err != nil {
		gprint.PrintError("%+v", err)
	}
	sleepStr := os.Getenv(BubbleteaProgramSleepEnv)
	sleepInt := gconv.Int(sleepStr)
	// Wait for bubbletea to exit.
	if sleepInt > 0 {
		time.Sleep(time.Duration(sleepInt) * time.Second)
	} else {
		time.Sleep(time.Second)
	}
}
