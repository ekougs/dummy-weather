package util

import (
	"os"
	"testing"
	"time"
)

var tempDataUtils, _ = NewDataUtils("../resources/cities.json")
var tempProvider, _ = NewTempProvider(tempDataUtils)

func TestMain(m *testing.M) {
	retCode := m.Run()

	tearDown()

	os.Exit(retCode)
}

func tearDown() {
	os.Remove("../resources/DKR.json")
	os.Remove("../resources/NYC.json")
	os.Remove("../resources/PAR.json")
}

func TestShouldReturnARealisticTempForDKR(t *testing.T) {
	location, _ := time.LoadLocation("Africa/Dakar")
	temp, error := tempProvider.generate("DKR", time.Date(2015, 4, 16, 13, 00, 00, 00, location))
	if error != nil {
		t.Error(error)
	}
	// 17 23
	if temp < 15 || temp > 26 {
		t.Errorf("Actual generated temp '%d' should be between 15 and 26", temp)
	}
}

func TestShouldReturnSameTempFor2CallsWithSameParameters(t *testing.T) {
	location, _ := time.LoadLocation("Africa/Dakar")
	var temp1, temp2 int
	var err error
	temp1, err = tempProvider.Get("DKR", time.Date(2015, 4, 21, 13, 00, 00, 00, location))
	if err != nil {
		t.Error(err)
	}
	temp2, err = tempProvider.Get("DKR", time.Date(2015, 4, 21, 13, 00, 00, 00, location))
	if err != nil {
		t.Error(err)
	}
	if temp1 != temp2 {
		t.Errorf("Value should be generated only once. First temp %d should be equal to second %d", temp1, temp2)
	}
}

func TestShouldReturnSameTempFor2CallsWithSameParametersBeforeDayLightSavingTimeIncrement(t *testing.T) {
	location, _ := time.LoadLocation("America/New_York")
	var temp1, temp2 int
	var err error
	temp1, err = tempProvider.Get("NYC", time.Date(2014, 2, 21, 10, 00, 00, 00, location))
	if err != nil {
		t.Error(err)
	}
	temp2, err = tempProvider.Get("NYC", time.Date(2014, 2, 21, 10, 00, 00, 00, location))
	if err != nil {
		t.Error(err)
	}
	if temp1 != temp2 {
		t.Errorf("Value should be generated only once. First temp %d should be equal to second %d", temp1, temp2)
	}
}

func TestShouldReturnSameTempFor2CallsWithSameParametersAfterDayLightSavingTimeIncrement(t *testing.T) {
	location, _ := time.LoadLocation("Europe/Paris")
	var temp1, temp2 int
	var err error
	temp1, err = tempProvider.Get("PAR", time.Date(2013, 4, 16, 10, 00, 00, 00, location))
	if err != nil {
		t.Error(err)
	}
	temp2, err = tempProvider.Get("PAR", time.Date(2013, 4, 16, 10, 00, 00, 00, location))
	if err != nil {
		t.Error(err)
	}
	if temp1 != temp2 {
		t.Errorf("Value should be generated only once. First temp %d should be equal to second %d", temp1, temp2)
	}
}
