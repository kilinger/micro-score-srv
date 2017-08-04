package models

import (
	"math"

	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	
	proto "code.xxxxxx.com/micro/proto/scores"
)

type Scores struct {
	gorm.Model
	UserId int64 `gorm:"unique_index"`
	Score int64 `gorm:"default:0"`
}

func GetUserScore (db *gorm.DB, user_id int64) int64 {
	score := Scores{UserId: user_id, Score: 0}
	db.Where(score).FirstOrInit(&score)
	return score.Score
}

type ScoresRecord struct {
	gorm.Model
	UserId int64
	Score int32
	Note string
	GiftId int64
	GiftName string
}

func CreateOrUpdateScore(db *gorm.DB, rscore int32, user_id int64) (change_faild bool) {
	tx := db.Begin()
	score := Scores{UserId: user_id}
	tx.Set("gorm:query_option", "FOR UPDATE").Where(score).First(&score)
	if rscore < 0 {
		rscore = int32(math.Abs(float64(rscore)))
		expr := gorm.Expr("score - ?", rscore)
		tx.Model(&score).Where("score > ?", rscore).UpdateColumn("score", expr)
	} else {
		tx.Model(&score).UpdateColumn("score", gorm.Expr("score + ?", rscore))
	}
	tx.Commit()
	score_after := Scores{UserId: user_id}
	db.Where(score_after).FirstOrCreate(&score_after)
	return score.Score == score_after.Score
}

func (record *ScoresRecord) AfterSave(db *gorm.DB) (err error) {
	// CreateOrUpdateScore(db, record.Score, record.UserId)
	return nil
}

func (record *ScoresRecord) AfterUpdate(db *gorm.DB) (err error) {
	// CreateOrUpdateScore(db, record.Score, record.UserId)
	return nil
}

func (record *ScoresRecord) ToProto() *proto.ScoreResponse {
	return &proto.ScoreResponse{
		UserId: record.UserId,
		Score: record.Score,
		Note: record.Note,
		GiftId: record.GiftId,
		GiftName: record.GiftName,
		CreatedAt: TimeToString(record.CreatedAt),
	}
}

func GetRecordDB (db *gorm.DB, req *proto.ListRequest) (*gorm.DB, int32) {
	var total int32
	db = db.Model(&ScoresRecord{})
	db = GetRecordUserDB(db, req.GetUserId())
	db = GetRecordTimeDB(db, req.GetStart(), req.GetEnd())
	db = GetRecordScoreDB(db, req.GetScoreStart(), req.GetScoreEnd())
	db.Count(&total)
	return db, total
}

func GetRecordUserDB (db *gorm.DB, user_id int64) *gorm.DB {
	if user_id != 0 {
		db = db.Where(&ScoresRecord{UserId: user_id})
	}
	return db
}

func GetRecordTimeDB (db *gorm.DB, start string, end string) *gorm.DB {
	if start != "" {
		start := StringToTime(start)
		if start != nil {
			db = db.Where("created_at >= ?", start)
		}
	}

	if end != "" {
		end := StringToTime(end)
		if end != nil {
			db = db.Where("created_at < ?", end)
		}
	}
	return db
}

func GetRecordScoreDB (db *gorm.DB, start int64, end int64) *gorm.DB {
	if start > 0 {
		db = db.Where("score >= ?", start)
	}

	if end > 0 {
		db = db.Where("score < ?", end)
	}
	return db
}
