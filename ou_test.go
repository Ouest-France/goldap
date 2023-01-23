package goldap

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateOrganizationalUnit(t *testing.T) {
	client, err := NewClientHelper()
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}

	err = client.CreateOrganizationalUnit(os.Getenv("GOLDAP_TESTOU"), os.Getenv("GOLDAP_TESTDESC"), "")
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}
}

func TestReadOrganizationalUnit(t *testing.T) {

	client, err := NewClientHelper()
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}

	_, err = client.ReadOrganizationalUnit(os.Getenv("GOLDAP_TESTOU"))
	if err != nil {
		fmt.Printf("fail: %s", err)
		t.FailNow()
	}
}

func TestUpdateOrganizationalUnitDescription(t *testing.T) {

	client, err := NewClientHelper()
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}

	err = client.UpdateOrganizationalUnitDescription(os.Getenv("GOLDAP_TESTOU"), "bar")
	if err != nil {
		fmt.Printf("fail: %s", err)
		t.FailNow()
	}
}

func TestUpdateOrganizationalUnitManagedBy(t *testing.T) {

	client, err := NewClientHelper()
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}

	err = client.UpdateOrganizationalUnitManagedBy(os.Getenv("GOLDAP_TESTOU"), os.Getenv("GOLDAP_TESTMANAGEDBY"))
	if err != nil {
		fmt.Printf("fail: %s", err)
		t.FailNow()
	}
}

func TestDeleteOrganizationalUnit(t *testing.T) {

	client, err := NewClientHelper()
	if err != nil {
		fmt.Printf("%s", err)
		t.FailNow()
	}

	err = client.DeleteOrganizationalUnit(os.Getenv("GOLDAP_TESTOU"))
	if err != nil {
		fmt.Printf("fail: %s", err)
		t.FailNow()
	}
}
