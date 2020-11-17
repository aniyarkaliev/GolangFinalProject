package main

import "fmt"

type Device interface {
	getDeviceID() int
	getDeviceName() string
}

type NewDevice struct {
	ID int
	name string
}

func (device *NewDevice) getDeviceID() int{
	return device.ID
}

func (device *NewDevice) getDeviceName() string{
	return device.name
}

type Action interface {
	execute(devices []Device, device Device) []Device
}

type addDevice struct{}

func (ad *addDevice) execute(devices []Device, device_to_add Device) []Device{
	devices = append(devices, device_to_add)
	fmt.Println("Added device "+device_to_add.getDeviceName())
	return devices
}

type deleteDevice struct{}

func (ad *deleteDevice) execute(devices []Device, device_to_delete Device) []Device{
	devicesListLength := len(devices)
	for i, device := range devices {
		if device.getDeviceID() == device_to_delete.getDeviceID() {
			devices[devicesListLength-1], devices[i] = devices[i], devices[devicesListLength-1]
			return devices[:devicesListLength-1]
		}
	}
	return devices
}

type showAllDevices struct {}

func (ad *showAllDevices) execute(devices []Device, device Device) []Device{
	len := len(devices)

	if len==0{
		fmt.Println("There are no devices")
	}
	if len!=0{
		for i, device := range devices {
			cnt:=i+1
			fmt.Printf("%v)Device with name: %v, and ID: %v", cnt, device.getDeviceName(), device.getDeviceID())
			fmt.Println()
		}
	}
	return devices
}

type ActionHelper struct{
	action Action
}

func (actionHelper *ActionHelper) setAction(action Action){
	actionHelper.action = action
}

func (actionHelper *ActionHelper) executeAction(devices []Device, device Device) []Device{
	return actionHelper.action.execute(devices, device)
}