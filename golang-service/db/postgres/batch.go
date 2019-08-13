package postgres

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/jinzhu/gorm"
)

const maxNumberOfBatchParameters = 65535

// BatchInsert from https://github.com/jinzhu/gorm/issues/255#issuecomment-507023750
func BatchInsert(db *gorm.DB, objArr []interface{}) (int64, error) {
	// If there is no data, nothing to do.
	if len(objArr) == 0 {
		return 0, errors.New("slice must not be empty")
	}

	numberOfFields := calculateNumberOfFields(db, (objArr)[0])
	return batchInsertChunks(db, splitIntoChunks(objArr, numberOfFields))
}

func splitIntoChunks(objArr []interface{}, numberOfFields int) [][]interface{} {
	var chunks [][]interface{}

	chunkSize := int(math.Floor(float64(maxNumberOfBatchParameters / float32(numberOfFields))))
	numberOfObjects := len(objArr)

	if numberOfObjects < chunkSize {
		return [][]interface{}{objArr}
	}

	for i := 0; i < numberOfObjects; i += chunkSize {
		end := i + chunkSize

		if end > numberOfObjects {
			end = numberOfObjects
		}

		chunks = append(chunks, objArr[i:end])
	}

	return chunks
}

func calculateNumberOfFields(db *gorm.DB, obj interface{}) int {
	return len(db.NewScope(obj).Fields())
}

func batchInsertChunks(db *gorm.DB, chunks [][]interface{}) (int64, error) {
	var rowsAffected int64 = 0
	for _, chunk := range chunks {
		chunkRowsAffected, err := batchInsert(db, chunk)
		if err != nil {
			return 0, err
		}

		rowsAffected += chunkRowsAffected
	}

	return rowsAffected, nil
}

func batchInsert(db *gorm.DB, objArr []interface{}) (int64, error) {
	// If there is no data, nothing to do.
	if len(objArr) == 0 {
		return 0, errors.New("slice must not be empty")
	}

	mainObj := objArr[0]
	mainScope := db.NewScope(mainObj)
	mainFields := mainScope.Fields()
	quoted := make([]string, 0, len(mainFields))
	for i := range mainFields {
		// If primary key has blank value (0 for int, "" for string, nil for interface ...), skip it.
		// If field is ignore field, skip it.
		if (mainFields[i].IsPrimaryKey && mainFields[i].IsBlank) || (mainFields[i].IsIgnored) {
			continue
		}
		quoted = append(quoted, mainScope.Quote(mainFields[i].DBName))
	}

	placeholdersArr := make([]string, 0, len(objArr))

	for _, obj := range objArr {
		scope := db.NewScope(obj)
		fields := scope.Fields()

		placeholders := make([]string, 0, len(fields))
		for i := range fields {
			if (fields[i].IsPrimaryKey && fields[i].IsBlank) || (fields[i].IsIgnored) {
				continue
			}
			var vars interface{}
			if (fields[i].Name == "CreatedAt" || fields[i].Name == "UpdatedAt") && fields[i].IsBlank {
				vars = gorm.NowFunc()
			} else {
				vars = fields[i].Field.Interface()
			}
			placeholders = append(placeholders, mainScope.AddToVars(vars))
		}

		placeholdersStr := "(" + strings.Join(placeholders, ", ") + ")"
		placeholdersArr = append(placeholdersArr, placeholdersStr)
	}

	mainScope.Raw(fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;",
		mainScope.QuotedTableName(),
		strings.Join(quoted, ", "),
		strings.Join(placeholdersArr, ", "),
	))

	// Execute and Log
	if err := mainScope.Exec().DB().Error; err != nil {
		return 0, err
	}
	return mainScope.DB().RowsAffected, nil
}
