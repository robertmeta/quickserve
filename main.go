package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var (
	dirs     servableDir
	certName string
	keyName  string
	userName string
	password string
	address  string
	port     int
	newCerts bool
	osSep    = string(filepath.Separator)
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2) // Enable multithreaded.
	flag.Var(&dirs, "d", "List of directories to serve (use multiple -d flags)")
	flag.StringVar(&address, "a", "", "The address to serve https on, blank means all local addresses")
	flag.IntVar(&port, "po", 443, "The port to serve https on")
	flag.StringVar(&certName, "c", "cert.pem", "The name of the cert to use or generate")
	flag.StringVar(&keyName, "k", "key.pem", "The name of the key to use or generate")
	flag.BoolVar(&newCerts, "n", false, "Force generation of new certs")
}

func main() {
	completeAddress := parseFlags()

	http.HandleFunc("/", serveIndex)
	doPerDir(func(l, s string) {
		http.Handle("/"+s, http.StripPrefix("/"+s, http.FileServer(http.Dir(l))))
	})
	if strings.HasPrefix(completeAddress, ":") {
		fmt.Println("Starting server at https://(all)" + completeAddress)
	} else {
		fmt.Println("Starting server at https://" + completeAddress)
	}
	log.Fatal(http.ListenAndServeTLS(completeAddress, certName, keyName, nil))
}

type perDir func(string, string)

func doPerDir(pd perDir) {
	for _, v := range dirs {
		path, err := filepath.Abs(v)
		if err != nil {
			log.Fatal(err)
		}
		path = strings.TrimRight(path, osSep)
		parts := strings.Split(path, string(filepath.Separator))
		shortName := parts[len(parts)-1] + "/"
		pd(path, shortName)
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<pre>`)
	doPerDir(func(l, s string) {
		fmt.Fprintf(w, "<a title=\"%s\" href=\"%s\">%s</a>\n", l, s, s)
	})
	fmt.Fprintf(w, `</pre>`)
}

func parseFlags() string {
	flag.Parse()
	if !certsExist() || newCerts {
		generateCerts()
	}
	if len(dirs) == 0 {
		dirs = append(dirs, ".")
	}
	if port == 0 {
		port = 443
	}
	return address + ":" + strconv.Itoa(port)
}
