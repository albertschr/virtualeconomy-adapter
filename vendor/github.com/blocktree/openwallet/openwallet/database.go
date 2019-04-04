/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwallet

import (
	"github.com/asdine/storm"
	"fmt"
)

type StormDB struct {
	*storm.DB
	FileName string
	Opened bool
}

//OpenStormDB
func OpenStormDB(filename string, stormOptions ...func(*storm.Options) error) (*StormDB, error) {

	db, err := storm.Open(filename, stormOptions...)
	//fmt.Println("open app db")
	if err != nil {
		return nil, fmt.Errorf("can not open dbfile: '%s', unexpected error: %v", filename, err)
	}

	// Check the metadata.
	stormDB := &StormDB{
		FileName: filename,
		DB:       db,
		Opened: true,
	}

	return stormDB, nil
}

// Close closes the database.
func (db *StormDB) Close() error {
	err := db.DB.Close()
	if err != nil {
		return err
	}
	db.Opened = false
	return nil
}

