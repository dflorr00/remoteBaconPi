package tools

import (
	models "backend/Models"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/clbanning/mxj"
)

func ReloadStatus30s() {
	var wg sync.WaitGroup
	for {
		wg.Add(1)
		go func() {
			refreshStates()
			wg.Done()
		}()
		wg.Wait()
		time.Sleep(30 * time.Second)
	}
}

func refreshStates() {
	var device models.Device
	devicesList := device.GetDevices(1)

	for _, elem := range devicesList {
		elem.Status = realiceServiceability(elem.Ip)
		elem.SetStatus()
	}
}

func realiceServiceability(printerIP string) int {

	// URL del archivo XML de Serviceability en el EWS de la impresora
	ip := printerIP
	ip = strings.Trim(ip, "[]")
	log.Printf("Realizando petición: " + "http://" + ip + "/hp/device/webAccess/open_status.xml")
	rand.Seed(time.Now().UnixNano())
	estados := []string{"correct", "warning", "failure"}
	estado := estados[rand.Intn(len(estados))]
	url := "http://127.0.0.1:8081/" + estado

	// Enviar una solicitud HTTP GET al EWS de la impresora
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error al realizar la petición: " + url)
		return -1
	}
	defer resp.Body.Close()

	// Leer el contenido del archivo XML de Serviceability
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error al leer la respuesta: " + ip)
		return -1
	}

	// Convertir los datos XML en un objeto JSON
	jsonData, err := mxj.NewMapXml(body)
	if err != nil {
		log.Printf("Error al parsear la respuesta: " + ip)
		return -1
	}

	//Obtener el valor del status general
	statusMap, ok := jsonData["Serviceability"].(map[string]interface{})
	if !ok {
		log.Println("No se encontró la cabecera 'Serviceability'")
		return -1
	}
	printerStatusMap, ok := statusMap["Printer_Status"].(map[string]interface{})
	if !ok {
		log.Println("No se encontró la cabecera 'Printer_Status'")
		return -1
	}
	printEngineMap, ok := printerStatusMap["PrintEngine"].(map[string]interface{})
	if !ok {
		log.Println("No se encontró la cabecera 'PrintEngine'")
		return -1
	}
	status, ok := printEngineMap["Status"].(string)
	if !ok {
		log.Println("No se encontró la cabecera 'Status'")
		return -1
	}
	log.Println(status)
	if status == "Idle" {
		log.Println("Esta impresora funciona correctamente")
		return 0
	} else if status == "CartridgeExpired" {
		log.Println("Esta impresora esta en aviso")
		return 1
	} else {
		log.Println("Esta impresora ha fallado")
		return 2
	}
}
