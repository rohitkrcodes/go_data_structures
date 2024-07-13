package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type addDataRequest struct {
	Data int `json:"data" binding:"required,min=1,max=100000"`
}

type listDataRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,max=10"`
}

type getDataRequest struct {
	Data int `uri:"id" binding:"required,min=1,max=100000"`
}

type removeDataRequest struct {
	Data int `uri:"id" binding:"required,min=1,max=100000"`
}

func (server *Server) addData(ctx *gin.Context) {

	var req addDataRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.db.Append(req.Data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "Add successful"})
}

func (server *Server) listData(ctx *gin.Context) {
	var req listDataRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := server.db.GetList(int(req.PageID), int(req.PageSize))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (server *Server) getData(ctx *gin.Context) {
	var req getDataRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := server.db.GetData(req.Data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (server *Server) removeData(ctx *gin.Context) {
	var req removeDataRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.db.Remove(req.Data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "Remove successful"})
}
