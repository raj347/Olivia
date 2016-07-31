package olilib

import (
	"testing"
)

func TestEncode(t *testing.T) {
	expectedReturn := "A3BZ5T3"
	retVal := Encode("AAABZZZZZTTT")

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}

func TestEncodeIntegers(t *testing.T) {
	expectedReturn := "052234"
	retVal := Encode("00000223333")

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}

func TestDecode(t *testing.T) {
	expectedReturn := "AAABZZZZZTTT"
	retVal := Decode("A3BZ5T3")

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}

func TestWriteOutputOver9(t *testing.T) {
	expectedReturn := "a9a9a9a9a9a9a9"
	retVal := writeOutput("", 'a', 63)

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}

func TestWriteOutputOver9WithExtra(t *testing.T) {
	expectedReturn := "a9a9a9a9a9a9a9a5"
	retVal := writeOutput("", 'a', 68)

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}


func TestDencodeIntegers(t *testing.T) {
	expectedReturn := "00000223333"
	retVal := Decode("052234")

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}

func TestEncodeDecode(t *testing.T) {
	expectedReturn := "AAABBZZZZZTTT"
	retVal := Encode(expectedReturn)
	retVal = Decode(retVal)

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}

func TestEncodeDecodeIntegers(t *testing.T) {
	expectedReturn := "00000223333"
	retVal := Encode(expectedReturn)
	retVal = Decode(retVal)

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}

func TestEncodeLongString(t *testing.T) {
	expectedReturn := "09150909090335524"
	retVal := Encode("0000000001111100000000000000000000000000000033333554")

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}

func TestDecodeLongString(t *testing.T) {
	expectedReturn := "0000000001111100000000000000000000000000000033333554"
	retVal := Decode("09150909090335524")

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}

func TestDecodeLongStringLongStringAtEnd(t *testing.T) {
	expectedReturn := "000000000111110000000000000000000000000000004444444"
	retVal := Decode("09150909090347")

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}

func TestEncodeDecodeLongString(t *testing.T) {

}
