package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ContextKey string
type LoginCredentials struct {
	Email    string `json:"email" example:"test@mail.com"`
	Password string `json:"password" example:"abc123"`
}

const (
	PhotographerNotVerifiedStatus = "NOT_VERIFIED"
	PhotographerPendingStatus     = "PENDING"
	PhotographerVerifiedStatus    = "VERIFIED"
	PhotographerRejectedStatus    = "REJECTED"
)

type User struct {
	bun.BaseModel      `bun:"table:users,alias:u"`
	Id                 uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	Name               string    `bun:"name,type:varchar" json:"name"`
	Email              string    `bun:"email,type:varchar" json:"email"`
	Provider           *string   `bun:"provider,type:varchar" json:"provider"`
	Password           *string   `bun:"password,type:varchar" json:"-"`
	LoggedOut          bool      `bun:"logged_out,type:boolean" json:"logged_out"`
	ProfilePictureKey  *string   `bun:"profile_picture_key,type:varchar" json:"profile_picture_key"`
	VerificationStatus string    `bun:"verification_status,type:varchar" json:"verification_status"`
}

type UserInput struct {
	Name     string  `json:"name" example:"test"`
	Email    string  `json:"email" example:"test@mail.com"`
	Password *string `json:"password" example:"root"`
}

type Administrator struct {
	bun.BaseModel `bun:"table:administrators,alias:admin"`
	Id            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	Email         string    `bun:"email,type:varchar" json:"email"`
	Password      string    `bun:"password,type:varchar" json:"password"`
	LoggedOut     bool      `bun:"logged_out,type:boolean" json:"logged_out"`
}

type Gallery struct {
	bun.BaseModel  `bun:"table:galleries,alias:galleries"`
	Id             uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	PhotographerId uuid.UUID `bun:"photographer_id,type:uuid" json:"photographer_id"`
	Location       string    `bun:"location,type:varchar" json:"location"`
	Name           string    `bun:"name,type:varchar" json:"name"`
	Price          int       `bun:"price,type:integer" json:"price"`
}

type GalleryInput struct {
	Name     *string `bun:"name,type:varchar" json:"name"`
	Location *string `bun:"name,type:varchar" json:"location"`
	Price    *int    `bun:"price,type:integer" json:"price"`
}

const (
	BookingPaidStatus                  = "USER_PAID"
	BookingCancelledStatus             = "CANCELLED"
	BookingCustomerReqCancelStatus     = "C_REQ_CANCEL"
	BookingPhotographerReqCancelStatus = "P_REQ_CANCEL"
	BookingCompletedStatus             = "COMPLETED"
	BookingPaidOutStatus               = "PAID_OUT"
)

type BookingProposal struct {
	GalleryId uuid.UUID `bun:"gallery_id,type:uuid" json:"gallery_id"`
	StartTime time.Time `bun:"start_time,type:timestamptz" json:"start_time"`
	EndTime   time.Time `bun:"end_time,type:timestamptz" json:"end_time"`
}

type Booking struct {
	bun.BaseModel `bun:"table:bookings,alias:bookings"`
	Id            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	CustomerId    uuid.UUID `bun:"customer_id,type:uuid" json:"customer_id"`
	GalleryId     uuid.UUID `bun:"gallery_id,type:uuid" json:"gallery_id"`
	StartTime     time.Time `bun:"start_time,type:timestamptz" json:"start_time"`
	EndTime       time.Time `bun:"end_time,type:timestamptz" json:"end_time"`
	Status        string    `bun:"status,type:varchar" json:"status"`
	CreatedAt     time.Time `bun:"created_at,type:timestamptz,default:now()" json:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at,type:timestamptz,default:now()" json:"updated_at"`
}

type SearchFilter struct {
	PhotographerId *uuid.UUID `form:"photographer_id"`
	Location       *string    `form:"location"`
	MinPrice       *int       `form:"min_price"`
	MaxPrice       *int       `form:"max_price"`
}

type Room struct {
	bun.BaseModel `bun:"table:rooms,alias:rooms"`
	Id            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	CreatedAt     time.Time  `bun:"created_at,type:timestamptz,default:now()" json:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at,type:timestamptz,default:now()" json:"updated_at"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete,nullzero,type:timestamptz" json:"deleted_at"`
}

type UserRoomLookup struct {
	bun.BaseModel `bun:"table:user_room_lookup,alias:urlookup"`
	Id            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	UserId        uuid.UUID  `bun:"user_id,type:uuid" json:"user_id"`
	RoomId        uuid.UUID  `bun:"room_id,type:uuid" json:"room_id"`
	CreatedAt     time.Time  `bun:"created_at,type:timestamptz,default:now()" json:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at,type:timestamptz,default:now()" json:"updated_at"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete,nullzero,type:timestamptz" json:"deleted_at"`
}

type Conversation struct {
	bun.BaseModel `bun:"table:conversations,alias:convs"`
	Id            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	Text          string     `bun:"text,type:varchar" json:"text"`
	UserId        uuid.UUID  `bun:"user_id,type:uuid" json:"user_id"`
	RoomId        uuid.UUID  `bun:"room_id,type:uuid" json:"room_id"`
	CreatedAt     time.Time  `bun:"created_at,type:timestamptz,default:now()" json:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at,type:timestamptz,default:now()" json:"updated_at"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete,nullzero,type:timestamptz" json:"deleted_at"`
}

type RoomMemberInput struct {
	MemberIds []uuid.UUID `binding:"required" json:"member_ids"`
}

type Photo struct {
	bun.BaseModel `bun:"table:photos,alias:photos"`
	Id            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	GalleryId     uuid.UUID `bun:"gallery_id,type:uuid" json:"gallery_id"`
	PhotoKey      string    `bun:"photo_key,type:varchar" json:"photo_key"`
}
