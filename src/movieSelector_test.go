package src

import(
	"testing"
	"../structs"
)

func TestSelectDrmMovies(t *testing.T) {
	var testMovies structs.StanRequest = structs.StanRequest{Payload: []structs.PayloadItem{
																	{Drm: true,EpisodeCount: 10},
																	{Drm: true,EpisodeCount: 0},
																	{Drm: false,EpisodeCount: 12},
																	{Drm: true,EpisodeCount: 0},
																	{Drm: true,EpisodeCount: 14},
																	{Drm: true,EpisodeCount: 0},
																	{Drm: false,EpisodeCount: 11},
																	{Drm: false,EpisodeCount: 12}}}

	test := SelectDrmMovies(testMovies)
	if len(test.Response) != 2 {
		t.Error("Incorrect selection of movies\nExpected\n2\nHave\n", len(test.Response))
	}
}