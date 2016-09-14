package config

import (
	"log"

	gcfg "gopkg.in/gcfg.v1"
)

// Storage contains needed access keys and enpoints for accessing
// the video and thumbnail storage service.
type StorageConfiguration struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	Database        string
	DatabaseLog     bool
}

// RenderingConfiguration specifies the directory template files
// will be located in.
type RenderingConfiguration struct {
	// TemplatesDir is the directory that template files can be found int.
	// Templates must end with the .tmpl file extension.
	TemplatesDir string
}

// SessionConfiguration defines the key that is used to encrypt all
// user sessions.
type SessionConfiguration struct {
	Key string
}

// Configuration encapsulates configuration of all submodules within the
// application.
type Configuration struct {
	Storage   StorageConfiguration
	Rendering RenderingConfiguration
	Session   SessionConfiguration
}

var conf Configuration

func init() {
	log.Println("Parsing configuration...")
	err := gcfg.ReadFileInto(&conf, "config.ini")
	if err != nil {
		log.Fatal("Could not read configuration : ", err)
	}
}

// GetStorageConfiguration returns the storage configuration specified in
// the config.ini file
func GetStorageConfiguration() StorageConfiguration {
	return conf.Storage
}

// GetRenderingConfiguration returns the Rednering storage configuration
// specified in the config.ini file
func GetRenderingConfiguration() RenderingConfiguration {
	return conf.Rendering
}

// GetSessionConfiguration returns the session configuration specified in
// the config.ini file
func GetSessionConfiguration() SessionConfiguration {
	return conf.Session
}
