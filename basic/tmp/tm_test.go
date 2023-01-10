package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestTm(t *testing.T){
var (
	str = "ok"
 ptrStr = &str
 wstr = "wrong"
 ptrWstr = &wstr
)
	tests := []struct{
		name string
		arg string
		want *string
	}{
		{
			name: "success",
			want: ptrStr,
		},
		{
			name: "fail",
			want: ptrWstr,
		},
	}
	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			got := tm("ok")
			if assert.Equal(t, tt.want, got){
				t.Errorf("want = %v got = %v\n", tt.want, got)
			}
		})
	}
}
