package main

import (
	"log"
	"testing"

	"github.com/99designs/aws-vault/keyring"
)

func TestExecCommandRun(t *testing.T) {
	exitCode := 0
	ui := MockUi(func(status int) {
		log.Printf("Exiting with %d\n", status)
		exitCode = status
	})

	kr := &keyring.ArrayKeyring{}
	storeCredentials(kr, "exec_test", "access key", "secret key")

	ExecCommand(ui, ExecCommandInput{
		Keyring: kr,
		Profile: "exec_test",
	})

	if exitCode != 0 {
		t.Fatalf("Bad exit code: %d. %#v", exitCode, ui.Error.(*MockLogger).String())
	}
}

// func TestExecCommandRunWithMissingProfile(t *testing.T) {
// 	ui := new(cli.MockUi)
// 	kr := &keyring.ArrayKeyring{}

// 	c := &ExecCommand{
// 		Ui:              ui,
// 		Keyring:         kr,
// 		sessionProvider: &testSessionProvider{},
// 		profileConfig:   vault.NewProfileConfig(),
// 	}

// 	code := c.Run([]string{"-profile", "llamas", "true"})
// 	if code != 1 {
// 		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
// 	}

// 	if !strings.Contains(ui.OutputWriter.String(), "Profile 'llamas' not found") {
// 		t.Fatalf("bad: %#v", ui.OutputWriter.String())
// 	}
// }
