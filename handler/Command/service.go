package command

import steamApp "github.com/ausro/game-of-the-week/handler/SteamApp"

type CommandHandler struct {
	appHandler *steamApp.SteamAppHandler
}

func NewCommandHandler(handler *steamApp.SteamAppHandler) {
	commandHandler := &CommandHandler{
		appHandler: handler,
	}

	go commandHandler.run()
}
