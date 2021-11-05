package glog

import "errors"

var ErrNotImplemented error = errors.New("not implemented")

var MaxSize uint64 = 0

type Level int32

func (l *Level) Get() interface{}       { return 0 }
func (l *Level) Set(value string) error { return ErrNotImplemented }
func (l *Level) String() string         { return "" }

type OutputStats struct{}

func (s *OutputStats) Bytes() int64 { return 0 }
func (s *OutputStats) Lines() int64 { return 0 }

type Verbose bool

func (v Verbose) Info(args ...interface{})                 {}
func (v Verbose) Infof(format string, args ...interface{}) {}
func (v Verbose) Infoln(args ...interface{})               {}

func V(level Level) Verbose { return false }

func CopyStandardLogTo(name string)               {}
func Error(args ...interface{})                   {}
func ErrorDepth(depth int, args ...interface{})   {}
func Errorf(format string, args ...interface{})   {}
func Errorln(args ...interface{})                 {}
func Exit(args ...interface{})                    {}
func ExitDepth(depth int, args ...interface{})    {}
func Exitf(format string, args ...interface{})    {}
func Exitln(args ...interface{})                  {}
func Fatal(args ...interface{})                   {}
func FatalDepth(depth int, args ...interface{})   {}
func Fatalf(format string, args ...interface{})   {}
func Fatalln(args ...interface{})                 {}
func Flush()                                      {}
func Info(args ...interface{})                    {}
func InfoDepth(depth int, args ...interface{})    {}
func Infof(format string, args ...interface{})    {}
func Infoln(args ...interface{})                  {}
func Warning(args ...interface{})                 {}
func WarningDepth(depth int, args ...interface{}) {}
func Warningf(format string, args ...interface{}) {}
func Warningln(args ...interface{})               {}
