package tasks

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"homeservice/src/constant"
// 	"homeservice/src/entities"

// 	"github.com/appleboy/go-fcm"
// 	"github.com/hibiken/asynq"
// 	"go.uber.org/zap"
// )

// type ActionType string

// const (
// 	TypeOrderChangeStatus = "employee:order_change_status"
// 	AssignToEmployee      = ActionType("assign_to_employee")
// 	CustomerCreateOrder   = ActionType("customer_create_order")
// 	EmployeeAcceptOrder   = ActionType("employee_accept_order")
// )

// type IAuthService interface {
// 	FindByTypeable(ctx context.Context, typeableId int, typeable int) (*entities.User, error)
// }

// type IOrderService interface {
// 	FindById(ctx context.Context, id int) (*entities.Order, error)
// }

// type FCMNotifyPayload struct {
// 	EmployeeID int
// 	CustomerID int
// 	OrderID    int
// 	ActionType ActionType
// }

// type FCMNotifyProcessor struct {
// 	AuthService  IAuthService
// 	OrderService IOrderService
// 	fcmClient    IFcmClient
// }

// type IFcmClient interface {
// 	Send(msg *fcm.Message) (*fcm.Response, error)
// }

// func NewFcmNotifyTask(payload *FCMNotifyPayload) (*asynq.Task, error) {
// 	b := new(bytes.Buffer)
// 	err := json.NewEncoder(b).Encode(&payload)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return asynq.NewTask(TypeOrderChangeStatus, b.Bytes()), nil
// }

// func (processor *FCMNotifyProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
// 	fmt.Println("Đã nhận task")
// 	var p FCMNotifyPayload
// 	if err := json.Unmarshal(t.Payload(), &p); err != nil {
// 		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
// 	}
// 	var order, err = processor.OrderService.FindById(ctx, p.OrderID)
// 	if err != nil {
// 		return err
// 	}
// 	var user *entities.User
// 	switch order.Status {
// 	case entities.AccpetedOrderStatus:
// 		user, _ = processor.AuthService.FindByTypeable(ctx, order.CustomerId, constant.TypeCustomer)
// 	case entities.WaitingOrderStatus:
// 		if order.EmployeeId != 0 {
// 			user, _ = processor.AuthService.FindByTypeable(ctx, order.EmployeeId, constant.TypeEmployee)
// 		}
// 	}
// 	// if p.EmployeeID != 0 {
// 	// 	user, _ = processor.AuthService.FindByTypeable(ctx, p.EmployeeID, constant.TypeEmployee)
// 	// } else {
// 	// 	user, _ = processor.AuthService.FindByTypeable(ctx, p.CustomerID, constant.TypeCustomer)
// 	// }
// 	if user != nil {
// 		if user.PushToken != "" {
// 			switch order.Status {
// 			case entities.WaitingOrderStatus:
// 				if order.EmployeeId != 0 {
// 					fmt.Println("Đơn hàng đã đc assign", user.PushToken)
// 					if _, err := processor.fcmClient.Send(&fcm.Message{
// 						To: user.PushToken,
// 						Data: map[string]interface{}{
// 							"orderID": p.OrderID,
// 						},
// 						Notification: &fcm.Notification{
// 							Title: "Chúc mừng",
// 							Body:  "Bạn đã nhận được đơn hàng mới",
// 						},
// 					}); err != nil {
// 						zap.S().Error(err)
// 						return err
// 					}
// 				}

// 			case entities.AccpetedOrderStatus:
// 				fmt.Println("Đơn hàng đã đc thợ nhận", user.PushToken)
// 				if _, err := processor.fcmClient.Send(&fcm.Message{
// 					To: user.PushToken,
// 					Data: map[string]interface{}{
// 						"orderID": p.OrderID,
// 					},
// 					Notification: &fcm.Notification{
// 						Title: "Chúc mừng",
// 						Body:  "Đơn hàng của bạn đã được thợ nhận",
// 					},
// 				}); err != nil {
// 					zap.S().Error(err)
// 					return err
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

// func NewFCMNotifyProcessor(client IFcmClient, authService IAuthService, orderService IOrderService) *FCMNotifyProcessor {
// 	return &FCMNotifyProcessor{
// 		AuthService:  authService,
// 		OrderService: orderService,
// 		fcmClient:    client,
// 	}
// }
