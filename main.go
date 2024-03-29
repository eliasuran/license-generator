package main

import (
	"fmt"
	"os"

	"github.com/eliasuran/license-generator/lic"
	"github.com/manifoldco/promptui"
)

func main() {
	licenses := lic.GetLicenses()

	key := selector(licenses)

	license := lic.GetLicenseByKey(licenses, key)

	licenseInfo := lic.GetLicenseInfo(license)

	lic.MakeLicense(licenseInfo)
}

func selector(parsedLicenses []lic.License) string {
	var licenses []string
	for _, license := range parsedLicenses {
		licenses = append(licenses, license.Name)
	}

	prompt := promptui.Select{
		Label: "Choose your license:",
		Items: licenses,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return result
}
