//** This file is code generated by gopp. Do not edit.

package db

import (
	"database/sql"
	"database/sql/driver"
	//"cmd/pprof/internal/report"
	"context"
	"fmt"
	gopp "github.com/spekary/gopp"
	"reflect"
	"strings"
)

type EqualityOperation int

const (
	EqNone EqualityOperation = iota
	Eq
	NotEq
)

type SqlDbI interface {
	gopp.BaseI

	Begin()
	Commit()
	Rollback()
	Exec(ctx context.Context, sql string, args ...interface{}) (r sql.Result, err error)
	Prepare(sql string) (r *sql.Stmt, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (r *sql.Rows, err error)
	SqlValue(data interface{}, eq EqualityOperation) string
	EscapeIdentifierBegin() string
	EscapeIdentifierEnd() string
	OnlyFullGroupBy() bool
	EscapeString(s string) string
	DbKey() string
	SetTypeTableSuffix(s string)
	SetAssociationTableSuffix(s string)
	TypeTableSuffix() string
	AssociationTableSuffix() string
	SetGoStructPrefix(s string)
	SetAssociatedObjectPrefix(s string)
	GoStructPrefix() string
	AssociatedObjectPrefix() string
	IdSuffix() string
	NewBuilder() QueryBuilderI
	generateSelectSql(b *sqlBuilder) (sql string, args []interface{})
	generateDeleteSql(b *sqlBuilder) (sql string, args []interface{})
}

type SqlDb struct {
	gopp.Base
	dbKey                 string  // key of the database as used in the global database map
	db                    *sql.DB // Internal copy of golang database
	escapeIdentifierBegin string
	escapeIdentifierEnd   string
	tx                    *sql.Tx
	txCount               int
	// codegen options
	typeTableSuffix        string // Primarily for sql tables
	associationTableSuffix string // Primarily for sql tables
	idSuffix               string // suffix to strip off the ends of names of foreign keys when converting them to internal names
	// These codegen options may be moved higher up in hierarchy some day
	goStructPrefix         string // Helps differentiate objects when different databases have the same name.
	associatedObjectPrefix string // Helps differentiate between objects and local values

}

// New SqlDb creates a new SqlDb object and returns its matching interface
func NewSqlDb(dbKey string) SqlDbI {
	s_ := SqlDb{}
	s_.Init(&s_)
	s_.Construct(dbKey)
	return s_.I().(SqlDbI)
}

func (s_ *SqlDb) Construct(dbKey string) {
	s_.dbKey = dbKey
	s_.typeTableSuffix = "_type"
	s_.associationTableSuffix = "_assn"
	s_.idSuffix = "_id"
}

func (s_ *SqlDb) Begin() {
	s_.txCount++

	if s_.txCount == 1 {
		var err error

		s_.tx, err = s_.db.Begin()
		if err != nil {
			panic(err.Error())
		}
	}
}

func (s_ *SqlDb) Commit() {
	s_.txCount--
	if s_.txCount < 0 {
		panic("Called Commit without a matching Begin")
	}
	if s_.txCount == 0 {
		err := s_.tx.Commit()
		if err != nil {
			panic(err.Error())
		}
		s_.tx = nil
	}
}

func (s_ *SqlDb) Rollback() {
	if s_.tx != nil {
		err := s_.tx.Rollback()
		if err != nil {
			panic(err.Error())
		}
		s_.tx = nil
		s_.txCount = 0
	}
}

func (s_ *SqlDb) Exec(ctx context.Context, sql string, args ...interface{}) (r sql.Result, err error) {
	if s_.tx != nil {
		r, err = s_.tx.ExecContext(ctx, sql, args...)
	} else {
		r, err = s_.db.ExecContext(ctx, sql, args...)
	}
	return
}

func (s_ *SqlDb) Prepare(sql string) (r *sql.Stmt, err error) {
	if s_.tx != nil {
		r, err = s_.tx.Prepare(sql)
	} else {
		r, err = s_.db.Prepare(sql)
	}
	return
}

func (s_ *SqlDb) Query(ctx context.Context, sql string, args ...interface{}) (r *sql.Rows, err error) {
	if s_.tx != nil {
		r, err = s_.tx.QueryContext(ctx, sql, args...)
	} else {
		r, err = s_.db.QueryContext(ctx, sql, args...)
	}
	return
}

func (s_ *SqlDb) SqlValue(data interface{}, eq EqualityOperation) string {
	var eqOp string

	switch eq {
	case Eq:
		eqOp = "= "
	case NotEq:
		eqOp = "!= "
	}

	if value, ok := data.(driver.Valuer); ok {
		data, _ = value.Value()
	}

	switch d := data.(type) {
	case nil:
		switch eq {
		case EqNone:
			return "NULL"
		case Eq:
			return "IS NULL"
		case NotEq:
			return "IS NOT NULL"
		}
	case bool:
		switch eq {
		case EqNone:
			if d {
				return "1"
			} else {
				return "0"
			}
		case Eq:
			if d {
				return "!= 0"
			} else {
				return "= 0"
			}
		case NotEq:
			if d {
				return "= 0"
			} else {
				return "!= 0"
			}
		}
	case int, float64:
		return eqOp + fmt.Sprint("%v", d)

	case string, []byte:
		return eqOp + s_.I().(SqlDbI).EscapeString(data.(string))

	default:
		var items []string
		var v = reflect.ValueOf(data)
		if k := v.Kind(); k == reflect.Array || k == reflect.Slice {
			for i := 0; i < v.Len(); i++ {
				items = append(items, s_.I().(SqlDbI).SqlValue(v.Index(i), eq))
			}
		}
		return "(" + strings.Join(items, ",") + ")"
	}
	return ""
}

func (s_ *SqlDb) EscapeIdentifierBegin() string {
	return ""
}

func (s_ *SqlDb) EscapeIdentifierEnd() string {
	return ""
}

func (s_ *SqlDb) OnlyFullGroupBy() bool {
	return false
}

func (s_ *SqlDb) EscapeString(s string) string {
	return ""
}

func (s_ *SqlDb) DbKey() string {
	return s_.dbKey
}

func (s_ *SqlDb) SetTypeTableSuffix(s string) {
	s_.typeTableSuffix = s
}

func (s_ *SqlDb) SetAssociationTableSuffix(s string) {
	s_.associationTableSuffix = s
}

func (s_ *SqlDb) TypeTableSuffix() string {
	return s_.typeTableSuffix
}

func (s_ *SqlDb) AssociationTableSuffix() string {
	return s_.associationTableSuffix
}

func (s_ *SqlDb) SetGoStructPrefix(s string) {
	s_.goStructPrefix = s
}

func (s_ *SqlDb) SetAssociatedObjectPrefix(s string) {
	s_.associatedObjectPrefix = s
}

func (s_ *SqlDb) GoStructPrefix() string {
	return s_.goStructPrefix
}

func (s_ *SqlDb) AssociatedObjectPrefix() string {
	return s_.associatedObjectPrefix
}

func (s_ *SqlDb) IdSuffix() string {
	return s_.idSuffix
}

func (s_ *SqlDb) NewBuilder() QueryBuilderI {
	return NewSqlBuilder(s_.I().(SqlDbI))
}

func (s_ *SqlDb) generateSelectSql(b *sqlBuilder) (sql string, args []interface{}) {
	return
}

func (s_ *SqlDb) generateDeleteSql(b *sqlBuilder) (sql string, args []interface{}) {
	return
}

func (s_ *SqlDb) IsA(className string) bool {
	if className == "SqlDb" {
		return true
	}
	return s_.Base.IsA(className)
}

func (s_ *SqlDb) Class() string {
	return "SqlDb"
}
