package main

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import (
	"database/sql"
	"log" // TODO change to traffic_monitor_golang/common/log
)

func preparePrivLevelStmt(db *sql.DB) (*sql.Stmt, error) {
	return db.Prepare("select r.priv_level from tm_user as u join role as r on u.role = r.id where u.username = $1")
}

func hasPrivLevel(privLevelStmt *sql.Stmt, user string, level int) bool {
	var privLevel int
	err := privLevelStmt.QueryRow(user).Scan(&privLevel)
	switch {
	case err == sql.ErrNoRows:
		log.Println("Error checking user " + user + " priv level: user not in database")
		return false
	case err != nil:
		log.Println("Error checking user " + user + " priv level: " + err.Error())
		return false
	default:
		return privLevel >= level
	}
}
