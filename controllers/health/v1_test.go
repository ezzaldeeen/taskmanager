package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterRouteGrp(t *testing.T) {
	type args struct {
		g *echo.Group
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Case 1",
			args: args{g: echo.New().Group("")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterRouteGrp(tt.args.g)
		})
	}
}

func Test_get(t *testing.T) {
	type args struct {
		c echo.Context
	}
	// fake request, and recorder
	req := httptest.NewRequest(http.MethodGet, "/path", nil)
	rec := httptest.NewRecorder()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			args:    args{c: echo.New().NewContext(req, rec)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := get(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
