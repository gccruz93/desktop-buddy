package emotes

import (
	"bytes"
	"desktop-buddy/pkg/helpers"
	"image/gif"
	"os"
)

type Emote struct {
	name      string
	frametime int
	gif       *helpers.CustomGif
}

var emotesList map[string]*Emote

func loadEmote(name string, frametime int) bool {
	if _, ok := emotesList[name]; ok {
		return true
	}
	file, err := os.ReadFile("./assets/emotes/" + name + ".gif")
	if err != nil {
		return false
	}
	g, _ := gif.DecodeAll(bytes.NewReader(file))
	emotesList[name] = &Emote{
		name:      name,
		frametime: frametime,
		gif:       helpers.SplitAnimatedGIF(g),
	}
	return true
}

func loadEmotesConfig() {

}
