package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Users struct which contains
type Users struct {
	Users []User `json:"users"`
}

// User struct
type User struct {
	Website string `json:"website "`
	Outcome string `json:"outcome "`
	IP      string `json:"ip"`
	Flavor  string `json:"flavor"`
	Open    string `json:"open "`
	OS      string `json:"os"`
}

func main() {

	filePath := os.Args[1]
	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened ex.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and

	/*for i := 0; i < len(users.Users); i++ {
		fmt.Println("website name: " + users.Users[i].Website)
		fmt.Println("outcome url: " + users.Users[i].Outcome)
		fmt.Println("User Name: " + users.Users[i].IP)
		fmt.Println("Flavor: " + users.Users[i].Flavor)
	}*/
	for i := 0; i < len(users.Users); i++ {
		sampledata := []string{"website name:" + users.Users[i].Website,
			"outcome url :" + users.Users[i].Outcome,
			"user ip:" + users.Users[i].IP,
			"flavor:" + users.Users[i].Flavor,
		}

		file, err := os.OpenFile("myfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}

		datawriter := bufio.NewWriter(file)

		for _, data := range sampledata {
			_, _ = datawriter.WriteString(data + "\n")
		}

		datawriter.Flush()
		file.Close()
	}
	/*erro := ioutil.WriteFile("myfile.txt", byteValue, 0777)
	// handle this error
	if erro != nil {
		// print it out
		fmt.Println(erro)
	}*/

}
