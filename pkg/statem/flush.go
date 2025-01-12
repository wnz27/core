/*
Copyright 2021 The tKeel Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package statem

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/tkeel-io/core/pkg/constraint"
	"github.com/tkeel-io/core/pkg/logger"
	"github.com/tkeel-io/kit/log"
	"go.uber.org/zap"
)

func (s *statem) Flush(ctx context.Context) error {
	return s.flush(ctx)
}

func (s *statem) FlushState() error {
	return errors.Wrap(s.flushState(s.ctx), "flush state-marchine state")
}

func (s *statem) FlushSearch() error {
	return errors.Wrap(s.flushSeatch(s.ctx), "flush state-marchine state")
}

func (s *statem) FlushTimeSeries() error {
	panic("not implement")
}

func (s *statem) flush(ctx context.Context) error {
	var err error
	// flush state properties to es.
	if err = s.flushSeatch(ctx); nil == err {
		log.Debug("entity flush Search completed", logger.EntityID(s.ID))
	}
	// flush state properties to state.
	if err = s.flushState(ctx); nil == err {
		log.Debug("entity flush State completed", logger.EntityID(s.ID))
	}
	return errors.Wrap(err, "entity flush data failed")
}

func (s *statem) flushState(ctx context.Context) error {
	bytes, _ := EncodeBase(&s.Base)
	s.stateManager.GetDaprClient().SaveState(ctx, "core-state", s.ID, bytes)
	return nil
}

func (s *statem) flushSeatch(ctx context.Context) error {
	var err error
	var flushData = make(map[string]interface{})
	for _, JSONPath := range s.searchConstraints {
		var val constraint.Node
		var ct *constraint.Constraint
		if val, err = s.getProperty(s.KValues, JSONPath); nil != err {
		} else if ct, err = s.getConstraint(JSONPath); nil != err {
		} else if val, err = constraint.ExecData(val, ct); nil != err {
		} else {
			flushData[JSONPath] = val.Value()
		}
		log.Error("patch.copy entity property failed", logger.EntityID(s.ID), zap.String("property_key", JSONPath), zap.Error(err))
	}

	// flush all.
	for key, val := range s.KValues {
		flushData[key] = val.String()
	}

	// basic fields.
	flushData["id"] = s.ID
	flushData["type"] = s.Type
	flushData["owner"] = s.Owner
	flushData["source"] = s.Source
	flushData["version"] = s.Version
	flushData["last_time"] = s.LastTime
	if err = s.stateManager.SearchFlush(ctx, flushData); nil != err {
		log.Error("flush state Search.", zap.Any("data", flushData), zap.Error(err))
	}

	log.Debug("flush state Search.", zap.Any("data", flushData))
	return errors.Wrap(err, "Search flush failed")
}

func (s *statem) getConstraint(jsonPath string) (*constraint.Constraint, error) {
	arr := strings.Split(jsonPath, ".")
	if len(arr) == 0 {
		return nil, errInvalidJSONPath
	} else if len(arr) == 1 {
		return s.constraints[arr[0]], nil
	}

	var ct *constraint.Constraint
	if ct = s.constraints[arr[0]]; nil != ct {
		return nil, nil
	}

	var index int
	for indx, key := range arr[1:] {
		var nextCt *constraint.Constraint
		for _, childCt := range ct.ChildNodes {
			if key == childCt.ID {
				nextCt, index = childCt, indx+1
				break
			}
		}
		if nextCt == nil {
			break
		}
		ct = nextCt
	}

	if index != len(arr)-1 {
		return nil, nil
	}

	return ct, nil
}
