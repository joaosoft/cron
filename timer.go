package cron

import (
	"math"
	"strconv"
	"strings"
	"time"
)

/*
Settings

MODE YEAR MONTH DAY HOUR MINUTE SECOND
0    1    1     1   0    0      0

Modes
0 - starts on midnight
1 - starts from now
*/

type setting string
type settingsMode int

const (
	settingMode   setting = "MODE"
	settingYear   setting = "YEAR"
	settingMonth  setting = "MONTH"
	settingDay    setting = "DAY"
	settingHour   setting = "HOUR"
	settingMinute setting = "MINUTE"
	settingSecond setting = "SECOND"

	settingsModeMidnight  settingsMode = 0
	settingsModeRecurring settingsMode = 1
)

var (
	confs = []setting{
		0: settingMode,
		1: settingYear,
		2: settingMonth,
		3: settingDay,
		4: settingHour,
		5: settingMinute,
		6: settingSecond,
	}
)

type settings map[setting]int

type timer struct {
	location *time.Location
	start    time.Time
	settings settings
}

func newTimer(location *time.Location, strSettings string) *timer {
	var start time.Time

	return &timer{
		location: location,
		start:    start,
		settings: parse(strSettings),
	}
}

func parse(strSettings string) settings {
	settings := make(settings)

	splitConfs := strings.Split(strSettings, " ")
	lenStrSettings := len(splitConfs)

	var index int
	for _, key := range confs {
		if index < lenStrSettings {
			settings[key], _ = strconv.Atoi(splitConfs[index])
		} else {
			settings[key] = 0
		}
		index++
	}

	return settings
}

func (timer *timer) Next() time.Duration {
	now := time.Now().In(timer.location)

	switch settingsMode(timer.settings[settingMode]) {
	case settingsModeMidnight:
		year, month, day := now.Date()
		timer.start = time.Date(year, month, day, 0, 0, 0, 0, timer.location).In(timer.location)

		missingTimeFromStartTime := timer.start.AddDate(timer.settings[settingYear], timer.settings[settingMonth], timer.settings[settingDay]).
			Add(time.Hour * time.Duration(timer.settings[settingHour])).
			Add(time.Minute * time.Duration(timer.settings[settingMinute])).
			Add(time.Second * time.Duration(timer.settings[settingSecond]))

		if missingTimeFromStartTime.Before(now) {
			missingTimeFromStartTime = missingTimeFromStartTime.AddDate(0, 0, 1)
		}

		missingTimeFromNow := missingTimeFromStartTime.Sub(now)

		return time.Duration(math.Abs(float64(missingTimeFromNow)))

	case settingsModeRecurring:
		timer.start = now

		return timer.start.
			AddDate(timer.settings[settingYear], timer.settings[settingMonth], timer.settings[settingDay]).
			Add(time.Hour * time.Duration(timer.settings[settingHour])).
			Add(time.Minute * time.Duration(timer.settings[settingMinute])).
			Add(time.Second * time.Duration(timer.settings[settingSecond])).
			Sub(timer.start)
	}

	return 0
}
