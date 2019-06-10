package meteomatics

// A FormatString is a string that represents a format.
type FormatString string

// A FormatStringer is something that can behave as a format
type FormatStringer interface {
	ContentType() string
	FormatString() FormatString
}

// A Format is a format.
type Format struct {
	formatString FormatString
	contentType  string
}

// Formats.
//nolint:gochecknoglobals
var (
	FormatGrads = Format{
		formatString: "grads",
		contentType:  "application/grads", // FIXME check content type
	}
	FormatCSV = Format{
		formatString: "csv",
		contentType:  "text/csv",
	}
	FormatHTML = Format{
		formatString: "html",
		contentType:  "text/html",
	}
	FormatHTMLMap = Format{
		formatString: "html-map",
		contentType:  "text/html",
	}
	FormatJSON = Format{
		formatString: "json",
		contentType:  "application/json",
	}
	FormatNetCDF = Format{
		formatString: "netcdf",
		contentType:  "application/x-netcdf4", // FIXME check content type
	}
	FormatPNG = Format{
		formatString: "png",
		contentType:  "image/png",
	}
	FormatXML = Format{
		formatString: "xml",
		contentType:  "application/xml",
	}
)

// ContentType returns f's content type.
func (f Format) ContentType() string {
	return f.contentType
}

// FormatString returns f as a FormatString.
func (f Format) FormatString() FormatString {
	return f.formatString
}
