package indicators

import (
	"testing"
)

func TestRSI(t *testing.T) {
	expectedOutput := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 44.45812807881774, 57.81971547743184, 66.33224703902086, 62.410570378951434, 55.129414672030386, 62.851691995121485, 57.84132833408279, 60.85043218545786, 58.77070214596548, 69.32413626576206, 69.219841514168, 66.46456741514962, 68.9198126180077, 69.13187699257008, 62.9650257819882, 60.99024718971327, 62.62870270755676, 60.79121291585303, 53.255313760441034, 54.2176172324083, 51.78358556264556, 57.564017655992586, 63.11257944941977, 70.10929881845185, 67.10404667822849, 69.75969074308522, 65.13560635396576, 65.83434013771698, 64.48598977483661, 63.279559284023556, 59.540761453486745, 67.02365684509118, 68.86379860505821, 69.51990106843671, 70.85617433516062, 74.214832026121, 64.65641688522238, 70.03203756033204, 70.93868400798598, 67.62254141999718, 69.36095550174697, 57.88254804234278, 58.09344708906193, 53.35836305063451, 61.0852776169437, 50.25569652546954, 50.52874223675009, 51.91929579143416, 51.39677503109441, 56.409425443648566, 50.506205181146356, 44.88824485352728, 37.992508451545184, 38.70841558348461, 37.30155209135433, 35.15095998456748, 35.17861412438783, 43.50629397955871, 41.13413007594025, 41.10177982709933, 47.52048579736041, 49.17905230600252, 44.647501933864085, 33.60762346215599, 32.62349991858598, 32.10067562134204, 33.21483119378425, 15.64202581002264, 26.253147633049778, 24.339367867693923, 26.55564236928525, 25.14370724684083, 28.96648427155869, 30.047516376779853, 39.428944871096924, 39.75439736577097, 39.5775780908228, 36.63733183563066, 44.67575541338243, 47.57231186398134, 46.838768586351065, 47.61845153818703, 43.73878228764954, 42.54568121807077, 38.985757961259026, 40.061357253493235}

	ret := RSI(testClose, 14)
	if len(ret) != len(expectedOutput) {
		t.Fatalf("unexpected length of return slice %v", len(ret))
	}

	for x := range ret {
		if ret[x] != expectedOutput[x] {
			t.Fatalf("unexpected value returned %v", ret[x])
		}
	}

	ret = RSI(testClose, 1)
	for x := range ret {
		if ret[x] != 0 {
			t.Fatalf("unexpected value returned %v", ret[x])
		}
	}
}