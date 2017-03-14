package system

import (
	"fmt"
	"strings"
	syscall "golang.org/x/sys/unix"

	ps "github.com/mitchellh/go-ps"
	//"os"
	"errors"
	"log"
	"strconv"
	"time"
	"os"
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
			err = procs[0].Kill()
			if err != nil {
				log.Println(err)
			}
		}

		time.Sleep(20 * time.Second)
	}
}

func (p *Processes) FindProcessByPid(pid int) *Process {
	for _, proc := range *p {
		if proc.Pid() == pid {
			return &proc
		}
	}
	return nil
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
	var (
		err       error
		pidString string
	)

	pidString = strconv.Itoa(p.Pid())
	err = syscall.Kill(p.Pid(), syscall.SIGKILL)
	//fmt.Println("killed ", p.Pid(), " with ", syscall.SIGKILL)
	if err != nil {
		return errors.New("Can't kill process with pid " + pidString + ": " + err.Error())
	}
	return nil
}

func (p *Process) Terminate() error {
	var (
		err       error
		pidString string
		pr *os.Process
	)

	pidString = strconv.Itoa(p.Pid())

	pr, err = os.FindProcess(p.Pid())
	if err != nil {
		return errors.New("Can't find process with pid " + pidString + ": " + err.Error())
	}
	err = pr.Signal(os.Interrupt)
	if err != nil {
		return errors.New("Can't terminate process with pid " + pidString + ": " + err.Error())
	}
	return nil
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
