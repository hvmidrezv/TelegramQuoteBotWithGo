package config

import (
	"encoding/json"
	"os"
	"sort"
	"strconv"
)

type Config struct {
	AdminID   int64          `json:"admin_id"`
	Groups    []int64        `json:"groups"`
	Intervals map[string]int `json:"intervals"`
}

var (
	ConfigData     Config
	ConfigFilePath = "config.json"
	DefaultMinutes = 1
)

func Load() error {
	data, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			ConfigData = Config{AdminID: 0, Groups: []int64{}, Intervals: map[string]int{}}
			return Save()
		}
		return err
	}
	return json.Unmarshal(data, &ConfigData)
}

func Save() error {
	data, _ := json.MarshalIndent(ConfigData, "", "  ")
	return os.WriteFile(ConfigFilePath, data, 0644)
}

func HasGroup(chatID int64) bool {
	for _, id := range ConfigData.Groups {
		if id == chatID {
			return true
		}
	}
	return false
}

func AddGroup(chatID int64) error {
	if !HasGroup(chatID) {
		ConfigData.Groups = append(ConfigData.Groups, chatID)
		sort.Slice(ConfigData.Groups, func(i, j int) bool { return ConfigData.Groups[i] < ConfigData.Groups[j] })
		if ConfigData.Intervals == nil {
			ConfigData.Intervals = map[string]int{}
		}
		key := strconv.FormatInt(chatID, 10)
		if _, ok := ConfigData.Intervals[key]; !ok {
			ConfigData.Intervals[key] = DefaultMinutes
		}
		return Save()
	}
	return nil
}

func IntervalFor(chatID int64) int {
	if ConfigData.Intervals == nil {
		return DefaultMinutes
	}
	if v, ok := ConfigData.Intervals[strconv.FormatInt(chatID, 10)]; ok && v > 0 {
		return v
	}
	return DefaultMinutes
}

func SetIntervalFor(chatID int64, minutes int) error {
	if ConfigData.Intervals == nil {
		ConfigData.Intervals = map[string]int{}
	}
	ConfigData.Intervals[strconv.FormatInt(chatID, 10)] = minutes
	return Save()
}
