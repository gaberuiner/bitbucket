package events

const msgHelp = `I can save and keep your data. Also I can offer you to read them.

In order to save the data, just send me data type(link, film, book, image) and link or name of it. 

As Exmaple:"book Lord of the Rings"

In order to get a random file from your list choose file type from button list.
Images URL must be in ".jpg", ".jpeg", ".png", ".gif", ".webp" format, otherwise bot will not accept it
Press /start to open buttons`

const (
	msgHello          = "Hello world! 🙃"
	msgUnknownCommand = "Unknown command 🤔"
	msgNoSavedFiles   = "You have no saved pages 🙊"
	msgSaved          = "Saved! 👌"
	msgAlreadyExists  = "You have already have this page in your list 🤗"
	msgUnknownType    = "Unknown type 🤔"
	msgIsNotLink      = "You need to send URL 😘"
	msgIsNotImageURL  = "You need to send image URL 😘"
)
