package poker

import (
	"fmt"
)

func ErrorParseLeague(err error) error {
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return err
}

func ErrorFileOpening(err error, fileName string) error {
	if err != nil {
		err = fmt.Errorf("problem opening %s %v", fileName, err)
	}
	return err
}

func ErrorFileCreation(err error) error {
	if err != nil {
		fmt.Errorf("problem creating file system player store, %v ", err)
	}
	return err
}

func ErrorListenAndServe(err error) error {
	if err != nil {
		fmt.Errorf("could not listen on port 5000 %v", err)
	}
	return err
}
