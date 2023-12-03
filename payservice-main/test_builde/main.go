package main

import (
	"fmt"
	"strings"
)

type SQLQueryBuilder interface {
	Select(fields ...string) SQLQueryBuilder
	From(table string) SQLQueryBuilder
	Where(condition string) SQLQueryBuilder
	OrderBy(field string) SQLQueryBuilder
	Like(word string) SQLQueryBuilder
	Build() string
	In(data string) SQLQueryBuilder
	BracketIn() SQLQueryBuilder
	BracketOut() SQLQueryBuilder

	Insert(table string, fields ...string) SQLQueryBuilder
	Values(fields ...string) SQLQueryBuilder
}

type ConcreteSQLQueryBuilder struct {
	queryParts []string
}

func NewSQLQueryBuilder() SQLQueryBuilder {
	return &ConcreteSQLQueryBuilder{}
}

func (qb *ConcreteSQLQueryBuilder) Select(fields ...string) SQLQueryBuilder {
	qb.queryParts = append(qb.queryParts, "SELECT "+strings.Join(fields, ", "))
	return qb
}

func (qb *ConcreteSQLQueryBuilder) Insert(table string, fields ...string) SQLQueryBuilder {
	qb.queryParts = append(qb.queryParts, "INSERT INTO "+table+"("+strings.Join(fields, ", ")+")")
	return qb
}

func (qb *ConcreteSQLQueryBuilder) Values(fields ...string) SQLQueryBuilder {
	qb.queryParts = append(qb.queryParts, "VALUES "+"("+strings.Join(fields, ", ")+") ")
	return qb
}

func (qb *ConcreteSQLQueryBuilder) From(table string) SQLQueryBuilder {
	qb.queryParts = append(qb.queryParts, "FROM "+table)
	return qb
}

func (qb *ConcreteSQLQueryBuilder) In(data string) SQLQueryBuilder {
	qb.queryParts = append(qb.queryParts, "IN ("+data+") ")
	return qb
}

func (qb *ConcreteSQLQueryBuilder) BracketIn() SQLQueryBuilder {
	qb.queryParts = append(qb.queryParts, "( ")
	return qb
}

func (qb *ConcreteSQLQueryBuilder) BracketOut() SQLQueryBuilder {
	qb.queryParts = append(qb.queryParts, ") ")
	return qb
}

func (qb *ConcreteSQLQueryBuilder) Where(condition string) SQLQueryBuilder {
	if condition == "" {
		return qb
	}
	qb.queryParts = append(qb.queryParts, "WHERE "+condition)
	return qb
}

func (qb *ConcreteSQLQueryBuilder) OrderBy(field string) SQLQueryBuilder {
	if field == "" {
		return qb
	}
	qb.queryParts = append(qb.queryParts, "ORDER BY "+field)
	return qb
}

func (qb *ConcreteSQLQueryBuilder) Like(word string) SQLQueryBuilder {
	if word == "" {
		return qb
	}
	qb.queryParts = append(qb.queryParts, "ALIKE '%"+word+"%'")
	return qb
}

func (qb *ConcreteSQLQueryBuilder) Build() string {
	query := strings.Join(qb.queryParts, " ")
	qb.queryParts = nil
	return query
}

func main() {
	queryBuilder := NewSQLQueryBuilder()

	// INSERT INTO user_balance (user_id, currency, balance) VALUES ($1, $2, $3);
	query := queryBuilder.
		Insert("user_balance", "user_id", "currency", "balance").
		Values("$1", "$2", "$3").
		Build()

	fmt.Println(query)

	subquery := queryBuilder.
		Select("card_id").
		From("user_cards").
		Where("user_id = $1").
		Build()

	query = queryBuilder.
		Select("*").
		From("card_transaction").
		Where("fromcard_id").
		In(subquery).
		Build()

}
