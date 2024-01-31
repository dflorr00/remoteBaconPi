package tools

import (
	models "backend/Models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/grandcat/zeroconf"
)

func scanner(results <-chan *zeroconf.ServiceEntry) {

	for entry := range results {

		// Crear un dispositivo con los datos escaneados
		var device models.Device
		mystring := fmt.Sprint("", entry.AddrIPv4)
		device.DeviceName = entry.Instance
		device.Ip = mystring
		device.OwnerID = 1
		device.Port = entry.Port
		device.Service = entry.Service
		code, err := device.AddDevice()
		if err != nil {
			log.Println("Fallo al insertar el dispositivo (Función -- scanner()) codigo: ", code)
		}

		data, err := json.Marshal(device)
		if err != nil {
			log.Println("Fallo al crear el json (Función -- scanner())")
		}
		log.Print(string(data) + "\n\n")

	}
}

func Discovery() {
	log.Print("Iniciando el discovery:\n\n")
	resolver, _ := zeroconf.NewResolver(nil)

	entries := make(chan *zeroconf.ServiceEntry)
	entries2 := make(chan *zeroconf.ServiceEntry)
	go scanner(entries)
	go scanner(entries2)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(15))
	resolver.Browse(ctx, "_http._tcp", "local", entries)
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second*time.Duration(15))
	resolver.Browse(ctx2, "_printer._tcp", "local", entries2)

	defer cancel()
	<-ctx.Done()
	defer cancel2()
	<-ctx2.Done()
	log.Println("No more entries.")
}
