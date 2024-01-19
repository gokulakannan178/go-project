package handlers

import (
	"fmt"
	"municipalproduct1-service/app"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// OnlinePayment : ""
func (h *Handler) GetOnlinePaymentTL(w http.ResponseWriter, r *http.Request) {
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

	payment := new(models.MakeTradeLicensePaymentReq)
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

		pymt, err := h.Service.GetSingleTradeLicensePayment(ctx, payment.TnxID)
		if err != nil {
			globalErr = "Wrong Payment - " + err.Error()
			goto Failure
		}
		if pymt.Demand.Total.Total != key {
			globalErr = "Wrong Payment - Payment is Vulnerable"
			goto Failure
		}
		payment.Details.Amount = key
		payment.Details.MadeAt = new(models.TradeLicenceMadeAt)
		payment.Details.MadeAt.At = "Online"
		payment.Details.MOP.Mode = constants.MOPNETBANKING
		payment.Details.MOP.No = data["bank_ref_no"]
		payment.Details.MOP.Vendor = "HDFC_CC_AVENUE"
		payment.Details.MOP.VendorType = "PaymentGateway"
		payment.Details.MOP.CardRNet.TxnID = data["bank_ref_no"]
		payment.Details.MOP.CardRNet.TrackingID = data["tracking_id"]
		payment.Details.MOP.CardRNet.Vendor = "HDFC_CC_AVENUE"
		payment.Details.MOP.CardRNet.VendorType = "PaymentGateway"
		payment.Details.MOP.CardRNet.CardType = data["payment_mode"]
		payment.Details.MOP.CardRNet.CardName = data["card_name"]
		payment.Details.Collector.By = constants.SYSTEM
		payment.Details.Collector.ByType = constants.SYSTEM
		payment.Details.MOP.VendorInfo.HDFC = *hpgcpsr
		propertyId, err := h.Service.MakeTradeLicensePayment(ctx, payment)
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

        <td><a href="` + resPD.UIURL + `#/consumerpropertyList"><button type= "button" value= "Hi" id= "btnOK" onclick= "ok.performClick(this.value); " > Go back to Properties </button></a>
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
		  <td><a href="` + resPD.UIURL + `#/consumerpropertyList"><button type= "button" value= "Hi" id= "btnOK" onclick= "ok.performClick(this.value); " > Go back to Properties </button></a>
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
		  <td><a href="` + resPD.UIURL + `#/consumerpropertyList"><button type= "button" value= "Hi" id= "btnOK" onclick= "ok.performClick(this.value); " > Go back to Properties </button></a>
		  </td>
		</tr>
		  
		</table>
	</div>
	
		</body>
		</html>`
	fmt.Fprintf(w, res)
}

func (h *Handler) GetFailedOnlinePaymentTL(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	//Params := r.URL.Query()
	fmt.Println(platform)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	// filter := make(map[string]interface{})
	// for i, v := range Params {
	// 	filter[i] = v
	// 	fmt.Println(i, "value is", v)

	// }
	// fmt.Println(filter)
	// encResp := r.FormValue("encResp")
	// workingKey := "2775EA71B1D1F115FF568EC1BF6D4653"
	// decryptedData := decrypt(encResp, workingKey)
	// fmt.Println(decryptedData)
	// // Split the string into key-value pairs
	// pairs := strings.Split(decryptedData, "&")
	// // Create a new map
	// data := make(map[string]string)

	// for _, pair := range pairs {
	// 	parts := strings.Split(pair, "=")
	// 	if len(parts) >= 2 {
	// 		key := parts[0]
	// 		value := parts[1]
	// 		data[key] = value
	// 		fmt.Println("key ====>", key)
	// 		fmt.Println("value ====>", value)
	// 	}
	// }

	resPD, err := h.Service.GetSingleDefaultProductConfiguration(ctx)
	if err != nil {
		response.With500mV2(w, "error in getting product configuration - ", platform)
		return
	}
	fmt.Println("resPD ======>", resPD)

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	res := `<html>
	<body>
	

	<div>
    <table style="margin: auto;">
      <tr>
        <td><img src="` + resPD.APIURL + resPD.Logo + `"></td>
      </tr>
      <tr style="text-align: center; font-weight: bold; font-size: 24px;">
        <td>Payment Failed</td>
      </tr>
      <tr style="text-align: center; font-size: 20px;">
        <td><a href="` + resPD.UIURL + `#/consumerpropertyList">Go to Home Page</a></td>
      </tr>
    </table>
</div>

	</body>
	</html>`
	fmt.Fprintf(w, res)

}

func (h *Handler) GetOnlinePaymentTLtest(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	//Params := r.URL.Query()
	fmt.Println(platform)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	// filter := make(map[string]interface{})
	// for i, v := range Params {
	// 	filter[i] = v
	// 	fmt.Println(i, "value is", v)

	// }
	// fmt.Println(filter)
	// encResp := r.FormValue("encResp")
	// workingKey := "2775EA71B1D1F115FF568EC1BF6D4653"
	// decryptedData := decrypt(encResp, workingKey)
	// fmt.Println(decryptedData)
	// // Split the string into key-value pairs
	// pairs := strings.Split(decryptedData, "&")
	// // Create a new map
	// data := make(map[string]string)

	// for _, pair := range pairs {
	// 	parts := strings.Split(pair, "=")
	// 	if len(parts) >= 2 {
	// 		key := parts[0]
	// 		value := parts[1]
	// 		data[key] = value
	// 		fmt.Println("key ====>", key)
	// 		fmt.Println("value ====>", value)
	// 	}
	// }
	resPayment, err := h.Service.GetSingleDefaultHDFCPaymentGateway(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	resPD, err := h.Service.GetSingleDefaultProductConfiguration(ctx)
	if err != nil {
		response.With500mV2(w, "error in getting product configuration - ", platform)
		return
	}
	fmt.Println("resPD ======>", resPD)
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	t := time.Now().Unix()
	res := `<html>

	<form method="POST" ngNoForm name="customerData" action="https://bhagalpur.biharmunicipal.com/api/hdfc/pg">
                 
                        
                       
                           
                           <input type="text" name="merchant_id" id="merchant_id" value="` + resPayment.MerchantID + `" />  
                        <input type="text" name="order_id" value="` + fmt.Sprintf("%dd", t) + `" />
                         <input type="text" name="merchant_param1" value="` + fmt.Sprintf("%dd", t) + `" />
                           <input type="text" name="merchant_param2" value="Bala" />
                            <input type="text" name="merchant_param3" value="Rupees one" />
                                <input type="text" name="currency" value="INR" />
                   
                            
                           <input type="text" name="amount" value="1" />
                             <input type="text" name="redirect_url"
                                value="https://bhagalpur.biharmunicipal.com/api/tradelicense/online/payment" />
                           <input type="text" name="cancel_url"
                                value="https://bhagalpur.biharmunicipal.com/api/tradelicense/online/payment/failed" />
                           <input type="text" name="language" id="language" value="EN" />
         
                       
                      
                            <button type="submit" class="btn btn-primary">Pay via HDFC</button>
                    
                        <br>
                    </table>
                </form>`

	// 	w.Header().Add("Content-Type", "text/html")
	// 	w.WriteHeader(200)
	// 	res := <html>
	// 	<body>

	// 	<div>
	//     <table style="margin: auto;">
	//       <tr>
	//         <td><img src="` + resPD.APIURL + resPD.Logo + `"></td>
	//       </tr>
	//       <tr style="text-align: center; font-weight: bold; font-size: 24px;">
	//         <td>Payment Failed</td>
	//       </tr>
	//       <tr style="text-align: center; font-size: 20px;">
	//         <td><a href="` + resPD.UIURL + `#/consumerpropertyList">Go to Home Page</a></td>
	//       </tr>
	//     </table>
	// </div>

	// 	</body>
	// 	</html>`
	fmt.Fprintf(w, res)

}
