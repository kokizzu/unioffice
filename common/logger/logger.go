//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package logger ;import (_bd "fmt";_ee "io";_g "os";_e "path/filepath";_d "runtime";);

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};

// Error logs error message.
func (_bdb WriterLogger )Error (format string ,args ...interface{}){if _bdb .LogLevel >=LogLevelError {_ecb :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_bdb .logToWriter (_bdb .Output ,_ecb ,format ,args ...);};};

// Trace logs trace message.
func (_ec ConsoleLogger )Trace (format string ,args ...interface{}){if _ec .LogLevel >=LogLevelTrace {_gbcg :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_ec .output (_g .Stdout ,_gbcg ,format ,args ...);};};func (_bg WriterLogger )logToWriter (_gaa _ee .Writer ,_bge string ,_ebd string ,_ac ...interface{}){_ad (_gaa ,_bge ,_ebd ,_ac );};

// Warning logs warning message.
func (_ba ConsoleLogger )Warning (format string ,args ...interface{}){if _ba .LogLevel >=LogLevelWarning {_dd :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_ba .output (_g .Stdout ,_dd ,format ,args ...);};};

// LogLevel is the verbosity level for logging.
type LogLevel int ;

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _ee .Writer )*WriterLogger {logger :=WriterLogger {Output :writer ,LogLevel :logLevel };return &logger ;};

// Warning logs warning message.
func (_ag WriterLogger )Warning (format string ,args ...interface{}){if _ag .LogLevel >=LogLevelWarning {_bde :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_ag .logToWriter (_ag .Output ,_bde ,format ,args ...);};};const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);var Log Logger =DummyLogger {};

// Info logs info message.
func (_bce ConsoleLogger )Info (format string ,args ...interface{}){if _bce .LogLevel >=LogLevelInfo {_dg :="\u005bI\u004e\u0046\u004f\u005d\u0020";_bce .output (_g .Stdout ,_dg ,format ,args ...);};};

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};

// Notice logs notice message.
func (_gff WriterLogger )Notice (format string ,args ...interface{}){if _gff .LogLevel >=LogLevelNotice {_eg :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_gff .logToWriter (_gff .Output ,_eg ,format ,args ...);};};

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// Trace logs trace message.
func (_fd WriterLogger )Trace (format string ,args ...interface{}){if _fd .LogLevel >=LogLevelTrace {_da :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_fd .logToWriter (_fd .Output ,_da ,format ,args ...);};};

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };

// Debug logs debug message.
func (_c ConsoleLogger )Debug (format string ,args ...interface{}){if _c .LogLevel >=LogLevelDebug {_cf :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_c .output (_g .Stdout ,_cf ,format ,args ...);};};func _ad (_eag _ee .Writer ,_agg string ,_fba string ,_df ...interface{}){_ ,_ge ,_gc ,_cb :=_d .Caller (3);if !_cb {_ge ="\u003f\u003f\u003f";_gc =0;}else {_ge =_e .Base (_ge );};_ff :=_bd .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_agg ,_ge ,_gc )+_fba +"\u000a";_bd .Fprintf (_eag ,_ff ,_df ...);};

// DummyLogger does nothing.
type DummyLogger struct{};

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _ee .Writer ;};

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_eb string ,_gb ...interface{});Warning (_ed string ,_edb ...interface{});Notice (_ga string ,_f ...interface{});Info (_a string ,_bc ...interface{});Debug (_ae string ,_gbc ...interface{});Trace (_af string ,_gad ...interface{});IsLogLevel (_ea LogLevel )bool ;};

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};func (_bcd ConsoleLogger )output (_dda _ee .Writer ,_gg string ,_afg string ,_ebb ...interface{}){_ad (_dda ,_gg ,_afg ,_ebb ...);};

// Debug logs debug message.
func (_ef WriterLogger )Debug (format string ,args ...interface{}){if _ef .LogLevel >=LogLevelDebug {_de :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_ef .logToWriter (_ef .Output ,_de ,format ,args ...);};};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_fb ConsoleLogger )IsLogLevel (level LogLevel )bool {return _fb .LogLevel >=level };

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_afc WriterLogger )IsLogLevel (level LogLevel )bool {return _afc .LogLevel >=level };

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};

// Error logs error message.
func (_fg ConsoleLogger )Error (format string ,args ...interface{}){if _fg .LogLevel >=LogLevelError {_bb :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_fg .output (_g .Stdout ,_bb ,format ,args ...);};};

// Notice logs notice message.
func (_dde ConsoleLogger )Notice (format string ,args ...interface{}){if _dde .LogLevel >=LogLevelNotice {_gf :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_dde .output (_g .Stdout ,_gf ,format ,args ...);};};

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };

// Info logs info message.
func (_fc WriterLogger )Info (format string ,args ...interface{}){if _fc .LogLevel >=LogLevelInfo {_fce :="\u005bI\u004e\u0046\u004f\u005d\u0020";_fc .logToWriter (_fc .Output ,_fce ,format ,args ...);};};