package src

import(
	"testing"
)

func TestParseReqBody(t *testing.T) {
	_, err := ParseReqBody(nil)
	if err == nil {
		t.Error("Parsed empty body without an error")
	}
	test, err := ParseReqBody([]byte(testBody))
	if err != nil {
		t.Error(err)
	}
	if len(test.Payload) != 3 {
		t.Error("Wrong size of the parsed payload")
	}
}

var testBody = `{"payload": [{"country": "UK","description": "What's life like when you have enough children to field your own football team?","drm": true,"episodeCount": 3,"genre": "Reality","image": {"showImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/16KidsandCounting1280.jpg"},"language": "English","nextEpisode": null,"primaryColour": "#ff7800","seasons": [{"slug": "show/16kidsandcounting/season/1"}],"slug": "show/16kidsandcounting","title": "16 Kids and Counting","tvChannel": "GEM"},{"slug": "show/seapatrol","title": "Sea Patrol","tvChannel": "Channel 9"},{"country": " USA","description": "The Taste puts 16 culinary competitors in the kitchen, where four of the World's most notable culinary masters of the food world judges their creations based on a blind taste. Join judges Anthony Bourdain, Nigella Lawson, Ludovic Lefebvre and Brian Malarkey in this pressure-packed contest where a single spoonful can catapult a contender to the top or send them packing.","drm": true,"episodeCount": 2,"genre": "Reality","image": {"showImage": "http://catchup.ninemsn.com.au/img/jump-in/shows/TheTaste1280.jpg"},"language": "English","nextEpisode": {"channel": null,"channelLogo": "http://catchup.ninemsn.com.au/img/player/logo_go.gif","date": null,"html": "<br><span class=\"visit\">Visit the Official Website</span></span>","url": "http://go.ninemsn.com.au/"},"primaryColour": "#df0000","seasons": [{"slug": "show/thetaste/season/1"}],"slug": "show/thetaste","title": "The Taste (Le Goût)","tvChannel": "GEM"}],"skip": 0,"take": 10,"totalRecords": 75}`



