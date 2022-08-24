package main

import (
	"flag"
	"fmt"
	"github.com/polynetwork/distribute-check/store/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var caseNum uint64

func init() {
	flag.Uint64Var(&caseNum, "case", 0, "case number")
	flag.Parse()
}

func main() {
	db, _ := gorm.Open(postgres.Open(fmt.Sprintf("postgresql://postgres:zmh630323@localhost:5432/zion_%d?sslmode=disable", caseNum)), &gorm.Config{})
	err := CleanDB(db)
	if err != nil {
		panic(err)
	}
}

func CleanDB(db *gorm.DB) error {
	err := db.Migrator().DropTable(&models.TrackHeight{})
	if err != nil {
		return fmt.Errorf("failed to auto drop TrackHeight")
	}

	err = db.Migrator().DropTable(&models.Validator{})
	if err != nil {
		return fmt.Errorf("failed to auto drop Validator")
	}

	err = db.Migrator().DropTable(&models.EpochInfo{})
	if err != nil {
		return fmt.Errorf("failed to auto drop EpochInfo")
	}

	err = db.Migrator().DropTable(&models.StakeInfo{})
	if err != nil {
		return fmt.Errorf("failed to auto drop StakeInfo")
	}

	err = db.Migrator().DropTable(&models.DoneTx{})
	if err != nil {
		return fmt.Errorf("failed to auto drop DoneTx")
	}

	err = db.Migrator().DropTable(&models.TotalGas{})
	if err != nil {
		return fmt.Errorf("failed to auto drop TotalGas")
	}

	err = db.Migrator().DropTable(&models.Rewards{})
	if err != nil {
		return fmt.Errorf("failed to auto drop Rewards")
	}

	err = db.Migrator().DropTable(&models.AccumulatedRewards{})
	if err != nil {
		return fmt.Errorf("failed to auto drop AccumulatedRewards")
	}

	return nil
}