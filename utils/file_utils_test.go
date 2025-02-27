package utils

import (
	"fmt"
	"testing"
	test "the-wedding-game-api/_tests"
)

func TestGenerateRandomFileName1(t *testing.T) {
	generatedFileName, err := generateRandomFileName("test.txt")
	if err != nil {
		t.Errorf("Error while generating random file name")
		return
	}

	fmt.Println(generatedFileName)

	if !test.IsUUID(generatedFileName[:36]) {
		t.Errorf("Incorrect UUID")
		return
	}

	if test.GetFileExtension(generatedFileName) != ".txt" {
		t.Errorf("Incorrect file extension")
		return
	}
}

func TestGenerateRandomFileName2(t *testing.T) {
	generatedFileName, err := generateRandomFileName("random.jpg")
	if err != nil {
		t.Errorf("Error while generating random file name")
		return
	}

	fmt.Println(generatedFileName)

	if !test.IsUUID(generatedFileName[:36]) {
		t.Errorf("Incorrect UUID")
		return
	}

	if test.GetFileExtension(generatedFileName) != ".jpg" {
		t.Errorf("Incorrect file extension")
		return
	}
}

func TestGenerateRandomFileName3(t *testing.T) {
	generatedFileName, err := generateRandomFileName("abcd.efgh")
	if err != nil {
		t.Errorf("Error while generating random file name")
		return
	}

	fmt.Println(generatedFileName)

	if !test.IsUUID(generatedFileName[:36]) {
		t.Errorf("Incorrect UUID")
		return
	}

	if test.GetFileExtension(generatedFileName) != ".efgh" {
		t.Errorf("Incorrect file extension")
		return
	}
}

func TestGenerateRandomFileName4(t *testing.T) {
	generatedFileName, err := generateRandomFileName("abcd.ef")
	if err != nil {
		t.Errorf("Error while generating random file name")
		return
	}

	fmt.Println(generatedFileName)

	if !test.IsUUID(generatedFileName[:36]) {
		t.Errorf("Incorrect UUID")
		return
	}

	if test.GetFileExtension(generatedFileName) != ".ef" {
		t.Errorf("Incorrect file extension")
		return
	}
}

func TestGenerateRandomFileName5(t *testing.T) {
	generatedFileName, err := generateRandomFileName("abcd.def.ghi")
	if err != nil {
		t.Errorf("Error while generating random file name")
		return
	}

	fmt.Println(generatedFileName)

	if !test.IsUUID(generatedFileName[:36]) {
		t.Errorf("Incorrect UUID")
		return
	}

	if test.GetFileExtension(generatedFileName) != ".ghi" {
		t.Errorf("Incorrect file extension")
		return
	}
}

func TestGenerateRandomFileNameWithNoExtension1(t *testing.T) {
	generatedFileName, err := generateRandomFileName("abcd")
	if err != nil {
		t.Errorf("Error while generating random file name")
		return
	}

	fmt.Println(generatedFileName)

	if !test.IsUUID(generatedFileName[:36]) {
		t.Errorf("Incorrect UUID")
		return
	}

	if test.GetFileExtension(generatedFileName) != "" {
		t.Errorf("Incorrect file extension")
		return
	}
}

func TestGenerateRandomFileNameWithNoExtension2(t *testing.T) {
	generatedFileName, err := generateRandomFileName("abcd.")
	if err != nil {
		t.Errorf("Error while generating random file name")
		return
	}

	fmt.Println(generatedFileName)

	if !test.IsUUID(generatedFileName[:36]) {
		t.Errorf("Incorrect UUID")
		return
	}

	if test.GetFileExtension(generatedFileName) != "." {
		t.Errorf("Incorrect file extension")
		return
	}
}
