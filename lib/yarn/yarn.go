package yarn

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

// If there is a open ticket, does nothing, else kinits a new session
func kCheck() {
	if err := exec.Command("klist", "-s").Run(); err != nil { // returns no err if there's an open ticket
		log.Println("Kerberos login required")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("User: ")
		user, _ := reader.ReadString('\n')
		user = user[:len(user)-1]

		kInit(user)
	}
}

func kInit(user string) {
	cmd := exec.Command("kinit", user)
	// stdout, err := cmd.StdoutPipe()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln("InPipe")
	}

	err = cmd.Start()
	if err != nil {
		log.Fatalln("Could not execute")
	}

	go io.Copy(stdin, os.Stdin)

	err = cmd.Wait()
	if err != nil {
		log.Fatalln("Failed to authenticate.")
	}
}

// StartCloud starts a yarn cloud by shelling out to hadoop
//
// This process needs to store the job-ID to kill the process in the future
func StartCloud(size int) {

	kCheck()

	cmdArgs := []string{
		"jar",                //
		"h2odriver.jar",      //FIXME: This should be a pack method
		"-n",                 //
		strconv.Itoa(size),   //
		"-mapperXmx",         //
		"10g",                // FIXME: This may be modifialbe down the road
		"-output",            //
		"steam_temp_out_001", // FIXME: This should be random and stored with the cloud
		"-disown",            //
	}

	log.Println("Attempting to start cloud...")
	cmdOut, err := exec.Command("hadoop", cmdArgs...).CombinedOutput()

	if err != nil {
		log.Println("Failed to launch hadoop.")
		log.Println("\n" + string(cmdOut)) // This captures error from the drive.jar
		log.Fatalln(os.Stderr, err)        // This captures erros from Stderr
	}
	hpOut := (string(cmdOut))
	// Capture only the address and ID respectively
	reNode := regexp.MustCompile(`H2O node (\d+\.\d+\.\d+\.\d+:\d+)`)
	reApID := regexp.MustCompile(`application_(\d+_\d+)`)

	for _, node := range reNode.FindAllStringSubmatch(hpOut, size) {
		address := node[1]
		log.Println("Node started at:", address)
	}
	apID := reApID.FindStringSubmatch(hpOut)[1]

	fmt.Println("")
	log.Println("Started cloud with ID:", apID)
}

// StopCloud kills a hadoop cloud by shelling out a command based on the job-ID
func StopCloud(id string) {

	kCheck()

	log.Println("Attempting to stop cloud...")
	cmdStop := exec.Command("hadoop", "job", "-kill", "job_"+id)
	if out, err := cmdStop.CombinedOutput(); err != nil {
		log.Println("Failed to shutdown hadoop.")
		log.Println("\n" + string(out))
		log.Fatalln(os.Stderr, err)
	}
	cmdClean := exec.Command("hadoop", "fs", "-rmdir", "steam_temp_out_001") //FIXME: this should use above saved dir
	log.Println("Stopped cloud:", "job_"+id)
	if out, err := cmdClean.Output(); err != nil {
		log.Fatalln("Failed to remove outdir.")
		log.Println("\n" + string(out))
		log.Fatalln(os.Stderr, err)
	}
}
