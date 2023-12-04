package usecase

import (
	"booking/internal/booking/entity"
	"booking/internal/booking/kafka"
	"booking/internal/booking/repository"
	"booking/internal/booking/server/consumer/dto"
	"booking/pkg/metrics"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"strconv"
)

type BookingUC struct {
	l *zap.SugaredLogger
	r *repository.Repo
	k *kafka.Producer
}

func NewBookingUC(l *zap.SugaredLogger, r *repository.Repo, k *kafka.Producer) *BookingUC {
	return &BookingUC{
		l: l,
		r: r,
		k: k,
	}
}

func (b *BookingUC) BookRoom() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		param := ctx.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			b.l.Infof("error with convertaion %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		randNum := rand.Intn(10)
		randNum2 := rand.Intn(10)
		randNum3 := rand.Intn(10)
		randNum4 := rand.Intn(10)
		strNum := fmt.Sprintf("%d%d%d%d", randNum, randNum3, randNum4, randNum2)
		num, _ := strconv.Atoi(strNum)
		reqInt, err := b.r.BookRoom(strconv.Itoa(id), num)
		roomBusyError := errors.New("all rooms are busy")
		if errors.Is(roomBusyError, err) {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "all rooms are busy"})
			return
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "cannot book room"})
			return
		}

		msg := dto.BookCode{Code: fmt.Sprintf("%d", num)}

		bm, err := json.Marshal(&msg)
		if err != nil {
			b.l.Errorf("Failed to marshal BookCode: %s", err)
			return
		}
		b.k.ProduceMessage(bm)
		metrics.HttpBookTotal.Inc()
		ctx.JSON(http.StatusOK, gin.H{"message": "room is booked, verify your book", "requestID": reqInt})
	}
}

func (b *BookingUC) ConfirmBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bookReq entity.BookRequest
		if err := ctx.ShouldBindJSON(&bookReq); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		err := b.r.ConfirmBook(&bookReq)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		repository.Jobs <- bookReq.Id
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func (b *BookingUC) CreateHotel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var hotel entity.Hotel
		if err := ctx.ShouldBindJSON(&hotel); err != nil {
			b.l.Infof("cannot unmarshall hotel with err:%v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		id, err := b.r.CreateHotel(context.Background(), &hotel)
		if err != nil {
			b.l.Infof("cannot unmarshall hotel with err:%v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, id)
	}
}

func (b *BookingUC) GetHotels() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hotels, err := b.r.GetHotels(context.Background())
		if err != nil {
			b.l.Infof("cannot get hotels with err:%v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, hotels)
	}
}

func (b *BookingUC) UpdateHotel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			b.l.Infof("error with convertaion %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		var hotel entity.Hotel
		if err := ctx.ShouldBindJSON(&hotel); err != nil {
			b.l.Infof("cannot unmarshall hotel with err:%v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		err = b.r.UpdateHotel(context.Background(), strconv.Itoa(id), &hotel)
		if err != nil {
			b.l.Infof("cannot update hotel with err:%v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, id)
	}
}

func (b *BookingUC) DeleteHotel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			b.l.Infof("error with convertaion %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if id == 0 {
			b.l.Info("user get empty id")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not be empty"})
			return
		}
		err = b.r.DeleteHotel(context.Background(), strconv.Itoa(id))
		if err != nil {
			b.l.Infof("cannot delete hotel with err:%v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": id})
	}
}

func (b *BookingUC) GetHotelById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			b.l.Infof("error with convertaion %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		hotel, err := b.r.GetHotelById(context.Background(), strconv.Itoa(id))
		if err != nil {
			b.l.Infof("cannot get hotel by with id %s with err:%v", id, err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, hotel)
	}
}
