package exchange

import (
	"context"
	"database/sql"
	"encoding/csv"
	"io"
	"os"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries/qm"
	"github.com/aaabigfish/goex/common"
	"github.com/aaabigfish/goex/common/cache"
	"github.com/aaabigfish/goex/database"
	modelPSQL "github.com/aaabigfish/goex/database/models/postgres"
	modelSQLite "github.com/aaabigfish/goex/database/models/sqlite3"
	"github.com/aaabigfish/goex/database/repository"
	"github.com/aaabigfish/gopkg/log"
)

// One returns one exchange by Name
func One(in string) (Details, error) {
	return one(in, "name")
}

// OneByUUID returns one exchange by UUID
func OneByUUID(in uuid.UUID) (Details, error) {
	return one(in.String(), "id")
}

// one returns one exchange by clause
func one(in, clause string) (out Details, err error) {
	if database.DB.SQL == nil {
		return out, database.ErrDatabaseSupportDisabled
	}

	whereQM := qm.Where(clause+"= ?", in)
	if repository.GetSQLDialect() == database.DBSQLite3 {
		ret, errS := modelSQLite.Exchanges(whereQM).One(context.TODO(), database.DB.SQL)
		if errS != nil {
			return out, errS
		}
		out.Name = ret.Name
		out.UUID, errS = uuid.FromString(ret.ID)
		if errS != nil {
			return out, errS
		}
	} else {
		ret, errS := modelPSQL.Exchanges(whereQM).One(context.TODO(), database.DB.SQL)
		if errS != nil {
			return out, errS
		}
		out.Name = ret.Name
		out.UUID, errS = uuid.FromString(ret.ID)
		if errS != nil {
			return out, errS
		}
	}

	return out, err
}

// Insert writes a single entry into database
func Insert(in Details) error {
	if database.DB.SQL == nil {
		return database.ErrDatabaseSupportDisabled
	}

	ctx := context.TODO()
	tx, err := database.DB.SQL.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if repository.GetSQLDialect() == database.DBSQLite3 {
		err = insertSQLite(ctx, tx, []Details{in})
	} else {
		err = insertPostgresql(ctx, tx, []Details{in})
	}

	if err != nil {
		errRB := tx.Rollback()
		if errRB != nil {
			log.Errorf("DatabaseMgr err(%v)", errRB)
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// InsertMany writes multiple entries into database
func InsertMany(in []Details) error {
	if database.DB.SQL == nil {
		return database.ErrDatabaseSupportDisabled
	}

	ctx := context.TODO()
	tx, err := database.DB.SQL.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if repository.GetSQLDialect() == database.DBSQLite3 {
		err = insertSQLite(ctx, tx, in)
	} else {
		err = insertPostgresql(ctx, tx, in)
	}

	if err != nil {
		errRB := tx.Rollback()
		if errRB != nil {
			log.Errorf("DatabaseMgr err(%v)", errRB)
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		errRB := tx.Rollback()
		if errRB != nil {
			log.Errorf("DatabaseMgr err(%v)", errRB)
		}
		return err
	}
	return nil
}

func insertSQLite(ctx context.Context, tx *sql.Tx, in []Details) (err error) {
	for x := range in {
		tempUUID, errUUID := uuid.NewV4()
		if errUUID != nil {
			return errUUID
		}
		var tempInsert = modelSQLite.Exchange{
			Name: strings.ToLower(in[x].Name),
			ID:   tempUUID.String(),
		}

		err = tempInsert.Insert(ctx, tx, boil.Infer())
		if err != nil {
			errRB := tx.Rollback()
			if errRB != nil {
				log.Errorf("DatabaseMgr err(%v)", errRB)
			}
			return err
		}
	}

	return nil
}

func insertPostgresql(ctx context.Context, tx *sql.Tx, in []Details) (err error) {
	for x := range in {
		var tempInsert = modelPSQL.Exchange{
			Name: strings.ToLower(in[x].Name),
		}

		err = tempInsert.Upsert(ctx, tx, true, []string{"name"}, boil.Infer(), boil.Infer())
		if err != nil {
			errRB := tx.Rollback()
			if errRB != nil {
				log.Errorf("DatabaseMgr err(%v)", errRB)
			}
			return
		}
	}
	return nil
}

// UUIDByName returns UUID of exchange
func UUIDByName(exchange string) (uuid.UUID, error) {
	exchange = strings.ToLower(exchange)
	v := exchangeCache.Get(exchange)
	if v != nil {
		u, ok := v.(uuid.UUID)
		if !ok {
			return uuid.UUID{}, common.GetTypeAssertError("uuid.UUID", v)
		}
		return u, nil
	}
	ret, err := One(exchange)
	if err != nil {
		if err != sql.ErrNoRows {
			return uuid.UUID{}, err
		}
		return uuid.UUID{}, ErrNoExchangeFound
	}

	exchangeCache.Add(exchange, ret.UUID)
	return ret.UUID, nil
}

// ResetExchangeCache reinitialise cache to blank state used to clear cache for testing
func ResetExchangeCache() {
	exchangeCache = cache.New(10)
}

// LoadCSV loads & parses a CSV list of exchanges
func LoadCSV(file string) (out []Details, err error) {
	csvFile, err := os.Open(file)
	if err != nil {
		return out, err
	}

	defer func() {
		err = csvFile.Close()
		if err != nil {
			log.Errorf("LoadCSV err(%v)", err)
		}
	}()

	csvData := csv.NewReader(csvFile)
	for {
		row, errCSV := csvData.Read()
		if errCSV != nil {
			if errCSV == io.EOF {
				return out, err
			}
			return out, errCSV
		}

		out = append(out, Details{
			Name: row[0],
		})
	}
}
