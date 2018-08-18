package handler

import (
	"reflect"
	"testing"
)

func Test_ping(t *testing.T) {
	type args struct {
		req pingReq
	}
	tests := []struct {
		name    string
		args    args
		want    *pingResp
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "t01", args: args{req: pingReq{}}, want: &pingResp{}},
	}
	for _, tt := range tests {
		got, err := ping(tt.args.req)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. ping() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ping() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
