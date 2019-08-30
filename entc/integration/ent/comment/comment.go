// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package comment

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUniqueInt holds the string denoting the unique_int vertex property in the database.
	FieldUniqueInt = "unique_int"
	// FieldUniqueFloat holds the string denoting the unique_float vertex property in the database.
	FieldUniqueFloat = "unique_float"

	// Table holds the table name of the comment in the database.
	Table = "comments"
)

// Columns holds all SQL columns are comment fields.
var Columns = []string{
	FieldID,
	FieldUniqueInt,
	FieldUniqueFloat,
}
