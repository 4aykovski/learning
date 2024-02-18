package db

import "errors"

var (
	ErrCantCreateDatabase  = errors.New("can't create database")
	ErrCantPingDatabase    = errors.New("can't ping database")
	ErrCantPrepareDatabase = errors.New("can't prepare database")
	ErrURLNotFound         = errors.New("url not found")
	ErrUrlExists           = errors.New("url exists")
)
