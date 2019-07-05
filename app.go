package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type Profile struct {
	Total     float64
	Consumida float64
}
type CPU struct {
	Total float64
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	w.Write([]byte(message))
}

func testw(w http.ResponseWriter, r *http.Request) {
	message := " Hola Mundo test "
	w.Write([]byte(message))
}

func getRAM(w http.ResponseWriter, r *http.Request) {
	nombreArchivo := "/proc/meminfo"
	bytesLeidos, err := ioutil.ReadFile(nombreArchivo)
	if err != nil {
		fmt.Printf("Error leyendo archivo: %v", err)
	}
	contenido := string(bytesLeidos)
	lineas := strings.Split(contenido, "\n")

	totalMem := 0.0
	MemAvailable := 0.0
	MemConsumida := 0.0
	porsentajeMem := 0.0
	for _, linea := range lineas {
		campos := strings.SplitN(linea, ":", 2)
		if len(campos) < 2 {
			continue
		}
		valorCampos := strings.Fields(campos[1])
		val, _ := strconv.ParseFloat(valorCampos[0], 64)

		if campos[0] == "MemTotal" {
			totalMem = val
			totalMem = totalMem / 1024
		}
		if campos[0] == "MemAvailable" {
			MemAvailable = val
			MemAvailable = MemAvailable / 1024
		}
	}
	MemConsumida = totalMem - MemAvailable
	porsentajeMem = (MemConsumida * 100) / totalMem

	u := Profile{totalMem, porsentajeMem}
	js, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getCPU(w http.ResponseWriter, r *http.Request) {
	percent, _ := cpu.Percent(time.Second, false)
	u := CPU{percent[0]}
	js, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func viewRam(w http.ResponseWriter, r *http.Request) {
	bytesLeidos, err := ioutil.ReadFile("temas/ram.html")
	if err != nil {
		fmt.Printf("Error leyendo archivo: %v", err)
	}
	contenido := string(bytesLeidos)
	w.Write([]byte(contenido))
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/test", testw)
	http.HandleFunc("/getRam", getRAM)
	http.HandleFunc("/getCPU", getCPU)
	http.HandleFunc("/ram", viewRam)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
