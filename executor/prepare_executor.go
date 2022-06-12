/*
 * Copyright (c) 2022, OpeningO
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package executor

import (
	"context"
	"github.com/xfali/gobatis/common"
	"github.com/xfali/gobatis/errors"
	"github.com/xfali/gobatis/reflection"
	"github.com/xfali/gobatis/transaction"
)

type PrepareExecutor struct {
	transaction transaction.Transaction
	closed      bool
}

func NewPrepareExecutor(transaction transaction.Transaction) *PrepareExecutor {
	return &PrepareExecutor{transaction: transaction}
}

func (exec *PrepareExecutor) Close(rollback bool) {
	defer func() {
		if exec.transaction != nil {
			exec.transaction.Close()
		}
		exec.transaction = nil
		exec.closed = true
	}()

	if rollback {
		exec.Rollback(true)
	}
}

func (exec *PrepareExecutor) Query(ctx context.Context, result reflection.Object, sql string, params ...interface{}) error {
	if exec.closed {
		return errors.EXECUTOR_QUERY_ERROR
	}

	conn := exec.transaction.GetConnection()
	if conn == nil {
		return errors.EXECUTOR_GET_CONNECTION_ERROR
	}

	//FIXME: stmt must be close, use stmtCache instead
	stmt, err := conn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	return stmt.Query(ctx, result, params...)
}

func (exec *PrepareExecutor) Exec(ctx context.Context, sql string, params ...interface{}) (common.Result, error) {
	if exec.closed {
		return nil, errors.EXECUTOR_QUERY_ERROR
	}

	conn := exec.transaction.GetConnection()
	if conn == nil {
		return nil, errors.EXECUTOR_GET_CONNECTION_ERROR
	}

	//FIXME: stmt must be close, use stmtCache instead
	stmt, err := conn.Prepare(sql)
	defer stmt.Close()

	if err != nil {
		return nil, err
	}
	return stmt.Exec(ctx, params...)
}

func (exec *PrepareExecutor) Begin() error {
	if exec.closed {
		return errors.EXECUTOR_BEGIN_ERROR
	}

	return exec.transaction.Begin()
}

func (exec *PrepareExecutor) Commit(require bool) error {
	if exec.closed {
		return errors.EXECUTOR_COMMIT_ERROR
	}

	if require {
		return exec.transaction.Commit()
	}

	return nil
}

func (exec *PrepareExecutor) Rollback(require bool) error {
	if !exec.closed {
		if require {
			return exec.transaction.Rollback()
		}
	}
	return nil
}
