package gridscale

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gridscale/gsclient-go/v3"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

type stepShutdownServer struct {
	client gsclient.ServerOperator
	ui     packer.Ui
}

func (s *stepShutdownServer) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	client := s.client
	ui := s.ui
	// Get server UUID
	serverUUID, ok := state.Get("server_uuid").(string)
	if !ok {
		err := errors.New("cannot convert server_uuid to string")
		ui.Error(err.Error())
		state.Put("error", err)
		return multistep.ActionHalt
	}
	if serverUUID == "" {
		err := errors.New("server_uuid is empty")
		ui.Error(err.Error())
		state.Put("error", err)
		return multistep.ActionHalt
	}
	ui.Say(fmt.Sprintf("Gracefully shutting down server (%s)...", serverUUID))
	//set the shutdown timeout specifically
	shutdownCtx, cancel := context.WithTimeout(context.Background(), serverShutdownTimeoutSecs*time.Second)
	defer cancel()
	err := suppressHTTPErrorCodes(
		client.ShutdownServer(shutdownCtx, serverUUID),
		http.StatusBadRequest,
	)
	if err != nil && err != shutdownCtx.Err() {
		err := fmt.Errorf("Error shutting down server: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	// if the server cannot be shutdown gracefully, try to turn it off
	if err != nil && err == shutdownCtx.Err() {
		ui.Say(fmt.Sprintf("Could not gracefully shutdown server (%s). Trying to turn it off instead...", serverUUID))
		err := suppressHTTPErrorCodes(
			client.StopServer(context.Background(), serverUUID),
			http.StatusBadRequest,
		)
		if err != nil {
			state.Put("error", err)
			ui.Error(fmt.Sprintf(
				"Error shutdown server: %s", err))
			return multistep.ActionHalt
		}
		ui.Say(fmt.Sprintf("Successfully turned off server (%s).", serverUUID))
		return multistep.ActionContinue
	}
	ui.Say(fmt.Sprintf("Gracefully shut down server (%s)", serverUUID))
	return multistep.ActionContinue

}

func (s *stepShutdownServer) Cleanup(state multistep.StateBag) {
}
