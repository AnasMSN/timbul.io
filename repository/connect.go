package repository

type Database interface {
	QueryOne(database, collection, key, value string) interface{}
}
