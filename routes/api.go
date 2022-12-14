package routes

import (
	"go-alqolam/auth"
	"go-alqolam/handler"
	"go-alqolam/helper"
	"go-alqolam/member"
	"go-alqolam/region"
	transferWallet "go-alqolam/transfer_wallet"
	"go-alqolam/user"
	"go-alqolam/wallet"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	// Repository
	userRepository := user.NewRepository(db)
	regionRepository := region.NewRepository(db)
	memberRepository := member.NewRepository(db)
	walletRepository := wallet.NewRepository(db)
	transferWalletRepository := transferWallet.NewRepository(db)

	// Service
	authService := auth.NewService()
	userService := user.NewService(userRepository)
	regionService := region.NewService(regionRepository)
	memberService := member.NewService(memberRepository)
	walletService := wallet.NewService(walletRepository)
	transferWalletService := transferWallet.NewService(transferWalletRepository)

	// Handler
	authHandler := handler.NewAuthHandler(userService, authService)
	userHandler := handler.NewUserHandler(userService)
	regionHandler := handler.NewRegionHandler(regionService)
	memberHandler := handler.NewMemberHandler(memberService)
	walletHandler := handler.NewWalletHandler(walletService)
	transferWalletHandler := handler.NewTransferWalletHandler(transferWalletService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1/")
	api.POST("register", authHandler.RegisterUser)
	api.POST("check-email", authHandler.CheckEmailAvailability)
	api.POST("login", authHandler.Login)

	// User
	api.GET("users", authMiddleware(authService, userService), userHandler.Index)
	api.GET("user/:id", authMiddleware(authService, userService), userHandler.Show)
	api.PUT("user/:id", authMiddleware(authService, userService), userHandler.Update)
	api.DELETE("user/:id", authMiddleware(authService, userService), userHandler.Destroy)

	// Region
	api.GET("regions", authMiddleware(authService, userService), regionHandler.Index)
	api.POST("region", authMiddleware(authService, userService), regionHandler.Store)
	api.GET("region/:id", authMiddleware(authService, userService), regionHandler.Show)
	api.PUT("region/:id", authMiddleware(authService, userService), regionHandler.Update)
	api.DELETE("region/:id", authMiddleware(authService, userService), regionHandler.Destroy)

	// Member
	api.GET("members", authMiddleware(authService, userService), memberHandler.Index)
	api.POST("member", authMiddleware(authService, userService), memberHandler.Store)
	api.GET("member/:id", authMiddleware(authService, userService), memberHandler.Show)
	api.PUT("member/:id", authMiddleware(authService, userService), memberHandler.Update)
	api.DELETE("member/:id", authMiddleware(authService, userService), memberHandler.Destroy)

	// Wallet
	api.GET("wallets", authMiddleware(authService, userService), walletHandler.Index)
	api.POST("wallet", authMiddleware(authService, userService), walletHandler.Store)
	api.GET("wallet/:id", authMiddleware(authService, userService), walletHandler.Show)
	api.PUT("wallet/:id", authMiddleware(authService, userService), walletHandler.Update)
	api.DELETE("wallet/:id", authMiddleware(authService, userService), walletHandler.Destroy)

	// Transfer Wallets
	api.GET("transfer-wallets", authMiddleware(authService, userService), transferWalletHandler.Index)
	api.POST("transfer-wallet", authMiddleware(authService, userService), transferWalletHandler.Store)
	api.GET("transfer-wallet/:id", authMiddleware(authService, userService), transferWalletHandler.Show)
	api.PUT("transfer-wallet/:id", authMiddleware(authService, userService), transferWalletHandler.Update)
	api.DELETE("transfer-wallet/:id", authMiddleware(authService, userService), transferWalletHandler.Destroy)

	return router
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.FindUser(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
