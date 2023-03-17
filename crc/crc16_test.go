package crc

import (
	"reflect"
	"testing"
)

func Test_CRC16XMODEM(t *testing.T) {
	tests := []struct {
		name string
		args []byte
		want uint16
	}{
		{"test1", []byte("123"), 0x9752},
		{"test2", []byte("123456789"), 0x31C3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenCRC16XMODEM(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genCrc16XMODEM() = % X, want % X", got, tt.want)
			}
		})
	}
}

func Test_CRC16MODBUS(t *testing.T) {
	tests := []struct {
		name string
		args []byte
		want uint16
	}{
		{"test1", []byte("123"), 0x7A75},
		{"test2", []byte("123456789"), 0x4B37},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenCRC16MODBUS(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genCrc16XMODEM() = % X, want % X", got, tt.want)
			}
		})
	}
}
