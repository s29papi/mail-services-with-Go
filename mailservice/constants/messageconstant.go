package constants

import "net/mail"

const (
	// message file: index.html
	First = "/nonparsedmessage/index.html"
	
)

var (
	// message sender
	From = mail.Address{"", ""}
	// message sender password
	Password = ""
	// message header
	


)







func MessageFile(name string) string {  return "message" + name  }

