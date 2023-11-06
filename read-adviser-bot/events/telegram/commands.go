package telegram

var (
	commandsWarnings = [...]string{
		"/disablewarnings",
		"/enablewarnings",
		"/w-",
		"/w+",
		"/dw",
		"/ew",
	}

	commandsShortcuts = []string{
		"/rnd",
		// "/rndset %d",
		"/h",
		"/cmd",
		"/g",
		// "/l",
		// "/t",
		// "/sett",
		// "/rm %s",
		"/w-",
		"/w+",
		"/dw",
		"/ew",
	}
	commandsFull = []string{
		"/random",
		// "randomset %d",
		"/help",
		"/commands",
		"/greet",
		"/start",
		"/stop",
		// "/list",
		// "/time",
		// "/settime",
		// "/remove %s",
		"/disablewarnings",
		"/enablewarnings",
		"/enableshortcuts",
		"/disableshortcuts",
		"IAmAbsolutelySureIWantToDeletemyHistory",
	}
	CommandsAll     = append(commandsFull, commandsShortcuts...)
	MutableCommands = make(map[string]bool, len(CommandsAll))
)

func DisableMostCommandsOnBotStop() {
	for _, cmd := range commandsFull {
		MutableCommands[cmd] = false
	}
	MutableCommands["/start"] = true
	MutableCommands["/time"] = true
}

func InitMutableCommands() {
	MutableCommands["/rndset %d"] = false
	MutableCommands["/rnd"] = false
	MutableCommands["/random"] = false
	MutableCommands["/randomset %d"] = false
	MutableCommands["/enablewarnings"] = false
	MutableCommands["/ew"] = false
	MutableCommands["/w+"] = false
	MutableCommands["/disablewarnings"] = true
	MutableCommands["/dw"] = true
	MutableCommands["/w-"] = true
	MutableCommands["/enableshortcuts"] = false
	MutableCommands["/disableshortcuts"] = true
	MutableCommands["/help"] = true
	MutableCommands["/h"] = true
	MutableCommands["/commands"] = true
	MutableCommands["/cmd"] = true
	MutableCommands["/start"] = true
	MutableCommands["/stop"] = false
	MutableCommands["/list"] = false
	MutableCommands["/l"] = false
	MutableCommands["/time"] = false
	MutableCommands["/t"] = false
	MutableCommands["/remove %s"] = false
	MutableCommands["/rm %s"] = false
	MutableCommands["/greet"] = true
	MutableCommands["/g"] = true
	MutableCommands["IAmAbsolutelySureIWantToDeletemyHistory"] = false
}

type Commands struct {
	CommandsMap map[string]bool
}

func NewCommands() *Commands {
	return &Commands{
		CommandsMap: make(map[string]bool, len(CommandsAll)),
	}
}
