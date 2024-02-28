package entities

import "testing"

func TestTradeGoodIllegal(t *testing.T) {
	tg := TradeGood{
		Id: "61",
	}

	if !tg.IsIllegal() {
		t.Error("tradegood '66' expected to be illegal, is not")
	}

	tg.Id = "41"
	if tg.IsIllegal() {
		t.Error("tradegood '41' expected to be legal, is not")
	}
}

func TestGetAvailableForCodes(t *testing.T) {
	tradeCodes := GetAvailableForCodes(false, 0, 0, "Pi")
	for _, tradeCode := range tradeCodes {
		if tradeCode.IsIllegal() {
			t.Errorf("trade code %q has id %q, but returned in non-illegal availability list", tradeCode.Type, tradeCode.Id)
		}
	}
	if len(tradeCodes) != 6 {
		t.Errorf("expected 6 trade codes, got %d", len(tradeCodes))
	}

	tradeCodes = GetAvailableForCodes(true, 0, 0, "Pi")
	foundIllegal := false
	for _, tradeCode := range tradeCodes {
		if tradeCode.IsIllegal() {
			foundIllegal = true
			break
		}
	}
	if !foundIllegal {
		t.Error("looking for illegal codes, found none")
	}

	if len(tradeCodes) != 7 {
		t.Errorf("expected 7 trade codes, got %d", len(tradeCodes))
	}
}
