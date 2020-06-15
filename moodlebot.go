package main

import (
	"bytes"
	"strings"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/emersion/go-smtp"
	"github.com/jhillyerd/enmime"
)

// The Backend implements SMTP server methods.
type Backend struct{}

// Login returns a session after login. Here we do not demand any login
func (bkd *Backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	return &Session{}, nil
}

// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return &Session{}, nil
}

// A Session is returned after successful login.
type Session struct{}

func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	log.Println("Mail from:", from)
	return nil
}

func (s *Session) Rcpt(to string) error {
	log.Println("Rcpt to:", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {

		r := strings.NewReader(string(b))
		env, err := enmime.ReadEnvelope(r)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(env.Text, "\n")

		url := "https://discord.com/api/webhooks/719401315910025229/DQwLU76-dm18r52fV8ztLTH68HdwuxRnV7MG4es8jXjy6ShDNtsn-Hmr0tL_kKLVF0yP?username=Flavio&content=Acabou%20de%20entrar!"
		method := "POST"
		payload := &bytes.Buffer{}
		writer := multipart.NewWriter(payload)
		_ = writer.WriteField("username", "MOODLE-HTO")
		_ = writer.WriteField("content", env.Text)

		err = writer.Close()
		if err != nil {
			fmt.Println(err)
		}

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Cookie", "__cfduid=d6498a7f20355e0632cb7500cbfe46be51591589127; __cfruid=8ffd542f0f49d35a80f0e69b2488f40b44813423-1591589127")

		req.Header.Set("Content-Type", writer.FormDataContentType())
		res, err := client.Do(req)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)

		log.Println("Data:", string(body))
	}
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

// A message is sent to Discord
type Message struct {
	url      string
	method   string
	username string
	message  string
}

func main() {
	be := &Backend{}

	s := smtp.NewServer(be)

	s.Addr = ":25"
	s.Domain = "209.182.235.117"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 20 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
