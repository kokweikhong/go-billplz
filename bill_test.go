package billplz_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kokweikhong/go-billplz"
)

func TestGetBill(t *testing.T) {
	ids := []string{
		"sl3giadb",
		"ixv7u4mo",
	}
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		t.Error("No .env file found")
	}

	bp := billplz.New(
		os.Getenv("BILLPLZ_API_KEY"),
		os.Getenv("BILLPLZ_API_URL"),
		"", "", "",
	)

	for _, id := range ids {
		bill, err := bp.GetBill(id)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(bill)
	}
}

func TestDeleteBill(t *testing.T) {
	ids := []string{
		"gsnrxpxn",
		"bjk2ilew",
	}
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		t.Error("No .env file found")
	}

	bp := billplz.New(
		os.Getenv("BILLPLZ_API_KEY"),
		os.Getenv("BILLPLZ_API_URL"),
		"", "", "",
	)

	for _, id := range ids {
		err := bp.DeleteBill(id)
		if err != nil {
			t.Error(err)
		}
	}
}
