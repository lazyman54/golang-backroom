package generic

import (
	"reflect"
	"testing"
)

func TestGetMyData(t *testing.T) {
	type args struct {
		dataType int64
	}
	tests := []struct {
		name string
		args args
		want DataInterface
	}{
		{
			name: "test1",
			args: args{dataType: 1},
			want: &DataType1{data: 0},
		},
		{
			name: "test2",
			args: args{dataType: 2},
			want: &DataType2{data: 0},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMyData(tt.args.dataType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMyData() = %v, want %v", got, tt.want)
			}
		})
	}
}
