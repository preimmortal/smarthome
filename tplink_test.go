package smarthome

import (
	"fmt"
	"testing"
)

func setupSubTest() {
}

func TestEncrypt(t *testing.T) {
	tp := tplink{}
	encodedString := tp.encrypt("hello")
	expectedBytes := string([]byte{0, 0, 0, 5, 195, 166, 202, 166, 201})

	if string(encodedString) != string(expectedBytes) {
		t.Fatalf("Encoded string is incorrect ")
	}
}

func TestDecrypt(t *testing.T) {
	tp := tplink{}
	encodedBytes := []byte{195, 166, 202, 166, 201}
	expectedString := "hello"

	decryptedString := tp.decrypt(encodedBytes)

	if decryptedString != expectedString {
		t.Fatalf("Decrypted string is incorrect")
	}
}

// Test_Positive_Send implements a positive send ip test
func Test_Positive_Send(t *testing.T) {
	tp := tplink{}
	sn := scan{}
	ip, err := sn.FindFirstIP()
	if err != nil {
		t.Fatalf("Could not find any IP")
	}
	result, err := tp.Send(ip, "{\"system\":{\"get_sysinfo\":{}}}")
	if err != nil {
		t.Fatalf("Could not execute info command")
	}

	fmt.Println(result)

}

// Test_Negative_Send_BadIP implements a negative bad ip test
func Test_Negative_Send_BadIP(t *testing.T) {
	tp := tplink{}
	t.Log("Negative Test - Bad IP")
	// Test  Bad IP
	bad_ip_result, err := tp.Send("999.999.999.999", "{\"system\":{\"get_sysinfo\":{}}}")
	if bad_ip_result != "" || err == nil {
		t.Fatalf("Expected command not to work")
	}
}

// Test_Negative_Send_BadCall implements a negative bad call test
func Test_Negative_Send_BadCall(t *testing.T) {
	tp := tplink{}
	sn := scan{}
	ip, err := sn.FindFirstIP()
	if err != nil {
		t.Fatalf("Could not find any IP")
	}
	t.Log("Negative Test - Bad Call")
	// Test  Bad Call
	bad_data_result, err := tp.Send(ip, "")
	if bad_data_result != "" || err == nil {
		t.Fatalf("Expected command not to work")
	}

}
