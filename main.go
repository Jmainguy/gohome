package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bwesterb/go-zonefile"
)

func updateHome(zf *zonefile.Zonefile, IP string) (*zonefile.Zonefile, bool) {
	found := false
	update := false
	for _, record := range zf.Entries() {
		if bytes.Equal(record.Domain(), []byte("home")) {
			if !bytes.Equal(record.Values()[0], []byte(IP)) {
				record.SetValue(0, []byte(IP))
				update = true
			}
			found = true
		}
	}

	if !found {
		record := fmt.Sprintf("%s IN A %s", "home", IP)
		entry, _ := zonefile.ParseEntry([]byte(record))
		zf.AddEntry(entry)
		update = true
	}
	return zf, update
}

// Increments the serial of a zonefile
func incrementZone(zf *zonefile.Zonefile) *zonefile.Zonefile {

	// Find SOA entry
	ok := false
	for _, e := range zf.Entries() {
		if !bytes.Equal(e.Type(), []byte("SOA")) {
			continue
		}
		vs := e.Values()
		if len(vs) != 7 {
			fmt.Println("Wrong number of parameters to SOA line")
			os.Exit(4)
		}
		serial, err := strconv.Atoi(string(vs[2]))
		if err != nil {
			fmt.Println("Could not parse serial:", err)
			os.Exit(5)
		}
		t := time.Now()
		today := t.Format("20060102")
		if strings.Contains(strconv.Itoa(serial), today) {
			e.SetValue(2, []byte(strconv.Itoa(serial+1)))
		} else {
			e.SetValue(2, []byte(today+"01"))
		}
		ok = true
		break
	}
	if !ok {
		fmt.Println("Could not find SOA entry")
		os.Exit(6)
	}
	return zf
}

func saveFile(zoneFilename string, zf *zonefile.Zonefile) {
	fh, err := os.OpenFile(zoneFilename, os.O_WRONLY, 0)
	if err != nil {
		fmt.Println(zoneFilename, err)
		os.Exit(7)
	}

	_, err = fh.Write(zf.Save())
	if err != nil {
		fmt.Println(zoneFilename, err)
		os.Exit(8)
	}
}

func loadZone(zoneFilename string) *zonefile.Zonefile {
	// Load zonefile
	data, ioerr := ioutil.ReadFile(zoneFilename)
	if ioerr != nil {
		fmt.Println(zoneFilename, ioerr)
		os.Exit(2)
	}

	zf, perr := zonefile.Load(data)
	if perr != nil {
		fmt.Println(zoneFilename, perr.LineNo(), perr)
		os.Exit(3)
	}
	return zf

}

func getIP() string {
	resp, err := http.Get("https://ip.jmainguy.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	answer := strings.Split(string(body), " ")
	IP := answer[3]
	count := strings.Split(IP, ".")
	if len(count) != 4 {
		fmt.Println("Could not determine IP")
		fmt.Println(body)
		os.Exit(1)
	}
	return IP
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<path to zonefile>")
		os.Exit(1)
	}
	IP := getIP()
	zoneFilename := os.Args[1]
	zf := loadZone(zoneFilename)
	zf, update := updateHome(zf, IP)
	if update {
		zf = incrementZone(zf)
		saveFile(zoneFilename, zf)
	}
}
