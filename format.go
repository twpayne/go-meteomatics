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
		contentType:  "application/csv", // FIXME check content type
	}
	FormatHTML = Format{
		formatString: "html",
		contentType:  "text/html",
	}
	FormatHTMLMap = Format{
		formatString: "html-map",
		contentType:  "application/html-map", // FIXME check content type
	}
	FormatJSON = Format{
		formatString: "json",
		contentType:  "application/json", // FIXME check content type
	}
	FormatNetCDF = Format{
		formatString: "netcdf",
		contentType:  "application/netcdf", // FIXME check content type
	}
	FormatPNG = Format{
		formatString: "png",
		contentType:  "image/png",
	}
	FormatXML = Format{
		formatString: "xml",
		contentType:  "application/xml", // FIXME check content type
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
