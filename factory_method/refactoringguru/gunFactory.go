package main

import "errors"

func getGun(guntype string) (iGun, error) {
	if guntype == "ak47" {
		return newAk47(), nil
	}
	if guntype == "musket" {
		return newMusket(), nil
	}
	return nil, errors.New("Wrong gun type passed")
}
