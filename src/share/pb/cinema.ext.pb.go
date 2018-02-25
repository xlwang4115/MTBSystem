// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cinema.ext.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	cinema.ext.proto
	cms.ext.proto
	comment.ext.proto
	film.ext.proto
	order.ext.proto
	place.ext.proto
	user.ext.proto

It has these top-level messages:
	LocationCinemaReq
	LocationCinemaRsp
	Cinema
	GetCinemaMessageReq
	GetCinemaMessageRsp
	GetCinemaMessageByCidReq
	FilmMessage
	GetCinemaMessageByCidRsp
	GetMovieHallByMHIdReq
	GetMovieHallByMHIdRsp
	UserLoginReq
	UserLoginRsp
	UpdateMessageReq
	UpdateMessageRsp
	AllFilmsReq
	AllFilmsRsp
	Film
	AllUsersReq
	AllUsersRsp
	User
	AllAdminUsersReq
	AllAdminUsersRsp
	AdminUser
	AllCommentsReq
	AllCommentsRsp
	Comment
	AllOrdersReq
	AllOrdersRsp
	OrderAll
	AllAddressReq
	AllAddressRsp
	PlaceAll
	AddFilmReq
	AddFilmRsp
	UpdateFilmReq
	UpdateFilmRsp
	DeleteFilmReq
	DeleteFilmRsp
	AddAdminUserReq
	AddAdminUserRsp
	AddAddressReq
	AddAddressRsp
	UpdateAddressReq
	UpdateAddressRsp
	DeleteAddressReq
	DeleteAddressRsp
	DeleteAdminUserReq
	DeleteAdminUserRsp
	AllMovieHallReq
	MovieHall
	AllMovieHallRsp
	AddMovieHallReq
	AddMovieHallRsp
	UpdateMovieHallReq
	UpdateMovieHallRsp
	DeleteMovieHallReq
	DeleteMovieHallRsp
	AllCinemaFilmsReq
	CinemaFilm
	AllCinemaFilmsRsp
	AddCinemaFilmReq
	AddCinemaFilmRsp
	UpdateCinemaFilmReq
	UpdateCinemaFilmRsp
	DeleteCinemaFilmReq
	DeleteCinemaFilmRsp
	RegisterCinemaReq
	RegisterCinemaRsp
	AllCinemaHallReq
	AllCinemaHallRsp
	HallAddressList
	HotCommentReq
	HotCommentRsp
	CommentData
	CommentMini
	CommentPlus
	CommentRecord
	MakeCommentReq
	MakeCommentRsp
	HotPlayMoviesReq
	HotPlayMoviesRep
	Movie
	HotMovie
	FilmAllMessage
	MovieDetailReq
	MovieDetailRep
	Release
	MovieCreditsWithTypesReq
	MovieCreditsWithTypesRep
	Type
	Persons
	ImageAllReq
	ImageAllRep
	Image
	LocationMoviesReq
	LocationMoviesRep
	MovieComingNewReq
	MovieComingNewRep
	SearchRep
	SearchReq
	SearchMovie
	Genres
	Rating
	Images
	DayMovie
	GetFilmsByCidADayReq
	GetFilmsByCidADayRsp
	WantTicketReq
	WantTicketRsp
	TicketReq
	TicketRsp
	Order
	MovieTicket
	LookOrdersReq
	LookOrdersRsp
	PayOrderReq
	PayOrderRsp
	UndoOrderReq
	UndoOrderRsp
	GetOrderMessageReq
	GetOrderMessageRsp
	TicketDetail
	LookAlreadyOrdersReq
	LookAlreadyOrdersRsp
	AlreadyMovie
	OrderCommentReq
	OrderCommentRsp
	HotCitiesByCinemaReq
	HotCitiesByCinemaRep
	Place
	RegistAccountReq
	RegistAccountRsp
	LoginAccountReq
	LoginAccountRsp
	ResetAccountReq
	ResetAccountRsp
	WantScoreReq
	WantScoreRsp
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LocationCinemaReq struct {
	LocationId int64 `protobuf:"varint,1,opt,name=locationId" json:"locationId,omitempty"`
}

func (m *LocationCinemaReq) Reset()                    { *m = LocationCinemaReq{} }
func (m *LocationCinemaReq) String() string            { return proto.CompactTextString(m) }
func (*LocationCinemaReq) ProtoMessage()               {}
func (*LocationCinemaReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LocationCinemaReq) GetLocationId() int64 {
	if m != nil {
		return m.LocationId
	}
	return 0
}

type LocationCinemaRsp struct {
	Cinemas []*Cinema `protobuf:"bytes,1,rep,name=cinemas" json:"cinemas,omitempty"`
}

func (m *LocationCinemaRsp) Reset()                    { *m = LocationCinemaRsp{} }
func (m *LocationCinemaRsp) String() string            { return proto.CompactTextString(m) }
func (*LocationCinemaRsp) ProtoMessage()               {}
func (*LocationCinemaRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LocationCinemaRsp) GetCinemas() []*Cinema {
	if m != nil {
		return m.Cinemas
	}
	return nil
}

type Cinema struct {
	CinemaName     string `protobuf:"bytes,1,opt,name=cinemaName" json:"cinemaName,omitempty"`
	CinemaAddress  string `protobuf:"bytes,2,opt,name=cinemaAddress" json:"cinemaAddress,omitempty"`
	CinemaSupport  string `protobuf:"bytes,3,opt,name=cinemaSupport" json:"cinemaSupport,omitempty"`
	CinemaCard     int64  `protobuf:"varint,4,opt,name=cinemaCard" json:"cinemaCard,omitempty"`
	CinemaMinPrice int64  `protobuf:"varint,5,opt,name=cinemaMinPrice" json:"cinemaMinPrice,omitempty"`
	CinemaDiscount int64  `protobuf:"varint,6,opt,name=cinemaDiscount" json:"cinemaDiscount,omitempty"`
	CinemaId       int64  `protobuf:"varint,7,opt,name=cinemaId" json:"cinemaId,omitempty"`
}

func (m *Cinema) Reset()                    { *m = Cinema{} }
func (m *Cinema) String() string            { return proto.CompactTextString(m) }
func (*Cinema) ProtoMessage()               {}
func (*Cinema) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Cinema) GetCinemaName() string {
	if m != nil {
		return m.CinemaName
	}
	return ""
}

func (m *Cinema) GetCinemaAddress() string {
	if m != nil {
		return m.CinemaAddress
	}
	return ""
}

func (m *Cinema) GetCinemaSupport() string {
	if m != nil {
		return m.CinemaSupport
	}
	return ""
}

func (m *Cinema) GetCinemaCard() int64 {
	if m != nil {
		return m.CinemaCard
	}
	return 0
}

func (m *Cinema) GetCinemaMinPrice() int64 {
	if m != nil {
		return m.CinemaMinPrice
	}
	return 0
}

func (m *Cinema) GetCinemaDiscount() int64 {
	if m != nil {
		return m.CinemaDiscount
	}
	return 0
}

func (m *Cinema) GetCinemaId() int64 {
	if m != nil {
		return m.CinemaId
	}
	return 0
}

type GetCinemaMessageReq struct {
	MovieId    int64  `protobuf:"varint,1,opt,name=movieId" json:"movieId,omitempty"`
	LocationId int64  `protobuf:"varint,2,opt,name=locationId" json:"locationId,omitempty"`
	Day        string `protobuf:"bytes,3,opt,name=day" json:"day,omitempty"`
}

func (m *GetCinemaMessageReq) Reset()                    { *m = GetCinemaMessageReq{} }
func (m *GetCinemaMessageReq) String() string            { return proto.CompactTextString(m) }
func (*GetCinemaMessageReq) ProtoMessage()               {}
func (*GetCinemaMessageReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetCinemaMessageReq) GetMovieId() int64 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

func (m *GetCinemaMessageReq) GetLocationId() int64 {
	if m != nil {
		return m.LocationId
	}
	return 0
}

func (m *GetCinemaMessageReq) GetDay() string {
	if m != nil {
		return m.Day
	}
	return ""
}

type GetCinemaMessageRsp struct {
	CinemaName     string `protobuf:"bytes,1,opt,name=cinemaName" json:"cinemaName,omitempty"`
	CinemaAddress  string `protobuf:"bytes,2,opt,name=cinemaAddress" json:"cinemaAddress,omitempty"`
	CinemaSupport  string `protobuf:"bytes,3,opt,name=cinemaSupport" json:"cinemaSupport,omitempty"`
	CinemaCard     int64  `protobuf:"varint,4,opt,name=cinemaCard" json:"cinemaCard,omitempty"`
	CinemaMinPrice int64  `protobuf:"varint,5,opt,name=cinemaMinPrice" json:"cinemaMinPrice,omitempty"`
	CinemaDiscount int64  `protobuf:"varint,6,opt,name=cinemaDiscount" json:"cinemaDiscount,omitempty"`
}

func (m *GetCinemaMessageRsp) Reset()                    { *m = GetCinemaMessageRsp{} }
func (m *GetCinemaMessageRsp) String() string            { return proto.CompactTextString(m) }
func (*GetCinemaMessageRsp) ProtoMessage()               {}
func (*GetCinemaMessageRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetCinemaMessageRsp) GetCinemaName() string {
	if m != nil {
		return m.CinemaName
	}
	return ""
}

func (m *GetCinemaMessageRsp) GetCinemaAddress() string {
	if m != nil {
		return m.CinemaAddress
	}
	return ""
}

func (m *GetCinemaMessageRsp) GetCinemaSupport() string {
	if m != nil {
		return m.CinemaSupport
	}
	return ""
}

func (m *GetCinemaMessageRsp) GetCinemaCard() int64 {
	if m != nil {
		return m.CinemaCard
	}
	return 0
}

func (m *GetCinemaMessageRsp) GetCinemaMinPrice() int64 {
	if m != nil {
		return m.CinemaMinPrice
	}
	return 0
}

func (m *GetCinemaMessageRsp) GetCinemaDiscount() int64 {
	if m != nil {
		return m.CinemaDiscount
	}
	return 0
}

type GetCinemaMessageByCidReq struct {
	CinemaId int64 `protobuf:"varint,1,opt,name=cinemaId" json:"cinemaId,omitempty"`
}

func (m *GetCinemaMessageByCidReq) Reset()                    { *m = GetCinemaMessageByCidReq{} }
func (m *GetCinemaMessageByCidReq) String() string            { return proto.CompactTextString(m) }
func (*GetCinemaMessageByCidReq) ProtoMessage()               {}
func (*GetCinemaMessageByCidReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetCinemaMessageByCidReq) GetCinemaId() int64 {
	if m != nil {
		return m.CinemaId
	}
	return 0
}

type FilmMessage struct {
	FilmName    string   `protobuf:"bytes,1,opt,name=filmName" json:"filmName,omitempty"`
	RatingFinal float32  `protobuf:"fixed32,2,opt,name=ratingFinal" json:"ratingFinal,omitempty"`
	Length      int64    `protobuf:"varint,3,opt,name=length" json:"length,omitempty"`
	Type        string   `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	ActorName   []string `protobuf:"bytes,5,rep,name=actorName" json:"actorName,omitempty"`
	MovieId     int64    `protobuf:"varint,6,opt,name=movieId" json:"movieId,omitempty"`
	Img         string   `protobuf:"bytes,7,opt,name=img" json:"img,omitempty"`
}

func (m *FilmMessage) Reset()                    { *m = FilmMessage{} }
func (m *FilmMessage) String() string            { return proto.CompactTextString(m) }
func (*FilmMessage) ProtoMessage()               {}
func (*FilmMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FilmMessage) GetFilmName() string {
	if m != nil {
		return m.FilmName
	}
	return ""
}

func (m *FilmMessage) GetRatingFinal() float32 {
	if m != nil {
		return m.RatingFinal
	}
	return 0
}

func (m *FilmMessage) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *FilmMessage) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FilmMessage) GetActorName() []string {
	if m != nil {
		return m.ActorName
	}
	return nil
}

func (m *FilmMessage) GetMovieId() int64 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

func (m *FilmMessage) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

type GetCinemaMessageByCidRsp struct {
	Cinema      *Cinema        `protobuf:"bytes,1,opt,name=cinema" json:"cinema,omitempty"`
	FilmMessage []*FilmMessage `protobuf:"bytes,2,rep,name=filmMessage" json:"filmMessage,omitempty"`
}

func (m *GetCinemaMessageByCidRsp) Reset()                    { *m = GetCinemaMessageByCidRsp{} }
func (m *GetCinemaMessageByCidRsp) String() string            { return proto.CompactTextString(m) }
func (*GetCinemaMessageByCidRsp) ProtoMessage()               {}
func (*GetCinemaMessageByCidRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetCinemaMessageByCidRsp) GetCinema() *Cinema {
	if m != nil {
		return m.Cinema
	}
	return nil
}

func (m *GetCinemaMessageByCidRsp) GetFilmMessage() []*FilmMessage {
	if m != nil {
		return m.FilmMessage
	}
	return nil
}

type GetMovieHallByMHIdReq struct {
	MhId int64 `protobuf:"varint,1,opt,name=mhId" json:"mhId,omitempty"`
}

func (m *GetMovieHallByMHIdReq) Reset()                    { *m = GetMovieHallByMHIdReq{} }
func (m *GetMovieHallByMHIdReq) String() string            { return proto.CompactTextString(m) }
func (*GetMovieHallByMHIdReq) ProtoMessage()               {}
func (*GetMovieHallByMHIdReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetMovieHallByMHIdReq) GetMhId() int64 {
	if m != nil {
		return m.MhId
	}
	return 0
}

type GetMovieHallByMHIdRsp struct {
	MhAddress string `protobuf:"bytes,1,opt,name=mhAddress" json:"mhAddress,omitempty"`
}

func (m *GetMovieHallByMHIdRsp) Reset()                    { *m = GetMovieHallByMHIdRsp{} }
func (m *GetMovieHallByMHIdRsp) String() string            { return proto.CompactTextString(m) }
func (*GetMovieHallByMHIdRsp) ProtoMessage()               {}
func (*GetMovieHallByMHIdRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetMovieHallByMHIdRsp) GetMhAddress() string {
	if m != nil {
		return m.MhAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*LocationCinemaReq)(nil), "pb.LocationCinemaReq")
	proto.RegisterType((*LocationCinemaRsp)(nil), "pb.LocationCinemaRsp")
	proto.RegisterType((*Cinema)(nil), "pb.Cinema")
	proto.RegisterType((*GetCinemaMessageReq)(nil), "pb.GetCinemaMessageReq")
	proto.RegisterType((*GetCinemaMessageRsp)(nil), "pb.GetCinemaMessageRsp")
	proto.RegisterType((*GetCinemaMessageByCidReq)(nil), "pb.GetCinemaMessageByCidReq")
	proto.RegisterType((*FilmMessage)(nil), "pb.FilmMessage")
	proto.RegisterType((*GetCinemaMessageByCidRsp)(nil), "pb.GetCinemaMessageByCidRsp")
	proto.RegisterType((*GetMovieHallByMHIdReq)(nil), "pb.GetMovieHallByMHIdReq")
	proto.RegisterType((*GetMovieHallByMHIdRsp)(nil), "pb.GetMovieHallByMHIdRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CinemaServiceExt service

type CinemaServiceExtClient interface {
	// 地点影城
	LocationCinema(ctx context.Context, in *LocationCinemaReq, opts ...client.CallOption) (*LocationCinemaRsp, error)
	// 根据位置查看有销售对应电影的影院信息
	// rpc GetCinemaMessage(GetCinemaMessageReq) returns(GetCinemaMessageRsp) {}
	// 根据id查看影院的信息和即将上映的影片信息
	GetCinemaMessageByCid(ctx context.Context, in *GetCinemaMessageByCidReq, opts ...client.CallOption) (*GetCinemaMessageByCidRsp, error)
	// 根据mh_id获取影厅信息
	GetMovieHallByMHId(ctx context.Context, in *GetMovieHallByMHIdReq, opts ...client.CallOption) (*GetMovieHallByMHIdRsp, error)
}

type cinemaServiceExtClient struct {
	c           client.Client
	serviceName string
}

func NewCinemaServiceExtClient(serviceName string, c client.Client) CinemaServiceExtClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "pb"
	}
	return &cinemaServiceExtClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *cinemaServiceExtClient) LocationCinema(ctx context.Context, in *LocationCinemaReq, opts ...client.CallOption) (*LocationCinemaRsp, error) {
	req := c.c.NewRequest(c.serviceName, "CinemaServiceExt.LocationCinema", in)
	out := new(LocationCinemaRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cinemaServiceExtClient) GetCinemaMessageByCid(ctx context.Context, in *GetCinemaMessageByCidReq, opts ...client.CallOption) (*GetCinemaMessageByCidRsp, error) {
	req := c.c.NewRequest(c.serviceName, "CinemaServiceExt.GetCinemaMessageByCid", in)
	out := new(GetCinemaMessageByCidRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cinemaServiceExtClient) GetMovieHallByMHId(ctx context.Context, in *GetMovieHallByMHIdReq, opts ...client.CallOption) (*GetMovieHallByMHIdRsp, error) {
	req := c.c.NewRequest(c.serviceName, "CinemaServiceExt.GetMovieHallByMHId", in)
	out := new(GetMovieHallByMHIdRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CinemaServiceExt service

type CinemaServiceExtHandler interface {
	// 地点影城
	LocationCinema(context.Context, *LocationCinemaReq, *LocationCinemaRsp) error
	// 根据位置查看有销售对应电影的影院信息
	// rpc GetCinemaMessage(GetCinemaMessageReq) returns(GetCinemaMessageRsp) {}
	// 根据id查看影院的信息和即将上映的影片信息
	GetCinemaMessageByCid(context.Context, *GetCinemaMessageByCidReq, *GetCinemaMessageByCidRsp) error
	// 根据mh_id获取影厅信息
	GetMovieHallByMHId(context.Context, *GetMovieHallByMHIdReq, *GetMovieHallByMHIdRsp) error
}

func RegisterCinemaServiceExtHandler(s server.Server, hdlr CinemaServiceExtHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&CinemaServiceExt{hdlr}, opts...))
}

type CinemaServiceExt struct {
	CinemaServiceExtHandler
}

func (h *CinemaServiceExt) LocationCinema(ctx context.Context, in *LocationCinemaReq, out *LocationCinemaRsp) error {
	return h.CinemaServiceExtHandler.LocationCinema(ctx, in, out)
}

func (h *CinemaServiceExt) GetCinemaMessageByCid(ctx context.Context, in *GetCinemaMessageByCidReq, out *GetCinemaMessageByCidRsp) error {
	return h.CinemaServiceExtHandler.GetCinemaMessageByCid(ctx, in, out)
}

func (h *CinemaServiceExt) GetMovieHallByMHId(ctx context.Context, in *GetMovieHallByMHIdReq, out *GetMovieHallByMHIdRsp) error {
	return h.CinemaServiceExtHandler.GetMovieHallByMHId(ctx, in, out)
}

func init() { proto.RegisterFile("cinema.ext.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x94, 0xdf, 0x6a, 0xdb, 0x3e,
	0x14, 0xc7, 0x7f, 0xb6, 0x5b, 0xf7, 0x97, 0x13, 0xd6, 0x65, 0x67, 0x74, 0x68, 0xa1, 0x8c, 0x20,
	0xca, 0x08, 0x0c, 0x02, 0x6b, 0xd9, 0x60, 0x77, 0x5b, 0xb3, 0xb5, 0x0d, 0x34, 0x63, 0xb8, 0xec,
	0x01, 0x14, 0x5b, 0x4d, 0x04, 0xfe, 0xa3, 0x5a, 0x6a, 0x69, 0xee, 0xf6, 0x6a, 0x7b, 0xa4, 0x5d,
	0xee, 0x6e, 0x48, 0x72, 0x62, 0xe7, 0xdf, 0x1e, 0x60, 0x77, 0x3a, 0xdf, 0xf3, 0x95, 0x8f, 0xce,
	0x47, 0xc7, 0x82, 0x4e, 0x2c, 0x72, 0x9e, 0xb1, 0x01, 0x7f, 0xd4, 0x03, 0x59, 0x16, 0xba, 0x40,
	0x5f, 0x4e, 0xe8, 0x19, 0x3c, 0xbb, 0x2e, 0x62, 0xa6, 0x45, 0x91, 0x0f, 0x6d, 0x3e, 0xe2, 0x77,
	0xf8, 0x0a, 0x20, 0xad, 0xc4, 0x51, 0x42, 0xbc, 0x9e, 0xd7, 0x0f, 0xa2, 0x86, 0x42, 0x3f, 0x6c,
	0x6c, 0x52, 0x12, 0x4f, 0xe0, 0xc0, 0x55, 0x50, 0xc4, 0xeb, 0x05, 0xfd, 0xf6, 0x29, 0x0c, 0xe4,
	0x64, 0x50, 0xe5, 0x17, 0x29, 0xfa, 0xc3, 0x87, 0xd0, 0x69, 0xa6, 0x8a, 0x53, 0xbf, 0xb2, 0x8c,
	0xdb, 0x2a, 0xad, 0xa8, 0xa1, 0xe0, 0x09, 0x3c, 0x71, 0xd1, 0xa7, 0x24, 0x29, 0xb9, 0x52, 0xc4,
	0xb7, 0x96, 0x55, 0xb1, 0x76, 0xdd, 0xdc, 0x4b, 0x59, 0x94, 0x9a, 0x04, 0x4d, 0x57, 0x25, 0xd6,
	0xb5, 0x86, 0xac, 0x4c, 0xc8, 0x9e, 0xeb, 0xa8, 0x56, 0xf0, 0x35, 0x1c, 0xba, 0x68, 0x2c, 0xf2,
	0x6f, 0xa5, 0x88, 0x39, 0xd9, 0xb7, 0x9e, 0x35, 0xb5, 0xf6, 0x7d, 0x16, 0x2a, 0x2e, 0xee, 0x73,
	0x4d, 0xc2, 0xa6, 0x6f, 0xa1, 0x62, 0x17, 0xfe, 0x77, 0xca, 0x28, 0x21, 0x07, 0xd6, 0xb1, 0x8c,
	0x29, 0x83, 0xe7, 0x97, 0x5c, 0x3b, 0x08, 0x63, 0xae, 0x14, 0x9b, 0x72, 0x03, 0x9d, 0xc0, 0x41,
	0x56, 0x3c, 0x08, 0xbe, 0x24, 0xbe, 0x08, 0xd7, 0xae, 0xc3, 0x5f, 0xbf, 0x0e, 0xec, 0x40, 0x90,
	0xb0, 0x79, 0xd5, 0xb8, 0x59, 0xd2, 0x5f, 0xde, 0x96, 0x1a, 0x4a, 0xfe, 0xcb, 0xc8, 0xe9, 0x7b,
	0x20, 0xeb, 0x2d, 0x9f, 0xcf, 0x87, 0x22, 0x31, 0x6c, 0x9b, 0xd7, 0xe1, 0xad, 0x5d, 0xc7, 0x4f,
	0x0f, 0xda, 0x17, 0x22, 0xcd, 0xaa, 0x3d, 0xc6, 0x7b, 0x2b, 0xd2, 0xac, 0x41, 0x68, 0x19, 0x63,
	0x0f, 0xda, 0x25, 0xd3, 0x22, 0x9f, 0x5e, 0x88, 0x9c, 0xa5, 0x96, 0x8e, 0x1f, 0x35, 0x25, 0x7c,
	0x01, 0x61, 0xca, 0xf3, 0xa9, 0x9e, 0x59, 0x28, 0x41, 0x54, 0x45, 0x88, 0xb0, 0xa7, 0xe7, 0x92,
	0x5b, 0x0e, 0xad, 0xc8, 0xae, 0xf1, 0x18, 0x5a, 0x2c, 0xd6, 0x45, 0x69, 0x4b, 0xed, 0xf7, 0x82,
	0x7e, 0x2b, 0xaa, 0x85, 0xe6, 0x3c, 0x84, 0xab, 0xf3, 0xd0, 0x81, 0x40, 0x64, 0x53, 0x3b, 0x57,
	0xad, 0xc8, 0x2c, 0xe9, 0xdd, 0xae, 0xde, 0x95, 0x44, 0x0a, 0xa1, 0xeb, 0xd5, 0x76, 0xb3, 0xfa,
	0x5b, 0x56, 0x19, 0x7c, 0x0b, 0xed, 0xdb, 0x1a, 0x01, 0xf1, 0xed, 0xff, 0xfb, 0xd4, 0x18, 0x1b,
	0x64, 0xa2, 0xa6, 0x87, 0xbe, 0x81, 0xa3, 0x4b, 0xae, 0xc7, 0xe6, 0x48, 0x57, 0x2c, 0x4d, 0xcf,
	0xe7, 0xe3, 0xab, 0x91, 0x65, 0x8d, 0xb0, 0x97, 0xcd, 0x96, 0x9c, 0xed, 0x9a, 0xbe, 0xdb, 0x6a,
	0x56, 0xd2, 0x20, 0xc8, 0x66, 0x8b, 0x61, 0x73, 0xb4, 0x6b, 0xe1, 0xf4, 0xb7, 0x07, 0x1d, 0x77,
	0xd2, 0x1b, 0x5e, 0x3e, 0x88, 0x98, 0x7f, 0x79, 0xd4, 0xf8, 0x11, 0x0e, 0x57, 0x1f, 0x1f, 0x3c,
	0x32, 0x07, 0xdd, 0x78, 0xc5, 0xba, 0xdb, 0x64, 0x25, 0xe9, 0x7f, 0xf8, 0xdd, 0x9e, 0x66, 0x93,
	0x16, 0x1e, 0x9b, 0x1d, 0xbb, 0x86, 0xa8, 0xfb, 0x97, 0xac, 0xfd, 0xec, 0x35, 0xe0, 0x66, 0x93,
	0xf8, 0xb2, 0xda, 0xb5, 0x49, 0xaa, 0xbb, 0x2b, 0x65, 0xbe, 0x36, 0x09, 0xed, 0x1b, 0x7d, 0xf6,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x78, 0xca, 0x53, 0xb7, 0x05, 0x00, 0x00,
}
