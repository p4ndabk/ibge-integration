package cmd

import (
	"fmt"
	"math/rand"
	"os"
)

type Command struct {
	Inspire      string
	ListCommands string
	RunMigration string
	ImportCities string
	ImportSolarEfficiency string
}

var Commands = Command{
	Inspire:      "inspire",
	ListCommands: "list:commands",
	RunMigration: "run:migration",
	ImportCities: "import:cities",
	ImportSolarEfficiency: "import:solar:efficiency",
}

func InitExec() {
	if len(os.Args) > 1 {
		runCommands()
		os.Exit(0)
	}
}

func runCommands() {
	if os.Args[1] == Commands.Inspire {
		inspire()
	}

	if os.Args[1] == Commands.ListCommands {
		listCommands()
	}

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

	if os.Args[1] == Commands.ImportSolarEfficiency {
		fmt.Println("Importing solar efficiency...")
		_, err := ImportSolarEfficiency()
		if err != nil {
			panic(err)
		}
		fmt.Println("Solar efficiency imported!")
	}
}

func migrations() (bool, error) {
	_, err := CreateCitiesTable()
	if err != nil {
		return false, err
	}

	_, err = CreateSolarEfficiencyTable()
	if err != nil {
		return false, err
	}

	return true, nil
}

func inspire() {
	quotes := []string{
		//motivational
		"Believe in yourself and all that you are.",
		"Success is the sum of small efforts repeated day in and day out.",
		"The journey of a thousand miles begins with one step.",
		"It is never too late to be what you might have been.",
		"The greater the obstacle, the more glory in overcoming it.",
		"Optimism is the faith that leads to achievement. Nothing can be done without hope and confidence.",
		"If you want something you've never had, you must be willing to do something you've never done.",
		"Don't let the fear of failing stop you from trying.",
		"Persistence can change failure into extraordinary achievement.",
		"The secret of getting ahead is getting started.",
		//desmotivational
		"If something can go wrong, it will.",
		"Success is temporary; failure is permanent.",
		"Nothing is worthwhile, so why bother?",
		"No matter how hard you try, you'll always fail.",
		"The world is already full of disappointments, why bother trying harder?",
		"Dreams are just unattainable illusions.",
		"Life is a series of never-ending letdowns.",
		"Never expect much from anything or anyone to avoid disappointment.",
		"The only certainty in life is failure.",
		"Why bother striving when the end result will be the same?",
	}
	randomIndex := rand.Intn(len(quotes))
	fmt.Println(quotes[randomIndex])
}

func listCommands() {
	fmt.Println("Available commands:")
	fmt.Println("Inspire:", Commands.Inspire)
	fmt.Println("ListCommands:", Commands.ListCommands)
	fmt.Println("RunMigration:", Commands.RunMigration)
	fmt.Println("ImportCities:", Commands.ImportCities)
}