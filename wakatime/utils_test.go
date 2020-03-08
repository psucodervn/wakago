package wakatime

import "testing"

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: string("success"), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetApiKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApiKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				t.Logf("api_key: %v\n", got)
			}
		})
	}
}

func Test_getApiKeyFromFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: string("success"), args: args{filePath: "~/.wakatime.cfg"}, wantErr: false},
		{name: string("failure"), args: args{filePath: "~/fail.cfg"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getApiKeyFromFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("getApiKeyFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				t.Logf("api_key: %v\n", got)
			}
		})
	}
}
