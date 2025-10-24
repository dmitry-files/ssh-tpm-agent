package keyring

import (
	"errors"
	"syscall"
	"testing"
)

/*
func TestSaveAndGetData(t *testing.T) {
	keyring, err := SessionKeyring.CreateKeyring()
	if err != nil {
		t.Fatalf("failed getting keyring: %v", err)
	}

	b := []byte("test string")
	if err := keyring.AddKey("test", b); err != nil {
		t.Fatalf("err: %v", err)
	}

	bb, err := keyring.ReadKey("test")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if !bytes.Equal(b, bb.Read()) {
		t.Fatalf("strings not equal")
	}
}
*/

func TestNoKey(t *testing.T) {
    t.Log("Starting TestNoKey: testing behavior when reading non-existent key")
    
    t.Log("Creating keyring...")
    keyring, err := SessionKeyring.CreateKeyring()
    if err != nil {
        t.Logf("ERROR: Failed to create keyring: %v", err)
        t.Fatalf("failed getting keyring: %v", err)
    }
    t.Log("Keyring created successfully")
    
    keyName := "this.key.does.not.exist"
    t.Logf("Attempting to read non-existent key: '%s'", keyName)
    
    result, err := keyring.ReadKey(keyName)
    if err != nil {
        t.Logf("ReadKey returned error as expected: %v", err)
        t.Logf("Error type: %T", err)
    } else {
        t.Logf("WARNING: ReadKey succeeded unexpectedly, result: %v", result)
    }
    
    t.Log("Checking if error is syscall.ENOKEY...")
    if !errors.Is(err, syscall.ENOKEY) {
        t.Logf("ERROR: Expected syscall.ENOKEY but got: %v", err)
        t.Logf("Error type: %T", err)
        t.Fatalf("err: %v", err)
    }
    
    t.Log("SUCCESS: Test passed - correctly returned syscall.ENOKEY for non-existent key")
}

/*
func TestRemoveKey(t *testing.T) {
	keyring, err := SessionKeyring.CreateKeyring()
	if err != nil {
		t.Fatalf("failed getting keyring: %v", err)
	}
	b := []byte("test string")
	if err := keyring.AddKey("test-2", b); err != nil {
		t.Fatalf("err: %v", err)
	}

	bb, err := keyring.ReadKey("test-2")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if !bytes.Equal(b, bb.Read()) {
		t.Fatalf("strings not equal")
	}

	if err = keyring.RemoveKey("test-2"); err != nil {
		t.Fatalf("failed removing key: %v", err)
	}
	_, err = keyring.ReadKey("test-2")
	if !errors.Is(err, syscall.ENOKEY) {
		t.Fatalf("we can still read the key")
	}
}
*/
