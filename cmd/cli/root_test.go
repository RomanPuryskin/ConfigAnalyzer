package cli

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRootCmd_JSONFileSuccess(t *testing.T) {

	file, err := os.Open("../../test_data/test_json.json")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	cmd := RootCmd()

	cmd.SetArgs([]string{file.Name()})

	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)

	err = cmd.Execute()

	require.NotNil(t, err)
	require.ErrorContains(t, err, "проблем обнаружено")
}

func TestRootCmd_YAMLFileSuccess(t *testing.T) {

	file, err := os.Open("../../test_data/test_yaml.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	cmd := RootCmd()

	cmd.SetArgs([]string{file.Name()})

	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)

	err = cmd.Execute()

	require.NotNil(t, err)
	require.ErrorContains(t, err, "проблем обнаружено")
}

func TestRootCmd_SilentFileSuccess(t *testing.T) {

	file, err := os.Open("../../test_data/test_yaml.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	cmd := RootCmd()

	cmd.SetArgs([]string{"--silent", file.Name()})

	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)

	err = cmd.Execute()

	require.Nil(t, err)
}
