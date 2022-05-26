package constants

import "testing"


func TestMessageFile(t *testing.T)  {

	if MessageFile(first) !=  "message/index.html"  {
		t.Error("ms.SendMail() == messages/index.html")
	}
}