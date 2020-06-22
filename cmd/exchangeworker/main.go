package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ngoline/quantocustaobitcoin/internal/constants"
	"github.com/ngoline/quantocustaobitcoin/internal/exchange"
	"github.com/ngoline/quantocustaobitcoin/internal/exchange/mercadobitcoin"
	"github.com/takama/daemon"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	// default exchange name
	defaultExchange = "mercadobitcoin"

	// name of the service
	name        = "qcobtcexchange"
	description = "QCOBTC Bitcoin Exchange Service"

	// port which daemon should be listen
	defaultPort = ":9978"

	usage = `Usage: exchangeworker <exchangename> <action>
	<exchangename>:
		mercadobitcoin
	<action>:
		install	|	Installs the service
		remove	|	Uninstall the service
		star	|	Start the service
		stop	|	Stop the service
		status	|	Displays the service's status
	
	pass no argument <action> to simply run the app.`
)

var stdlog, errlog *log.Logger

// Service has embedded daemon
type Service struct {
	daemon.Daemon
}

// GetTradeData using Mercado Bitcoin API
func GetTradeData(worker exchange.Worker) {
	db := worker.Init()
	defer db.Close()

	fmt.Printf("Collecting trade data for %s...\n", worker.GetName())

	for true {
		worker.SyncData(db)
		time.Sleep(30 * time.Second)
	}
}

// Manage by daemon commands or run the daemon
func (service *Service) Manage() (string, error) {
	// Set defaults
	exchangeName := defaultExchange
	port := defaultPort
	var worker exchange.Worker
	worker = mercadobitcoin.NewExchange()

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		exchangeName = os.Args[1]

		switch exchangeName {
		case "mercadobitcoin":
			break
		// Not used for this sample
		// case "bitcointrade":
		// 	port = ":9979"
		// 	worker = bitcointrade.NewExchange()
		// 	break
		default:
			return usage, nil
		}

		if len(os.Args) > 2 {
			command := os.Args[2]

			switch command {
			case "install":
				service.SetTemplate(constants.SystemDConfig)
				return service.Install(exchangeName)
			case "remove":
				return service.Remove()
			case "start":
				return service.Start()
			case "stop":
				return service.Stop()
			case "status":
				return service.Status()
			default:
				return usage, nil
			}
		}
	}

	go GetTradeData(worker)

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Set up listener for defined host and port
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return "Possibly was a problem with the port binding", err
	}

	// set up channel on which to send accepted connections
	listen := make(chan net.Conn, 100)
	go acceptConnection(listener, listen)

	// loop work cycle with accept cdefaultPonnections or interrupt
	// by system signal
	for {
		select {
		case conn := <-listen:
			go handleClient(conn)
		case killSignal := <-interrupt:
			stdlog.Println("Got signal:", killSignal)
			stdlog.Println("Stoping listening on ", listener.Addr())
			listener.Close()
			if killSignal == os.Interrupt {
				return "Daemon was interruped by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}

	// never happen, but need to complete code
	return usage, nil
}

// Accept a client connection and collect it in a channel
func acceptConnection(listener net.Listener, listen chan<- net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		listen <- conn
	}
}

func handleClient(client net.Conn) {
	for {
		buf := make([]byte, 4096)
		numbytes, err := client.Read(buf)
		if numbytes == 0 || err != nil {
			return
		}
		client.Write(buf[:numbytes])
	}
}

func init() {
	stdlog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	errlog = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}

func main() {
	if len(os.Args) == 1 {
		errlog.Println(usage)
		os.Exit(1)
	}

	srv, err := daemon.New(name+os.Args[1], description+" "+os.Args[1])
	if err != nil {
		errlog.Println("Error: ", err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		errlog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
