package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Config struct {
	Easy   int    `json:"easy"`
	Medium int    `json:"medium"`
	Hard   int    `json:"hard"`
	HScore HScore `json:"heighest_score`
	file   *os.File
}

type HScore struct {
	Time string `json:"time"`
	RC   string `json:"remaining_chances"`
}

func (c *Config) initConfigFile() {
	_, err := os.Stat("ngg.conf")
	if os.IsNotExist(err) {
		c.file, err = os.Create("ngg_config.json")
		if err != nil {
			log.Fatalln(err.Error())
		}
		c.saveChanges()
	} else {
		c.file, err = os.OpenFile("ngg_config.json", os.O_RDWR, 0644)
		if err != nil {
			log.Fatalln(err.Error())
		}
		content, err := io.ReadAll(c.file)
		if err != nil {
			log.Fatalln(err.Error())
		}
		err = json.Unmarshal(content, &c)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}

func (c *Config) saveChanges() {
	if c.Easy == 0 || c.Medium == 0 || c.Hard == 0 {
		c.Easy = 4
		c.Medium = 7
		c.Hard = 10
	}

	jsonData, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Fatalln(err.Error())
	}

	c.file.Truncate(0)
	c.file.Seek(0, 0)

	_, err = c.file.Write(jsonData)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
