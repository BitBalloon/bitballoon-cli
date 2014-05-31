package main

import (
  "github.com/BitBalloon/bitballoon-go"
  "github.com/spf13/cobra"
  "log"
	"os"
	"bufio"
  "strings"
)

var deleteCmd = &cobra.Command{
  Use:   "delete",
  Short: "delete a BitBalloon site permanently",
  Long:  "permanently deletes a site and all related deploys. No undo!",
}

func init() {
  deleteCmd.Run = delete
  deleteCmd.Flags().StringVarP(&SiteId, "site", "s", "", "site domain or id")
}

func delete(cmd *cobra.Command, args []string) {
  client := newClient()

  if SiteId == "" {
    log.Fatalln("No site id specified. Use the --site options")
  }

  site, _, err := client.Sites.Get(SiteId)

  if err != nil {
    log.Fatalf("Error deleting site: %v", err)
  }

  confirmDeletion(site)

  _, err = site.Destroy()

  if err != nil {
    log.Fatalf("Error deleting site: %v", err)
  }

  log.Printf("Site deleted.")
}

func confirmDeletion(site *bitballoon.Site) {
  log.Printf("Permanently Delete Site %v (Y/N)? ", site.Url)
	reader := bufio.NewReader(os.Stdin)
	result,err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

  if strings.ToLower(strings.TrimSpace(result)) != "y" {
    log.Fatalln("Aborted")
  }
}
