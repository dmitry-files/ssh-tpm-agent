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
    
    // Проверяем что обе ошибки равны и соответствуют ENOKEY
    t.Log("=== Comparing results from both attempts ===")
    t.Logf("First error: %v (type: %T)", err1, err1)
    t.Logf("Second error: %v (type: %T)", err2, err2)
    
    // Проверяем что обе ошибки равны друг другу
    if err1 != err2 {
        t.Logf("ERROR: Errors from two attempts are not equal")
        t.Logf("First error: %v", err1)
        t.Logf("Second error: %v", err2)
        t.Fatalf("errors from two attempts are not equal: %v vs %v", err1, err2)
    }
    t.Log("Both errors are equal")
    
    // Проверяем что ошибка соответствует ENOKEY
    t.Log("Checking if error is syscall.ENOKEY...")
    if !errors.Is(err1, syscall.ENOKEY) {
        t.Logf("ERROR: Expected syscall.ENOKEY but got: %v", err1)
        t.Logf("Error type: %T", err1)
        t.Fatalf("err: %v", err1)
    }
    
    t.Log("SUCCESS: Test passed - both attempts returned identical syscall.ENOKEY errors for non-existent key")
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
