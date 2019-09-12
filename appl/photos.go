package appl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
)

func (rnds *renditions) ParsePhotoComponent(pc pubcomponent, mt mediaType) {
	key := strings.ToLower(pc.Role)
	chars := pc.Node.Node("Characteristics")

	switch key {
	case "main":
		jo := rnds.GetRendition("Full Resolution (JPG)", pc.Role, mediaTypePhoto, pc.Node, chars)
		rnds.AddRendition(key, jo, mt != mediaTypePhoto)
	case "preview", "thumbnail":
		if mt == mediaTypeVideo {
			n, k := getBinaryName(chars, "", true)
			key = key + k

			title := fmt.Sprintf("%s (%s)", pc.Role, n)

			jo := rnds.GetRendition(title, pc.Role, mediaTypePhoto, pc.Node, chars)
			rnds.AddRendition(key, jo, true)
		} else {
			jo := rnds.GetRendition(pc.Role+" (JPG)", pc.Role, mediaTypePhoto, pc.Node, chars)
			rnds.AddRendition(key, jo, false)
		}
	}
}

func (rnds *renditions) ParsePhotoCollection(pc pubcomponent, duration int64) {
	if pc.Node.Nodes == nil || !strings.EqualFold(pc.Role, "Thumbnail") {
		return
	}

	const (
		TokenDot rune = 46
		Token0   rune = 48
	)

	var (
		ext, sep string
		href     bool
		w, h     int
		zeros    []rune
		times    []int64
	)

	url := pc.Node.Attribute("BaseFileName")
	if url != "" {
		runes := []rune(url)
		i := len(runes) - 3

		for i >= 0 {
			r := runes[i]
			if href {
				if r == Token0 {
					if zeros == nil {
						zeros = []rune{r}
					} else {
						zeros = append(zeros, r)
					}
				} else {
					sep = string(r)
					url = string(runes[0:i])
					break
				}
			} else if r == TokenDot {
				ext = string(runes[i+1:])
				href = true
			}
			i--
		}
	}

	if sep == "" {
		sep = "/"
	}

	chars := pc.Node.Node("Characteristics")
	if chars.Nodes != nil {
		for _, n := range chars.Nodes {
			switch n.Name {
			case "Width":
				w, _ = strconv.Atoi(n.Text)
			case "Height":
				h, _ = strconv.Atoi(n.Text)
			}
		}
	}

	for _, n := range pc.Node.Nodes {
		if n.Name == "File" {
			v := n.Attribute("TimeOffSetMilliseconds")
			i, _ := strconv.ParseInt(v, 0, 64)
			if times == nil {
				times = []int64{i}
			} else {
				times = append(times, i)
			}
		}
	}

	if times != nil {
		shots := json.Array{}
		last := len(times) - 1
		var end int64

		for i, ms := range times {
			jo := json.Object{}
			jo.AddInt("seq", i+1)

			if href {
				s := strconv.Itoa(i)
				z := string(zeros[0 : len(zeros)-len(s)])
				s = fmt.Sprintf("%s%s%s%s.%s", url, sep, z, s, ext)
				jo.AddString("href", s)
			}

			if w > 0 {
				jo.AddInt("width", w)
			}

			if h > 0 {
				jo.AddInt("height", h)
			}

			jo.AddString("start", formatTime(ms))

			if i < last {
				end = times[i+1]
			} else {
				end = duration
			}

			jo.AddString("end", formatTime(end))

			jo.AddString("timeunit", "normalplaytime")

			shots.AddObject(jo)
		}

		jp := json.NewArrayProperty("shots", shots)

		if rnds.NonRenditions == nil {
			rnds.NonRenditions = []json.Property{jp}
		} else {
			rnds.NonRenditions = append(rnds.NonRenditions, jp)
		}
	}
}
