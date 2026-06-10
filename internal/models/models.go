package models

import "time"

type User struct {
	UserID      string `gorm:"primaryKey;type:varchar(20)" json:"user_id"`
	FullName    string `gorm:"type:varchar(100);not null" json:"full_name"`
	Email       string `gorm:"type:varchar(100);unique;not null" json:"email"`
	PhoneNumber string `gorm:"type:varchar(20)" json:"phone_number"`
	Password    string `gorm:"type:varchar(255);not null" json:"-"`
}

type Bus struct {
	BusID       string `gorm:"primaryKey;type:varchar(20)" json:"bus_id"`
	BusCode     string `gorm:"type:varchar(20);not null" json:"bus_code"`
	ChassisType string `gorm:"type:varchar(50)" json:"chassis_type"`
	BodyType    string `gorm:"type:varchar(50)" json:"body_type"`
	PlateNumber string `gorm:"type:varchar(20);unique" json:"plate_number"`
	BusClass    string `gorm:"type:varchar(50)" json:"bus_class"`
	BusStatus   string `gorm:"type:varchar(20)" json:"bus_status"`
}

type Seat struct {
	SeatID     string `gorm:"primaryKey;type:varchar(20)" json:"seat_id"`
	SeatNumber string `gorm:"type:varchar(10);not null" json:"seat_number"`
	SeatStatus string `gorm:"type:varchar(20)" json:"seat_status"`
	BusID      string `gorm:"type:varchar(20)" json:"bus_id"`
}

type Route struct {
	RouteID     string `gorm:"primaryKey;type:varchar(20)" json:"route_id"`
	Origin      string `gorm:"type:varchar(100);not null" json:"origin"`
	Destination string `gorm:"type:varchar(100);not null" json:"destination"`
}

type Schedule struct {
	ScheduleID    string    `gorm:"primaryKey;type:varchar(20)" json:"schedule_id"`
	DepartureDate time.Time `gorm:"type:date;not null" json:"departure_date"`
	DepartureTime string    `gorm:"type:time" json:"departure_time"`
	ArrivalTime   string    `gorm:"type:time" json:"arrival_time"`
	RouteID       string    `gorm:"type:varchar(20)" json:"route_id"`
	BusID         string    `gorm:"type:varchar(20)" json:"bus_id"`
	Route         Route     `gorm:"foreignKey:RouteID" json:"route,omitempty"`
	Bus           Bus       `gorm:"foreignKey:BusID" json:"bus,omitempty"`
}

type Booking struct {
	BookingID     string    `gorm:"primaryKey;type:varchar(20)" json:"booking_id"`
	BookingDate   time.Time `gorm:"type:timestamp;not null" json:"booking_date"`
	BookingStatus string    `gorm:"type:varchar(20)" json:"booking_status"`
	TotalPrice    float64   `gorm:"type:decimal(10,2)" json:"total_price"`
	UserID        string    `gorm:"type:varchar(20)" json:"user_id"`
	ScheduleID    string    `gorm:"type:varchar(20)" json:"schedule_id"`
	User          User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Schedule      Schedule  `gorm:"foreignKey:ScheduleID" json:"schedule,omitempty"`
	Payment       Payment   `gorm:"foreignKey:BookingID" json:"payment,omitempty"`
	Ticket        Ticket    `gorm:"foreignKey:BookingID" json:"ticket,omitempty"`
	SeatID        string    `gorm:"type:varchar(20)" json:"seat_id"`
}

type Payment struct {
	PaymentID     string    `gorm:"primaryKey;type:varchar(20)" json:"payment_id"`
	PaymentMethod string    `gorm:"type:varchar(50)" json:"payment_method"`
	PaymentStatus string    `gorm:"type:varchar(20)" json:"payment_status"`
	PaymentDate   time.Time `gorm:"type:timestamp" json:"payment_date"`
	BookingID     string    `gorm:"type:varchar(20);unique" json:"booking_id"`
}

type Ticket struct {
	TicketID      string    `gorm:"primaryKey;type:varchar(20)" json:"ticket_id"`
	TicketCode    string    `gorm:"type:varchar(50);unique;not null" json:"ticket_code"`
	DepartureDate time.Time `gorm:"type:date;not null" json:"departure_date"`
	QRCode        string    `gorm:"type:text" json:"qr_code"`
	BookingID     string    `gorm:"type:varchar(20);unique" json:"booking_id"`
}
