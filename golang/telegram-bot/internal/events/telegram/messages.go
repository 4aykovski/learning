package telegram

const msgHelp = `I can save and keep your pages. Also i can offer you them to read.

	In order to save the page, just send me link to it.

	In order to get a random page from your list, send me command /rnd.
	Caution! After that, this page will be removed from your list!`

const msgHello = "He there! \n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command"
	msgNoSavedPages   = "You have no saved pages"
	msgSaved          = "Saved!"
	msgAlreadyExists  = "You have already saved this page!"
)