package crons

// import (
// 	"context"

// 	"srating/domain"
// 	"srating/utils"

// 	"github.com/robfig/cron/v3"
// )

// type Cron struct {
// 	reservationService domain.ReservationService
// 	*cron.Cron
// }

// func NewCron(reservationService domain.ReservationService) *Cron {
// 	return &Cron{
// 		reservationService: reservationService,
// 		Cron:               cron.New(),
// 	}
// }

// func (c *Cron) StartC(ctx context.Context) {
// 	_, err := c.AddFunc("0 0 * * *", func() {
// 		// if err := c.reservationService.UpdateReservationStatus(ctx); err != nil {
// 		// 	utils.LogError(err, "Failed")
// 		// 	return
// 		// }
// 		// utils.LogInfo("Success")
// 	})
// 	if err != nil {
// 		utils.LogError(err, "Failed")
// 		return
// 	}
// 	c.Start()
// }

// func (c *Cron) Reload(ctx context.Context) {
// 	c.Stop()
// 	c.StartC(ctx)
// }
