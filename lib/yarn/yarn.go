package yarn

import (
	"log"
)

func StartCloud(name, version string) {
	log.Println("Started cloud:", name, version)
}

func StopCloud(name string) {
	log.Println("Stopped cloud:", name)
}
