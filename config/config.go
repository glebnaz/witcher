package config

import (
	"context"
	"os"
	"strings"

	"github.com/glebnaz/witcher/log"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ENV map[string]string `json:"config" yaml:"config"`

	// add more config here in future
	// some k8s config and also
}

// MustSetEnvFromFile sets env variables from file, and return error of set env
// file should be in yaml format
// example:
// config:
//
//	ENV:
//	  key: value
//	  key2: value2
//	  key3: value3
//	  key4: value4 //nolint:dupl
func MustSetEnvFromFile(ctx context.Context, filePath string) error {
	// adjust ctx

	ctx = log.AddEntryToCTX(ctx, logrus.WithFields(map[string]interface{}{
		"file": filePath,
		"from": "MustSetEnvFromFile",
	}))

	log.Infof(ctx, "setting env from file %s", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf(ctx, "error opening file %s", err.Error())
		return errors.Wrap(err, "error opening file")
	}

	defer func() {
		errClose := file.Close()
		if errClose != nil {
			log.Debugf(ctx, "error closing file %s", errClose.Error())
		}
	}()

	// read file

	// parse file

	var cfg Config
	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		log.Errorf(ctx, "error decoding file %s", err.Error())
		return errors.Wrap(err, "error decoding file")
	}

	err = setEnvFromMap(ctx, cfg.ENV, true)
	if err != nil {
		return errors.Wrap(err, "error setting env from file "+filePath)
	}

	return nil
}

// SetEnvFromFile sets env variables from file, ignore error of set env, return only read from file error
// file should be in yaml format
// example:
// config:
//
//	ENV:
//	  key: value
//	  key2: value2
//	  key3: value3
//	  key4: value4
func SetEnvFromFile(ctx context.Context, filePath string) error {
	// adjust ctx

	ctx = log.AddEntryToCTX(ctx, logrus.WithFields(map[string]interface{}{
		"file": filePath,
		"from": "MustSetEnvFromFile",
	}))

	log.Infof(ctx, "setting env from file %s", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf(ctx, "error opening file %s", err.Error())
		return errors.Wrap(err, "error opening file")
	}

	defer func() {
		errClose := file.Close()
		if errClose != nil {
			log.Debugf(ctx, "error closing file %s", errClose.Error())
		}
	}()

	// read file

	// parse file

	var cfg Config
	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		log.Errorf(ctx, "error decoding file %s", err.Error())
		return errors.Wrap(err, "error decoding file")
	}

	err = setEnvFromMap(ctx, cfg.ENV, false)
	if err != nil {
		return errors.Wrap(err, "error setting env from file "+filePath)
	}

	return nil
}

func setEnvFromMap(ctx context.Context, env map[string]string, returnError bool) error {
	for k, v := range env {
		k = strings.ToUpper(k)

		err := os.Setenv(k, v)
		if err != nil {
			if returnError {
				log.Errorf(ctx, "error setting env %s", err.Error())
				return errors.Wrap(err, "error setting env")
			}
			log.Debugf(ctx, "error setting env %s", err.Error())
		}

		log.Debugf(ctx, "set env %s", k)
	}

	return nil
}
