// Code generated by ent, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldUserName holds the string denoting the user_name field in the database.
	FieldUserName = "user_name"
	// FieldOptionalNullableBool holds the string denoting the optional_nullable_bool field in the database.
	FieldOptionalNullableBool = "optional_nullable_bool"
	// EdgeRequiredCity holds the string denoting the required_city edge name in mutations.
	EdgeRequiredCity = "required_city"
	// EdgeOptionalCity holds the string denoting the optional_city edge name in mutations.
	EdgeOptionalCity = "optional_city"
	// EdgeFriendList holds the string denoting the friend_list edge name in mutations.
	EdgeFriendList = "friend_list"
	// Table holds the table name of the user in the database.
	Table = "users"
	// RequiredCityTable is the table that holds the required_city relation/edge.
	RequiredCityTable = "users"
	// RequiredCityInverseTable is the table name for the City entity.
	// It exists in this package in order to avoid circular dependency with the "city" package.
	RequiredCityInverseTable = "cities"
	// RequiredCityColumn is the table column denoting the required_city relation/edge.
	RequiredCityColumn = "user_required_city"
	// OptionalCityTable is the table that holds the optional_city relation/edge.
	OptionalCityTable = "users"
	// OptionalCityInverseTable is the table name for the City entity.
	// It exists in this package in order to avoid circular dependency with the "city" package.
	OptionalCityInverseTable = "cities"
	// OptionalCityColumn is the table column denoting the optional_city relation/edge.
	OptionalCityColumn = "user_optional_city"
	// FriendListTable is the table that holds the friend_list relation/edge. The primary key declared below.
	FriendListTable = "user_friend_list"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldFirstName,
	FieldLastName,
	FieldUserName,
	FieldOptionalNullableBool,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_required_city",
	"user_optional_city",
}

var (
	// FriendListPrimaryKey and FriendListColumn2 are the table columns denoting the
	// primary key for the friend_list relation (M2M).
	FriendListPrimaryKey = []string{"user_id", "friend_list_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
)
