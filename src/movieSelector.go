package src

import(
	"../structs"
)

func SelectDrmMovies(movies structs.StanRequest) structs.StanResponse {
	var selectedMovies structs.StanResponse
	for _, movie := range movies.Payload {
		if movie.Drm == true {//&& movie.EpisodeCount > 0 {
			var buf structs.ResponseMovie = structs.ResponseMovie{}
			buf.Image = movie.Image.ShowImage
			buf.Slug = movie.Slug
			buf.Title = movie.Title
			selectedMovies.Response = append(selectedMovies.Response, buf)
		}
	}
	return selectedMovies
}