package main

import (
	"fmt"
	"os"
)

type Room interface {
	Bathroom()
	Kitchen()
	Bedroom()
	Hall()
}

type CustomSize struct {
	size string
}
func (cu *CustomSize) Bathroom() {
	fmt.Printf("Bathroom with your custom design and with size %s", cu.size)
}
func (cu *CustomSize) Kitchen() {
	fmt.Printf("Kitchen with your custom design and with size %s", cu.size)
}
func (cu *CustomSize) Bedroom() {
	fmt.Printf("Bedroom with your custom design and with size %s", cu.size)
}
func (cu *CustomSize) Hall() {
	fmt.Printf("Hall with your custom design and with size %s", cu.size)
}



type CatalogSize struct {}
func (ca *CatalogSize) Bathroom(){
	fmt.Println("Bathroom with catalog design and with size Small")
}
func (ca *CatalogSize) Kitchen(){
	fmt.Println("Kitchen with catalog design and with size Medium")
}
func (ca *CatalogSize) Bedroom(){
	fmt.Println("Bedroom with catalog design and with size Large")
}
func (ca *CatalogSize) Hall(){
	fmt.Println("Hall with catalog design and with size Large")
}

//Bathroom______________________________________________________________________________________________________________
type Bathroom struct {
	Room Room
	helper *ActionHelper
	roomDevices []Device
	DesignInfo string
}

func (b *Bathroom) getDesignInfo() string{
	return b.DesignInfo
}

func (b *Bathroom) Assemble() {
	b.Room.Bathroom()
}

func (b *Bathroom) makeAction(n int){
	if n==1{
		b.helper.setAction(&addDevice{})
		var device_name string
		var device_id int
		fmt.Println("\nEnter device name: ")
		fmt.Fscan(os.Stdin, & device_name)
		fmt.Println("Enter device ID: ")
		fmt.Fscan(os.Stdin, & device_id)
		device := &NewDevice{name: device_name, ID: device_id}
		b.roomDevices = b.helper.executeAction(b.roomDevices, device)
	}
	if n==2{
		b.helper.setAction(&deleteDevice{})
		var device_id int
		fmt.Println("Enter device ID to delete: ")
		fmt.Fscan(os.Stdin, & device_id)
		device := &NewDevice{name: "", ID: device_id}
		b.roomDevices = b.helper.executeAction(b.roomDevices, device)
	}
	if n==3{
		b.helper.setAction(&showAllDevices{})
		device := &NewDevice{name: "", ID: 0}
		b.roomDevices = b.helper.executeAction(b.roomDevices, device)
	}
}

func NewBathroom(room Room) *Bathroom {
	helper := &ActionHelper{}
	return &Bathroom{room, helper ,[]Device{}, "Bathroom "}
}

//Kitchen______________________________________________________________________________________________________________
type Kitchen struct {
	Room Room
	helper *ActionHelper
	roomDevices []Device
	DesignInfo string
}

func (k *Kitchen) getDesignInfo() string{
	return k.DesignInfo
}

func (k *Kitchen) Assemble() {
	k.Room.Kitchen()
}

func (b *Kitchen) makeAction(n int){
	if n==1{
		b.helper.setAction(&addDevice{})
		var device_name string
		var device_id int
		fmt.Println("\nEnter device name: ")
		fmt.Fscan(os.Stdin, & device_name)
		fmt.Println("Enter device ID: ")
		fmt.Fscan(os.Stdin, & device_id)
		device := &NewDevice{name: device_name, ID: device_id}
		b.roomDevices = b.helper.executeAction(b.roomDevices, device)
	}
	if n==2{
		b.helper.setAction(&deleteDevice{})
		var device_id int
		fmt.Println("Enter device ID to delete: ")
		fmt.Fscan(os.Stdin, & device_id)
		device := &NewDevice{name: "", ID: device_id}
		b.roomDevices = b.helper.executeAction(b.roomDevices, device)
	}
	if n==3{
		b.helper.setAction(&showAllDevices{})
		device := &NewDevice{name: "", ID: 0}
		b.roomDevices = b.helper.executeAction(b.roomDevices, device)
	}
}

func NewKitchen(room Room) *Kitchen {
	helper := &ActionHelper{}
	return &Kitchen{room, helper ,[]Device{},"Kitchen "}
}

//Bedroom______________________________________________________________________________________________________________
type Bedroom struct {
	Room Room
	helper *ActionHelper
	roomDevices []Device
	DesignInfo string
}

func (be *Bedroom) getDesignInfo() string{
	return be.DesignInfo
}

func (be *Bedroom) Assemble() {
	be.Room.Bedroom()
}

func (be *Bedroom) makeAction(n int){
	if n==1{
		be.helper.setAction(&addDevice{})
		var device_name string
		var device_id int
		fmt.Println("\nEnter device name: ")
		fmt.Fscan(os.Stdin, & device_name)
		fmt.Println("Enter device ID: ")
		fmt.Fscan(os.Stdin, & device_id)
		device := &NewDevice{name: device_name, ID: device_id}
		be.roomDevices = be.helper.executeAction(be.roomDevices, device)
	}
	if n==2{
		be.helper.setAction(&deleteDevice{})
		var device_id int
		fmt.Println("Enter device ID to delete: ")
		fmt.Fscan(os.Stdin, & device_id)
		device := &NewDevice{name: "", ID: device_id}
		be.roomDevices = be.helper.executeAction(be.roomDevices, device)
	}
	if n==3{
		be.helper.setAction(&showAllDevices{})
		device := &NewDevice{name: "", ID: 0}
		be.roomDevices = be.helper.executeAction(be.roomDevices, device)
	}
}

func NewBedroom(room Room) *Bedroom {
	helper := &ActionHelper{}
	return &Bedroom{room, helper ,[]Device{}, "Bedroom "}
}

//Hall______________________________________________________________________________________________________________
type Hall struct {
	Room Room
	helper *ActionHelper
	roomDevices []Device
	DesignInfo string
}

func (h *Hall) getDesignInfo() string{
	return h.DesignInfo
}

func (h *Hall) Assemble() {
	h.Room.Hall()
}

func (b *Hall) makeAction(n int){
	if n==1{
		b.helper.setAction(&addDevice{})
		var device_name string
		var device_id int
		fmt.Println("\nEnter device name: ")
		fmt.Fscan(os.Stdin, & device_name)
		fmt.Println("Enter device ID: ")
		fmt.Fscan(os.Stdin, & device_id)
		device := &NewDevice{name: device_name, ID: device_id}
		b.roomDevices = b.helper.executeAction(b.roomDevices, device)
	}
	if n==2{
		b.helper.setAction(&deleteDevice{})
		var device_id int
		fmt.Println("Enter device ID to delete: ")
		fmt.Fscan(os.Stdin, & device_id)
		device := &NewDevice{name: "", ID: device_id}
		b.roomDevices = b.helper.executeAction(b.roomDevices, device)
	}
	if n==3{
		b.helper.setAction(&showAllDevices{})
		device := &NewDevice{name: "", ID: 0}
		b.roomDevices = b.helper.executeAction(b.roomDevices, device)
	}
}

func NewHall(room Room) *Hall {
	helper := &ActionHelper{}
	return &Hall{room, helper ,[]Device{}, "Hall "}
}
