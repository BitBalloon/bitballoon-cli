package main

import (
  "github.com/spf13/cobra"
  "log"
)

var publishCmd = &cobra.Command{
  Use:   "publish",
  Short: "publish a deploy",
  Long:  "publish a deploy to make it the active deploy for a site",
}

var deployId string

func init() {
  publishCmd.Run = publish
  publishCmd.Flags().StringVarP(&deployId, "deploy", "d", "", "Deploy ID")
}

func publish(cmd *cobra.Command, args []string) {
  client := newClient()

  if deployId == "" {
    log.Fatalln("No deploy id. Use --deploy=<id-of-deploy-to-publish>")
  }

  deploy, _, err := client.Deploys.Get(deployId)

  if err != nil {
    log.Fatalf("Could not load deploy: %v", err)
  }

  _, err = deploy.Publish()

  if err != nil {
    log.Fatalf("Failed to publish deploy: %v", err)
  }

  log.Printf("Deploy %v published at %v", deploy.SiteUrl)
}
