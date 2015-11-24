package cmd

import (
	"github.com/pritunl/pritunl-link/constants"
	"github.com/pritunl/pritunl-link/utils"
	"os"
)

type options struct {
	Id     string
	Host   string
	Token  string
	Secret string
}

func getOptions() (opts *options) {
	id := os.Getenv("ID")
	if id == "" {
		id = utils.RandName()
	}

	constants.Host = id

	opts = &options{
		Id:     id,
		Host:   os.Getenv("HOST"),
		Token:  os.Getenv("TOKEN"),
		Secret: os.Getenv("SECRET"),
	}

	return
}