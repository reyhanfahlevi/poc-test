package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/reyhanfahlevi/pkg/go/filexporter"
)

func TestInit(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name   string
		args   args
		mockFn func(a args)
	}{
		{
			name: "Test Success",
			args: args{
				opts: []Option{
					WithConfigFile("config.json", ""),
				},
			},
			mockFn: func(a args) {
				_ = filexporter.ExportJSON(Config{}, "", "config.json")
			},
		}, {
			name: "Test Failed Parse Config - File Not Found",
			args: args{
				opts: []Option{
					WithConfigFile("configs.json", ""),
				},
			},
			mockFn: func(a args) {
				_ = ioutil.WriteFile("config.json", []byte(`{`), os.ModePerm)
			},
		},
		{
			name: "Test Failed Parse Vault Config",
			args: args{
				opts: []Option{
					WithConfigFile("config.json", ""),
				},
			},
			mockFn: func(a args) {
				_ = ioutil.WriteFile("config.json", []byte(`{}`), os.ModePerm)
			},
		},
		{
			name: "Test Failed Parse Vault Config",
			args: args{
				opts: []Option{
					WithConfigFile("config.json", ""),
				},
			},
			mockFn: func(a args) {
				_ = ioutil.WriteFile("config.json", []byte(`{}`), os.ModePerm)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			_ = Init(tt.args.opts...)

			Get()
			config = nil
			Get()

			_ = os.RemoveAll("config.json")
		})
	}
}
