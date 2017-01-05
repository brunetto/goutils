package system

import (
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"strings"
	"syscall"
	//"os"
	"time"
	"errors"
	"strconv"
)

type Processes []Process
type Process struct {
	ps.Process
}

func List() (Processes, error) {
	var (
		procs     []ps.Process
		processes Processes
		err       error
	)
	procs, err = ps.Processes()
	if err != nil {
		return processes, err
	}
	for _, proc := range procs {
		processes = append(processes, Process{proc})
	}
	return processes, err
}

func MonitorAndKill(processName string) error {
	var (
		processes Processes
		procs     Processes
		err       error
	)

	for {
		processes, err = List()
		if err != nil {
			return errors.New("Can't retrieve processes list: " + err.Error())
		}
		procs = processes.FindProcessByName(processName)
		//procs.Print()

		switch len(procs) {
		// Here just as a temp placeholder
		//case -1:
		//	fmt.Println("No process found.")
		//	os.Exit(0)
		case 1:
			procs[0].Kill()
		}

		time.Sleep(20 * time.Second)
	}
}

func (p *Processes) FindProcessByName(searchString string) Processes {
	var procs Processes
	for _, proc := range *p {
		if strings.Contains(strings.ToLower(proc.Executable()), strings.ToLower(searchString)) {
			procs = append(procs, proc)
		}
	}
	return procs
}

func (p *Processes) Print() {
	fmt.Println("PID\t\tExecutable name")
	fmt.Println("===\t\t===============")
	for _, proc := range *p {
		fmt.Println(proc.Pid(), "\t\t", proc.Executable())
	}
}

func (p *Process) Kill() error {
	var err error
	err = syscall.Kill(p.Pid(), syscall.SIGKILL)
	//fmt.Println("killed ", p.Pid(), " with ", syscall.SIGKILL)
	return errors.New("Can't kill process with pid " + strconv.Itoa(p.Pid()) + ": " + err.Error())
}

func (p *Process) Pid() int {
	return p.Process.Pid()
}

func (p *Process) PPid() int {
	return p.Process.PPid()
}

func (p *Process) Executable() string {
	return p.Process.Executable()
}


