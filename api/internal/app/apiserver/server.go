package apiserver

import (
	"api/internal/app/model"
	store "api/internal/app/stores"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	resize "github.com/nfnt/resize"
	"github.com/sirupsen/logrus"
)

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionField sessions.Store
}

var (
	ErrorNotAuthenticate = errors.New("Not Authenticate")
	ErrorEmailorPasswd   = errors.New("email or password incorrect")
	NotFound             = errors.New("not found any user")
)

const (
	contextky ctxKEY = iota
	ctxUserRequest
)

const SessionName = "Google"
const CookieName = "JWT_token"

var SecretKey = []byte("secret")

type ctxKEY int8

func NewServer(str store.Store, session sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        str,
		sessionField: session,
	}

	s.NastroitRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) NastroitRouter() {
	s.router.Use(s.LogRequest)
	s.router.Use(s.SetRequestID)
	// s.router.Use(s.ResfreshToken)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	// s.router.HandleFunc("/", s.handleHello())
	s.router.HandleFunc("/create", s.HandleCreateUser()).Methods("POST")
	s.router.HandleFunc("/sessions", s.HandleUsersSession()).Methods("POST")
	s.router.HandleFunc("/getuser", s.GetAllUsers()).Methods("GET")
	s.router.HandleFunc("/getby_email", s.GetByEmail())
	s.router.HandleFunc("/img.jpg", s.HandleAwtorizasia()).Methods("GET")
	s.router.HandleFunc("/019_Desktop Wallpapers  HD Part (161).jpg", s.HandleAwtorizasiaSecond()).Methods("GET")

	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.Authentication)
	// private.Use(s.ResfreshToken)
	private.HandleFunc("/whoamI", s.WhoamI()).Methods("GET")
	private.HandleFunc("/update", s.UpdateUser()).Methods("POST")
}

func (s *server) GetAllUsers() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		u, err := s.store.Users().GetAll()
		if err != nil {
			s.error(w, r, http.StatusBadRequest, NotFound)
			return
		}

		snbs := make([]model.Users, 0)

		defer u.Close()

		for u.Next() {
			snb := model.Users{}
			err := u.Scan(&snb.ID, &snb.Name, &snb.Email, &snb.EncryptedPassword)

			if err != nil {
				s.Respond(w, r, http.StatusBadRequest, NotFound)
			}

			snbs = append(snbs, snb)
		}
		for _, u := range snbs {
			s.Respond(w, r, http.StatusOK, u)
		}
	}
}

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	}
}

func (s *server) HandleAwtorizasia() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request){
		var path = "./static/img.jpg"
		c, err := r.Cookie(CookieName)

		if err != nil {
			if err == http.ErrNoCookie {
				s.error(w, r, http.StatusUnauthorized, err)
				return
			}

			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		cookie := c.Value

		claims := &jwt.StandardClaims{}

		token, err := jwt.ParseWithClaims(cookie, claims, func(t *jwt.Token) (interface{}, error) {

			return SecretKey, nil
		})

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if !token.Valid {
			s.error(w, r, http.StatusUnauthorized, errors.New("token is not valid !!!!!"))
			return
		}
		
		id, err := strconv.Atoi(claims.Issuer)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		u, err := s.store.Users().Find(id)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, ErrorNotAuthenticate)
			return
		}

		http.SetCookie(w, c)
		
		if u.Email == "allayar@gmail.com" {
		
		Razmer(w, path, 500)
		img , err := os.Open(path)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		defer img.Close()

		io.Copy(w, img)

		} else {
			s.Respond(w, r, http.StatusOK, "bu sahypa siz ucin dal")
		}
	}
}

func (s *server) HandleAwtorizasiaSecond() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request){

		var Path = "./static/019_Desktop Wallpapers  HD Part (161).jpg"

		c, err := r.Cookie(CookieName)

		if err != nil {
			if err == http.ErrNoCookie {
				s.error(w, r, http.StatusUnauthorized, err)
				return
			}

			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		cookie := c.Value

		claims := &jwt.StandardClaims{}

		token, err := jwt.ParseWithClaims(cookie, claims, func(t *jwt.Token) (interface{}, error) {

			return SecretKey, nil
		})

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if !token.Valid {
			s.error(w, r, http.StatusUnauthorized, errors.New("token is not valid !!!!!"))
			return
		}
		
		id, err := strconv.Atoi(claims.Issuer)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		u, err := s.store.Users().Find(id)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, ErrorNotAuthenticate)
			return
		}
		
		http.SetCookie(w, c)

		if u.Email == "yakup@gmai.com" {
		Razmer(w, Path, 500)

		img, err := os.Open(Path)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		defer img.Close()

		io.Copy(w, img)
			} else {
			s.Respond(w, r, http.StatusOK, "bu sahypa siz ucin dal")
		}

	}
}

func (s *server) Authentication(p http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// session, err := s.sessionField.Get(r, SessionName)

		// if err != nil {
		// 	s.error(w, r, http.StatusInternalServerError, err)
		// 	return
		// }

		c, err := r.Cookie(CookieName)

		if err != nil {
			if err == http.ErrNoCookie {
				s.error(w, r, http.StatusUnauthorized, err)
				return
			}

			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		cookie := c.Value

		claims := &jwt.StandardClaims{}

		token, err := jwt.ParseWithClaims(cookie, claims, func(t *jwt.Token) (interface{}, error) {

			return SecretKey, nil
		})

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if !token.Valid {
			s.error(w, r, http.StatusUnauthorized, errors.New("token is not valid !!!!!"))
			return
		}

		id, err := strconv.Atoi(claims.Issuer)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		u, err := s.store.Users().Find(id)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, ErrorNotAuthenticate)
			return
		}

		p.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), contextky, u)))
	})
}

func (s *server) ResfreshToken(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(CookieName)
		if err != nil {
			if err == http.ErrNoCookie {
				s.error(w, r, http.StatusUnauthorized, errors.New("yalnyslyk su yerde"))
				return
			}
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		cookie := c.Value

		claims := &jwt.StandardClaims{}

		token, err := jwt.ParseWithClaims(cookie, claims, func(t *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				s.error(w, r, http.StatusUnauthorized, err)
			}
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if !token.Valid {
			s.error(w, r, http.StatusUnauthorized, errors.New("User not found"))
			return
		}

		// if time.Unix(claims.ExpiresAt,0).Sub(time.Now()) > 30 * time.Second{
		// 	s.error(w, r, http.StatusBadRequest, errors.New("not enough time to expire of token"))
		// }
		if claims.ExpiresAt == 0 {
			expirationtime := time.Now().Add(24*time.Hour)
			claims.ExpiresAt = expirationtime.Unix()

			tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			tokenString, err := tkn.SignedString([]byte(SecretKey))
			if err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
			}

			http.SetCookie(w, &http.Cookie{
				Name: "JWT_token",
				Value:    tokenString,
				Expires:  expirationtime,
				HttpOnly: true,

			})
		}
	})

}

func (s *server) WhoamI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Respond(w, r, http.StatusOK, r.Context().Value(contextky).(*model.Users))
	}
}

func (s *server) HandleCreateUser() http.HandlerFunc {
	type Request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.Users{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.store.Users().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.ParolGizle()

		s.Respond(w, r, http.StatusCreated, u)

	}
}

func (s *server) HandleUsersSession() http.HandlerFunc {
	type Request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.Users().FindBYemail(req.Email)
		if err != nil || !u.ComparePassWord(req.Password) {
			s.error(w, r, http.StatusUnauthorized, ErrorEmailorPasswd)
			return
		}

		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer: strconv.Itoa(int(u.ID)),
			// Subject: strconv.Itoa(u.ID),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		})
		token, err := claims.SignedString([]byte(SecretKey))

		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
		}

		cookie := http.Cookie{
			Name:     "JWT_token",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24),
			HttpOnly: true,
		}

		http.SetCookie(w, &cookie)

		// session, err := s.sessionField.Get(r, SessionName)
		// if err != nil {
		// 	s.error(w, r, http.StatusInternalServerError, err)
		// 	return
		// }
		// session.Values["user_id"] = u.ID
		// if err := s.sessionField.Save(r, w, session); err != nil {
		// 	s.error(w, r, http.StatusInternalServerError, err)
		// 	return
		// }

		s.Respond(w, r, http.StatusOK, token)
	}
}

func (s *server) GetByEmail() http.HandlerFunc {
	type Request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.Users().FindBYemail(req.Email)
		if err != nil {
			s.Respond(w, r, http.StatusInternalServerError, "not Found")
			return
		}

		s.Respond(w, r, http.StatusOK, u)
	}
}

func (s *server) UpdateUser() http.HandlerFunc {
	type Request struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u, err := s.store.Users().Update(req.Name, req.Email, req.ID)
		fmt.Println(u)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.Respond(w, r, http.StatusOK, u)
	}
}

func (s *server) SetRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxUserRequest, id)))
	})
}

func (s *server) LogRequest(param http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remoute_adress": r.RemoteAddr,
			"request_id":     r.Context().Value(ctxUserRequest),
		})
		logger.Infof("started %v %v", r.Method, r.RequestURI)

		start := time.Now()

		rw := &REsponseWriter{w, http.StatusOK}
		param.ServeHTTP(rw, r)
		logger.Infof("completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.Respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func Razmer(w io.Writer, path string, size uint) {
	var ImgExtn = strings.Split(path, ".jpg")
	var Imgnum = strings.Split(ImgExtn[0], "/")
	var ImgNam = Imgnum[len(Imgnum) - 1]
	fmt.Println(ImgNam)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	IMG, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	file.Close()
	a := resize.Resize(size, 0, IMG, resize.Lanczos3)

	jpeg.Encode(w, a, nil)
}