package consumption

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// OperatorType enumerates the values for operator type.
type OperatorType string

const (
	// EqualTo ...
	EqualTo OperatorType = "EqualTo"
	// GreaterThan ...
	GreaterThan OperatorType = "GreaterThan"
	// GreaterThanOrEqualTo ...
	GreaterThanOrEqualTo OperatorType = "GreaterThanOrEqualTo"
)

// PossibleOperatorTypeValues returns an array of possible values for the OperatorType const type.
func PossibleOperatorTypeValues() []OperatorType {
	return []OperatorType{EqualTo, GreaterThan, GreaterThanOrEqualTo}
}

// TimeGrainType enumerates the values for time grain type.
type TimeGrainType string

const (
	// Annually ...
	Annually TimeGrainType = "Annually"
	// Monthly ...
	Monthly TimeGrainType = "Monthly"
	// Quarterly ...
	Quarterly TimeGrainType = "Quarterly"
)

// PossibleTimeGrainTypeValues returns an array of possible values for the TimeGrainType const type.
func PossibleTimeGrainTypeValues() []TimeGrainType {
	return []TimeGrainType{Annually, Monthly, Quarterly}
}
