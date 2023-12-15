/*
 * Copyright 2023 The RuleGo Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rulego

import (
	"github.com/rulego/rulego/api/types"
	"github.com/rulego/rulego/utils/json"
)

// JsonParser Json
type JsonParser struct {
}

func (p *JsonParser) DecodeRuleChain(config types.Config, dsl []byte) (types.Node, error) {
	rootRuleChainDef, err := ParserRuleChain(dsl)
	if err != nil {
		return nil, err
	}
	return InitRuleChainCtx(config, &rootRuleChainDef)
}
func (p *JsonParser) DecodeRuleNode(config types.Config, dsl []byte) (types.Node, error) {
	node, err := ParserRuleNode(dsl)
	if err != nil {
		return nil, err
	}
	return InitRuleNodeCtx(config, &node)
}
func (p *JsonParser) EncodeRuleChain(def interface{}) ([]byte, error) {
	// 缩进符为两个空格
	v, err := json.Marshal(def)
	if err != nil {
		return nil, err
	}
	return json.Format(v)
}
func (p *JsonParser) EncodeRuleNode(def interface{}) ([]byte, error) {
	// 缩进符为两个空格
	v, err := json.Marshal(def)
	if err != nil {
		return nil, err
	}
	return json.Format(v)
}
