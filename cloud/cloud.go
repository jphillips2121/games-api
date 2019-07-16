//Package cloud handles integration with cloud services, e.g. AWS.
package cloud

import "github.com/jphillips2121/games-api/models"

// CLOUD interface declares how to interact with cloud service regardless of which service is used.
type CLOUD interface {
	IsValidDeveloper(game *models.Game) (bool, error)
}
