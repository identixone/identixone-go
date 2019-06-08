package person

import (
	"reflect"
	"testing"

	"github.com/identixone/identixone-go/api/const/conf"
)

func TestPersonaCreateRequest_IsValid(t *testing.T) {
	tests := []struct {
		name    string
		pc      *PersonaCreateRequest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pc.IsValid(); (err != nil) != tt.wantErr {
				t.Errorf("PersonaCreateRequest.IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPersonaCreateRequest_SetSource(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		pc   *PersonaCreateRequest
		args args
		want *PersonaCreateRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.SetSource(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PersonaCreateRequest.SetSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersonaCreateRequest_SetFacesize(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		pc   *PersonaCreateRequest
		args args
		want *PersonaCreateRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.SetFacesize(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PersonaCreateRequest.SetFacesize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersonaCreateRequest_SetCreateOnHa(t *testing.T) {
	type args struct {
		val bool
	}
	tests := []struct {
		name string
		pc   *PersonaCreateRequest
		args args
		want *PersonaCreateRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.SetCreateOnHa(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PersonaCreateRequest.SetCreateOnHa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersonaCreateRequest_SetCreateOnJunk(t *testing.T) {
	type args struct {
		val bool
	}
	tests := []struct {
		name string
		pc   *PersonaCreateRequest
		args args
		want *PersonaCreateRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.SetCreateOnJunk(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PersonaCreateRequest.SetCreateOnJunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersonaCreateRequest_SetAsm(t *testing.T) {
	type args struct {
		val bool
	}
	tests := []struct {
		name string
		pc   *PersonaCreateRequest
		args args
		want *PersonaCreateRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.SetAsm(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PersonaCreateRequest.SetAsm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersonaCreateRequest_SetLiveness(t *testing.T) {
	type args struct {
		val bool
	}
	tests := []struct {
		name string
		pc   *PersonaCreateRequest
		args args
		want *PersonaCreateRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.SetLiveness(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PersonaCreateRequest.SetLiveness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersonaCreateRequest_SetCreateLivenessOnly(t *testing.T) {
	type args struct {
		val bool
	}
	tests := []struct {
		name string
		pc   *PersonaCreateRequest
		args args
		want *PersonaCreateRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.SetCreateLivenessOnly(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PersonaCreateRequest.SetCreateLivenessOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPersonaCreateRequest(t *testing.T) {
	type args struct {
		photoPath  string
		sourceName string
	}
	tests := []struct {
		name    string
		args    args
		want    PersonaCreateRequest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPersonaCreateRequest(tt.args.photoPath, tt.args.sourceName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPersonaCreateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPersonaCreateRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSearch(t *testing.T) {
	type args struct {
		photoPath string
		asm       bool
		liveness  bool
	}
	tests := []struct {
		name    string
		args    args
		want    Search
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSearch(tt.args.photoPath, tt.args.asm, tt.args.liveness)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchResult_Success(t *testing.T) {
	tests := []struct {
		name string
		sr   SearchResult
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sr.Success(); got != tt.want {
				t.Errorf("SearchResult.Success() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReinitImageRequest_IsValid(t *testing.T) {
	tests := []struct {
		name    string
		ri      *ReinitImageRequest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ri.IsValid(); (err != nil) != tt.wantErr {
				t.Errorf("ReinitImageRequest.IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewReinitImageRequest(t *testing.T) {
	type args struct {
		filePath string
		idxid    string
	}
	tests := []struct {
		name    string
		args    args
		want    ReinitImageRequest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewReinitImageRequest(tt.args.filePath, tt.args.idxid)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewReinitImageRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReinitImageRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReinitImageRequest_SetFacesize(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		ri   *ReinitImageRequest
		args args
		want *ReinitImageRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ri.SetFacesize(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReinitImageRequest.SetFacesize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReinitImageRequest_SetLiveness(t *testing.T) {
	type args struct {
		val bool
	}
	tests := []struct {
		name string
		ri   *ReinitImageRequest
		args args
		want *ReinitImageRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ri.SetLiveness(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReinitImageRequest.SetLiveness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReinitImageRequest_SetReinitLivenessOnly(t *testing.T) {
	type args struct {
		val bool
	}
	tests := []struct {
		name string
		ri   *ReinitImageRequest
		args args
		want *ReinitImageRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ri.SetReinitLivenessOnly(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReinitImageRequest.SetReinitLivenessOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReinitImageRequest_SetConf(t *testing.T) {
	type args struct {
		val conf.Conf
	}
	tests := []struct {
		name string
		ri   *ReinitImageRequest
		args args
		want *ReinitImageRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ri.SetConf(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReinitImageRequest.SetConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
