package cmd

import (
	"fmt"
	"os"
)

type Command struct {
	RunMigration string
	ImportCities string
}

var Commands = Command{
	RunMigration: "run:migration",
	ImportCities: "import:cities",
}

func InitExec() {
	if len(os.Args) > 1 {
		if os.Args[1] == Commands.RunMigration {
			fmt.Println("Running migrations...")
			_, err := migrations()
			if err != nil {
				panic(err)
			}
			fmt.Println("Migrations finished!")
		}

		if os.Args[1] == Commands.ImportCities {
			fmt.Println("Importing cities...")
			_, err := ImportCities()
			if err != nil {
				panic(err)
			}
			fmt.Println("Cities imported!")
		}
	}

	fmt.Println(true)
	os.Exit(0)
}


func migrations() (bool, error) {
	_, err := CreateCitiesTable()
	if err != nil {
		return false, err
	}

	return true, nil
}


