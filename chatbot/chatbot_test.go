package chatbot

import (
	"io"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	r, w, _ := os.Pipe()

	oldStdout := os.Stdout
	defer func() {
		os.Stdout = oldStdout
	}()
	os.Stdout = w

	Greet("Jacek", "2023")

	w.Close()

	output, _ := io.ReadAll(r)
	expected := "Hello! My name is Jacek.\nI was created in 2023.\n"

	if string(output) != expected {
		t.Errorf("expected %q, got %q", expected, output)
	}
}

func TestGoodbye(t *testing.T) {
	r, w, _ := os.Pipe()

	oldStdout := os.Stdout
	defer func() {
		os.Stdout = oldStdout
	}()
	os.Stdout = w

	SayGoodbye()

	w.Close()

	output, _ := io.ReadAll(r)
	expected := "Congratulations, have a nice day!\n"

	if string(output) != expected {
		t.Errorf("expected %q, got %q", expected, output)
	}
}

func TestShowName(t *testing.T) {
	r, w, _ := os.Pipe()

	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()
	os.Stdin = r
	os.Stdout = w
	input := "Jacek\n"
	io.WriteString(w, input)

	ShowName()

	w.Close()

	output, _ := io.ReadAll(r)
	expected := "Please, remind me of your name.\nWhat a great name you have, Jacek!\n"

	if string(output) != expected {
		t.Errorf("expected %q, got %q", expected, output)
	}
}
