package strcase

<<<<<<< HEAD
var uppercaseAcronym = map[string]string{
	"ID": "id",
}

// ConfigureAcronym allows you to add additional words which will be considered acronyms
func ConfigureAcronym(key, val string) {
	uppercaseAcronym[key] = val
=======
import (
	"sync"
)

var uppercaseAcronym = sync.Map{}
	//"ID": "id",

// ConfigureAcronym allows you to add additional words which will be considered acronyms
func ConfigureAcronym(key, val string) {
	uppercaseAcronym.Store(key, val)
>>>>>>> main
}
