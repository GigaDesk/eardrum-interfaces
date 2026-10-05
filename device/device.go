package device

// Device a logged in device in the system.
type Device interface {
	
	GetDeviceIMEI() string //GetDeviceIMEI returns the IMEI device ID of the device 
	GetDeviceModel() *string //GetDeviceModel returns the model of the device
}