package argorithm

import "testing"

func Test_buyAndSellSolution(t *testing.T) {
	type args struct {
		priceList []float32
	}
	tests := []struct {
		name  string
		args  args
		want  float32
		want1 float32
	}{
		{"case1", args{[]float32{1663.38, 1642.99, 1604.20, 1611.11, 1645.26, 1639.75, 1641.71, 1625.86}}, 1604.2, 1645.26}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := buyAndSellSolution(tt.args.priceList)
			if got != tt.want {
				t.Errorf("buyAndSellSolution() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("buyAndSellSolution() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
