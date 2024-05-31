package main

import (
	"fmt"
	"github.com/muka/go-bluetooth/api"
	"log"
)

func main() {
	// Initialize the Bluetooth adapter
	manager, err := api.NewManager()
	if err != nil {
		log.Fatalf("Failed to initialize Bluetooth manager: %s", err)
	}
	defer manager.Close()

	// Start the discovery of nearby Bluetooth devices
	discovery, cancel, err := manager.Discover()
	if err != nil {
		log.Fatalf("Failed to start discovery: %s", err)
	}
	defer cancel()

	fmt.Println("Scanning for nearby Bluetooth devices...")

	// Listen for discovered devices
	go func() {
		for {
			select {
			case ev := <-discovery:
				if ev.Type == api.DeviceRemoved {
					// Device removed
					fmt.Printf("Device removed: %s\n", ev.Path)
				} else if ev.Type == api.DeviceAdded {
					// Device added
					device, err := api.GetDeviceByPath(ev.Path)
					if err != nil {
						log.Printf("Failed to get device by path %s: %s", ev.Path, err)
						continue
					}
					props, err := device.GetProperties()
					if err != nil {
						log.Printf("Failed to get properties for device %s: %s", device.Path, err)
						continue
					}
					name, _ := props.GetString("Name")
					address, _ := props.GetString("Address")
					fmt.Printf("Device added: %s (%s)\n", name, address)
				}
			}
		}
	}()

	// Wait indefinitely
	select {}
}
