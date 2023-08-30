package execute

import (
	"log"
	"os"
	"os/exec"
)

type Executor interface {
	Execute(cmd string, options []string, environment []string) error
	GetCommand() []string
}
type SystemExecutor struct {
}

func (s *SystemExecutor) Execute(cmd string, options []string, environment []string) error {
	command := exec.Command(cmd, options...)
    envs := command.Environ()
    envs = append(envs, environment...)
    // fmt.Println(envs)
    command.Env = envs
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	log.Printf("Executing '%s' on your operating system", command.String())
	err := command.Run()
	if err != nil {
		return err
	}
	return nil
}
func (s *SystemExecutor) GetCommand() []string {
	return []string{}
}

func GetExecutor() Executor {
	return &SystemExecutor{}
}
