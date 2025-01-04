package controllers

import (
	"context"
	"database/sql"
	"errors"
	"github.com/germatpreis/eldie/server/db/sqlc"
	"github.com/germatpreis/eldie/server/schemas"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

type ContactController struct {
	db  *db.Queries
	ctx context.Context
}

func NewContactController(db *db.Queries, ctx context.Context) *ContactController {
	return &ContactController{
		db,
		ctx,
	}
}

func (cc *ContactController) CreateContact(ctx *gin.Context) {
	var payload *schemas.CreateContact

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		// this validates the request against the schema
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed Payload", "error": err.Error()})
		return
	}

	now := time.Now()
	args := &db.CreateContactParams{
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		PhoneNumber: payload.PhoneNumber,
		Street:      payload.Street,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	contact, err := cc.db.CreateContact(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed to retrieve contact", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "successfully created contact", "contact": contact})
}

func (cc *ContactController) UpdateContact(ctx *gin.Context) {
	var payload *schemas.UpdateContact
	contactId := ctx.Param("contactId")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed Payload", "error": err.Error()})
		return
	}

	now := time.Now()
	args := &db.UpdateContactParams{
		ContactID:   uuid.MustParse(contactId),
		FirstName:   sql.NullString{String: payload.FirstName, Valid: payload.FirstName != ""},
		LastName:    sql.NullString{String: payload.LastName, Valid: payload.LastName != ""},
		PhoneNumber: sql.NullString{String: payload.PhoneNumber, Valid: payload.PhoneNumber != ""},
		Street:      sql.NullString{String: payload.Street, Valid: payload.Street != ""},
		UpdatedAt:   sql.NullTime{Time: now, Valid: true},
	}

	contact, err := cc.db.UpdateContact(ctx, *args)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to retrieve contact with this ID"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving contact", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "successfully updated contact", "contact": contact})
}

func (cc *ContactController) GetContactById(ctx *gin.Context) {
	contactId := ctx.Param("contactId")

	contact, err := cc.db.GetContactById(ctx, uuid.MustParse(contactId))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to retrieve contact with this ID"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving contact", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully retrieved contact", "contact": contact})
}

func (cc *ContactController) GetAllContacts(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "0")
	var limit = ctx.DefaultQuery("limit", "10")

	reqPageId, _ := strconv.Atoi(page)
	reqLimit, _ := strconv.Atoi(limit)

	params := &db.ListContactsParams{Limit: int32(reqLimit), Offset: int32(reqPageId)}

	contacts, err := cc.db.ListContacts(ctx, *params)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed to retrieve contacts", "error": err.Error()})
		return
	}

	if contacts == nil {
		contacts = []db.Contact{}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully retrieved contacts", "result": contacts})
}

func (cc *ContactController) DeleteContact(ctx *gin.Context) {
	var contactId = ctx.Param("contactId")

	_, err := cc.db.GetContactById(ctx, uuid.MustParse(contactId))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to retrieve contact with this ID"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving contact", "error": err.Error()})
		return
	}

	err = cc.db.DeleteContact(ctx, uuid.MustParse(contactId))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "failed", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "successfully deleted"})
}
