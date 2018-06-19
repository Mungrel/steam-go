package test

import (
	"io/ioutil"
)

const (
	partialPath        = "payloads/"
	news               = "news.json"
	globalAchievements = "global_achievements.json"
)

func getJSON(fileName string) string {
	bytes, err := ioutil.ReadFile(partialPath + fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
