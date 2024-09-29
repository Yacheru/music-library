package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"music-library/internal/entities"
	"music-library/internal/service"
	"music-library/pkg/constants"
	"net/http"
	"strconv"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// EditSong
// @Summary      Edit Song
// @Description  Edit song
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        body body entities.Song true "Body"
// @Param        title   query    string  false  "edit song by title"
// @Param        link    query    string  false  "edit song by link"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.Response
// @Failure      404  {object}  handlers.Response
// @Failure      500  {object}  handlers.Response
// @Router       / [patch]
func (h *Handler) EditSong(ctx *gin.Context) {
	var entitySong = new(entities.Song)
	if err := ctx.ShouldBindJSON(entitySong); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	title := ctx.Query("title")
	link := ctx.Query("link")

	newSong, err := h.service.Music.EditSong(ctx, title, link, entitySong)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NewErrorResponse(ctx, http.StatusNotFound, "Song not found")
			return
		}

		NewErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Song edited successfully", newSong)
}

// DeleteSong
// @Summary      Delete Song
// @Description  Delete song
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        title   query    string  false  "delete song by title"
// @Param        link    query    string  false  "delete song by link"
// @Success      400  {object}  handlers.Response
// @Success      200  {object}  handlers.Response
// @Failure      404  {object}  handlers.Response
// @Failure      500  {object}  handlers.Response
// @Router       / [delete]
func (h *Handler) DeleteSong(ctx *gin.Context) {
	title := ctx.Query("title")
	link := ctx.Query("link")

	if err := h.service.Music.DeleteSong(ctx, title, link); err != nil {
		if errors.Is(err, constants.SongNotFoundError) {
			NewErrorResponse(ctx, http.StatusNotFound, "song not found")
			return
		}

		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Song successfully deleted", nil)
	return
}

// GetVerse
// @Summary      Get Verses
// @Description  Get Verses
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        title   query    string  false  "get verses by title"
// @Param        link    query    string  false  "get verses by link"
// @Param        limit   query    int  false  "set limit"
// @Param        offset    query    int  false  "set offset"
// @Failure      400  {object}  handlers.Response
// @Success      200  {object}  handlers.Response
// @Failure      404  {object}  handlers.Response
// @Failure      500  {object}  handlers.Response
// @Router       /verse [get]
func (h *Handler) GetVerse(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	title := ctx.Query("title")
	link := ctx.Query("title")

	verses, err := h.service.Music.GetVerse(ctx, title, link, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NewErrorResponse(ctx, http.StatusNotFound, "song not found")
			return
		}

		NewErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Verses", verses)
}

// GetAllSongs
// @Summary      Get All Songs
// @Description  Get All Songs
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        limit   query    int  false  "set limit"
// @Param        offset    query    int  false  "set offset"
// @Param        filter    query    string  false  "set filter"
// @Failure      400  {object}  handlers.Response
// @Success      200  {object}  handlers.Response
// @Failure      404  {object}  handlers.Response
// @Failure      500  {object}  handlers.Response
// @Router       /all [get]
func (h *Handler) GetAllSongs(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	var (
		songs []*entities.Song
		err   error
	)

	filter, exists := ctx.Get("filter")
	if exists {
		songs, err = h.service.Music.GetAllSongs(ctx, limit, offset, filter.(*entities.Filter))
	} else {
		songs, err = h.service.Music.GetAllSongs(ctx, limit, offset, nil)
	}

	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if songs == nil {
		NewErrorResponse(ctx, http.StatusNotFound, "No songs found")
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "All songs here!", songs)
	return
}

// StorageNewSong
// @Summary      Storage New Song
// @Description  Storage New Song
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        body body entities.NewSong true "Body"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.Response
// @Failure      404  {object}  handlers.Response
// @Failure      408  {object}  handlers.Response
// @Failure      409  {object}  handlers.Response
// @Failure      500  {object}  handlers.Response
// @Router       /new [post]
func (h *Handler) StorageNewSong(ctx *gin.Context) {
	var newSongEntity = new(entities.NewSong)
	if err := ctx.ShouldBindJSON(newSongEntity); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Music.StorageNewSong(ctx, newSongEntity)
	if err != nil {
		if errors.Is(err, constants.NoLyricsFoundError) {
			NewErrorResponse(ctx, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, constants.TimeOutError) {
			NewErrorResponse(ctx, http.StatusRequestTimeout, err.Error())
			return
		}
		if errors.Is(err, constants.SongAlreadyExistsError) {
			NewErrorResponse(ctx, http.StatusConflict, err.Error())
			return
		}

		NewErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "song saved successfully", nil)
	return
}
