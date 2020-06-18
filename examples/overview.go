package main

import (
	"flag"
	"fmt"
	"log"

	xclient "github.com/thebinary/go-xen-client"
)

func main() {
	var hostname, username, password string
	flag.StringVar(&hostname, "hostname", "localhost", "XenServer hostname")
	flag.StringVar(&username, "username", "root", "XenServer authentication username")
	flag.StringVar(&password, "password", "root", "XenServer authentication password")
	flag.Parse()

	client := xclient.NewXenClient(hostname, username, password)

	err := client.Login()
	if err != nil {
		log.Fatalf("Login failure: %v", err)
	}

	hosts, err := client.HostGetAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total Hosts: ", len(hosts))

	vms, err := client.VMGetAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total VMs: ", len(vms))
	fmt.Println()

	fmt.Println("Resident VMs per Host")
	fmt.Println("-----------------------")
	fmt.Println("")
	for _, host := range hosts {
		nameLabel, _ := client.HostGetNameLabel(host)
		fmt.Printf("======== %v =========\n", nameLabel)

		residentVMs, _ := client.HostGetResidentVMs(host)
		for j, vm := range residentVMs {
			nameLabel, _ := client.VMGetNameLabel(vm)
			fmt.Printf("%d. %v\n", j+1, nameLabel)
		}
		fmt.Println()
	}
	fmt.Println()
}
