package configservers

import "admin-go/data"

//StartUp config server db
func StartUp() {
	// Start a MongoDB session
	data.CreateDbSession()
}
