package mailservice 

import (
	"crypto/tls"
	"services/mailservice/constants"
	"net/smtp"
)

// Refresh knowledge on smtp, if need be: https://www.ietf.org/rfc/rfc821.txt
// there are diffrences, as i could infer, the above could have been written
// with mail requests from the browser in mind. compare the above with the
// code implementation below.
// Here we first establish a two way transmission channel to a receiver
// smtp server (an endpoint) which may be the final or an intermediate server
// when we call tls.Dail(network, hostandport, tlsconfig), using tcp (transmission control protocol) 
// which is a transport layer protocol or transport, the host and port we establish connection
// and create a communication channel between this client as an endpoint and the smtp server as another 
// endpoint, after this we use tls to secure and encrypt message exchanged
// https://www.cloudflare.com/learning/ssl/what-happens-in-a-tls-handshake/#:~:text=A%20TLS%20handshake%20is%20the,and%20agree%20on%20session%20keys.

type MailService struct {
		tlsconnection	*tls.Conn
		smtpclient		*smtp.Client
}


func EstablishConn()                 (*MailService, error)                 {
		tlsconfig := &tls.Config { 
						InsecureSkipVerify: true,
						ServerName: constants.Host,
		}
		conn, err := tls.Dial("tcp", constants.HostAndPort, tlsconfig)
		if err != nil { return nil, err }

		return &MailService { tlsconnection: conn } , nil
}


// authenticates, if no error, assigns a client to mail service
func (ms *MailService)    Auth()                    (error)                 { 
		client, err := smtp.NewClient(ms.tlsconnection, constants.Host)
		if err != nil { return err }
		auth := smtp.PlainAuth("", constants.From.Address, constants.Password, constants.Host)
		err = client.Auth(auth); 
		if err != nil {  return err }
		ms.smtpclient = client

		return nil
}


// the SMTP client sends a MAIL command indicating the sender of the mail,
// as you can see below. If the receiver SMTP server can accept the mail it responds
// with an OK reply and it doesn't error.
// The SMTP client then sends a RCPT command identifying a recipient of the mail. If 
// the receiver SMTP server can accept mail for that recipient it responds with an
// OK reply and it doesn't error. This errors is a response with a reply rejecting 
// the specified recipient (but not the whole mail transaction).
// Note this is a dialog and everything happens a step at a time
// if it hasn't errored till now then we write the message string as a slice of bytes
//
// we quit the smtp client
   

func (ms MailService)    SendMail(to, message string) 		 (error)        {	
		err := ms.smtpclient.Mail(constants.From.Address)
		if err != nil {   return err   }
		err = ms.smtpclient.Rcpt(to)
		if err != nil {   return err   }
		w, err := ms.smtpclient.Data()
		if err != nil {   return err   }
		_, err = w.Write([]byte(message))
		if err != nil {   return err   }
		err = w.Close()
		if err != nil {   return err   }
		ms.smtpclient.Quit()

		return nil
} 