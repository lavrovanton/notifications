package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavrovanton/notifications/internal/api/request"
	"github.com/lavrovanton/notifications/internal/model"
)

type NotificationRepository interface {
	Fetch() (res []model.Notification, err error)
	Store(m *model.Notification) error
}

type NotificationController struct {
	repo NotificationRepository
}

func NewNotificationController(repo NotificationRepository) *NotificationController {
	return &NotificationController{repo}
}

type NotificationsResponse struct {
	Notifications []model.Notification `json:"notifications"`
}

// @Summary get notifications
// @Tags notifications
// @Accept json
// @Produce json
// @Success 200 {array} NotificationsResponse
// @Router /notifications [get]
func (c *NotificationController) Index(ctx *gin.Context) {
	notifications, err := c.repo.Fetch()
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, NotificationsResponse{notifications})
}

// @Summary create notification
// @Tags notifications
// @Accept json
// @Produce json
// @Param request body request.CreateNotification true "request body"
// @Success 200 {array} model.Notification
// @Router /notifications [post]
func (c *NotificationController) Create(ctx *gin.Context) {
	request := request.CreateNotification{}

	err := ctx.BindJSON(&request)

	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, ErrBadParamInput)
		return
	}

	notification := request.ToModel()

	if err = c.repo.Store(&notification); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, notification)
}
