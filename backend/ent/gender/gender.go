// Code generated by entc, DO NOT EDIT.

package gender

const (
	// Label holds the string label denoting the gender type in the database.
	Label = "gender"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGenderName holds the string denoting the gendername field in the database.
	FieldGenderName = "gender_name"

	// EdgeCustomers holds the string denoting the customers edge name in mutations.
	EdgeCustomers = "customers"

	// Table holds the table name of the gender in the database.
	Table = "genders"
	// CustomersTable is the table the holds the customers relation/edge.
	CustomersTable = "customers"
	// CustomersInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomersInverseTable = "customers"
	// CustomersColumn is the table column denoting the customers relation/edge.
	CustomersColumn = "Gender_ID"
)

// Columns holds all SQL columns for gender fields.
var Columns = []string{
	FieldID,
	FieldGenderName,
}
