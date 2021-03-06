package compose

import (
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/docker/libcompose/project/options"
)

func (project *ComposeProject) Up(NoRecreate, ForceRecreate, NoBuild bool) {
	optionsUp := options.Up{
		Create: options.Create{
			NoRecreate:    NoRecreate,
			ForceRecreate: ForceRecreate,
			NoBuild:       NoBuild,
		},
	}

	if err := project.APIProject.Up(context.Background(), optionsUp); err != nil {
		log.WithError(err).Fatal("Could not up the project.")
	}
}
