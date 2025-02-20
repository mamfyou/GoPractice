package nilerr

type CustomError struct {
	Message string
}

func (cerr *CustomError) Error() string {
	return cerr.Message
}

func IsCustomErrNil(err error) bool {
	if err == nil {
		return true
	} 
	
	if cerr, ok := err.(*CustomError); ok {
		if cerr == nil {
			return true
		}
		return cerr.Message == ""
	}

	return false
}