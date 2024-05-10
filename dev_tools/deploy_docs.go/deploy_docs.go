// deploy_docs tool for deploying ciscoconfparse docs to pennington.net
package main

import (
	"flag"
	"fmt"

	// goph is an ssh client...
	"github.com/melbahja/goph"

	// logoru provides fancy logging similar to python loguru
	"github.com/gleich/logoru"
)

// Declare all variables here...
var (
	dochost string
)

func main() {

	// define CLI flags here... '--dochost' is saved as string variable -> dochost
	flag.StringVar(&dochost, "dochost", "chestnut.he.net", "host to upload docs to.")
	flag.Parse()

	fmt.Printf("A script to upload CiscoConfParse docs to %s.\n", dochost)

	logoru.Info("Starting new CiscoConfParse documentation upload to " + dochost + ".")

	// Start a new goph ssh connection with private key auth...
	logoru.Info("    Initialize ssh key-auth")
	auth, err := goph.Key("/home/mpenning/.ssh/id_rsa", "")
	if err != nil {
		logoru.Critical(err.Error())
	}

	// Get an ssh client instance with private-key / public-key auth...
	logoru.Info("    ssh into remote webhost " + dochost + ".")
	client, err := goph.New("mpenning", dochost, auth)
	if err != nil {
		logoru.Critical(err.Error())
	}

	// Run a shell command via the 'client' ssh connection to dochost...
	logoru.Info("    Remove old documentation files")
	_, err = client.Run("bash -c 'cd public_html/py/ciscoconfparse2/ && rm -rf *'")
	if err != nil {
		logoru.Critical(err.Error())
	}

	// Upload a file to the remote server...
	logoru.Info("    Start copying ciscoconfparse2 docs to remote ssh server")
	err = client.Upload("/home/mpenning/ccp2.tar.gz", "public_html/py/ciscoconfparse2/ccp2.tar.gz")
	if err != nil {
		logoru.Critical(err.Error())
	}
	logoru.Info("    Finished copying tarball to remote ssh server")

	logoru.Info("    extract CiscoConfParse documentation tarball")
	_, err = client.Run("bash -c 'cd public_html/py/ciscoconfparse2/ && tar xvfz ./ccp2.tar.gz'")
	if err != nil {
		logoru.Critical(err.Error())
	}

	// NOT_NEEDED -> defer client.Close()
}
