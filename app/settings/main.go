package settings

import (
	"log"

	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type settings struct {
	Name string
}

var Settings settings

func BootSettings() {
	fileBuffer, err := ioutil.ReadFile("./settings.yaml")

	if err != nil {
		path, _ := filepath.Abs("./")
		log.Fatalf("Arquivo de configurações não encontrado em %s", path)
	}

	settings := settings{}

	err = yaml.Unmarshal(fileBuffer, &settings)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	Settings = settings
}
