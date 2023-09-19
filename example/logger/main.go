package main

import (
	"context"

	"github.com/glebnaz/witcher/log"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetReportCaller(true)
	entry := logrus.WithFields(logrus.Fields{"id": 1})
	ctx := log.AddEntryToCTX(context.Background(), entry)

	log.Infof(ctx, "test with id")
	log.Infof(context.Background(), "test without id")
}
