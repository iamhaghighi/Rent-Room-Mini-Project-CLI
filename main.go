// TODO: ADD TUI - REFACTOR ALL CODE

package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
)

type RoomStatus string
type RoomType string

const (
	RoomAvailable RoomStatus = "available"
	RoomBooked    RoomStatus = "booked"
)

const (
	RoomSingle   RoomType = "single"
	RoomDouble   RoomType = "double"
	RoomStandard RoomType = "standard"
	RoomSuit     RoomType = "suit"
)

type Room struct {
	Id       int
	Type     RoomType
	BedCount int
	Price    int
	Status   RoomStatus
}

var Rooms []Room = generateRoom()

func mainUserInput() string {
	var userInput = ""
	fmt.Scan(&userInput)
	return userInput
}

func conditionMain(userInput string) {
	switch userInput {
	case "1":
		GetRooms()
	case "2":
		RentRoom()
	case "3":
		AddRoom()
	case "0":
		fmt.Print("Exiting... I hope to see you again. ^^")
		return
	default:
		print("wrong command\n")
	}
}

func runMain() {
	var userInput string
	for userInput != "0" {
		print(menu)
		userInput = mainUserInput()
		conditionMain(userInput)
	}
}

func main1() {
	runMain()
}

func handlerAllRoomsInputUser() string {
	print(allRoomsMenu)
	var input = ""
	fmt.Scan(&input)
	return input
}

func handlerAllRoomsCommandCondition(allRoomsUserInput string) {
	switch allRoomsUserInput {
	case "1":
		PrintRooms(Rooms)
	case "2":
		AvailableRooms()
	case "3":
		BookedRooms()
	default:
		print("wrong command!\n")
	}
}

func GetRooms() {
	for {
		command := handlerAllRoomsInputUser()
		if command != "0" {
			handlerAllRoomsCommandCondition(command)
		} else {
			return
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

func DisplayRooms(rooms []Room) {
	if isRoomListEmpty(rooms) {
		fmt.Println("<-----NO ROOM BOOKED----->")
	} else {
		PrintRooms(rooms)
	}
}

func AvailableRooms() {
	rooms := FilterRooms(func(item Room) bool { return item.Status == RoomAvailable })
	DisplayRooms(rooms)
}

func BookedRooms() {
	rooms := FilterRooms(func(item Room) bool { return item.Status == RoomBooked })
	DisplayRooms(rooms)
}

func scanInt() (int, error) {
	var val int
	_, err := fmt.Scan(&val)
	return val, err
}

// * add room by admin
func AddRoomUserInput() (roomType RoomType, bedCount int, price int) {
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
	*rooms = append(*rooms, Room{Id: len(*rooms) + 1, Type: roomType, BedCount: bedCount, Price: price, Status: RoomAvailable})
}

func getInfoFromUser() (id int, nights int, personCount int) {
	fmt.Println("Enter room info line by line:")
	fmt.Print("1) Room ID: ")
	fmt.Scan(&id)
	fmt.Println("[SAVED] ✅")
	fmt.Print("2) Nights: ")
	fmt.Scan(&nights)
	fmt.Println("[SAVED] ✅")
	fmt.Print("3) Person count: ")
	fmt.Scan(&personCount)
	fmt.Println("Done ✅")

	return
}

func GetRoomById(id int) *Room {
	for i := range Rooms {
		if Rooms[i].Id == id {
			return &Rooms[i]
		}
	}
	return nil
}

func RentRoom() {
	id, nights, countPerson := getInfoFromUser()

	room := GetRoomById(id)

	if room == nil {
		fmt.Println("room not found")
		return
	}

	if room.Status == RoomBooked {
		fmt.Println("room isn't available")
		return
	}

	price, tax, discount, finalPrice := calculateRoomPrice(*room, nights, countPerson)

	fmt.Printf(
		"Price: %s | Tax: %s | Discount: %s | Final price: %s\n",
		humanize.Comma(int64(price)),
		humanize.Comma(int64(tax)),
		humanize.Comma(int64(discount)),
		humanize.Comma(int64(finalPrice)),
	)

	room.Status = RoomBooked
	fmt.Println("room reserved successfully ✅")
}

func calculateRoomPrice(room Room, nights int, countPerson int) (price int, tax float64, discount float64, finalPrice int) {

	switch room.Type {
	case "standard":
		price = (nights * room.Price) * countPerson
	case "double":
		price = (nights * room.Price) * countPerson
	case "suit":
		price = (nights * room.Price) * countPerson
	case "single":
		price = (nights * room.Price) * countPerson
	}

	discountPercentage := 0.0
	if nights >= 7 && nights <= 15 {
		discountPercentage = 0.1
	} else if nights >= 15 && nights <= 30 {
		discountPercentage = 0.15
	} else if nights > 30 {
		discountPercentage = 0.2
	}

	tax = float64(price) * 0.09
	discount = float64(price) * discountPercentage
	finalPrice = (price + int(tax)) - int(discount)

	return
}

func generateRoom() []Room {
	rooms := []Room{}

	rooms = append(rooms, Room{Id: 1, Type: "single", BedCount: 1, Price: 250, Status: RoomAvailable})
	rooms = append(rooms, Room{Id: 2, Type: "single", BedCount: 2, Price: 350, Status: RoomAvailable})
	rooms = append(rooms, Room{Id: 3, Type: "double", BedCount: 4, Price: 550, Status: RoomAvailable})
	rooms = append(rooms, Room{Id: 4, Type: "double", BedCount: 3, Price: 380, Status: RoomAvailable})
	rooms = append(rooms, Room{Id: 6, Type: "suit", BedCount: 2, Price: 780, Status: RoomAvailable})
	rooms = append(rooms, Room{Id: 7, Type: "suit", BedCount: 1, Price: 680, Status: RoomAvailable})
	rooms = append(rooms, Room{Id: 5, Type: "standard", BedCount: 4, Price: 480, Status: RoomAvailable})
	rooms = append(rooms, Room{Id: 8, Type: "standard", BedCount: 3, Price: 490, Status: RoomAvailable})
	rooms = append(rooms, Room{Id: 9, Type: "single", BedCount: 1, Price: 200, Status: RoomAvailable})
	rooms = append(rooms, Room{Id: 10, Type: "double", BedCount: 2, Price: 320, Status: RoomAvailable})
	rooms = append(rooms, Room{Id: 11, Type: "suit", BedCount: 3, Price: 850, Status: RoomAvailable})
	rooms = append(rooms, Room{Id: 12, Type: "standard", BedCount: 2, Price: 300, Status: RoomAvailable})
	rooms = append(rooms, Room{Id: 13, Type: "single", BedCount: 1, Price: 180, Status: RoomAvailable})
	rooms = append(rooms, Room{Id: 14, Type: "double", BedCount: 3, Price: 420, Status: RoomAvailable})

	return rooms
}

const menu string = `━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
       🏨  Hotel Rental System  
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
 [1]  All Rooms
 [2]  Rent a Room
 [3]  Add New Room (Admin)
 [0]  Exit
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
 Your choice: `

const allRoomsMenu string = `=== Rooms Menu ===
1) View All Rooms
2) Available Rooms
3) Booked Rooms
0) Back
➜ `

const tableHeader string = "id       type        bed   price   status"
const tableHeaderLine string = "--  ---------------  ---   -----   ------"
const tableUnderLine string = "-----------------------------------------\n"
