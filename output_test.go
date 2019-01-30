package logging

import "testing"

func TestLog(t *testing.T) {
	SetLevel(TRACE)

	Trace("Trace")
	Tracef("Tracef")
	Debug("Debug")
	Debugf("Debugf")
	Info("Info")
	Infof("Infof")
	Warning("Warning")
	Warningf("Warningf")
	Error("Error")
	Errorf("Errorf")
}
