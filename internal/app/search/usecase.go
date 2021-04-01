package search

import "2021_1_Noskool_team/internal/app/search/models"

type Usecase interface {
	SearchContent(searchQuery string) *models.Search
}
