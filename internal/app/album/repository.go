package album

import "2021_1_Noskool_team/internal/app/album/models"

type Repository interface {
	GetAlbumByID(albumID int) (*models.Album, error)
}
