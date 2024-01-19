package handlers

import (
	"fmt"
	"io/ioutil"
	"municipalproduct1-service/app"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) PostReq(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	payment := new(models.RefHDFCPaymentGateway)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	payment, err := h.Service.GetSingleDefaultHDFCPaymentGateway(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	// workingKey := "2775EA71B1D1F115FF568EC1BF6D4653" // Put in the 32-Bit key shared by CCAvenues.
	// accessCode := "AVNV28KC60AW66VNWA"               // Put in the Access Code shared by CCAvenues.
	formbody := ""
	/*`{

		"merchant_id":2153776,
		"order_id":"1",
		"currency":"INR",
		"amount":200,
		"redirect_url":"https://bhagalpur.biharmunicipal.com/api/online/payment",
		"cancel_url":"https://bhagalpur.biharmunicipal.com/api/online/payment",
		"payment_option":"OPTDBCRD",
		"card_type":"CRDC",
		"card_name":"TEST",
		"data_accept":"Y",
		"card_number":"4012001037141112",
		"expiry_month":12,
		"expiry_year":2023,
		"cvv_number":123,
		"issuing_bank":"HDFC",
		"mobile_no":7299424027



	}`*/
	// payload := &bytes.Buffer{}
	// writer := multipart.NewWriter(payload)
	// _ = writer.WriteField("merchant_id", "2153776")
	// _ = writer.WriteField("order_id", "1")
	// _ = writer.WriteField("currency", "INR")
	// _ = writer.WriteField("amount", "200")
	// _ = writer.WriteField("redirect_url", "https://bhagalpur.biharmunicipal.com/api/online/payment")
	// _ = writer.WriteField("cancel_url", "https://bhagalpur.biharmunicipal.com/api/online/payment")
	// _ = writer.WriteField("payment_option", "OPTDBCRD")
	// _ = writer.WriteField("card_type", "CRDC")
	// _ = writer.WriteField("card_name", "TEST")
	// _ = writer.WriteField("data_accept", "Y")
	// _ = writer.WriteField("card_number", "4012001037141112")
	// _ = writer.WriteField("expiry_month", "12")
	// _ = writer.WriteField("expiry_year", "2023")
	// _ = writer.WriteField("cvv_number", "123")
	// _ = writer.WriteField("issuing_bank", "HDFC")

	// formbody = payload.String()

	// Generate MD5 hash for the key and then convert to base64 string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// altBody := "merchant_id=2153776&order_id=SM1003ORD&amount=1.00&currency=INR&redirect_url=https://bhagalpur.biharmunicipal.com/api/online/payment&cancel_url=https://bhagalpur.biharmunicipal.com/api/online/payment/failed&language=EN"
	encRequest := encrypt(body, payment.WorkingKey)
	// fmt.Println(string(altBody))
	// encRequest, err := encryptAES([]byte(workingKey), []byte(altBody))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// encRequest = "cba0a19c5da2cc8d5d202340d6593f6f713f0b76056ec98ea214e002e3ed8707180c0462525d319d8a525ab2605dfe2badf65451f5e041cef514e4a2198d5be23031939c908e9dd8e48052157a8ad83344b85381359a83fd1544d5f0dc0ba4b17b5e1273e2312e3496cad5de6e366c4de0cf2bb3a6a64d9f62baeaf6c204776a5c43f795b946fb6b57c5847a0cbeebdf6cdf1539151da2c0f0f2c7dbe5305a5a46b1e5563109de79f9f69eb8feb3c1ff444bbd1fd75f196509a54ee8e535a64342c07e7ffc5da611a86898b093cad56630a2ad93f261c8f193ce7096390801d4"

	// https://test.ccavenue.com/transaction/transaction.do?command=initiateTransaction&nbsp
	formbody = "<html><body><form id=\"nonseamless\" method=\"post\" name=\"redirect\" action=\"" + payment.BaseURL + "\"/> <input type=\"hidden\" id=\"encRequest\" name=\"encRequest\" value=\"" + encRequest + "\"><input type=\"hidden\" name=\"access_code\" id=\"access_code\" value=\"" + payment.AccessCode + "\"><script language=\"javascript\">document.redirect.submit();</script></form></body></html>"

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(formbody))
}

// OnlinePayment : ""
func (h *Handler) GetOnlinePayment(w http.ResponseWriter, r *http.Request) {
	globalErr := ""

	platform := r.URL.Query().Get("platform")
	Params := r.URL.Query()
	fmt.Println(platform)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := make(map[string]interface{})
	for i, v := range Params {
		filter[i] = v
		fmt.Println(i, "value is", v)

	}
	fmt.Println(filter)
	encResp := r.FormValue("encResp")
	resPayment, err := h.Service.GetSingleDefaultHDFCPaymentGateway(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	// workingKey := "2775EA71B1D1F115FF568EC1BF6D4653"
	workingKey := resPayment.WorkingKey
	decryptedData := decrypt(encResp, workingKey)
	fmt.Println(decryptedData)
	// Split the string into key-value pairs
	pairs := strings.Split(decryptedData, "&")
	// Create a new map
	data := make(map[string]string)

	for _, pair := range pairs {
		parts := strings.Split(pair, "=")
		if len(parts) >= 2 {
			key := parts[0]
			value := parts[1]
			data[key] = value
		}
	}

	payment := new(models.PropertyMakePayment)
	resPD, err := h.Service.GetSingleDefaultProductConfiguration(ctx)
	if err != nil {
		response.With500mV2(w, "error in getting product configuration - ", platform)
		return
	}
	fmt.Println("resPD ======>", resPD)
	payment.TnxID = data["merchant_param1"]
	payment.Details.PayeeName = data["merchant_param2"]
	payment.Details.AmountInWords = data["merchant_param3"]
	fmt.Println("payment.TnxID,payment.Details.PayeeName,payment.Details.AmountInWords", payment.TnxID, payment.Details.PayeeName, payment.Details.AmountInWords)
	if data["order_status"] == "Success" {
		cps := new(models.HDFCPaymentGatewayCheckPaymentStatus)
		cps.OrderNo = data["order_id"]
		hpgcpsr, statusAPIErr := h.Service.CheckPaymentStatus(ctx, cps)
		if statusAPIErr != nil {
			response.With500mV2(w, "status api failed - "+statusAPIErr.Error(), platform)
			return
		}
		if hpgcpsr == nil {
			response.With500mV2(w, "status api failed - hpgcpsr is nil", platform)
			return
		}
		key, err1 := strconv.ParseFloat(data["amount"], 32)
		if err1 != nil {
			return
		}

		pymt, err := h.Service.GetSinglePropertyPaymentTxtID(ctx, payment.TnxID)
		if err != nil {
			globalErr = "Wrong Payment - " + err.Error()
			goto Failure
		}
		if pymt.Demand.TotalTax != key {
			globalErr = "Wrong Payment - Payment is Vulnerable"
			goto Failure
		}
		payment.Details.Amount = key
		// payment.Details.AmountInWords = ""
		// payment.Details.PayeeName = ""
		// payment.Details.MadeAt = nil
		payment.Details.MadeAt = new(models.PropertyPaymentDetailsMadeAt)
		payment.Details.MadeAt.At = "Online"
		payment.Details.MOP.Mode = constants.MOPNETBANKING
		payment.Details.MOP.Cheque = nil
		payment.Details.MOP.DD = nil
		payment.Details.MOP.CardRNet = new(models.CardRNet)
		payment.Details.MOP.CardRNet.TxnID = data["bank_ref_no"]
		payment.Details.MOP.CardRNet.TrackingID = data["tracking_id"]
		payment.Details.MOP.CardRNet.Vendor = "HDFC_CC_AVENUE"
		// payment.Details.MOP.PropertyPaymentCardRNet.Vendor = "HDFC_CC_AVENUE"
		payment.Details.MOP.CardRNet.VendorType = "PaymentGateway"
		payment.Details.MOP.CardRNet.CardType = data["payment_mode"]
		payment.Details.MOP.CardRNet.CardName = data["card_name"]

		payment.Details.MOP.PropertyPaymentCardRNet.TxnID = data["bank_ref_no"]
		payment.Details.MOP.PropertyPaymentCardRNet.TrackingID = data["tracking_id"]
		payment.Details.MOP.PropertyPaymentCardRNet.Vendor = "HDFC_CC_AVENUE"
		// payment.Details.MOP.PropertyPaymentCardRNet.Vendor = "HDFC_CC_AVENUE"
		payment.Details.MOP.PropertyPaymentCardRNet.VendorType = "PaymentGateway"
		payment.Details.MOP.PropertyPaymentCardRNet.CardType = data["payment_mode"]
		payment.Details.MOP.PropertyPaymentCardRNet.CardName = data["card_name"]
		payment.Details.Collector.ID = constants.SYSTEM
		payment.Details.Collector.Type = constants.SYSTEM
		payment.Details.MOP.VendorInfo.HDFC = *hpgcpsr
		propertyId, err := h.Service.PropertyMakePayment(ctx, payment)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				response.With500mV2(w, "failed no data for this id", platform)
				return
			}
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		if propertyId == "" {
			response.With500mV2(w, "failed to get property id- ", platform)
			return
		}

		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// fmt.Print(string(body))

		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(200)
		res := `<html>
	<body>
	
	<script>
	sessionStorage.setItem('consLogin', 'true');
	</script>
	<div>
    <table style="margin: auto;">
      <tr>
        <td><img src="` + resPD.APIURL + resPD.Logo + `" style= "height: 120px; margin-left: 90px;"> </td>
      </tr>
      <tr style="text-align: center; font-weight: bold; font-size: 24px;">
        <td>Payment Received Successfully</td>
      </tr>
      <tr style="text-align: center; font-size: 20px;">
        <td>Order ID - ` + data["order_id"] + `</td>
      </tr>
      <tr style="text-align: center; font-size: 20px;">
        <td>Amount - ` + data["amount"] + `</td>
      </tr>
      <tr style="text-align: center; font-size: 20px;">

        <td><a href="` + resPD.UIURL + `#/consumersaflistview"><button type= "button" value= "Hi" id= "btnOK" onclick= "ok.performClick(this.value); " > Go back to Properties </button></a>
		</td>
      </tr>
	  
    </table>
</div>

	

	</body>
	</html>`
		fmt.Fprintf(w, res)
		return

	} else if data["order_status"] == "Failure" {
		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(200)
		res := `<html>
		<body>
		
	
		<div>
		<table style="margin: auto;">
		  <tr>
			<td><img src="` + resPD.APIURL + resPD.Logo + `" style= "height: 120px; margin-left: 23px;"></td>
		  </tr>
		  <tr style="text-align: center; font-weight: bold; font-size: 24px;">
			<td>Payment Failed</td>
		  </tr>
		  <tr style="text-align: center; font-size: 20px;">
		  <td><a href="` + resPD.UIURL + `#/consumersaflistview"><button type= "button" value= "Hi" id= "btnOK" onclick= "ok.performClick(this.value); " > Go back to Properties </button></a>
		  </td>
		</tr>
		  
		</table>
	</div>
	
		</body>
		</html>`
		fmt.Fprintf(w, res)
		return
	}
Failure:
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	res := `<html>
		<body>
		
	
		<div>
		<table style="margin: auto;">
		  <tr>
			<td><img src="` + resPD.APIURL + resPD.Logo + `" style= "height: 120px; margin-left: 23px;"></td>
		  </tr>
		  <tr style="text-align: center; font-weight: bold; font-size: 24px;">
			<td>Payment Failed due to ` + globalErr + `</td>
		  </tr>
		  <tr style="text-align: center; font-size: 20px;">
		  <td><a href="` + resPD.UIURL + `#/consumersaflistview"><button type= "button" value= "Hi" id= "btnOK" onclick= "ok.performClick(this.value); " > Go back to Properties </button></a>
		  </td>
		</tr>
		  
		</table>
	</div>
	
		</body>
		</html>`
	fmt.Fprintf(w, res)
}
