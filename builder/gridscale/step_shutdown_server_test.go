package gridscale

import (
	"context"
	"testing"

	"github.com/gridscale/gsclient-go/v3"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

func Test_stepShutdownServer_Run(t *testing.T) {
	type fields struct {
		client gsclient.ServerOperator
		ui     packer.Ui
	}
	type args struct {
		ctx   context.Context
		state multistep.StateBag
	}
	ui := &uiMock{}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   multistep.StepAction
	}{
		{
			name: "success",
			fields: fields{
				client: ServerOperatorMock{},
				ui:     ui,
			},
			args: args{
				state: StateBagMock{state: map[string]interface{}{
					"server_uuid": "ShutdownSuccess",
				}},
			},
			want: multistep.ActionContinue,
		},
		{
			name: "API call fail",
			fields: fields{
				client: ServerOperatorMock{},
				ui:     ui,
			},
			args: args{
				state: StateBagMock{state: map[string]interface{}{
					"server_uuid": "fail",
				}},
			},
			want: multistep.ActionHalt,
		},
		{
			name: "convert server_uuid to string fail",
			fields: fields{
				client: ServerOperatorMock{},
				ui:     ui,
			},
			args: args{
				state: StateBagMock{state: make(map[string]interface{})},
			},
			want: multistep.ActionHalt,
		},
		{
			name: "empty server_uuid",
			fields: fields{
				client: ServerOperatorMock{},
				ui:     ui,
			},
			args: args{
				state: StateBagMock{state: map[string]interface{}{
					"server_uuid": "",
				}},
			},
			want: multistep.ActionHalt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stepShutdownServer{
				client: tt.fields.client,
				ui:     tt.fields.ui,
			}
			if got := s.Run(tt.args.ctx, tt.args.state); got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
