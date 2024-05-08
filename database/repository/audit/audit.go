package audit

import (
	"context"
	"time"

	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries/qm"
	"github.com/aaabigfish/goex/database"
	modelPSQL "github.com/aaabigfish/goex/database/models/postgres"
	modelSQLite "github.com/aaabigfish/goex/database/models/sqlite3"
	"github.com/aaabigfish/goex/database/repository"
	"github.com/aaabigfish/gopkg/log"
)

// Event inserts a new audit event to database
func Event(id, msgtype, message string) {
	if database.DB.SQL == nil {
		return
	}

	ctx := context.TODO()
	ctx = boil.SkipTimestamps(ctx)

	tx, err := database.DB.SQL.BeginTx(ctx, nil)
	if err != nil {
		log.Errorf("Global Event transaction begin failed: %v", err)
		return
	}

	if repository.GetSQLDialect() == database.DBSQLite3 {
		var tempEvent = modelSQLite.AuditEvent{
			Type:       msgtype,
			Identifier: id,
			Message:    message,
		}
		err = tempEvent.Insert(ctx, tx, boil.Blacklist("created_at"))
	} else {
		var tempEvent = modelPSQL.AuditEvent{
			Type:       msgtype,
			Identifier: id,
			Message:    message,
		}
		err = tempEvent.Insert(ctx, tx, boil.Blacklist("created_at"))
	}

	if err != nil {
		log.Errorf("Global Event insert failed: %v", err)
		err = tx.Rollback()
		if err != nil {
			log.Errorf("Global Event Transaction rollback failed: %v", err)
		}
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Errorf("Global Event Transaction commit failed: %v", err)
		return
	}
}

// GetEvent () returns list of order events matching query
func GetEvent(startTime, endTime time.Time, order string, limit int) (interface{}, error) {
	if database.DB.SQL == nil {
		return nil, database.ErrDatabaseSupportDisabled
	}

	query := qm.Where("created_at BETWEEN ? AND ?", startTime, endTime)

	orderByQueryString := "id"
	if order == "desc" {
		orderByQueryString += " desc"
	}

	orderByQuery := qm.OrderBy(orderByQueryString)
	limitQuery := qm.Limit(limit)

	ctx := context.TODO()
	if repository.GetSQLDialect() == database.DBSQLite3 {
		return modelSQLite.AuditEvents(query, orderByQuery, limitQuery).All(ctx, database.DB.SQL)
	}

	return modelPSQL.AuditEvents(query, orderByQuery, limitQuery).All(ctx, database.DB.SQL)
}
