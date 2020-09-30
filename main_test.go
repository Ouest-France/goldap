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

	tlsInsecure, err := strconv.ParseBool(os.Getenv("GOLDAP_TLS_INSECURE"))
	if err != nil {
		return Client{}, err
	}

	client := Client{
		Host:         os.Getenv("GOLDAP_HOST"),
		Port:         port,
		BindUser:     os.Getenv("GOLDAP_BINDUSER"),
		BindPassword: os.Getenv("GOLDAP_BINDPASSWORD"),
		TLS:          tls,
		TLSCACert:    os.Getenv("GOLDAP_TLS_CA_CERT"),
		TLSInsecure:  tlsInsecure,
	}

	err = client.Connect()
	return client, err
}
