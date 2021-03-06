// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/entc/integration/ent/schema"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOptionalInt holds the string denoting the optional_int vertex property in the database.
	FieldOptionalInt = "optional_int"
	// FieldAge holds the string denoting the age vertex property in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldLast holds the string denoting the last vertex property in the database.
	FieldLast = "last"
	// FieldNickname holds the string denoting the nickname vertex property in the database.
	FieldNickname = "nickname"
	// FieldPhone holds the string denoting the phone vertex property in the database.
	FieldPhone = "phone"
	// FieldPassword holds the string denoting the password vertex property in the database.
	FieldPassword = "password"
	// FieldRole holds the string denoting the role vertex property in the database.
	FieldRole = "role"

	// Table holds the table name of the user in the database.
	Table = "users"
	// CardTable is the table the holds the card relation/edge.
	CardTable = "cards"
	// CardInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardInverseTable = "cards"
	// CardColumn is the table column denoting the card relation/edge.
	CardColumn = "owner_id"
	// PetsTable is the table the holds the pets relation/edge.
	PetsTable = "pets"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pets"
	// PetsColumn is the table column denoting the pets relation/edge.
	PetsColumn = "owner_id"
	// FilesTable is the table the holds the files relation/edge.
	FilesTable = "files"
	// FilesInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FilesInverseTable = "files"
	// FilesColumn is the table column denoting the files relation/edge.
	FilesColumn = "owner_id"
	// GroupsTable is the table the holds the groups relation/edge. The primary key declared below.
	GroupsTable = "user_groups"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// FriendsTable is the table the holds the friends relation/edge. The primary key declared below.
	FriendsTable = "user_friends"
	// FollowersTable is the table the holds the followers relation/edge. The primary key declared below.
	FollowersTable = "user_following"
	// FollowingTable is the table the holds the following relation/edge. The primary key declared below.
	FollowingTable = "user_following"
	// TeamTable is the table the holds the team relation/edge.
	TeamTable = "pets"
	// TeamInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	TeamInverseTable = "pets"
	// TeamColumn is the table column denoting the team relation/edge.
	TeamColumn = "team_id"
	// SpouseTable is the table the holds the spouse relation/edge.
	SpouseTable = "users"
	// SpouseColumn is the table column denoting the spouse relation/edge.
	SpouseColumn = "user_spouse_id"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "users"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_id"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "users"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldOptionalInt,
	FieldAge,
	FieldName,
	FieldLast,
	FieldNickname,
	FieldPhone,
	FieldPassword,
	FieldRole,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the User type.
var ForeignKeys = []string{
	"group_blocked_id",
	"user_spouse_id",
	"parent_id",
}

var (
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"user_id", "group_id"}
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"user_id", "friend_id"}
	// FollowersPrimaryKey and FollowersColumn2 are the table columns denoting the
	// primary key for the followers relation (M2M).
	FollowersPrimaryKey = []string{"user_id", "follower_id"}
	// FollowingPrimaryKey and FollowingColumn2 are the table columns denoting the
	// primary key for the following relation (M2M).
	FollowingPrimaryKey = []string{"user_id", "follower_id"}
)

var (
	mixin       = schema.User{}.Mixin()
	mixinFields = [...][]ent.Field{
		mixin[0].Fields(),
	}
	fields = schema.User{}.Fields()

	// descOptionalInt is the schema descriptor for optional_int field.
	descOptionalInt = mixinFields[0][0].Descriptor()
	// OptionalIntValidator is a validator for the "optional_int" field. It is called by the builders before save.
	OptionalIntValidator = descOptionalInt.Validators[0].(func(int) error)

	// descLast is the schema descriptor for last field.
	descLast = fields[2].Descriptor()
	// DefaultLast holds the default value on creation for the last field.
	DefaultLast = descLast.Default.(string)
)

// Role defines the type for the role enum field.
type Role string

// RoleUser is the default Role.
const DefaultRole = RoleUser

// Role values.
const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

func (s Role) String() string {
	return string(s)
}

// RoleValidator is a validator for the "r" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleUser, RoleAdmin:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}
