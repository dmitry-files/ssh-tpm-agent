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
    
    // Первая попытка чтения
    t.Logf("=== First attempt to read key ===")
    t.Logf("Attempting to read non-existent key: '%s'", keyName)
    
    result1, err1 := keyring.ReadKey(keyName)
    if err1 != nil {
        t.Logf("ReadKey returned error: %v", err1)
        t.Logf("Error type: %T", err1)
    } else {
        t.Logf("WARNING: ReadKey succeeded unexpectedly, result: %v", result1)
    }
    
    t.Log("Checking if error is syscall.ENOKEY...")
    if !errors.Is(err1, syscall.ENOKEY) {
        t.Logf("ERROR: Expected syscall.ENOKEY but got: %v", err1)
        t.Logf("Error type: %T", err1)
        t.Fatalf("err: %v", err1)
    }
    t.Log("First attempt: correctly returned syscall.ENOKEY")
    
    // Вторая попытка чтения того же ключа
    t.Logf("=== Second attempt to read the same key ===")
    t.Logf("Attempting to read non-existent key again: '%s'", keyName)
    
    result2, err2 := keyring.ReadKey(keyName)
    if err2 != nil {
        t.Logf("ReadKey returned error: %v", err2)
        t.Logf("Error type: %T", err2)
    } else {
        t.Logf("WARNING: ReadKey succeeded unexpectedly, result: %v", result2)
    }
    
    t.Log("Checking if error is syscall.ENOKEY...")
    if !errors.Is(err2, syscall.ENOKEY) {
        t.Logf("ERROR: Expected syscall.ENOKEY but got: %v", err2)
        t.Logf("Error type: %T", err2)
        t.Fatalf("err: %v", err2)
    }
    t.Log("Second attempt: correctly returned syscall.ENOKEY")
    
    t.Log("SUCCESS: Test passed - both attempts correctly returned syscall.ENOKEY for non-existent key")
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
