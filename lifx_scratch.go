package main

import (
	"fmt"
	"github.com/wkirk01/lifx"
)

func main() {
	lifx.InitClient("c4682413f1c4efa36095a51f580fc31d45a2f461b4649126e7061173b4984bc3")
	lights := lifx.GetLights()
	for _, light := range lights {
		fmt.Println(light.Label, "-", light.Group.Name, "-", light.Power)
	}
}
