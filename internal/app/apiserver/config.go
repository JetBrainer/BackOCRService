package apiserver

import "github.com/caarlos0/env/v6"

type Config struct {
	ApiKey   string `env:"API_KEY" envDefault:"7f628fab5f88957"`
	Language string `env:"LANGUAGE" envDefault:"rus"`
	Url      string `env:"URL" envDefault:"https://api.ocr.space/parse/image"`
	DBUrl    string `env:"DATABASE_URL" envDefault:"postgres://oedbfnojeglfik:945a538bc9cf38a29dad405949bc8694ff7b6b39d59f599d14befa304066635c@ec2-52-209-134-160.eu-west-1.compute.amazonaws.com:5432/d5i7mb4hehn571"`
	HttpPort string `env:"PORT" envDefault:"8081"`
	GRPCPort string `env:"grpc_port" envDefault:":4000"`
	JWTKey   string `env:"jwt_key" envDefault:"somatic-key"`
}

// A list of document values returns in the response
// swagger:response docStrRepoResp
//type swaggerDocStrRepoResp struct {
//	// All value is document
//	// in: body
//	DocStr
//}

func InitConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
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
