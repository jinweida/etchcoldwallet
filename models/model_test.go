package models

import "testing"

func TestCreateFund(t *testing.T) {
	Setup()
	CreateFund("0x28b546ba64d4099fbfc2bb0cbcfa7e637189b640")
}
