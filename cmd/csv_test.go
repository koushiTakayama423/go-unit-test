package cmd

import (
	"reflect"
	"testing"
)

func TestInput_csv(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name:    "正常系",
			args:    args{path: "/input/input.csv"},
			want:    [][]string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Input_csv(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input_csv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Input_csv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutput_csv(t *testing.T) {
	type args struct {
		records [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Output_csv(tt.args.records); (err != nil) != tt.wantErr {
				t.Errorf("Output_csv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
