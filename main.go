package main

import (
	"fmt"
	"log"
	"mini-project-evermos/configs"
	"mini-project-evermos/exceptions"
	"mini-project-evermos/handlers"
	"mini-project-evermos/models/entities/migration"
	"mini-project-evermos/models/responder"
	"mini-project-evermos/repositories"
	"mini-project-evermos/services"
	"mini-project-evermos/utils" // Add this line
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Add this line before starting the server
	if err := utils.InitializeDirectories(); err != nil {
		log.Fatal("Failed to initialize directories:", err)
	}

	// Setup Configuration
	configuration := configs.New()

	// Setup Database
	database := configs.NewMysqlDatabase(configuration)

	// Debug database connection
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Test connection
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Show tables
	var tables []string
	database.Raw("SHOW TABLES").Scan(&tables)
	fmt.Println("Available tables:", tables)

	// Run migration
	migration.RunMigration(database)

	// Initialize repositories
	userRepository := repositories.NewUserRepository(database)
	authRepository := repositories.NewAuthRepository(database)
	addressRepository := repositories.NewAddressRepository(database)
	categoryRepository := repositories.NewCategoryRepository(database)
	storeRepository := repositories.NewStoreRepository(database)
	storePhotoRepository := repositories.NewStorePhotoRepository(database)
	productRepository := repositories.NewProductRepository(database)
	fotoProdukRepository := repositories.NewFotoProdukRepository(database)
	transactionRepository := repositories.NewTransactionRepository(database)
	productLogRepository := repositories.NewProductLogRepository(database)
	trxDetailRepo := repositories.NewTransactionDetailRepository(database)
	keranjangBelanjaRepository := repositories.NewKeranjangBelanjaRepository(database)
	wishlistRepo := repositories.NewWishlistRepository(database)
	productReviewRepository := repositories.NewProductReviewRepository(database)
	notificationRepository := repositories.NewNotificationRepository(database)
	promoRepository := repositories.NewProductPromoRepository(database)
	diskonProdukRepo := repositories.NewDiskonProdukRepository(database)
	orderRepository := repositories.NewOrderRepository(database)
	couponRepository := repositories.NewProductCouponRepository(database)

	// Initialize services
	regionService := services.NewRegionService()
	userService := services.NewUserService(&userRepository)
	authService := services.NewAuthService(&authRepository, &userRepository)
	addressService := services.NewAddressService(&addressRepository, &userRepository, &regionService)
	categoryService := services.NewCategoryService(&categoryRepository)
	storeService := services.NewStoreService(&storeRepository, &storePhotoRepository)
	storePhotoService := services.NewStorePhotoService(storePhotoRepository)
	productService := services.NewProductService(productRepository, storeRepository, categoryRepository)
	transactionService := services.NewTransactionService(
		&transactionRepository,
		&productRepository,
		&addressRepository,
		&userRepository, // Add user repository
		&regionService,
		&productLogRepository, // Add this
	)
	productLogService := services.NewProductLogService(&productLogRepository)
	fotoProdukService := services.NewFotoProdukService(fotoProdukRepository, productRepository)
	trxDetailService := services.NewTransactionDetailService(trxDetailRepo)
	keranjangBelanjaService := services.NewKeranjangBelanjaService(&keranjangBelanjaRepository)
	wishlistService := services.NewWishlistService(&wishlistRepo, &storeRepository, &productRepository)
	productReviewService := services.NewProductReviewService(productReviewRepository, storeRepository)
	notificationService := services.NewNotificationService(notificationRepository)
	promoService := services.NewProductPromoService(promoRepository)
	diskonProdukService := services.NewDiskonProdukService(diskonProdukRepo)
	orderService := services.NewOrderService(orderRepository, trxDetailRepo)
	couponService := services.NewProductCouponService(couponRepository)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(&userService)
	authHandler := handlers.NewAuthHandler(&authService)
	addressHandler := handlers.NewAddressHandler(&addressService)
	regionHandler := handlers.NewRegionHandler(&regionService)
	categoryHandler := handlers.NewCategoryHandler(&categoryService)
	storeHandler := handlers.NewStoreHandler(&storeService)
	storePhotoHandler := handlers.NewStorePhotoHandler(storePhotoService)
	productHandler := handlers.NewProductHandler(&productService)
	transactionHandler := handlers.NewTransactionHandler(&transactionService)
	productLogHandler := handlers.NewProductLogHandler(&productLogService)
	fotoProdukHandler := handlers.NewFotoProdukHandler(&fotoProdukService)
	trxDetailHandler := handlers.NewTransactionDetailHandler(trxDetailService)
	keranjangBelanjaHandler := handlers.NewKeranjangBelanjaHandler(&keranjangBelanjaService)
	wishlistHandler := handlers.NewWishlistHandler(&wishlistService)
	productReviewHandler := handlers.NewProductReviewHandler(&productReviewService)
	notificationHandler := handlers.NewNotificationHandler(notificationService)
	promoHandler := handlers.NewProductPromoHandler(&promoService)
	diskonProdukHandler := handlers.NewDiskonProdukHandler(diskonProdukService)
	orderHandler := handlers.NewOrderHandler(orderService)
	couponHandler := handlers.NewProductCouponHandler(couponService)

	// Setup Fiber
	app := fiber.New(configs.NewFiberConfig())

	// Remove this line since Locals is not available
	// app.Locals("db", database)

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New())

	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll("./uploads", 0755); err != nil {
		log.Fatal("Failed to create uploads directory:", err)
	}

	// Serve static files from uploads directory
	app.Static("/uploads", "./uploads")

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(responder.ApiResponse{
			Status:  true,
			Message: configuration.Get("APP_NAME"),
			Error:   nil,
			Data:    nil,
		})
	})

	// Register routes
	authHandler.Route(app)
	userHandler.Route(app)
	addressHandler.Route(app)
	regionHandler.Route(app)
	categoryHandler.Route(app, database) // Pass db here
	storeHandler.Route(app)
	storePhotoHandler.Route(app)
	productHandler.Route(app)
	transactionHandler.Route(app)
	productLogHandler.Route(app)
	fotoProdukHandler.Route(app)
	trxDetailHandler.Route(app)
	keranjangBelanjaHandler.Route(app)
	wishlistHandler.Route(app)
	productReviewHandler.Route(app)
	notificationHandler.Route(app)
	promoHandler.Route(app)
	diskonProdukHandler.Route(app)
	orderHandler.Route(app)
	couponHandler.Route(app)

	// Not Found Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Status:  false,
			Message: "NOT FOUND",
			Error:   exceptions.NewString("Not Found"),
			Data:    nil,
		})
	})

	// Graceful Shutdown
	chanServer := make(chan os.Signal, 1)
	signal.Notify(chanServer, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	host := ":" + configuration.Get("APP_PORT")
	go func() {
		<-chanServer
		log.Printf("Server is shutting down...")
		if err := app.Shutdown(); err != nil {
			log.Printf("Error in shutting down the server: %v", err)
		}
	}()

	log.Printf("Server is running on port %s", host)
	if err := app.Listen(host); err != nil {
		log.Fatalf("Error in running the server: %v", err)
	}
}
