package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func windows(OperatingSystem string) {

	var OS string = OperatingSystem
	string_operating_system := strings.ToUpper(OS)

	fmt.Printf("Hello User, Your physical machine is running on %v, This GO / Executable will run a FLASK SERVER\n", string_operating_system)

	fmt.Println("Enter Your Flask running file (eg. flask_running.py): ")
	var flask_run_file string
	fmt.Scanln(&flask_run_file)
	fmt.Println("Enter Your Flask server running TCP/IP Port (eg. 5000): ")
	var flask_run_port string
	fmt.Scanln(&flask_run_port)

	os.Setenv("FLASK_APP", flask_run_file)
	os.Setenv("FLASK_RUN_PORT", flask_run_port)
	os.Setenv("FLASK_ENV", "development")

	var flask_program string = "flask"
	var flask_args string = "run"

	cmd := exec.Command(flask_program, flask_args)
	fmt.Printf("\nFlask back-end is running on: \n\n\t File: %v\n \t TCP/IP PORT: %v\n \t Flask environment: %v\n \n\t click on keyboard Command+C to stop TCP connection", os.Getenv("FLASK_APP"), os.Getenv("FLASK_RUN_PORT"), os.Getenv("FLASK_ENV"))

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}

func mac(OperatingSystem string) {

	var OS string = OperatingSystem
	string_operating_system := strings.ToUpper(OS)

	fmt.Printf("Hello User, Your physical machine is running on %v, This GO / Executable will run a FLASK SERVER\n", string_operating_system)

	fmt.Println("Enter Your Flask running file (eg. flask_running.py): ")
	var flask_run_file string
	fmt.Scanln(&flask_run_file)
	fmt.Println("Enter Your Flask server running TCP/IP Port (eg. 5000): ")
	var flask_run_port string
	fmt.Scanln(&flask_run_port)

	os.Setenv("FLASK_APP", flask_run_file)
	os.Setenv("FLASK_RUN_PORT", flask_run_port)
	os.Setenv("FLASK_ENV", "development")

	var flask_program string = "flask"
	var flask_args string = "run"

	cmd := exec.Command(flask_program, flask_args)
	fmt.Printf("\nFlask back-end is running on: \n\n\t File: %v\n \t TCP/IP PORT: %v\n \t Flask environment: %v\n \n\t click on keyboard Command+C to stop TCP connection", os.Getenv("FLASK_APP"), os.Getenv("FLASK_RUN_PORT"), os.Getenv("FLASK_ENV"))

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}

func main() {
	var os string = runtime.GOOS
	var osArgument string = os
	switch os {
	case "windows":
		windows(osArgument)
	case "darwin":
		mac(osArgument)
	}
}
