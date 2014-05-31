package main

import (
	"github.com/BitBalloon/bitballoon-go"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy a folder or a zip to a BitBalloon site",
	Long:  "deploy a folder or a zip to a BitBalloon site and waits for processing to finish. The --draft flag will prevent the new deploy from getting published once uploaded.",
}

var draft bool

func init() {
	deployCmd.Run = deploy
	deployCmd.Flags().BoolVar(&draft, "draft", false, "Deploy as a draft")
	deployCmd.Flags().StringVarP(&SiteId, "site", "s", "", "site domain or id")

}

func deploy(cmd *cobra.Command, args []string) {
	client := newClient()

	path, err := pathFromArgs(args)
	if err != nil {
		log.Fatalln("Bad directory path")
	}

	log.Printf("Deploying site: %v - dir: %v", SiteId, path)

	site, _, err := client.Sites.Get(SiteId)

	if err != nil {
		log.Fatalf("Error during deploy: %v", err)
	}

	deploy, _, err := deployPath(site, path)

	if err != nil {
		log.Fatalf("Deploy failed with error: ", err)
	}

	err = deploy.WaitForReady(0)
	if err != nil {
		log.Fatalf("Error dring site processing: ", err)
	}

	if draft {
		log.Printf("Deploy can be previewed at %v", deploy.DeployUrl)
		log.Printf("Deploy ID %v", deploy.Id)
	} else {
		log.Println("Site deployed to", site.Url)
	}
}

func pathFromArgs(args []string) (string, error) {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "."
	}

	return filepath.Abs(path)
}

func deployPath(site *bitballoon.Site, path string) (*bitballoon.Deploy, *bitballoon.Response, error) {
	if draft {
		return site.Deploys.CreateDraft(path)
	} else {
		return site.Deploys.Create(path)
	}
}
