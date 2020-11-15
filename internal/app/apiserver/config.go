package apiserver

type Config struct {
	ApiKey   string `toml:"apikey"`
	Language string `toml:"language"`
	Url		 string `toml:"url"`
	DBUrl	 string `toml:"database_url"`
	HttpPort string `toml:"port"`
}

// A list of document values returns in the response
// swagger:response docStrRepoResp
//type swaggerDocStrRepoResp struct {
//	// All value is document
//	// in: body
//	DocStr
//}



func InitConfig() *Config{
	return &Config{
		ApiKey:   "SomeKey",
		Language: "rus",
		Url:	  "https://api.ocr.space/parse/image",
	}
}

type OCRText struct {
	ParsedResults []struct {
		TextOverlay struct {
			Lines []struct {
				Words []struct {
					WordText string  `json:"WordText"`
					Left     float64 `json:"Left"`
					Top      float64 `json:"Top"`
					Height   float64 `json:"Height"`
					Width    float64 `json:"Width"`
				} `json:"Words"`

				MaxHeight float64 `json:"MaxHeight"`
				MinTop    float64 `json:"MinTop"`
			} `json:"Lines"`

			HasOverlay bool   `json:"HasOverlay"`
			Message    string `json:"Message"`
		} `json:"TextOverlay"`

		TextOrientation   string `json:"TextOrientation"`
		FileParseExitCode int    `json:"FileParseExitCode"`
		ParsedText        string `json:"ParsedText"`
		ErrorMessage      string `json:"ErrorMessage"`
		ErrorDetails      string `json:"ErrorDetails"`
	} `json:"ParsedResults"`

	OCRExitCode                  int      `json:"OCRExitCode"`
	IsErroredOnProcessing        bool     `json:"IsErroredOnProcessing"`
	ErrorMessage                 []string `json:"ErrorMessage"`
	ErrorDetails                 string   `json:"ErrorDetails"`
	ProcessingTimeInMilliseconds string   `json:"ProcessingTimeInMilliseconds"`
	SearchablePDFURL             string   `json:"SearchablePDFURL"`
}
