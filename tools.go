package main

import (
	"encoding/json"
	"os/exec"
)

type status struct {
	Status string                 `json:"status"`
	Out    string                 `json:"stdout"`
	Result map[string]interface{} `json:"jsonst"`
}

func (s *status) putExec(p string) error {

	cmd := exec.Command(EXEC_PATH+PUT_CMD, p)
	cmd.Dir = EXEC_PATH
	out, err := cmd.CombinedOutput()

	if err != nil {
		s.Out = err.Error()
		return err
	}

	s.Out = string(out)
	json.Unmarshal(out, &s.Result)

	return nil
}

func (s *status) getStatus(id string, key string, value string) error {

	cmdArguments := []string{id, key, value}
	cmd := exec.Command(EXEC_PATH+GET_CMD, cmdArguments...)
	cmd.Dir = EXEC_PATH
	out, err := cmd.CombinedOutput()

	if err != nil {
		s.Out = err.Error()
		return err
	}

	s.Out = string(out)
	json.Unmarshal(out, &s.Result)

	return nil
}
