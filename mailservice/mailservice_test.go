package mailservice

import (
	// "fmt"
	"services/mailservice/message"
	"services/mailservice/constants"
	"testing"
) 
      

// func TestEstablishConnection(t *testing.T)  {

// 	_, err := EstablishConn()
	
// 	if err !=  nil  {
// 		// fmt.Printf("Value: %+v\n", mailservice.tlsconnection) 

// 		t.Error("There was an error while establishing connection with the server")
// 	}
// }

// func TestAuthentication(t *testing.T)  {

// 	mailservice, _ := EstablishConn()
	
// 	if mailservice.Auth() !=  nil  {
// 		// fmt.Printf("Value: %+v\n", mailservice.tlsconnection) 

// 		t.Error("There was an error while Authenticating")
// 	}
// }

// Text message
func TestSendMail(t *testing.T)  {



	msgval := message.MessageValue{
					Filename: constants.MessageFile(constants.First),
					HV: message.HeaderValue{
							From: constants.From.Address,
							To: "usih.anselm@gmail.com",
							Subj: "Testing new Code",
							ContentType: "text/html; charset=UTF-8",
							MimeVersion: "1.0",
					},
	       }
	msg, _ := message.NonParsedMessage(msgval)
	
    mailservice, _ := EstablishConn()

	mailservice.Auth()
	
	mailservice.SendMail(msgval.HV.To, msg)
	
	// fmt.Println(msg)
	// if mailservice.SendMail(msg.HV.To) !=  nil  {
	// 	// fmt.Printf("Value: %+v\n", mailservice.tlsconnection) 

	// 	t.Error("Mail wasn't sent")
	// }
}

// non parsed message 

// text messagehvjh