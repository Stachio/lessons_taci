package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// World - VRChat world
type World struct {
	Name        string
	Description string
}

// Friend - VRChat friend
type Friend struct {
	Name           string
	Personality    string
	Avatar         string
	Drinks         bool
	FavoriteWorlds []World
}

func main() {

	log.Println("Hello world!")

	/* var taciturn string
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

	var maybe bool
	maybe = true
	fmt.Println("maybe:", maybe)

	var glue = ","

	var fortune = "not looking good"
	var appearance string
	appearance = "looks fine"

	var outlook string
	if maybe {
		outlook = appearance
	} else {
		outlook = fortune
	}

		if maybe {
			fmt.Println(stachioName, "is", stachioPersonality, glue, appearance)
			fmt.Println(adrianName, "is", adrianPersonality, glue, appearance)
			fmt.Println(riotName, "is", riotPersonality, glue, appearance)
			fmt.Println(taciturn, "is", taciPersonality, glue, appearance)
			fmt.Println(luciName, "is", luciPersonality, glue, appearance)
			fmt.Println(coboomsName, "is", coboomsPersonality, glue, appearance)
		} else {
			fmt.Println(stachioName, "is", stachioPersonality, glue, fortune)
			fmt.Println(adrianName, "is", adrianPersonality, glue, fortune)
			fmt.Println(riotName, "is", riotPersonality, glue, fortune)
			fmt.Println(taciturn, "is", taciPersonality, glue, fortune)
			fmt.Println(luciName, "is", luciPersonality, glue, fortune)
			fmt.Println(coboomsName, "is", coboomsPersonality, glue, fortune)
		}
	//NewFriend("hebiohime", "party time").Print()
	//NewFriend("taci2.0", "grand daddy of all time").Print()
	//NewFriend("wifeband", "cutiest cutie of the cutes").Print()
	//NewFriend("robin", "skull skull skull").Print()

	fmt.Println(stachioName, "is", stachioPersonality, glue, outlook)
	fmt.Println(adrianName, "is", adrianPersonality, glue, outlook)
	fmt.Println(riotName, "is", riotPersonality, glue, outlook)
	fmt.Println(taciturn, "is", taciPersonality, glue, outlook)
	fmt.Println(luciName, "is", luciPersonality, glue, outlook)
	fmt.Println(coboomsName, "is", coboomsPersonality, glue, outlook) */

	nierAutomataWorld := World{
		Name:        "Nier Flowers",
		Description: "Really cute Nier flower world",
	}

	lunaPlantsWorld := World{
		Name:        "Luna Plants",
		Description: "On the moooooooon but cute af",
	}

	stachio := Friend{
		Name:        "Stachio",
		Personality: "doesn't know when to stfu",
		Avatar:      "based onsie",
		Drinks:      true,
		FavoriteWorlds: []World{
			nierAutomataWorld,
			lunaPlantsWorld,
		},
	}

	adrian := Friend{
		Name:        "Adrian",
		Personality: "mad trol",
		Avatar:      "contains boyfriend hunter",
		Drinks:      true,
	}

	riot := Friend{
		Name:        "Riot",
		Personality: "awkward but hilarious",
		Avatar:      "super dubsteppy ball-spinny darkness",
		Drinks:      true,
	}

	tacиturn := Friend{
		Name:        "Tacinator",
		Personality: "cute af",
		Avatar:      "almost invisible pants, makes me think he's naked",
		Drinks:      false,
	}

	luci := Friend{
		Name:        "Luci",
		Personality: "sweet and silly drowning noises",
		Avatar:      "almost as lewd as cobooms",
		Drinks:      false,
	}

	notLuci := Friend{
		Name:        "Not Luci",
		Personality: "something",
		Avatar:      "nun costume",
		Drinks:      true,
	}

	friends := []Friend{
		stachio,
		adrian,
		riot,
		tacиturn,
		luci,
		notLuci,
	}

	fmt.Println(stachio)
	fmt.Println(adrian)

	/* fmt.Println(stachio.name, "is", stachio.personality, "and uses the avatar", stachio.avatar)
	fmt.Println(adrian.name, "is", adrian.personality, "and uses the avatar", adrian.avatar)
	fmt.Println(riot.name, "is", riot.personality, "and uses the avatar", riot.avatar) */

	for _, friend := range friends {
		var action string
		if friend.Drinks {
			action = "drink"
		} else {
			action = "do not drink"
		}
		fmt.Println(friend.Name, "is", friend.Personality, "and uses the avatar", friend.Avatar, "and they", action)
		//data, _ := json.Marshal(&friend)
		jsonData, _ := json.MarshalIndent(&friend, "", "    ")
		xmlData, _ := xml.MarshalIndent(&friend, "", "    ")
		ioutil.WriteFile(friend.Name+".json", jsonData, os.ModePerm)
		ioutil.WriteFile(friend.Name+".xml", xmlData, os.ModePerm)
	}

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

	//array of string
	/* var nameArray = []string{
		stachioName, //element at index 0
		adrianName,  //element at index 1
		riotName,    //element at index 2
		taciturn,    //element at index 3
		coboomsName,
	} */

	//var index = 0
	/* fmt.Println(nameArray[index])
	fmt.Println(nameArray[index+1])
	fmt.Println(nameArray[index+2])
	fmt.Println(nameArray[index+3]) */

	/* for index := range nameArray {
		fmt.Println(nameArray[index])
	} */

	// for (int index = 0; index < nameArray.length; index++)
	// 0 -> nameArray[index] = stachioName
	// 1 -> nameArray[index] = adrianName
	// 2 -> nameArray[index] = riotName
	// 3 -> nameArray[index] = taciturn
	// 4 -> nameArray[index] = coboomsName

	/*
		string[] nameArray = new []string {
			"Stachio",
			"Adrian",
			"Riot",
			"Taciturn",
			"Cabooms"
		}

		foreach (string name in nameArray) {
			console.WriteLine(name);
		}

		for (int index = 0; index < nameArray.length; index++) {
			console.WriteLine(nameArray[index]);
		}
	*/

	/* for _, name := range nameArray {
		fmt.Println(name)
	}

	for index := range nameArray {
		fmt.Println(index, nameArray[index])
	}

	nameArray[0] = "Supreme overlord"

	for index, name := range nameArray {
		fmt.Println(index, name)
	} */

	numberArray := []int{
		7,
		35,
		29,
		10,
		58,
		203,
		1,
		98,
		32137,
		743547,
		43,
		65474,
	}

	highestNumber := 0
	for index, number := range numberArray {
		fmt.Println("Testing", index, number)
		if number > highestNumber {
			fmt.Println("Number", number, "at index", index, "is higher than", highestNumber)
			highestNumber = number
		}
	}

	fmt.Println("Highest number:", highestNumber)

	// Data Types
	// Arrays
	// Loops
	// VRChat API

	var yes = "yes"
	character := yes[0]
	fmt.Println(character)
	if character == 'y' {
		fmt.Println("Found a yes bois")
	}
}
