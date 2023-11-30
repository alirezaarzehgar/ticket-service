package util

const (
	ALERT_SUCCESS               = "Success"
	ALERT_NOT_FOUND             = "Not found"
	ALERT_TOKEN_NOT_FOUND       = "Your JWT token not found"
	ALERT_USER_CONFLICT         = "Someone registered with this email"
	ALERT_ORG_CONFLICT          = "Organization already created with this name"
	ALERT_CONFLICT              = "Conflict"
	ALERT_INVALID_EMAIL         = "Invalid email address"
	ALERT_INSECURE_PASSWORD     = "Insecure password"
	ALERT_LOGIN_UNAUTHORIZED    = "Your password or email address is invalid"
	ALERT_ORG_EDIT_UNAUTHORIZED = "You are not member of this organization"
	ALERT_BAD_REQUEST           = "Bad request"
	ALERT_REQUIRED_FIELDS       = "Your request is in complete. You should fill all required fields"
	ALERT_INVALID_TOKEN         = "Invalid JWT token"
	ALERT_BAD_FILE              = "Invalid file format"
	ALERT_STORY_WRONG           = "Wrong story field"
	ALERT_DUP_ESTORY            = "You can't create more that one explore story"
	ALERT_USER_ONLY             = "You should be user to access this page"
	ALERT_SUPER_ADMIN_ONLY      = "You should be super admin to access this page"
	ALERT_ADMIN_ONLY            = "You should be admin or super admin to access this page"
	ALERT_INTERNAL              = "Internal server error"
)

// Response model
type Response struct {
	// say status of response.
	// false means to failure and true means to success.
	Status bool `json:"status" example:"true"`

	// message for client.
	Alert string `json:"alert" example:"suitable alert"`

	// will be null in case of status=false.
	// client needed data
	Data any `json:"data"`
}

// ReponseError model
type ResponseError struct {
	// say status of response.
	// false means to failure and true means to success.
	Status bool `json:"status" example:"false"`

	// message for client.
	Alert string `json:"alert" example:"some error"`

	// Will be null.
	Data int `json:"data" example:"0"`
}
