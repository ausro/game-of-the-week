package command

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ausro/game-of-the-week/util"
)

const ()

type Command struct {
	command string
	args    []string
}

func (c *CommandHandler) run() error {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		in := util.Must(r.ReadString('\n'))
		in = strings.TrimSuffix(in, "\n")
		in = strings.TrimSuffix(in, "\r")

		commandString := strings.Split(in, " ")
		command := &Command{
			command: commandString[0],
			args:    commandString[1:],
		}

		c.execute(command)
	}
}

func (c *CommandHandler) execute(in *Command) error {
	var appId int
	if len(in.args) > 0 {
		appId = util.Must(strconv.Atoi(in.args[0]))

		if appId == 0 {
			return nil
		}
	}

	switch in.command {
	case "add":
		c.appHandler.AddGameById(context.Background(), appId)
	case "remove":
		c.appHandler.DeleteGameById(context.Background(), appId)
	case "promote":
		app := c.appHandler.GetAppById(appId)
		if app == nil {
			log.Printf("App %d not found on list", appId)
			break
		}
		app.Promoted = !app.Promoted
		c.appHandler.Add(context.Background(), app)
	case "list":
		for _, app := range c.appHandler.ListApps(context.Background()) {
			log.Printf("%s, ID: %d\n", app.Name, app.ID)
		}
	case "fetch":
		c.appHandler.RequestGames(context.Background())
	case "exit":
		log.Fatalln("Exit called by user")
	default:
		log.Printf("Unknown command.\n")
	}

	return nil
}
