package main

import (
	"fmt"
	"github.com/wkirk01/lifx"
	"log"
)

func main() {
	client := lifx.NewClient("c4682413f1c4efa36095a51f580fc31d45a2f461b4649126e7061173b4984bc3")
	lights, err := client.GetLights()
	if err != nil {
		log.Fatal(err)
	}
	groups := map[string][]lifx.Light{}
	for _, light := range lights {
		groups[light.Group.Name] = append(groups[light.Group.Name], light)
	}
	for groupName, lights := range groups {
		fmt.Println("\n" + "== " + groupName + " ==")
		for _, light := range lights {
			fmt.Println(light.Label, "-", light.Group.Name, "-", light.Power)
		}
	}
}
