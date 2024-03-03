package main

import "testing"

var messageGenerated = []byte("protocol test")

func TestProtocol(t *testing.T) {
	generateProtocol, err := CreateProtocol(nil, nil, messageGenerated)
	if err != nil {
		t.Errorf("Error to initialize protocol object: %v", err)
	}

	t.Run("should generate a protocol valid", func(t *testing.T) {
		if generateProtocol.GetKey() == nil {
			t.Errorf("Expected protocol key to be not empty, but got %v", generateProtocol.GetKey())
		}

		if len(generateProtocol.GetKey()) != _keyLen {
			t.Errorf("Expected protocol key to be of length %v, but got %v", _keyLen, len(generateProtocol.GetKey()))
		}

		if generateProtocol.GetIV() == nil {
			t.Errorf("Expected protocol iv to be not empty, but got %v", generateProtocol.GetIV())
		}

		if len(generateProtocol.GetIV()) != _ivLen {
			t.Errorf("Expected protocol iv to be of length %v, but got %v", _ivLen, len(generateProtocol.GetIV()))
		}

		if generateProtocol.GetMessage() == nil {
			t.Errorf("Expected protocol message to be not empty, but got %v", generateProtocol.GetMessage())
		}

		if len(generateProtocol.GetMessage()) != len(messageGenerated) {
			t.Errorf("Expected protocol message to be of length %v, but got %v", len(messageGenerated), len(generateProtocol.GetMessage()))
		}

		if generateProtocol.GetVersion() == nil {
			t.Errorf("Expected protocol version to be not empty, but got %v", generateProtocol.GetVersion())
		}

		if generateProtocol.GetVersionString() != string(version) {
			t.Errorf("Expected protocol version to be %v, but got %v", string(version), generateProtocol.GetVersionString())
		}

		if len(generateProtocol.ToByteArray()) != (len(generateProtocol.GetKey()) + len(generateProtocol.GetIV()) + len(generateProtocol.GetVersion()) + len(generateProtocol.GetMessage())) {
			t.Errorf("Expected protocol byte array to be of length %v, but got %v", (len(generateProtocol.GetKey()) + len(generateProtocol.GetIV()) + len(generateProtocol.GetVersion()) + len(generateProtocol.GetMessage())), len(generateProtocol.ToByteArray()))
		}
	})

	importProtocol, err := ImportProtocol(generateProtocol.ToByteArray())
	if err != nil {
		t.Errorf("Error to import protocol object: %v", err)
	}

	t.Run("should import a protocol valid", func(t *testing.T) {
		if importProtocol.GetKey() == nil {
			t.Errorf("Expected protocol key to be not empty, but got %v", importProtocol.GetKey())
		}

		if len(importProtocol.GetKey()) != _keyLen {
			t.Errorf("Expected protocol key to be of length %v, but got %v", _keyLen, len(importProtocol.GetKey()))
		}

		if importProtocol.GetIV() == nil {
			t.Errorf("Expected protocol iv to be not empty, but got %v", importProtocol.GetIV())
		}

		if len(importProtocol.GetIV()) != _ivLen {
			t.Errorf("Expected protocol iv to be of length %v, but got %v", _ivLen, len(importProtocol.GetIV()))
		}

		if importProtocol.GetMessage() == nil {
			t.Errorf("Expected protocol message to be not empty, but got %v", importProtocol.GetMessage())
		}

		if len(importProtocol.GetMessage()) != len(messageGenerated) {
			t.Errorf("Expected protocol message to be of length %v, but got %v", len(messageGenerated), len(importProtocol.GetMessage()))
		}

		if importProtocol.GetVersion() == nil {
			t.Errorf("Expected protocol version to be not empty, but got %v", importProtocol.GetVersion())
		}

		if importProtocol.GetVersionString() != string(version) {
			t.Errorf("Expected protocol version to be %v, but got %v", string(version), importProtocol.GetVersionString())
		}

		if len(importProtocol.ToByteArray()) != (len(importProtocol.GetKey()) + len(importProtocol.GetIV()) + len(importProtocol.GetVersion()) + len(importProtocol.GetMessage())) {
			t.Errorf("Expected protocol byte array to be of length %v, but got %v", (len(importProtocol.GetKey()) + len(importProtocol.GetIV()) + len(importProtocol.GetVersion()) + len(importProtocol.GetMessage())), len(importProtocol.ToByteArray()))
		}
	})

	importProtocolBase64, err := ImportProtocolBase64(generateProtocol.ToBase64())
	if err != nil {
		t.Errorf("Error to import protocol object from base64: %v", err)
	}

	t.Run("should import a protocol valid from base64", func(t *testing.T) {
		if importProtocolBase64.GetKey() == nil {
			t.Errorf("Expected protocol key to be not empty, but got %v", importProtocolBase64.GetKey())
		}

		if len(importProtocolBase64.GetKey()) != _keyLen {
			t.Errorf("Expected protocol key to be of length %v, but got %v", _keyLen, len(importProtocolBase64.GetKey()))
		}

		if importProtocolBase64.GetIV() == nil {
			t.Errorf("Expected protocol iv to be not empty, but got %v", importProtocolBase64.GetIV())
		}

		if len(importProtocolBase64.GetIV()) != _ivLen {
			t.Errorf("Expected protocol iv to be of length %v, but got %v", _ivLen, len(importProtocolBase64.GetIV()))
		}

		if importProtocolBase64.GetMessage() == nil {
			t.Errorf("Expected protocol message to be not empty, but got %v", importProtocolBase64.GetMessage())
		}

		if len(importProtocolBase64.GetMessage()) != len(messageGenerated) {
			t.Errorf("Expected protocol message to be of length %v, but got %v", len(messageGenerated), len(importProtocolBase64.GetMessage()))
		}

		if importProtocolBase64.GetVersion() == nil {
			t.Errorf("Expected protocol version to be not empty, but got %v", importProtocolBase64.GetVersion())
		}

		if importProtocolBase64.GetVersionString() != string(version) {
			t.Errorf("Expected protocol version to be %v, but got %v", string(version), importProtocolBase64.GetVersionString())
		}

		if len(importProtocolBase64.ToByteArray()) != (len(importProtocolBase64.GetKey()) + len(importProtocolBase64.GetIV()) + len(importProtocolBase64.GetVersion()) + len(importProtocolBase64.GetMessage())) {
			t.Errorf("Expected protocol byte array to be of length %v, but got %v", (len(importProtocolBase64.GetKey()) + len(importProtocolBase64.GetIV()) + len(importProtocolBase64.GetVersion()) + len(importProtocolBase64.GetMessage())), len(importProtocolBase64.ToByteArray()))
		}
	})
}
