package http

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"payservice/internal/payment/entity"
	"payservice/internal/payment/payment"
)

type EndpointHandler struct {
	paymentService payment.UseCase
	logger         *zap.SugaredLogger
	wg             sync.WaitGroup
}

func NewEndpointHandler(paymentService payment.UseCase, l *zap.SugaredLogger) *EndpointHandler {
	return &EndpointHandler{
		paymentService: paymentService,
		logger:         l,
	}
}

// GetCardPayments		godoc
// @Summary			card to card payments
// @Description		Get Payments from card to card payments. Need Authorization Header
// @Success			200
// @Router			/card/payments [get]
func (h *EndpointHandler) GetCardPayments(ctx *gin.Context) {
	resp, err := ctx.Get("user")

	if !err {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	var params entity.QueryParams

	params.SortBy = ctx.Query("sort_by")
	params.SortOrder = ctx.Query("sort_order")
	params.Search = ctx.Query("search")
	params.SearchQuery = ctx.Query("q")

	payments, e := h.paymentService.UserPayments(resp.(int), params)
	if e != nil {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	h.logger.Info("Getting payment is completed")
	ctx.JSON(http.StatusOK, payments)
}

// GetCardPayment		godoc
// @Summary			card to card payment
// @Description		Get Payments from card to card payment. Need Authorization Header
// @Success			200
// @Router			/card/payment/:id [get]
func (h *EndpointHandler) GetCardPayment(ctx *gin.Context) {
	resp, err := ctx.Get("user")
	if !err {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}
	sid := ctx.Param("id")
	id, e := strconv.Atoi(sid)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, "Bad request")
	}

	payments, e := h.paymentService.PaymentByID(resp.(int), id)
	if e != nil {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	h.logger.Info("Getting payment is completed")
	ctx.JSON(http.StatusOK, payments)
}

// GetCards		godoc
// @Summary			get user's all cards
// @Description		Get user's all cards. Need Authorization Header
// @Success			200
// @Router			/cards [get]
func (h *EndpointHandler) GetCards(ctx *gin.Context) {
	resp, err := ctx.Get("user")
	if !err {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	cards, e := h.paymentService.UserCards(resp.(int))
	if e != nil {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	h.logger.Info("Getting cards is completed")
	ctx.JSON(http.StatusOK, cards)
}

// CreateCard		godoc
// @Summary			Create Card for user
// @Description		Get Card data from request and create card for user. Need Authorization Header
// @Param			request body Card true "CreateCard"
// @Produce			application/json
// @Success			200
// @Router			/card [post]
func (h *EndpointHandler) CreateCard(ctx *gin.Context) {
	resp, err := ctx.Get("user")
	if !err {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	var card entity.Card
	e := ctx.ShouldBindJSON(&card)
	if e != nil {
		h.logger.Error("Bad request")
		ctx.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	e = h.paymentService.ValidateCard(resp.(int), card)
	if e != nil {
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	h.logger.Info("Creating card is completed")
	ctx.JSON(http.StatusOK, "Card is created")
}

// CreateCardPayment		godoc
// @Summary			Create Payment from card to card
// @Description		Create Payment from card to card. Need Authorization Header and Card id
// @Param			request body Card true "CreateCardPayment"
// @Produce			application/json
// @Success			200
// @Router			/card/payment [post]
func (h *EndpointHandler) CreateCardPayment(ctx *gin.Context) {
	resp, err := ctx.Get("user")
	if !err {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not Authorized")
		return
	}

	var transaction entity.CardTransaction
	e := ctx.ShouldBindJSON(&transaction)
	if e != nil {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	e = h.paymentService.ValidateTransaction(transaction, resp.(int))
	if e != nil {
		h.logger.Error("Error with Validating transaction")
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	h.logger.Info("Creating payment is completed")
	ctx.JSON(http.StatusOK, "Payment is created")
}

// UpdateCard		godoc
// @Summary			Update Card for user
// @Description		Get Card data from request and card id for update user card. Need Authorization Header
// @Param			request body Card true "UpdateCardCard"
// @Produce			application/json
// @Success			200
// @Router			/card/:id [put]
func (h *EndpointHandler) UpdateCard(ctx *gin.Context) {
	resp, e := ctx.Get("user")
	if !e {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad param")
	}

	var card entity.Card
	err = ctx.ShouldBindJSON(&card)
	if err != nil {
		h.logger.Error("Bad request")
		ctx.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	err = h.paymentService.UpdateCard(id, card)
	if err != nil {
		h.logger.Error("Error with delete card")
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	ctx.JSON(http.StatusOK, "Card is updated")
}

// DeleteCard		godoc
// @Summary			Update Card for user
// @Description		Get Card id for delete user card. Need Authorization Header
// @Produce			application/json
// @Success			200
// @Router			/card/:id [delete]
func (h *EndpointHandler) DeleteCard(ctx *gin.Context) {
	resp, e := ctx.Get("user")
	if !e {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad param")
	}

	err = h.paymentService.DeleteCard(id)
	if err != nil {
		h.logger.Error("Error with delete card")
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	ctx.JSON(http.StatusOK, "Card is removed")
}

// GetAccounts		godoc
// @Summary			get user's all accounts
// @Description		Get user's all accounts. Need Authorization Header
// @Success			200
// @Router			/accounts [get]
func (h *EndpointHandler) GetAccounts(ctx *gin.Context) {
	resp, e := ctx.Get("user")
	if !e {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	accounts, err := h.paymentService.GetAccounts(resp.(int))
	if err != nil {
		h.logger.Errorf("Error with GetAccounts. User id is %d", resp.(int))
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

// GetAccount		godoc
// @Summary			get user's account by id
// @Description		Get user's account by id. Need Authorization Header
// @Success			200
// @Router			/account/:id [get]
func (h *EndpointHandler) GetAccount(ctx *gin.Context) {
	resp, e := ctx.Get("user")
	if !e {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad param")
	}

	account, err := h.paymentService.GetAccount(resp.(int), id)
	if err != nil {
		h.logger.Errorf("Error with GetAccounts. User id is %d", resp.(int))
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// CreateAccount		godoc
// @Summary			Create Account for user
// @Description		Create Account for user with balance and currency. Need Authorization Header and currency
// @Param			request body Account true "CreateAccount"
// @Produce			application/json
// @Success			200
// @Router			/account [post]
func (h *EndpointHandler) CreateAccount(ctx *gin.Context) {
	resp, e := ctx.Get("user")
	if !e {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	var account entity.Account
	err := ctx.ShouldBindJSON(&account)
	if err != nil {
		h.logger.Errorf("Error with Marshal JSON in CreateAccount. User id is %d", resp.(int))
		ctx.JSON(http.StatusBadRequest, "Bad request")
	}

	err = h.paymentService.CreateAccount(resp.(int), account)
	if err != nil {
		h.logger.Errorf("Error with CreateAccount. User id is %d", resp.(int))
		ctx.JSON(http.StatusInternalServerError, "Error with creating account")
	}

	ctx.JSON(http.StatusOK, "Account is created")
}

// DeleteAccount		godoc
// @Summary			Delete User Account
// @Description		Delete User Account with balance and currency. Need Authorization Header and account id
// @Success			200
// @Router			/account/:id [delete]
func (h *EndpointHandler) DeleteAccount(ctx *gin.Context) {
	resp, e := ctx.Get("user")
	if !e {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad param")
	}

	err = h.paymentService.DeleteAccount(id)
	if err != nil {
		h.logger.Errorf("Error with DeleteAccount. User id is %d", resp.(int))
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	ctx.JSON(http.StatusOK, "Account is deleted")
}

// CreateAccountPayment		godoc
// @Summary			Create Payment from account to account
// @Description		Create Payment from account to account. Need Authorization Header and account id
// @Param			request body AccountTransaction true "CreateAccountPayment"
// @Produce			application/json
// @Success			200
// @Router			/account/payment [post]
func (h *EndpointHandler) CreateAccountPayment(ctx *gin.Context) {
	resp, e := ctx.Get("user")
	if !e {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	var transaction entity.AccountTransaction
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		h.logger.Errorf("Error with JSON Marshal in CreateAccountPayment. User id is %d", resp.(int))
		ctx.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	h.wg.Add(1)
	go func() {
		err := h.paymentService.ValidateAccountTransaction(resp.(int), transaction, &h.wg)
		if err != nil {
			h.logger.Error("Error with Validating transaction")
			ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}()

	/*if err != nil {
		h.logger.Error("Error with Validating transaction")
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}*/

	h.logger.Info("Creating payment is completed")
	ctx.JSON(http.StatusOK, "Payment is created")
}

// GetAccountPayments		godoc
// @Summary			account to account payment
// @Description		Get All from account to account payments. Need Authorization Header
// @Success			200
// @Router			/account/payments [get]
func (h *EndpointHandler) GetAccountPayments(ctx *gin.Context) {
	resp, e := ctx.Get("user")
	if !e {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	transactions, err := h.paymentService.GetAccountPayments(resp.(int))
	if err != nil {
		h.logger.Error("Error with Get transactions")
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

// GetAccountPayment		godoc
// @Summary			account to account payment
// @Description		Get from account to account payment. Need Authorization Header
// @Success			200
// @Router			/account/payment/:id [get]
func (h *EndpointHandler) GetAccountPayment(ctx *gin.Context) {
	resp, e := ctx.Get("user")
	if !e {
		h.logger.Errorf("Not authorized. User id is %d", resp.(int))
		ctx.JSON(http.StatusUnauthorized, "Not authorized")
		return
	}

	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad param")
	}

	transactions, err := h.paymentService.GetAccountPayment(resp.(int), id)
	if err != nil {
		h.logger.Error("Error with Get transaction")
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

func (h *EndpointHandler) GetAdminCardPayments(ctx *gin.Context) {
	var params entity.QueryParams

	params.SortBy = ctx.Query("sort_by")
	params.SortOrder = ctx.Query("sort_order")
	params.Search = ctx.Query("search")
	params.SearchQuery = ctx.Query("q")

	payments, err := h.paymentService.UserPayments(0, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	h.logger.Info("Admin: Getting payment is completed")

	ctx.JSON(http.StatusOK, payments)
}

func (h *EndpointHandler) DeleteAdminCardPayment(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad param")
	}

	err = h.paymentService.RemoveCardPayment(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	h.logger.Info("Admin: Removing payment is completed")
	ctx.JSON(http.StatusOK, "Payment is removed")
}
