package constants

//SALE TRANSPORT status
const (
	PAYMENTSTATUSINIT                string = "Init"
	PAYMENTSTATUSACTIVE              string = "Active"
	PAYMENTSTATUSDEACTIVE            string = "Disabled"
	PAYMENTSTATUSDELETED             string = "Deleted"
	PAYMENTSTATUSPARTIAL             string = "Partial"
	PAYMENTSTATUSBLOCKED             string = "Blocked"
	PAYMENTSTATUSPENDING             string = "Pending"
	PAYMENTSTATUSDONE                string = "Done"
	PAYMENTSTATUSVERIFICATIONPENDING string = "VerificationPending"

	SALEPAYMENTSTATUSPENDING   string = "Pending"
	SALEPAYMENTSTATUSCOMPLETED string = "Completed"

	PAYMENTTYPECHEQUE     string = "Cheque"
	PAYMENTTYPECASH       string = "Cash"
	PAYMENTTYPENETBANKING string = "NetBanking"
	PAYMENTTYPECREDIT     string = "Credit"
	PAYMENTTYPEDD         string = "DD"
)

// Sale status
const (
	SALESTATUSACTIVE   = "Active"
	SALESTATUSDEACTIVE = "Disabled"
	SALESTATUSDELETED  = "Deleted"
)
