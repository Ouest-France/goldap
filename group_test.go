package goldap

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateGroup(t *testing.T) {

	client, err := NewClientHelper()
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}

	err = client.CreateGroup(os.Getenv("GOLDAP_TESTGROUP"), os.Getenv("GOLDAP_TESTNAME"), []string{})
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}
}

func TestReadGroup(t *testing.T) {

	client, err := NewClientHelper()
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}

	_, err = client.ReadGroup(os.Getenv("GOLDAP_TESTGROUP"))
	if err != nil {
		fmt.Printf("fail: %s", err)
		t.FailNow()
	}
}

func TestDeleteGroup(t *testing.T) {

	client, err := NewClientHelper()
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}

	_, err = client.ReadGroup(os.Getenv("GOLDAP_TESTGROUP"))
	if err != nil {
		fmt.Printf("fail: %s", err)
		t.FailNow()
	}
}
