package handler

import (
	"golang.org/x/net/context"
	"github.com/jinzhu/gorm"

	proto "code.xxxxxx.com/micro/proto/scores"
	m "code.xxxxxx.com/micro/scores-srv/models"
	"code.xxxxxx.com/micro/scores-srv/store"
)

const (
	success = "SUCCESS"
	fail = "FAIL"
)

// Handler the Scores Handler
type Handler struct {
	store store.Store
	DB *gorm.DB
}

func (h *Handler) Create(ctx context.Context, req *proto.CreateRequest, rsp *proto.ScoreResponse) error {

	msg, ok := CheckCreateRequest(req); if !ok { 
		rsp.Code = fail
		rsp.Msg = msg
		return nil
	}

	change_faild := m.CreateOrUpdateScore(h.DB, req.GetScore(), req.GetUserId()); if change_faild {
		rsp.Code = fail
		rsp.Msg = "user score deficiency"
	} else {
		record := m.ScoresRecord{
			UserId: req.GetUserId(),
			Score: req.GetScore(),
			Note: req.GetNote(),
		}
		h.DB.Create(&record)
		rsp.UserId = record.UserId
		rsp.Score = record.Score
		rsp.Note = record.Note
		rsp.CreatedAt = m.TimeToString(record.CreatedAt)
		rsp.Code = success
	}
	
	return nil
}

func (h *Handler) List(ctx context.Context, req *proto.ListRequest, rsp *proto.ListResponse) error {

	var (
		records []m.ScoresRecord
		results []*proto.ScoreResponse
	)

	db, total := m.GetRecordDB(h.DB, req)
	limit, offset := GetLimitOffset(req.GetPage(), req.GetCount(), total)
	db = db.Offset(offset).Limit(limit).Order("created_at desc")
	db.Find(&records)

	for _, record := range records {
		results = append(results, record.ToProto())
	}

	rsp.Results, rsp.Total = results, total
	return nil
}

func (h *Handler) Stat(ctx context.Context, req *proto.StatRequest, rsp *proto.StatResponse) error {

	msg, ok := CheckStatRequest(req); if !ok { 
		rsp.Code = fail
		rsp.Msg = msg
		return nil
	}

	rsp.Code = success
	rsp.Score = m.GetUserScore(h.DB, req.UserId)
	return nil
}

func (h *Handler) Exchange(ctx context.Context, req *proto.ExchangeRequest, rsp *proto.ExchangeResponse) error {

	msg, ok := CheckExchangeRequest(req); if !ok { 
		rsp.Code = fail
		rsp.Msg = msg
		return nil
	}

	score := req.GetScore()
	if score > 0 {
		score = -score
	}
	
	change_faild := m.CreateOrUpdateScore(h.DB, score, req.GetUserId()); if change_faild {
		rsp.Code = fail
		rsp.Msg = "user score deficiency"
	} else {
		note := "积分兑换"
		record := m.ScoresRecord{
			UserId: req.GetUserId(),
			Score: score,
			Note: note,
			GiftId: req.GetGiftId(),
			GiftName: req.GetGiftName(),
		}
		h.DB.Create(&record)
		rsp.UserId = record.UserId
		rsp.Score = record.Score
		rsp.Note = record.Note
		rsp.GiftId = record.GiftId
		rsp.GiftName = record.GiftName
		rsp.CreatedAt = m.TimeToString(record.CreatedAt)
		rsp.Code = success
	}

	return nil
}

// New create handler
func New(s store.Store, db *gorm.DB) *Handler {
	return &Handler{
		store: s,
		DB: db,
	}
}
