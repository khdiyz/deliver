package helper

import "deliver/internal/constants"

func IsValidOrderStatus(status string) bool {
	pickedUp := constants.OrderStatusPickedUp
	onDelivery := constants.OrderStatusOnDelivery
	delivered := constants.OrderStatusDelivered
	collected := constants.OrderStatusPaymentCollected

	if status != pickedUp && status != onDelivery && status != delivered && status != collected {
		return false
	}

	return true
}
