package smppconn

type SmppConnection struct {
	addr     string `json:"addr"`
	user     string `json:"user"`
	password string `json:"password"`
}

func (c *SmppConnection) Addr() string {
	return c.addr
}

func (c *SmppConnection) SetAddr(addr string) *SmppConnection {
	c.addr = addr
	return c
}

func (c *SmppConnection) User() string {
	return c.user
}

func (c *SmppConnection) SetUser(user string) *SmppConnection {
	c.user = user
	return c
}

func (c *SmppConnection) Password() string {
	return c.password
}

func (c *SmppConnection) SetPassword(password string) *SmppConnection {
	c.password = password
	return c
}
