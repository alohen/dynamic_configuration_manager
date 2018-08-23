package configuration

type ParsingError struct {
	err error
}

func NewParsingError(err error) ParsingError {
	return ParsingError{
		err: err,
	}
}

func (err ParsingError) Error() string {
	return err.Error()
}

type EditingError struct {
	err error
}

func NeweditingError(err error) EditingError {
	return EditingError{
		err: err,
	}
}

func (err EditingError) Error() string {
	return err.Error()
}

type PageBuildingError struct {
	err error
}

func NewPageBuildingError(err error) PageBuildingError {
	return PageBuildingError{
		err: err,
	}
}

func (err PageBuildingError) Error() string {
	return err.Error()
}
