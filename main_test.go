package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

const dstFilepath = "_testdata/dst.jpg"

func TestRun(t *testing.T) {
	var buf bytes.Buffer
	if err := run("test", &buf); err != nil {
		t.Fatalf("Should not be fail: %v.", err)
	}

	if os.Getenv("IS_CREATE_DST_FILE") == "true" {
		createDstFile(t, buf.Bytes())
	}

	dstFile, err := os.Open(dstFilepath)
	if err != nil {
		t.Fatalf("Should not be fail: %v.", err)
	}
	want, err := ioutil.ReadAll(dstFile)
	if err != nil {
		t.Fatalf("Should not be fail: %v.", err)
	}

	if got := buf.Bytes(); !bytes.Equal(got, want) {
		t.Fatalf("Should be same as %s.", dstFilepath)
	}
}

func createDstFile(t *testing.T, bs []byte) {
	t.Helper()

	f, err := os.Create(dstFilepath)
	if err != nil {
		t.Fatalf("Should not be fail: %v.", err)
	}
	defer f.Close()

	if _, err := f.Write(bs); err != nil {
		t.Fatalf("Should not be fail: %v.", err)
	}

	t.Skip("Create destination file.")
}
