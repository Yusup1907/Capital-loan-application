package middleware

// import (
// 	"context"
// 	"errors"
// 	"net/http"
// 	"strings"

// 	"pinjam-modal-app/model"
// 	"pinjam-modal-app/repository"

// 	"github.com/dgrijalva/jwt-go"
// )

// var userRepo = repository.UserRepository{}

// func AuthenticationMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Periksa token dari header Authorization
// 		authorizationHeader := r.Header.Get("Authorization")
// 		if authorizationHeader == "" {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		// Misalnya, format header Authorization: Bearer <token>
// 		splitToken := strings.Split(authorizationHeader, "Bearer ")
// 		if len(splitToken) != 2 {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		tokenString := splitToken[1]

// 		// Validasi token
// 		token, err := validateToken(tokenString)
// 		if err != nil {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		// Periksa apakah pengguna telah logout
// 		if isTokenLoggedOut(tokenString) {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		// Setel data pengguna pada konteks request
// 		user, err := getUserFromToken(token)
// 		if err != nil {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), "user", user)
// 		r = r.WithContext(ctx)

// 		next.ServeHTTP(w, r)
// 	})
// }

// func validateToken(tokenString string) (*jwt.Token, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Verifikasi metode tanda tangan token
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("Invalid token")
// 		}

// 		// Kembalikan kunci rahasia yang sama yang digunakan saat pembuatan token
// 		return []byte(repository.TokenKey), nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if !token.Valid {
// 		return nil, errors.New("Invalid token")
// 	}

// 	return token, nil
// }

// func isTokenLoggedOut(tokenString string) bool {
// 	// Implementasi logika untuk memeriksa apakah token telah logout
// 	// Misalnya, Anda dapat memeriksa apakah token ada dalam daftar token yang telah logout

// 	// Contoh sederhana: tidak ada tindakan yang dilakukan pada token
// 	return false
// }

// func getUserFromToken(token *jwt.Token) (*model.UserModel, error) {
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return nil, errors.New("Invalid token claims")
// 	}

// 	username, ok := claims["username"].(string)
// 	if !ok {
// 		return nil, errors.New("Invalid username in token claims")
// 	}

// 	user, err := userRepo.GetUserByUsername(username)
// 	if err != nil {
// 		return nil, errors.New("User not found")
// 	}

// 	return user, nil
// }
