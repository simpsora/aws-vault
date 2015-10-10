package main

import (
	"bytes"
	"log"
)

type Logger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type Ui struct {
	Logger
	Error, Debug Logger
	Exit         func(code int)
}

type logWriter struct{ Logger }

func (w logWriter) Write(b []byte) (int, error) {
	w.Printf("%s", b)
	return len(b), nil
}

func MockUi(exit func(code int)) *Ui {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	debug := &bytes.Buffer{}
	return &Ui{
		Logger: &MockLogger{log.New(stdout, "", 0), stdout, exit},
		Debug:  &MockLogger{log.New(debug, "", 0), debug, exit},
		Error:  &MockLogger{log.New(stderr, "", 0), stderr, exit},
		Exit:   exit,
	}
}

type MockLogger struct {
	*log.Logger
	*bytes.Buffer
	Exit func(code int)
}

func (l *MockLogger) Fatal(v ...interface{}) {
	l.Logger.Println(v...)
	l.Exit(1)
}

func (l *MockLogger) Fatalf(format string, v ...interface{}) {
	l.Logger.Printf(format, v...)
	l.Exit(1)
}

func (l *MockLogger) Fatalln(v ...interface{}) {
	l.Logger.Println(v...)
	l.Exit(1)
}
