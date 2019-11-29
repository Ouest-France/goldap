package goldap

import (
	"os"
	"strconv"
)

func NewClientHelper() (Client, error) {

	port, err := strconv.Atoi(os.Getenv("GOLDAP_PORT"))
	if err != nil {
		return Client{}, err
	}

	tls, err := strconv.ParseBool(os.Getenv("GOLDAP_TLS"))
	if err != nil {
		return Client{}, err
	}

	client := Client{
		Host:         os.Getenv("GOLDAP_HOST"),
		Port:         port,
		BindUser:     os.Getenv("GOLDAP_BINDUSER"),
		BindPassword: os.Getenv("GOLDAP_BINDPASSWORD"),
		TLS:          tls,
	}

	err = client.Connect()
	return client, err
}
