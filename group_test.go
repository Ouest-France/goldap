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

	err = client.CreateGroup(os.Getenv("GOLDAP_TESTGROUP"), os.Getenv("GOLDAP_TESTNAME"), os.Getenv("GOLDAP_TESTDESC"), []string{})
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}
}

func TestUpdateGroupMembers(t *testing.T) {

	client, err := NewClientHelper()
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}

	err = client.UpdateGroupMembers(os.Getenv("GOLDAP_TESTGROUP"), []string{os.Getenv("GOLDAP_TESTMEMBER")})
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}
}

func TestUpdateGroupDescription(t *testing.T) {

	client, err := NewClientHelper()
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}

	err = client.UpdateGroupDescription(os.Getenv("GOLDAP_TESTGROUP"), os.Getenv("GOLDAP_TESTDESCNEW"))
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}
}

func TestUpdateGroupType(t *testing.T) {

	client, err := NewClientHelper()
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}

	err = client.UpdateGroupType(os.Getenv("GOLDAP_TESTGROUP"), "-2147483640")
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

	_, err = client.ReadGroup(os.Getenv("GOLDAP_TESTGROUP"), 1500)
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

	_, err = client.ReadGroup(os.Getenv("GOLDAP_TESTGROUP"), 10)
	if err != nil {
		fmt.Printf("fail: %s", err)
		t.FailNow()
	}
}
