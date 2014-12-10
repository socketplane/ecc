package ecc

import "testing"

// Combining the tests under a single test case to

func TestStart(t *testing.T) {
	err := Start(true, true, "", "data-dir")
	if err != nil {
		t.Error("Error starting Consul ", err)
	}
}

func TestJoin(t *testing.T) {
	err := Join("1.1.1.1")
	if err == nil {
		t.Error("Join to unknown peer must fail")
	}
}

func TestGet(t *testing.T) {
	existingValue, _, ok := Get("ipam", "test")
	if ok {
		t.Fatal("Please cleanup the existing database and restart the test :", existingValue)
	}
}

func TestPut(t *testing.T) {
	existingValue, _, ok := Get("ipam", "test")
	if ok {
		t.Fatal("Please cleanup the existing database and restart the test")
	}

	err := Put("ipam", "test", "192.168.56.1", existingValue)
	if err != nil {
		t.Fatal("Error putting value into ipam store")
	}

	// Test with Old existingValue
	err = Put("ipam", "test", "192.168.56.1", existingValue)
	if err == nil {
		t.Fatal("Put must fail if the existingValue is NOT in sync with the db")
	}

	// Test with New existingValue
	existingValue, _, ok = Get("ipam", "test")
	if !ok {
		t.Fatal("test key is missing in ipam store")
	}

	err = Put("ipam", "test", "192.168.56.2", existingValue)
	if err != nil {
		t.Error("Error putting value into ipam store")
	}

	existingValue, _, ok = Get("ipam", "test")
	if !ok {
		t.Fatal("test key is missing in ipam store")
	}
	if existingValue != "192.168.56.2" {
		t.Fatal("Value for test key is not updated in DB")
	}
}

func TestCleanup(t *testing.T) {
	Delete("ipam", "test")
	Leave()
}
