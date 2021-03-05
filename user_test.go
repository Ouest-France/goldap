package goldap

import (
	"fmt"
	"os"
	"testing"
)

func TestReadUserName(t *testing.T) {

	client, err := NewClientHelper()
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}

	_, err = client.ReadUser(os.Getenv("GOLDAP_TESTUSER_OU"), os.Getenv("GOLDAP_TESTUSER_NAME"), "", "")
	if err != nil {
		fmt.Printf("fail: %s", err)
		t.FailNow()
	}
}

func TestReadUserSamAccountName(t *testing.T) {

	client, err := NewClientHelper()
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}

	_, err = client.ReadUser(os.Getenv("GOLDAP_TESTUSER_OU"), "", os.Getenv("GOLDAP_TESTUSER_SAM_ACCOUNT_NAME"), "")
	if err != nil {
		fmt.Printf("fail: %s", err)
		t.FailNow()
	}
}

func TestReadUserUPN(t *testing.T) {

	client, err := NewClientHelper()
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}

	_, err = client.ReadUser(os.Getenv("GOLDAP_TESTUSER_OU"), "", "", os.Getenv("GOLDAP_TESTUSER_UPN"))
	if err != nil {
		fmt.Printf("fail: %s", err)
		t.FailNow()
	}
}
