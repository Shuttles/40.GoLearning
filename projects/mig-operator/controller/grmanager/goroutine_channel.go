package grmanager

type chanInfo struct {
	trigger string
	msg     string
}

// goroutineChannel defines
type goroutineChannel struct {
	gid  uint64
	name string
	info chan chanInfo
}
