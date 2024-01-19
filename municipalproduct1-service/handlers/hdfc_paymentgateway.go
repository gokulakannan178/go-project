package handlers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

func pad(input []byte) []byte {
	padding := aes.BlockSize - (len(input) % aes.BlockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(input, padText...)
}

func encrypt(plainText []byte, workingKey string) string {
	iv := []byte("\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f")
	plainText = pad(plainText)
	hash := md5.New()
	hash.Write([]byte(workingKey))
	key := hash.Sum(nil)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, len(plainText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plainText)
	return hex.EncodeToString(ciphertext)
}

func (h *Handler) GetReq(w http.ResponseWriter, r *http.Request) {
	data := `<html>
	<head>
	</head>
	<body>
		<form method="POST" name="customerData" action="/api/hdfc/pg">
			<table width="40%" height="100" border='1' align="center">
				<caption>
					<font size="4" color="blue"><b>Integration Kit</b></font>
				</caption>
			</table>
			<table width="40%" height="100" border='1' align="center">
				<tr>
					<td>Parameter Name:</td>
					<td>Parameter Value:</td>
				</tr>
				<tr>
					<td colspan="2">Compulsory information</td>
				</tr>
				<tr>
					<td>Merchant Id</td>
					<td><input type="text" name="merchant_id" id="merchant_id" value="2153776" /> </td>
				</tr>
				<tr>
					<td>Order Id</td>
					<td><input type="text" name="order_id" value="21234124" /></td>
				</tr>
				<tr>
					<td>Currency</td>
					<td><input type="text" name="currency" value="INR" /></td>
				</tr>
				<tr>
					<td>Amount</td>
					<td><input type="text" name="amount" value="1.00" /></td>
				</tr>
				<tr>
					<td>Redirect URL</td>
					<td><input type="text" name="redirect_url"
						value="https://bhagalpur.biharmunicipal.com/api/online/payment" />
					</td>
				</tr>
				<tr>
					<td>Cancel URL</td>
					<td><input type="text" name="cancel_url"
						value="https://bhagalpur.biharmunicipal.com/api/online/payment/failed" />
					</td>
				</tr>
				<tr>
					<td>Language</td>
					<td><input type="text" name="language" id="language" value="EN" /></td>
				</tr>
				
					<td></td>
					<td><INPUT TYPE="submit" value="Checkout"></td>
				</tr>
			</table>
		</form>
	</body>
	</html>
	`
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(data))
}

// SaveHDFCPaymentGateway
func (h *Handler) SaveHDFCPaymentGateway(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	hdfc := new(models.HDFCPaymentGateway)
	err := json.NewDecoder(r.Body).Decode(&hdfc)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveHDFCPaymentGateway(ctx, hdfc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["hdfcPaymentGateway"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleMerchantIDHDFCPaymentGateway : ""
func (h *Handler) GetSingleMerchantIDHDFCPaymentGateway(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	merchantId, err := h.Service.GetSingleMerchantIDHDFCPaymentGateway(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["merchantId"] = merchantId
	response.With200V2(w, "Success", m, platform)
}

// GetSingleDefaultHDFCPaymentGateway : ""
func (h *Handler) GetSingleDefaultHDFCPaymentGateway(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	res, err := h.Service.GetSingleDefaultHDFCPaymentGateway(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["hdfcPayment"] = res
	response.With200V2(w, "Success", m, platform)
}
