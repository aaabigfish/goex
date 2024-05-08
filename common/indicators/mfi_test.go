package indicators

import (
	"testing"
)

func TestMFI(t *testing.T) {
	expectedOutput := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 57.391305628502195, 57.5838678932476, 71.43019990163879, 78.34344988616412, 70.58347417347551, 77.77145098640743, 78.48329056457762, 78.3955242181035, 83.09009292935414, 88.64865035567145, 90.6184164683864, 87.26182941863743, 87.63293817677746, 87.76031875334715, 80.11596021213859, 75.6091146899656, 73.62064381616979, 71.22782415011986, 68.90582863863374, 61.20644434020349, 57.782810110855365, 59.732842735493755, 62.18937897169424, 59.439583512567864, 58.66467168695388, 66.81971524169006, 59.11813131753377, 58.181853524080395, 62.27473417313991, 68.03648278437234, 62.49782587355779, 64.90957454975195, 74.38720203057159, 81.03708500097112, 84.32535095543355, 84.21914275845396, 73.98841421536208, 73.77696117032289, 74.0402128634972, 62.3122908309856, 69.06044933773742, 63.67663318382115, 62.721220127194535, 56.56893034021645, 62.79698546554747, 52.86178485288686, 45.20676507878647, 45.33102302850656, 42.28655280854022, 41.23127823336024, 42.106162736674996, 33.28504964174277, 23.990293579071885, 24.406666398939763, 16.988113964250427, 17.787646626676885, 17.902749361942284, 24.859550790586145, 16.865158082092822, 18.499231417088332, 27.045891878959488, 26.65791910571348, 26.465703802127177, 20.45177881046813, 19.41158621846312, 27.65971328635374, 29.39664882153999, 21.66843538859191, 17.69268244270957, 23.202361336237303, 29.036322125576515, 22.93487989714391, 28.286531494489157, 34.17037934202175, 37.67472850439017, 42.86370988981417, 41.78201383615478, 42.020674467968945, 47.80579816532109, 49.566971273675456, 53.765933843572164, 62.912696512736446, 71.6467177535723, 66.67802323393133, 62.092314913752766, 70.1625876246609}
	ret := MFI(testHigh, testLow, testClose, testVolume, 14)

	if len(ret) != len(expectedOutput) {
		t.Fatalf("unexpected length of return slice %v", len(ret))
	}

	for x := range ret {
		if ret[x] != expectedOutput[x] {
			t.Fatalf("unexpected value returned %v", ret[x])
		}
	}
}