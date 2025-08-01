package middlewares

import (
	dbConfig "AuthInGo/config/db"
	env "AuthInGo/config/env"
	repo "AuthInGo/db/repositories"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func JWTAuthMiddleware(nex http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		fmt.Println("Authorization header:", authHeader)
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Authorization header must start with Bearer", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			http.Error(w, "Token is required", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.GetString("JWT_SECRET", "TOKEN")), nil
		})

		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		userId, okId := claims["id"].(float64)

		email, okEmail := claims["email"].(string)

		if !okId || !okEmail {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		fmt.Println("Authenticated user ID:", int64(userId), "Email:", email)

		ctx := context.WithValue(r.Context(), "userID", strconv.FormatFloat(userId, 'f', 0, 64))
		ctx = context.WithValue(ctx, "email", email)
		nex.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequireAllRoles(roles ...string) func(http.Handler) http.Handler {
	// function that can create a middleware for checking the above set of roles
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userIdStr := r.Context().Value("userID").(string)
			userId, err := strconv.ParseInt(userIdStr, 10, 64)
			if err != nil {
				http.Error(w, "Invalid user ID", http.StatusUnauthorized)
				return
			}

			dbConn, dbErr := dbConfig.SetupDB()
			if dbErr != nil {
				http.Error(w, "Database connection error: "+dbErr.Error(), http.StatusInternalServerError)
				return
			}

			urr := repo.NewUserRoleRepository(dbConn)

			hasAllRoles, hasAllRolesErr := urr.HasAllRoles(userId, roles)
			if hasAllRolesErr != nil {
				http.Error(w, "Error checking user roles:"+hasAllRolesErr.Error(), http.StatusInternalServerError)
				return
			}

			if !hasAllRoles {
				http.Error(w, "Forbidden: You do not have the required roles", http.StatusForbidden)
				return
			}

			fmt.Println("User has all required roles: ", roles)
			next.ServeHTTP(w, r)
		})
	}
}

func RequiredAnyRole(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userIdStr := r.Context().Value("userID").(string)
			userId, err := strconv.ParseInt(userIdStr, 10, 64)
			if err != nil {
				http.Error(w, "Invalid user ID", http.StatusUnauthorized)
				return
			}

			dbConn, dbErr := dbConfig.SetupDB()
			if dbErr != nil {
				http.Error(w, "Database connection error:"+dbErr.Error(), http.StatusInternalServerError)
				return
			}

			urr := repo.NewUserRoleRepository(dbConn)

			hasAnyRole, hasAnyRoleErr := urr.HasAnyRole(userId, roles)
			fmt.Println("userid", userId, "roles", roles, "hasAnyRole", hasAnyRole)

			if hasAnyRoleErr != nil {
				http.Error(w, "Error checking user roles:"+hasAnyRoleErr.Error(), http.StatusInternalServerError)
				return
			}

			if !hasAnyRole {
				http.Error(w, "Forbidden: You do not have the required roles", http.StatusForbidden)
				return
			}

			fmt.Println("User has all required roles:", roles)
			next.ServeHTTP(w, r)
		})
	}
}
