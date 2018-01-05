package ldapserver

// logger is an interface that implements different log methods.
// The standard interface only supports Print* functions which means
// a default logger or logrus will work as a logger
type logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}
