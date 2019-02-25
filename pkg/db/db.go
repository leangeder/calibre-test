package db

type DB interface {
	Connect()
	Disconnect()
	Get()
	Delete()
	Add()
	Update()
	data [
}
