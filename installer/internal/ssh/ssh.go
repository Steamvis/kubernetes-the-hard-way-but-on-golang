package ssh

import (
	"golang.org/x/crypto/ssh"
)

type Client struct {
	client *ssh.Client
}

type Session struct {
	session *ssh.Session
}

func NewConnect(user, pass, host string) (*Client, error) {
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host+":22", sshConfig)
	if err != nil {
		return nil, err
	}
	return &Client{client: client}, nil
}

func (c *Client) CloseConnect() {
	c.client.Close()
}

func (c *Client) NewSession() (*Session, error) {
	session, err := c.client.NewSession()
	if err != nil {
		return nil, err
	}
	return &Session{session: session}, nil
}

func (s *Session) CloseSession() {
	s.session.Close()
}

func (s *Session) RunCommand(command string) (string, error) {
	output, err := s.session.CombinedOutput(command)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
