package message



import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/mail"
)



type HeaderValue     struct     {
		  From            string
		  To              string
		  Subj            string
		  MimeVersion     string
		  ContentType	  string
}


type MessageValue    struct     {
		  Filename        string
		  HV         HeaderValue
		  Data map[string]string
}

func messageHeader(hv HeaderValue)		          (*mail.Header)		    	{
				var header     =    make(mail.Header)
				header["From"]                   =   []string{      hv.From             } 
				header["To"]                     =   []string{      hv.To               }
				header["Subject"]                =   []string{      hv.Subj             }
				header["MIME-version"]           =   []string{      hv.MimeVersion      }
				header["Content-Type"]           =   []string{      hv.ContentType      }

				return &header
}

func nonParsedBody(filename  string)                    (string, error)  {
				bodyBytes, err := ioutil.ReadFile(filename)
				if err         != nil {   return "", err    }
				body           := string(bodyBytes)


				return body, nil
}

func parsedBody(filename string, data map[string]string) (string, error) {
				template, err := template.ParseFiles(filename)
				if        err != nil {   return "", err    }
				buffer        := new(bytes.Buffer)
				template.Execute(buffer, data)
				body          := buffer.String()

				return body, nil
}

func ParsedMessage(mv MessageValue)   (string, error)        {
				body, err := parsedBody(mv.Filename, mv.Data)
				if    err != nil {   return "", err     }
				headers   := messageHeader(mv.HV)
				message   := ""
				for key, value := range *headers {
					   message += fmt.Sprintf("%s: %s\r\n", key, value[0])
				}
				message   += "\r\n" + body

				return message, nil
}




func NonParsedMessage(mv MessageValue) (string, error) {
				body, err := nonParsedBody(mv.Filename)
				if    err != nil {   return "", err   }
				headers   := messageHeader(mv.HV)
				message   := ""
				for key, value := range *headers {
					   message += fmt.Sprintf("%s: %s\r\n", key, value[0])
				}
				message   += "\r\n" + body

				return message, nil
}



func TextMessage(body string, mv MessageValue)      (string, error) {
				headers   := messageHeader(mv.HV)
				message   := ""
				for key, value := range *headers {
					message += fmt.Sprintf("%s: %s\r\n", key, value[0])
				}
				message   += "\r\n" + body

				return message, nil
}



// year, month, day := time.Now().Date()
// hour, min, sec := time.Now().Clock()
// t := fmt.Sprintf("%d:%d:%d WAT (%s %d, %d)", hour, min, sec, month, day, year)


// func OTPMessage() (string) {
// 	return ""
// }







// from := mail.Address{"NairaRamp", "testsend2022@yahoo.com"}
// to   := mail.Address{"", "usih.anselm@gmail.com"}
// subj := "Your login code is 019827"
// // MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
// body := buffer.String()
// password := "nlikvcvrgtlmsmhr"






