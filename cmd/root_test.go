package cmd

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestMakeRootCmd(t *testing.T) {
	cmd := MakeRootCmd()

	assert.NotNil(t, cmd)
	assert.IsType(t, &cobra.Command{}, cmd)
	assert.Equal(t, "moducate", cmd.Use)
	assert.NotEmpty(t, cmd.Short)
}

func TestExecute(t *testing.T) {
	b := bytes.NewBufferString("")

	cmd := MakeRootCmd()
	cmd.SetOut(b)

	assert.NotPanics(t, func() {
		Execute(cmd, io.Discard)
	})

	out, err := ioutil.ReadAll(b)
	assert.NoError(t, err)

	assert.Equal(t, "Run and manage Moducate", strings.Split(string(out), "\n")[0])
}
