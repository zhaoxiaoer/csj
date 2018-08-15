// schema structs
package schestr

type InsuranceInfo struct {
	Name        string `schema:"name"`
	EmployeeID  string `schema:"employeeID"`
	NumChildren int    `schema:"numChildren"`
	IsMarried   bool   `schema:"married"`
}
