package system

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"github.com/pelletier/go-toml"
)

type Application struct {
	Config   *toml.TomlTree
	Template *template.Template
}

func (application *Application) Init(filename *string) {
	config, err := toml.LoadFile(*filename)
	if err != nil {
		glog.Fatalf("TOML load failed: %s\n", err)
	}

	application.Config = config
}

func (application *Application) LoadTemplates() error {
	var templates []string

	fn := func(path string, f os.FileInfo, err error) error {
		fmt.Println(f.Name())
		if f.IsDir() != true && strings.HasSuffix(f.Name(), ".html") {
			templates = append(templates, path)
		}
		return nil
	}

	err := filepath.Walk(application.Config.Get("general.template_path").(string), fn)

	if err != nil {
		return err
	}

	application.Template = template.Must(template.ParseFiles(templates...))
	return nil
}
