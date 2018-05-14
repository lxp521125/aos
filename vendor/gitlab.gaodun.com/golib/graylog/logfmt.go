package graylog

import (
	"io"
	"log"
	"os"
)

type LogStruct struct {
	Debug  bool
	out    io.Writer
	addLog func(info string)
}

func (nl LogStruct) Write(p []byte) (int, error) {
	if nl.Debug {
		return nl.out.Write(p)
	}
	go nl.addLog(string(p))
	return 0, nil
}

func (nl *LogStruct) SetOutput(f func(info string)) {
	nl.addLog = f
}


func New(debug bool, f func(info string)) *log.Logger {
	l := log.New(os.Stderr, "", log.LstdFlags|log.Llongfile)
	nl := LogStruct{out: os.Stderr, Debug: debug}
	nl.SetOutput(f)
	l.SetOutput(nl)
	return l
}

func NewSQL(debug bool, f func(info string)) LogStruct {
	nl := LogStruct{out: os.Stderr, Debug: debug}
	nl.SetOutput(f)
	return nl
}
