package regex

import "testing"

func TestIdsValidate(t *testing.T) {
	type args struct {
		idsStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test01",
			args:    args{idsStr: ""},
			wantErr: true,
		},
		{
			name:    "test02",
			args:    args{idsStr: "1 "},
			wantErr: true,
		},
		{
			name:    "test03",
			args:    args{idsStr: "1,2,++"},
			wantErr: true,
		},
		{
			name:    "test04",
			args:    args{idsStr: "1,2,2"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IdsValidate(tt.args.idsStr); (err != nil) != tt.wantErr {
				t.Errorf("IdsValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
