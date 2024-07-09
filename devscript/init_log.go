package devscript

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type logFormat struct{}

func (l logFormat) Format(entry *logrus.Entry) ([]byte, error) {
	tf := entry.Time.Format("15:04:05")
	return []byte(fmt.Sprintf("%s: %s\n", tf, entry.Message)), nil
}

func init() {
	logrus.SetFormatter(&logFormat{})
}
