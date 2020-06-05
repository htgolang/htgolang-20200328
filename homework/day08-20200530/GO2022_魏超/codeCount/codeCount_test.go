package codeCount

import (
	"os"
	"path/filepath"
	"testing"
)

/*
***************** TEST DATA ***********************
*                                                 *
*                                                 *
*                                                 *
*                                                 *
*                                                 *
*                                                 *
***************************************************
 */
func TestCodeCount(t *testing.T) {
	path, _ := os.Getwd()
	testfile := filepath.Join(path, "codeCount_test.go")
	count := CodeCount(testfile)
	if count != 14 {
		t.Errorf("实际行数:%d.", count)
	}
}
