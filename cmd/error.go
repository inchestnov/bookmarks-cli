package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func onError(_ *cobra.Command, err error) {
	log.Fatal(err)
}
