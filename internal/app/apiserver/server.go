package apiserver

import (
	"github.com/JetBrainer/BackOCRService/internal/app/model"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

type server struct {
	router  *mux.Router
	logger  zerolog.Logger
	config  *Config
	//doc		*model.DocStruct
}

func newServer(config *Config) *server{
	// Put Log Level to Debug
	logLevel :=  zerolog.InfoLevel
	logLevel = zerolog.DebugLevel
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Configure Router
	s := &server{
		router: mux.NewRouter(),
		logger: logger,
		config: config,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter,r *http.Request){
	s.router.ServeHTTP(w,r)
}

func (s *server) configureRouter(){
	s.router.HandleFunc("/", s.getDocHandler()).Methods(http.MethodPost)
}

func (s *server) getDocHandler() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse MultiPart File
		err := r.ParseMultipartForm(10 << 20)
		if err != nil{
			http.Error(w,err.Error(),http.StatusBadRequest)
		}

		jValue := &OCRText{}
		// Send Request to another Api and get text result
		err = s.config.ParseFromPost(r.Body, jValue)
		if err != nil{
			s.logger.Err(err).Msg("Error parsing from Local")
		}

		invAndData := model.InvNumAndData("(....|....(\\s|\\s\\s)..(\\s|\\s\\s).{4,5}.)(\\s|\\s\\s)N.(\\s|\\s\\s)[0-9]{2,4}(\\s|\\s\\s)..(\\s|\\s\\s)[1-3][1-9]((\\s|\\s\\s)\\D{3,8}|\\.(0|1)[1-2]\\.)(\\d{4}|\\d{2})")
		invAndData.Match(jValue.JustText())

		payer := model.Payer("(\\W{3}(\\s|\\s\\s)(«|\\\")\\W{1,}\\d+\\W+\\d+|(^\\W+:|^.{1,}:)(\\s|\\s\\s)\\W{1,}\\d{1,}\\s\\W{1,}\\d{1,}(\\s|\\s\\s)\\W+)")
		payer.Match(jValue.JustText())

		produce := model.Producer("(П.{1,10}|^.........:)\\\\s.*(”|“)")
		produce.Match(jValue.JustText())

		requis := model.Requisites("(?m)(^[Сс][Чч](,\\\\s|.\\\\s|\\\\s\\\\s).*\\\\sБ|^[Сс][Чч](,\\\\s|.\\\\s|\\\\s\\\\s).*\\\\s^[Сс][Чч](,\\\\s|.\\\\s|\\\\s\\\\s).*\\\\s.*\\\\s.*)")
		requis.Match(jValue.JustText())

		sumWTax := model.SumWithTax("....у(\\s|\\s\\s)(\\d\\,\\s|\\d\\'\\s|\\d\\s|\\d)\\d{3}\\,(\\d{2}|\\d{3}\\,|}|\\d{3}\\s)")
		sumWTax.Match(jValue.JustText())

		amount := model.Amount("[^\\\\d]{12}(\\\\s|\\\\s\\\\s)\\\\d(\\\\.|\\\\,)")
		amount.Match(jValue.JustText())

		follow := model.Followed("Ру.{1,}\\s\\W{1,}")
		follow.Match(jValue.JustText())

		fullSum := model.SumTax("Су\\W{1,}.*\\s.*\\s.*(\\s\\d.*)")
		fullSum.Match(jValue.JustText())

		prodN := model.ProdName("(?m)(^[Тт]о...(а|))\\s.*\\s.*\\s.*")
		prodN.Match(jValue.JustText())
	}
}