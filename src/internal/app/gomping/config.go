package gomping

import (
	"errors"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"time"
)

type StartupConfig struct {
	Settings     Settings      `yaml:"settings"`
	TargetGroups []TargetGroup `yaml:"target_group"`
}

type Settings struct {
	GroupPerPage int   `yaml:"group_per_page"`
	RedLine      int64 `yaml:"red_line"`
	YellowLine   int64 `yaml:"yellow_line"`
}

type TargetGroup struct {
	GroupName string `yaml:"group_name"`
	Hosts     []Host `yaml:"target_host"`
}

type Host struct {
	Order    int
	Hostname string `yaml:"hostname"`
	Ip       string `yaml:"ip"`
	Stats    Stats
}

type Stats struct {
	LastRTT string
	Count   float64
	Loss    float64
	sum     time.Duration
}

func NewConfig(c string) (*StartupConfig, error) {
	file, e := ioutil.ReadFile(c)
	if e != nil {
		return nil, e
	}

	var conf StartupConfig
	e = yaml.Unmarshal([]byte(file), &conf)
	if e != nil {
		return nil, e
	}

	if len(conf.TargetGroups) == 0 {
		return nil, errors.New("no target group")
	}

	return &conf, nil
}
