package message 

import (
	"fmt"
	"testing"
	"services/mailservice/constants"
) 

// func TestMessageHeader(t *testing.T)  { update for table driven test
// 	hv := HeaderValue{ from: "eiei"}

// 	p := MessageHeader(hv)

// 	fmt.Println(p)
	
// 	if MessageHeader("testing") !=  "message/index.html"  {
// 		t.Error("ms.SendMail() == messages/index.html")
// 	}
// }

// func TestParsedBody(t *testing.T)  { // update for table driven test
	
// 	filename :=  constants.MessageFile(constants.First)
// 	p , _ := parsedBody(filename,  map[string]string{ "sugba": "fillo"})

// 	fmt.Println(p)
	
// 	if MessageHeader("testing") !=  "message/index.html"  {
// 		t.Error("ms.SendMail() == messages/index.html")
// 	}
// }

// func TestNonParsedBody(t *testing.T)  { // update for table driven test
	
// 	filename :=  constants.MessageFile(constants.First)
// 	p , _ := nonParsedBody(filename)

// 	fmt.Println(p)
	
// 	if MessageHeader("testing") !=  "message/index.html"  {
// 		t.Error("ms.SendMail() == messages/index.html")
// 	}
// }

// func TestParsedMessage(t *testing.T)  { // update for table driven test
// 	p,_ := ParsedMessage(
// 				MessageValue{ 
// 					constants.MessageFile(constants.First), 
// 					HeaderValue{ from: "anselm", to: "james", subj: "go to school"},
// 					map[string]string{ "sugba": "fillo"},
// 				},
// 		)

// 	fmt.Println(p)
	
// 	if MessageHeader("testing") !=  "message/index.html"  {
// 		t.Error("ms.SendMail() == messages/index.html")
// 	}
// }

// func TestNonParsedMessage(t *testing.T)  { // update for table driven test
// 	p,_ := NonParsedMessage(
// 				MessageValue{ 
// 					constants.MessageFile(constants.First), 
// 					HeaderValue{ from: "anselm", to: "james", subj: "go to school"},
// 					map[string]string{ "sugba": "fillo"},
// 				},
// 		)

// 	fmt.Println(p)
	
// 	if MessageHeader("testing") !=  "message/index.html"  {
// 		t.Error("ms.SendMail() == messages/index.html")
// 	}
// }

// func TestTextMessage(t *testing.T)  { // update for table driven test
// 	p,_ := TextMessage(
// 				"Hello emeka,\r\n\nWe have received your complaint and we are sincerly sorry",
// 				MessageValue{ 
// 					constants.MessageFile(constants.First), 
// 					HeaderValue{ From: "anselm", To: "james", Subj: "go to school", ContentType: "text/plain;"},
// 					map[string]string{ "sugba": "fillo"},
// 				},
// 		)

// 	fmt.Println(p)
	
// 	if MessageHeader("testing") !=  "message/index.html"  {
// 		t.Error("ms.SendMail() == messages/index.html")
// 	}
// }