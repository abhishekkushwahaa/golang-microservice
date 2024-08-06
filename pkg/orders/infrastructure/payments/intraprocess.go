package payments

import (
	"github.com/abhishekkushwahaa/golang-microservice/pkg/common/price"
	"github.com/abhishekkushwahaa/golang-microservice/pkg/orders/domain/orders"
	"github.com/abhishekkushwahaa/golang-microservice/pkg/payments/interfaces/intraprocess"
)

type IntraprocessService struct {
	orders chan<- intraprocess.OrderToProcess
}

func NewIntraprocessService(ordersChannel chan<- intraprocess.OrderToProcess) IntraprocessService {
	return IntraprocessService{ordersChannel}
}

func (i IntraprocessService) InitializeOrderPayment(id orders.ID, price price.Price) error {
	i.orders <- intraprocess.OrderToProcess{string(id), price}
	return nil
}
