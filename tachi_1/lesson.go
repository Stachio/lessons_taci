package main

import (
	"fmt"
	"log"
)

type Friend struct {
	name        string
	personality string
}

func NewFriend(name string, personality string) *Friend {
	return &Friend{
		name:        name,
		personality: personality,
	}
}

func (friend *Friend) Print() {
	fmt.Println(friend.name, "is", friend.personality)
}

func RealTaci(friend *Friend) {
	friend.name = "Sike Taci is a lie"
}

func main() {

	log.Println("Hello world!")

	var taciturn string
	fmt.Println(taciturn)
	taciturn = "The cutest but not as cute as wifeband"
	fmt.Println(taciturn)

	taciturn = "taciturn"
	var age int
	age = 45
	var language string
	language = "Russian and English"
	fmt.Println(taciturn, "is", age, "and speaks", language)

	var stachioName string
	stachioName = "stachio"
	var stachioPersonality string
	stachioPersonality = "talker"

	var adrianName string
	adrianName = "adrian"
	var adrianPersonality string
	adrianPersonality = "youthful"

	var riotName string
	riotName = "riot"
	var riotPersonality string
	riotPersonality = "cautious"

	var taciPersonality string
	taciPersonality = "sexy as фюк"

	var luciName = "luci"
	var luciPersonality = "neizla"

	var coboomsName = "cobooms"
	var coboomsPersonality = "lewd master x"

	fmt.Println(stachioName, "is", stachioPersonality)
	fmt.Println(adrianName, "is", adrianPersonality)
	fmt.Println(riotName, "is", riotPersonality)
	fmt.Println(taciturn, "is", taciPersonality)
	fmt.Println(luciName, "is", luciPersonality)
	fmt.Println(coboomsName, "is", coboomsPersonality)
	NewFriend("hebiohime", "party time").Print()
	//NewFriend("taci2.0", "grand daddy of all time").Print()
	//NewFriend("wifeband", "cutiest cutie of the cutes").Print()
	//NewFriend("robin", "skull skull skull").Print()

	friend := Friend{
		name:        "The real Taci",
		personality: "Maybe kinda sorta cuter than wifeband but not really",
	}

	friend.Print()
	RealTaci(&friend)
	friend.Print()

	const MaxUint32 = ^uint32(0)
	const MinUint32 = 0
	const MaxInt32 = int32(MaxUint32 >> 1)
	const MinInt32 = -MaxInt32 - 1

	number := MaxInt32
	fmt.Println(number)

	//000000000000000000000000000000000
	//011111111111111111111111111111111 + 1
	//111111111111111111111111111111111 + 1
	//111111111111111111111111111111110 + 1
	//111111111111111111111111111111101 + 1

	//100000000000000000000000000000001 + 1
	//000000000000000000000000000000000 + 1
	//000000000000000000000000000000001
}
