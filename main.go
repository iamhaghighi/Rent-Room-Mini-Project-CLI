package main

import (
	"fmt"
)

type Room struct {
	Id       int
	Type     string
	BedCount int
	Price    int
	Status   bool // true means full or reserved and false mean empty or available
}

var Rooms []Room = generateRoom()

func main() {
	userInput := ""

	for userInput != "0" {
		print(menu)
		fmt.Scan(&userInput)

		switch userInput {
		case "1":
			GetRooms()
		case "2":
			RentRoom()
		case "3":
			AddRoom()
		case "0":
			fmt.Print("Exiting... \nI hope to see you again. ^^")
			return
		default:
			print("wrong command\n")
		}
	}
}

// * AllRoom (Input User)
func handlerAllRoomsInputUser() string {
	print(allRoomsMenu)
	var input = ""
	fmt.Scan(&input)
	return input
}

// * AllRoom (Command Condition)
func handlerAllRoomsCommandCondition(allRoomsUserInput string) {
	switch allRoomsUserInput {
	case "1":
		PrintRooms(Rooms)
	case "2":
		EmptyRooms()
	case "3":
		RentedRooms()
	default:
		print("wrong command!\n")
	}
}

// * AllRoom (Get Rooms)
func GetRooms() {
	for {
		command := handlerAllRoomsInputUser()
		if command != "0" {
			handlerAllRoomsCommandCondition(command)
		} else {
			return // return yani inke func ro tamoom mikone va mire jayi ke oon ro seda zadan
		}

	}
}

func PrintRooms(rooms []Room) {
	fmt.Println(tableHeader)
	fmt.Println(tableHeaderLine)

	for _, item := range rooms {
		fmt.Printf("%-7d %-13s %-5d %-6.0d %v\n",
			item.Id, item.Type, item.BedCount, item.Price, item.Status)
	}

	fmt.Print(tableUnderLine)
}

func FilterRooms(filter func(Room) bool) []Room {
	result := []Room{}
	for _, room := range Rooms {
		if filter(room) {
			result = append(result, room)
		}
	}
	return result
}

func isRoomListEmpty(rooms []Room) bool {
	return len(rooms) == 0
}

// ->
func DisplayRooms(rooms []Room) {
	if isRoomListEmpty(rooms) {
		fmt.Println("<-----WE DON'T HAVE ANY ROOM----->")
	} else {
		PrintRooms(rooms)
	}
}

func EmptyRooms() {
	rooms := FilterRooms(func(item Room) bool { return item.Status == false })
	DisplayRooms(rooms)
}

func RentedRooms() {
	rooms := FilterRooms(func(item Room) bool { return item.Status == true })
	DisplayRooms(rooms)
}

// * Add Room User Input
func AddRoomUserInput() (roomType string, bedCount int, price int) {
	print("input room line by line: (type - bed count - price) \n")

	fmt.Scan(&roomType)
	fmt.Println("Saved ✔️")
	fmt.Scan(&bedCount)
	fmt.Println("Saved ✔️")
	fmt.Scan(&price)
	fmt.Println("Done ✔️")

	return
}

func AddRoom() {
	roomType, bedCount, price := AddRoomUserInput()
	rooms := &Rooms
	*rooms = append(*rooms, Room{Id: len(*rooms) + 1, Type: roomType, BedCount: bedCount, Price: price, Status: false})
}

func GetRoomByIdUserInput() (input int) {
	fmt.Println("enter id for reserved room: ")
	fmt.Scan(&input)
	return
}

// *.1 -----------------------------------------
// func GetRoomById(id int) *Room {
// 	for i, _ := range Rooms {
// 		if Rooms[i].Id == id {
// 			return &Rooms[i]
// 		}
// 	}
// 	return nil
// }

// *2 -----------------------------------------
// اینجا فانکشنی داریم که میگه من مقدار آدرس خونه ای رو میگیرم که باید bool باشه
// و بعد پایینتر میگه که خب پس من آدرس خونه ای رو بهت میدم که bool داخلشه
func GetStatusRoom(id int) *bool {
	for i, _ := range Rooms {
		if Rooms[i].Id == id {
			return &Rooms[i].Status
		}
	}
	return nil
}

func RentRoom() {
	// *.1 -----------------------------------------
	// roomItem := GetRoomById(GetRoomByIdUserInput())
	// roomItem.Status = true

	// fmt.Printf("%p\n", &Rooms[GetRoomByIdUserInput() - 1].Status)
	// fmt.Printf("%p", &roomItem.Status)

	// *2 -----------------------------------------
	roomStatus := GetStatusRoom(GetRoomByIdUserInput())

	*roomStatus = true

	fmt.Printf("%p\n", &Rooms[GetRoomByIdUserInput()-1].Status)
	fmt.Printf("%p", roomStatus)

}

func generateRoom() []Room {
	rooms := []Room{}

	rooms = append(rooms, Room{Id: 1, Type: "Single", BedCount: 1, Price: 250, Status: false})
	rooms = append(rooms, Room{Id: 2, Type: "Single", BedCount: 2, Price: 350, Status: false})
	rooms = append(rooms, Room{Id: 3, Type: "double", BedCount: 4, Price: 550, Status: false})
	rooms = append(rooms, Room{Id: 4, Type: "double", BedCount: 3, Price: 380, Status: false})
	rooms = append(rooms, Room{Id: 6, Type: "suit", BedCount: 2, Price: 780, Status: false})
	rooms = append(rooms, Room{Id: 7, Type: "suit", BedCount: 1, Price: 680, Status: false})
	rooms = append(rooms, Room{Id: 5, Type: "standard", BedCount: 4, Price: 480, Status: false})
	rooms = append(rooms, Room{Id: 8, Type: "standard", BedCount: 3, Price: 490, Status: false})
	rooms = append(rooms, Room{Id: 9, Type: "Single", BedCount: 1, Price: 200, Status: false})
	rooms = append(rooms, Room{Id: 10, Type: "double", BedCount: 2, Price: 320, Status: true})
	rooms = append(rooms, Room{Id: 11, Type: "suit", BedCount: 3, Price: 850, Status: false})
	rooms = append(rooms, Room{Id: 12, Type: "standard", BedCount: 2, Price: 300, Status: false})
	rooms = append(rooms, Room{Id: 13, Type: "Single", BedCount: 1, Price: 180, Status: false})
	rooms = append(rooms, Room{Id: 14, Type: "double", BedCount: 3, Price: 420, Status: false})

	return rooms
}

const menu string = `[--Enter a Command--]
[1] Rooms
[2] Rent Room
[3] Add Room [Admin Access]
[0] Exit
Command: `

const allRoomsMenu string = `1: Rooms List
  2: Empty Rooms
    3: Rented Rooms
      0: Back
Command: `

const tableHeader string = "id       type        bed   price   status"
const tableHeaderLine string = "--  ---------------  ---   -----   ------"
const tableUnderLine string = "-----------------------------------------\n"
