package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-10-2

 tcp,udp, 4tcp, 6tcp network analyzer using golang and awk.
 it prepares the files for network struct that i will code
 tomorrow. The worst part of the nextwork file is that the
 file is not being well written using the linux file.io utils
 and you can implement the regular expression as the spaces
 between the space is not the same, so i implemented the awk
 for faster code write.

*/

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:  "flags",
	Long: "Creating the network structs",
	Run:  netFunc,
}

var inputfile string

func init() {
	rootCmd.Flags().StringVarP(&inputfile, "input", "i", "path to the networkfile", "network")
}

func netFunc(cmd *cobra.Command, args []string) {
	t := time.Now()
	tstring := t.String()
	command, err := exec.Command("awk", "{print $1}", inputfile).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(command) + tstring)
	pid, err := exec.Command("awk", "{print $2}", inputfile).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(pid) + tstring)
	user, err := exec.Command("awk", "{print $3}", inputfile).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(user) + tstring)
	typed, err := exec.Command("awk", "{print $5}", inputfile).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(typed) + tstring)
	device, err := exec.Command("awk", "{print $6}", inputfile).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(device) + tstring)
	node, err := exec.Command("awk", "{print $8}", inputfile).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(node) + tstring)
	name, err := exec.Command("awk", "{print $9}", inputfile).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(name) + tstring)
}
