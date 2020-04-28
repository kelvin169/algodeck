package main
// declaration of support documentation 
import (
	"log"
	"github.com/AndroidStudyOpenSource/mpesa-api-go"
)

const (
	appKey    = ""
	appSecret = ""
)
//the func main is where the core part of the code is added. 
func main() {

	svc, err := mpesa.New(appKey, appSecret, mpesa.SANDBOX)
	if err != nil {
		panic(err)
	}

	res, err := svc.B2CRequest(mpesa.B2C{
		InitiatorName:      "",
		SecurityCredential: "",
		CommandID:          "",
		Amount:             "",
		PartyA:             "",
		PartyB:             "",
		Remarks:            "",
		QueueTimeOutURL:    "",
		ResultURL:          "",
		Occassion:          "",
	})

	if err != nil {
		log.Println(err)
	}
	log.Println(res)

}