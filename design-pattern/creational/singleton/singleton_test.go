package singleton

import (
	"testing"
)

func TestSingletonLogger(t *testing.T) {
	logger1 := GetLogger()
	logger2 := GetLogger()

	// Kiểm tra xem hai logger có cùng một địa chỉ không (singleton)
	if logger1 != logger2 {
		t.Error("Expected both loggers to be the same instance (singleton), but they are different")
	}
}
