package service

import (
	"Go-Mall/app/payment/biz/dal/mysql"
	"Go-Mall/app/payment/biz/model"
	payment "Go-Mall/rpc_gen/kitex_gen/payment"
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"time"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// TODO 检查信用卡信息 生产环境时开启
	//card := creditcard.Card{
	//	Number: req.CreditCard.CreditCardNumber,
	//	Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
	//	Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
	//	Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	//}
	//err = card.Validate(true)
	//if err != nil {
	//	return nil, kerrors.NewGRPCBizStatusError(4004001, err.Error())
	//}

	// 生成transactionId
	transactionId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005001, err.Error())
	}

	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: transactionId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005002, err.Error())
	}
	return &payment.ChargeResp{
		TransactionId: transactionId.String(),
	}, nil
}
