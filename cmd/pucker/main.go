package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/WOT-Lemons/gumble/gumble"
	_ "github.com/WOT-Lemons/gumble/opus"
	"github.com/WOT-Lemons/pucker"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Command line flags
	server := flag.String("server", "comms.wotlemons.com:64738", "the server to connect to")
	username := flag.String("username", "252", "the username of the client")
	password := flag.String("password", "SECRET", "the password of the server")
	insecure := flag.Bool("insecure", true, "skip server certificate verification")
	certificate := flag.String("certificate", "", "PEM encoded certificate and private key")
	channel := flag.String("channel", "Race Ops", "mumble channel to join by default")

	flag.Parse()

	// Initialize
	b := pucker.Pucker{
		Config:      gumble.NewConfig(),
		Address:     *server,
		ChannelName: *channel,
	}

	// if no username specified, quit
	if len(*username) == 0 {
		fmt.Fprintf(os.Stderr, "%s\nNo username specified.\n")
		os.Exit(1)
	}

	b.Config.Username = *username
	b.Config.Password = *password

	if *insecure {
		b.TLSConfig.InsecureSkipVerify = true
	}
	if *certificate != "" {
		cert, err := tls.LoadX509KeyPair(*certificate, *certificate)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		b.TLSConfig.Certificates = append(b.TLSConfig.Certificates, cert)
	}

	b.Init()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	exitStatus := 0

	<-sigs
	b.CleanUp()

	os.Exit(exitStatus)
}
