package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kurmidev/webox/utils"
)

type BouqueList struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Mrp         float32    `json:"mrp"`
	IfFixNcf    bool       `json:"ifFixNCF"`
	Type        int        `json:"type"`
	Description string     `json:"description"`
	SortBy      int        `json:"sort_by"`
	BoxtypeLbl  string     `json:"boxtype_lbl"`
	TypeLbl     string     `json:"type_lbl"`
	Alacarte    []Alacarte `json:"alacarte"`
	Package     []Package  `json:"package"`
}

type Alacarte struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Code            string  `json:"code"`
	BroadcasterRate float32 `json:"broadcasterRate"`
	IsAlacarte      string  `json:"isAlacarte"`
	IsNcf           string  `json:"isNCF_lbl"`
	IsFta           int     `json:"isFta_lbl"`
	BroadcasterName string  `json:"broadcaster_lbl"`
	ChannelType     string  `json:"channel_type_lbl"`
}

type Package struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Code              string  `json:"code"`
	BroadcasterRate   float32 `json:"broadcasterRate"`
	PackageTypeLbl    string  `json:"package_type_lbl"`
	IsFta             string  `json:"isFta"`
	TotalChannelCount int     `json:"totalChannelCount"`
	NcfChannelCount   int     `json:"ncfChannelCount"`
	PayChannelCount   int     `json:"payChannelCount"`
	FtaChannelCount   int     `json:"ftaChannelCount"`
	ChannelIds        []int   `json:"channelIds"`
}

func (h *Handlers) BouqueList(w http.ResponseWriter, r *http.Request) {

	bouques, err := h.Models.GetBouques()
	if err != nil {
		message := map[string]string{"number": "No bouquet founds."}
		h.Common.WriteJSON(w, http.StatusBadRequest, message)
		return
	}

	var bouquList []BouqueList

	for _, b := range bouques {
		rb := BouqueList{
			ID:          b.ID,
			Name:        b.Name,
			Mrp:         float32(b.Mrp),
			IfFixNcf:    b.GetIsCutomNcf(),
			Type:        b.Type,
			Description: b.Description,
			SortBy:      b.SortBy,
			BoxtypeLbl:  utils.BoxTypeLbl(b.IsHd),
			TypeLbl:     utils.BouqueTypeslbl(b.Type),
		}

		var alacarte []Alacarte
		var packages []Package
		for _, assoc := range b.BouqueAssetAssoc {
			if assoc.PackageId > 0 {
				counts := assoc.Package.GetCount()
				packages = append(packages, Package{
					ID:                assoc.PackageId,
					Name:              assoc.Package.Name,
					Code:              assoc.Package.Code,
					BroadcasterRate:   assoc.Package.BroadcasterRate,
					PackageTypeLbl:    utils.BoxTypeLbl(assoc.Package.IsHd),
					IsFta:             utils.IsApp(assoc.Package.IsFta),
					TotalChannelCount: len(counts["total"]),
					NcfChannelCount:   len(counts["ncf"]),
					PayChannelCount:   len(counts["pay"]),
					FtaChannelCount:   len(counts["fta"]),
					ChannelIds:        counts["total"],
				})
			}

			if assoc.ChannelId > 0 {
				alacarte = append(alacarte, Alacarte{
					ID:              assoc.Channel.ID,
					Name:            assoc.Channel.Name,
					Code:            assoc.Channel.Code,
					BroadcasterRate: assoc.Channel.BroadcasterRate,
					IsAlacarte:      utils.IsApp(assoc.Channel.IsAlacarte),
					IsNcf:           utils.IsApp(assoc.Channel.IsNcf),
					IsFta:           assoc.Channel.IsFta,
					BroadcasterName: assoc.Channel.Broadcaster.Name,
					ChannelType:     utils.BoxTypeLbl(assoc.Channel.IsHd),
				})
			}
		}
		rb.Alacarte = alacarte
		rb.Package = packages
		bouquList = append(bouquList, rb)
	}
	h.Common.WriteJSON(w, http.StatusOK, bouquList)
}

type ChannelList struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	LogoUrl string `json:"logo_url"`
}

type GenreList struct {
	GenreId       int           `json:"genre_id"`
	GenreName     string        `json:"genre_name"`
	ChannelsCount int           `json:"channels_count"`
	ChannelList   []ChannelList `json:"channel_list"`
}

type SingleBouque struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Channels []GenreList `json:"channels"`
}

func (h *Handlers) Bouque(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	bouque, err := h.Models.GetBouque(id)

	if bouque.ID <= 0 || err != nil {
		message := map[string]string{"number": "No bouquet founds."}
		h.Common.WriteJSON(w, http.StatusBadRequest, message)
		return
	}

	var genreList = make(map[int]GenreList)

	for _, bsa := range bouque.BouqueAssetAssoc {
		if bsa.PackageId > 0 {
			for _, bchan := range bsa.Package.ChannelPackageAssoc {
				if _, ok := genreList[bchan.Channel.GenreId]; ok {
					r := genreList[bchan.Channel.GenreId]
					r.ChannelList = append(r.ChannelList, ChannelList{
						ID:      bchan.ChannelId,
						Name:    bchan.Channel.Name,
						Type:    utils.BoxTypeLbl(bchan.Channel.IsHd),
						LogoUrl: fmt.Sprintf("%s%d", os.Getenv("BASEURL_LOGO"), bchan.ChannelId),
					})
					r.ChannelsCount = len(r.ChannelList)
					genreList[bchan.Channel.GenreId] = r
				} else {
					var gl []ChannelList
					gl = append(gl, ChannelList{
						ID:      bchan.ChannelId,
						Name:    bchan.Channel.Name,
						Type:    utils.BoxTypeLbl(bchan.Channel.IsHd),
						LogoUrl: fmt.Sprintf("%s%d", os.Getenv("BASEURL_LOGO"), bchan.ChannelId),
					})
					genreList[bchan.Channel.GenreId] = GenreList{
						GenreId:       bchan.Channel.GenreId,
						GenreName:     bchan.Channel.Genre.Name,
						ChannelsCount: 1,
						ChannelList:   gl,
					}
				}
			}
		} else {
			if _, ok := genreList[bsa.Channel.GenreId]; ok {
				r := genreList[bsa.Channel.GenreId]
				r.ChannelList = append(r.ChannelList, ChannelList{
					ID:      bsa.ChannelId,
					Name:    bsa.Channel.Name,
					Type:    utils.BoxTypeLbl(bsa.Channel.IsHd),
					LogoUrl: fmt.Sprintf("%s%d", os.Getenv("BASEURL_LOGO"), bsa.ChannelId),
				})
				r.ChannelsCount = len(r.ChannelList)
				genreList[bsa.Channel.GenreId] = r
			} else {
				var gl []ChannelList
				gl = append(gl, ChannelList{
					ID:      bsa.ChannelId,
					Name:    bsa.Channel.Name,
					Type:    utils.BoxTypeLbl(bsa.Channel.IsHd),
					LogoUrl: fmt.Sprintf("%s%d", os.Getenv("BASEURL_LOGO"), bsa.ChannelId),
				})
				genreList[bsa.Channel.GenreId] = GenreList{
					GenreId:       bsa.Channel.GenreId,
					GenreName:     bsa.Channel.Genre.Name,
					ChannelsCount: 1,
					ChannelList:   gl,
				}
			}
		}
	}

	var gl []GenreList

	for _, lst := range genreList {
		gl = append(gl, lst)
	}

	data := SingleBouque{
		ID:       bouque.ID,
		Name:     bouque.Name,
		Channels: gl,
	}

	h.Common.WriteJSON(w, http.StatusOK, data)
}
