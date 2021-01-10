// Code generated by entc, DO NOT EDIT.

package customer

const (
	// Label holds the string label denoting the customer type in the database.
	Label = "customer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCustomerName holds the string denoting the customername field in the database.
	FieldCustomerName = "customer_name"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldPhoneNumber holds the string denoting the phonenumber field in the database.
	FieldPhoneNumber = "phone_number"

	// EdgeGenders holds the string denoting the genders edge name in mutations.
	EdgeGenders = "genders"
	// EdgePersonals holds the string denoting the personals edge name in mutations.
	EdgePersonals = "personals"
	// EdgeTitles holds the string denoting the titles edge name in mutations.
	EdgeTitles = "titles"

	// Table holds the table name of the customer in the database.
	Table = "customers"
	// GendersTable is the table the holds the genders relation/edge.
	GendersTable = "customers"
	// GendersInverseTable is the table name for the Gender entity.
	// It exists in this package in order to avoid circular dependency with the "gender" package.
	GendersInverseTable = "genders"
	// GendersColumn is the table column denoting the genders relation/edge.
	GendersColumn = "Gender_ID"
	// PersonalsTable is the table the holds the personals relation/edge.
	PersonalsTable = "customers"
	// PersonalsInverseTable is the table name for the Personal entity.
	// It exists in this package in order to avoid circular dependency with the "personal" package.
	PersonalsInverseTable = "personals"
	// PersonalsColumn is the table column denoting the personals relation/edge.
	PersonalsColumn = "Personal_ID"
	// TitlesTable is the table the holds the titles relation/edge.
	TitlesTable = "customers"
	// TitlesInverseTable is the table name for the Title entity.
	// It exists in this package in order to avoid circular dependency with the "title" package.
	TitlesInverseTable = "titles"
	// TitlesColumn is the table column denoting the titles relation/edge.
	TitlesColumn = "Title_ID"
)

// Columns holds all SQL columns for customer fields.
var Columns = []string{
	FieldID,
	FieldCustomerName,
	FieldAddress,
	FieldPhoneNumber,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Customer type.
var ForeignKeys = []string{
	"Gender_ID",
	"Personal_ID",
	"Title_ID",
}
