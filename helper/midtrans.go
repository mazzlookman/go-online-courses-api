package helper

import (
	"fmt"
	"github.com/veritrans/go-midtrans"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"os"
	"time"
)

func GetPaymentUrl(transaction domain.Transaction, user web.UserResponse) []string {
	midClient := midtrans.NewClient()
	midClient.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	midClient.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midClient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midClient,
	}

	transactionId := fmt.Sprintf("%d-%d", transaction.Id, time.Now().Unix())
	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transactionId,
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
	}

	snapResponse, err := snapGateway.GetToken(snapReq)
	PanicIfError(err)

	return []string{snapResponse.RedirectURL, transactionId}
}
