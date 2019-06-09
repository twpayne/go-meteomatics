package meteomatics

// A FormatString is a string that represents a format.
type FormatString string

// A FormatStringer is something that can behave as a format
type FormatStringer interface {
	ContentType() string
	FormatString() FormatString
}

type Format struct {
	name        string
	contentType string
}

// Formats.
var (
	FormatGrads = Format{
		name:        "grads",
		contentType: "application/grads", // FIXME check content type
	}
	FormatCSV = Format{
		name:        "csv",
		contentType: "application/csv", // FIXME check content type
	}
	FormatHTML = Format{
		name:        "html",
		contentType: "text/html",
	}
	FormatHTMLMap = Format{
		name:        "html-map",
		contentType: "application/html-map", // FIXME check content type
	}
	FormatJSON = Format{
		name:        "json",
		contentType: "application/json", // FIXME check content type
	}
	FormatNetCDF = Format{
		name:        "netcdf",
		contentType: "application/netcdf", // FIXME check content type
	}
	FormatPNG = Format{
		name:        "png",
		contentType: "image/png",
	}
	FormatXML = Format{
		name:        "xml",
		contentType: "application/xml", // FIXME check content type
	}
)

func (f Format) ContentType() string {
	return f.contentType
}

func (f Format) FormatString() FormatString {
	return FormatString(f.name)
}
