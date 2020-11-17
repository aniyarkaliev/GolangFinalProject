package main

import (
	"bufio"
	"fmt"
	"os"
)

var allRooms []BasicRoom
var notifier = Notifier{}

func showAllRooms(){
	for i, room := range allRooms{
		cnt := i+1
		fmt.Printf("%v)"+room.getDesignInfo(), cnt)
		fmt.Println()
	}
}

func addRoom(room BasicRoom) []BasicRoom{
	allRooms := append(allRooms, room)
	return allRooms
}

func displayMenu(){
	fmt.Println("-----Control panel-----")
	fmt.Println("1) Press 1 to add new family member")
	fmt.Println("2) Press 2 to add new Room")
	fmt.Println("3) Press 3 to choose available Rooms")
	fmt.Println("4) Press 4 to show all family members")
	fmt.Println("5) Press 5 to show all rooms")
	fmt.Println("6) Press 6 to show messages log")
	fmt.Println("7) Press 7 to send messages to all family members")
	fmt.Println("10) To exit from application")
	fmt.Println("-----------------------")
}

func createRoom(){
	fmt.Println("-----Room creator-----")
	fmt.Println("1) Press 1 to add new Bathroom")
	fmt.Println("2) Press 2 to add new Bedroom")
	fmt.Println("3) Press 3 to add new Kitchen")
	fmt.Println("4) Press 4 to add new Hall")
	fmt.Println("5) Press 5 to send message for all family members")
	fmt.Println("10) To exit from room creator")
	fmt.Println("-----------------------")
	var choice int
	fmt.Fscan(os.Stdin, &choice)
	switch choice{
	case 1:
		//BATHROOM CREATION
		fmt.Println("-----Choose custom or catalog(default) size-----")
		fmt.Println("1) Press 1 to custom")
		fmt.Println("2) Press 2 to catalog")
		fmt.Println("3) To exit from room creator")
		fmt.Println("-----------------------")
		var choice1 int
		fmt.Fscan(os.Stdin, &choice1)
		switch choice1 {
		//for custom size
		case 1:
			fmt.Println("Enter room size(small, medium, large): ")
			var size string
			fmt.Fscan(os.Stdin, &size)
			room := CustomSize{size}
			bathroom := NewBathroom(&room)
			fmt.Println("-----Design of room-----")
			fmt.Println("1) Press 1 to add Aquarium to room")
			fmt.Println("2) Press 2 to add Fireplace to room")
			fmt.Println("3) Press 3 to add Coffetable to room")
			fmt.Println("4) without design")
			fmt.Println("-----------------------")
			var choice2 int
			fmt.Fscan(os.Stdin, &choice2)
			switch choice2 {
			case 1:
				withAqua := withAquarium{bathroom}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Fireplace to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withFire := withFirePlace{&withAqua}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withAqua}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withAqua)
				}
			case 2:
				withFirePlace := withFirePlace{bathroom}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withFirePlace}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withFirePlace}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withFirePlace)
				}
			case 3:
				withCoffeeTable := withCoffeeTable{bathroom}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Fireplace to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withCoffeeTable}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withFire := withFirePlace{&withCoffeeTable}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withCoffeeTable)
				}
			case 4:
				allRooms = addRoom(bathroom)
			}
			fmt.Println("New room added")
		case 2:
			room := CatalogSize{}
			bathroom := NewBathroom(&room)
			fmt.Println("-----Design of room-----")
			fmt.Println("1) Press 1 to add Aquarium to room")
			fmt.Println("2) Press 2 to add Fireplace to room")
			fmt.Println("3) Press 3 to add Coffetable to room")
			fmt.Println("4) without design")
			fmt.Println("-----------------------")
			var choice2 int
			fmt.Fscan(os.Stdin, &choice2)
			switch choice2 {
			case 1:
				withAqua := withAquarium{bathroom}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Fireplace to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withFire := withFirePlace{&withAqua}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withAqua}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withAqua)
				}
			case 2:
				withFirePlace := withFirePlace{bathroom}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withFirePlace}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withFirePlace}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withFirePlace)
				}
			case 3:
				withCoffeeTable := withCoffeeTable{bathroom}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Fireplace to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withCoffeeTable}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withFire := withFirePlace{&withCoffeeTable}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withCoffeeTable)
				}
			case 4:
				allRooms = addRoom(bathroom)
			}
			fmt.Println("New room added")
		}
	case 2:
		//BEDROOM CREATION
		fmt.Println("-----Choose custom or catalog(default) size-----")
		fmt.Println("1) Press 1 to custom")
		fmt.Println("2) Press 2 to catalog")
		fmt.Println("3) To exit from room creator")
		fmt.Println("-----------------------")
		var choice1 int
		fmt.Fscan(os.Stdin, &choice1)
		switch choice1 {
		//for custom size
		case 1:
			fmt.Println("Enter room size(small, medium, large): ")
			var size string
			fmt.Fscan(os.Stdin, &size)
			room := CustomSize{size}
			bedroom := NewBedroom(&room)
			fmt.Println("-----Design of room-----")
			fmt.Println("1) Press 1 to add Aquarium to room")
			fmt.Println("2) Press 2 to add Fireplace to room")
			fmt.Println("3) Press 3 to add Coffetable to room")
			fmt.Println("4) without design")
			fmt.Println("-----------------------")
			var choice2 int
			fmt.Fscan(os.Stdin, &choice2)
			switch choice2 {
			case 1:
				withAqua := withAquarium{bedroom}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Fireplace to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withFire := withFirePlace{&withAqua}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withAqua}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withAqua)
				}
			case 2:
				withFirePlace := withFirePlace{bedroom}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withFirePlace}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withFirePlace}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withFirePlace)
				}
			case 3:
				withCoffeeTable := withCoffeeTable{bedroom}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Fireplace to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withCoffeeTable}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withFire := withFirePlace{&withCoffeeTable}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withCoffeeTable)
				}
			case 4:
				allRooms = addRoom(bedroom)
			}
			fmt.Println("New room added")
		case 2:
			room := CatalogSize{}
			bedroom := NewBedroom(&room)
			fmt.Println("-----Design of room-----")
			fmt.Println("1) Press 1 to add Aquarium to room")
			fmt.Println("2) Press 2 to add Fireplace to room")
			fmt.Println("3) Press 3 to add Coffetable to room")
			fmt.Println("4) without design")
			fmt.Println("-----------------------")
			var choice2 int
			fmt.Fscan(os.Stdin, &choice2)
			switch choice2 {
			case 1:
				withAqua := withAquarium{bedroom}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Fireplace to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withFire := withFirePlace{&withAqua}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withAqua}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withAqua)
				}
			case 2:
				withFirePlace := withFirePlace{bedroom}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withFirePlace}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withFirePlace}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withFirePlace)
				}
			case 3:
				withCoffeeTable := withCoffeeTable{bedroom}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Fireplace to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withCoffeeTable}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withFire := withFirePlace{&withCoffeeTable}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withCoffeeTable)
				}
			case 4:
				allRooms = addRoom(bedroom)
			}
			fmt.Println("New room added")
		}
	case 3:
		//KITCHEN CREATION
		fmt.Println("-----Choose custom or catalog(default) size-----")
		fmt.Println("1) Press 1 to custom")
		fmt.Println("2) Press 2 to catalog")
		fmt.Println("3) To exit from room creator")
		fmt.Println("-----------------------")
		var choice1 int
		fmt.Fscan(os.Stdin, &choice1)
		switch choice1 {
		//for custom size
		case 1:
			fmt.Println("Enter room size(small, medium, large): ")
			var size string
			fmt.Fscan(os.Stdin, &size)
			room := CustomSize{size}
			kitchen := NewKitchen(&room)
			fmt.Println("-----Design of room-----")
			fmt.Println("1) Press 1 to add Aquarium to room")
			fmt.Println("2) Press 2 to add Fireplace to room")
			fmt.Println("3) Press 3 to add Coffetable to room")
			fmt.Println("4) without design")
			fmt.Println("-----------------------")
			var choice2 int
			fmt.Fscan(os.Stdin, &choice2)
			switch choice2 {
			case 1:
				withAqua := withAquarium{kitchen}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Fireplace to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withFire := withFirePlace{&withAqua}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withAqua}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withAqua)
				}
			case 2:
				withFirePlace := withFirePlace{kitchen}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withFirePlace}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withFirePlace}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withFirePlace)
				}
			case 3:
				withCoffeeTable := withCoffeeTable{kitchen}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Fireplace to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withCoffeeTable}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withFire := withFirePlace{&withCoffeeTable}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withCoffeeTable)
				}
			case 4:
				allRooms = addRoom(kitchen)
			}
			fmt.Println("New room added")
		case 2:
			room := CatalogSize{}
			kitchen := NewKitchen(&room)
			fmt.Println("-----Design of room-----")
			fmt.Println("1) Press 1 to add Aquarium to room")
			fmt.Println("2) Press 2 to add Fireplace to room")
			fmt.Println("3) Press 3 to add Coffetable to room")
			fmt.Println("4) without design")
			fmt.Println("-----------------------")
			var choice2 int
			fmt.Fscan(os.Stdin, &choice2)
			switch choice2 {
			case 1:
				withAqua := withAquarium{kitchen}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Fireplace to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withFire := withFirePlace{&withAqua}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withAqua}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withAqua)
				}
			case 2:
				withFirePlace := withFirePlace{kitchen}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withFirePlace}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withFirePlace}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withFirePlace)
				}
			case 3:
				withCoffeeTable := withCoffeeTable{kitchen}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Fireplace to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withCoffeeTable}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withFire := withFirePlace{&withCoffeeTable}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withCoffeeTable)
				}
			case 4:
				allRooms = addRoom(kitchen)
			}
			fmt.Println("New room added")
		}
	case 4:
		//HALL CREATION
		fmt.Println("-----Choose custom or catalog(default) size-----")
		fmt.Println("1) Press 1 to custom")
		fmt.Println("2) Press 2 to catalog")
		fmt.Println("3) To exit from room creator")
		fmt.Println("-----------------------")
		var choice1 int
		fmt.Fscan(os.Stdin, &choice1)
		switch choice1 {
		//for custom size
		case 1:
			fmt.Println("Enter room size(small, medium, large): ")
			var size string
			fmt.Fscan(os.Stdin, &size)
			room := CustomSize{size}
			hall := NewHall(&room)
			fmt.Println("-----Design of room-----")
			fmt.Println("1) Press 1 to add Aquarium to room")
			fmt.Println("2) Press 2 to add Fireplace to room")
			fmt.Println("3) Press 3 to add Coffetable to room")
			fmt.Println("4) without design")
			fmt.Println("-----------------------")
			var choice2 int
			fmt.Fscan(os.Stdin, &choice2)
			switch choice2 {
			case 1:
				withAqua := withAquarium{hall}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Fireplace to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withFire := withFirePlace{&withAqua}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withAqua}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withAqua)
				}
			case 2:
				withFirePlace := withFirePlace{hall}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withFirePlace}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withFirePlace}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withFirePlace)
				}
			case 3:
				withCoffeeTable := withCoffeeTable{hall}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Fireplace to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withCoffeeTable}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withFire := withFirePlace{&withCoffeeTable}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withCoffeeTable)
				}
			case 4:
				allRooms = addRoom(hall)
			}
			fmt.Println("New room added")
		case 2:
			room := CatalogSize{}
			hall := NewHall(&room)
			fmt.Println("-----Design of room-----")
			fmt.Println("1) Press 1 to add Aquarium to room")
			fmt.Println("2) Press 2 to add Fireplace to room")
			fmt.Println("3) Press 3 to add Coffetable to room")
			fmt.Println("4) without design")
			fmt.Println("-----------------------")
			var choice2 int
			fmt.Fscan(os.Stdin, &choice2)
			switch choice2 {
			case 1:
				withAqua := withAquarium{hall}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Fireplace to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withFire := withFirePlace{&withAqua}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withAqua}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withAqua)
				}
			case 2:
				withFirePlace := withFirePlace{hall}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Coffetable to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withFirePlace}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withCoffee := withCoffeeTable{&withFirePlace}
					allRooms = addRoom(&withCoffee)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withFirePlace)
				}
			case 3:
				withCoffeeTable := withCoffeeTable{hall}
				fmt.Println("You can add 1 more")
				fmt.Println("1) Press 1 to add Aquarium to room")
				fmt.Println("2) Press 2 to add Fireplace to room")
				fmt.Println("3) close")
				fmt.Println("-----------------------")
				fmt.Println("Enter room size(small, medium, large): ")
				var choice3 int
				fmt.Fscan(os.Stdin, &choice3)
				if(choice3 == 1){
					withAqua := withAquarium{&withCoffeeTable}
					allRooms = addRoom(&withAqua)
				}
				if(choice3 == 2){
					withFire := withFirePlace{&withCoffeeTable}
					allRooms = addRoom(&withFire)
				}
				if(choice3 == 3){
					allRooms = addRoom(&withCoffeeTable)
				}
			case 4:
				allRooms = addRoom(hall)
			}
			fmt.Println("New room added")
		}
	}
}

func createUser(){
	fmt.Println("Enter firstname:")
	var firstname string
	fmt.Fscan(os.Stdin, &firstname)
	fmt.Println("Enter family lastname:")
	var lastname string
	fmt.Fscan(os.Stdin, &lastname)
	fmt.Println("Enter family status:")
	var status string
	fmt.Fscan(os.Stdin, &status)
	user := NewUserBuilder().PersonalInfo().withFirstName(firstname).withLastname(lastname).
		FamilyStatus().InFamily(status).Build()
	notifier.register(user)
	fmt.Println("New user added!")
}

func sendMessage(){
	fmt.Println("Enter message: ")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	notifier.addNewMessage(line)
}

func showMessagesLog(){
	notifier.showMessageLog()
}

func displayDeviceMenu(){
	fmt.Println("-----Device control panel-----")
	fmt.Println("1) Press 1 to add new Device in room")
	fmt.Println("2) Press 2 to delete Device from room")
	fmt.Println("3) Press 3 to show all Devices of room")
	fmt.Println("4) Close")
}

func chooseRoom(){
	fmt.Println("All available rooms")
	showAllRooms()
	fmt.Println("Choose room by num:")
	var choice int
	fmt.Fscan(os.Stdin, &choice)
	displayDeviceMenu()
	fmt.Println("Choose action by num:")
	var choice1 int
	fmt.Fscan(os.Stdin, &choice1)
	allRooms[choice-1].makeAction(choice1)
}

func main() {
	{
		//user1 := CustomDesign{"Large"}
		//bt1 := NewBathroom(&user1)
		//bt1.Assemble()
		//bt1.makeAction(1)
		//bt1.makeAction(3)
		//bt1.makeAction(2)
		//bt1.makeAction(3)

		//user1 := NewUserBuilder().PersonalInfo().withFirstName("Aza").withLastname("Saiduly").FamilyStatus().
		//	InFamily("Shegol").Build()
		//user2 := NewUserBuilder().PersonalInfo().withFirstName("Aniyar").withLastname("Kaliev").FamilyStatus().
		//	InFamily("Chmo").Build()
		//user3 := NewUserBuilder().PersonalInfo().withFirstName("Ainura").withLastname("Kursabayeva").FamilyStatus().
		//	InFamily("Sister").Build()
		//
		//notifier := &Notifier{}
		//notifier.register(user1)
		//notifier.register(user2)
		//notifier.register(user3)
		//
		//notifier.addNewMessage("HELLO EVERYONE")

		//custom := CustomSize{size: "Large"}
		//bathroom := NewBathroom(&custom)
		//
		//withFirePlace := withFirePlace{bathroom}
		//fmt.Println(withFirePlace.getDesignInfo())
		//withAqua := withAquarium{&withFirePlace}
		//fmt.Println(withAqua.getDesignInfo())

	}
	fmt.Println("Welcome to SMARTHOME 2.0")
	for i:=0; i<=100; i++{
		displayMenu()
		fmt.Println("Enter command: ")
		var choice int
		fmt.Fscan(os.Stdin, &choice)
		switch choice {
			case 1:
				createUser()
			case 2:
				createRoom()
			case 3:
				chooseRoom()
			case 4:
				notifier.showAllMembers()
			case 5:
				showAllRooms()
			case 6:
				showMessagesLog()
			case 7:
				sendMessage()
			case 10:
				os.Exit(0)
		}
	}

}

