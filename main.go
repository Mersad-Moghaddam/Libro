package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"libro/controllers/authController"
	"libro/controllers/bookController"
	"libro/controllers/mainController"
	"libro/controllers/readingController"
	"libro/controllers/userController"
	"libro/controllers/wishlistController"
	"libro/middleware/auth"
	"libro/models/book"
	"libro/models/purchaseLink"
	"libro/models/user"
	"libro/models/wishlist"
	"libro/pkg/cache"
	"libro/services/core"
	"libro/statics/configs"
)

func main() {
	cfg, err := configs.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(mysql.Open(cfg.MySQLDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	_ = db.AutoMigrate(&user.User{}, &book.Book{}, &wishlist.Wishlist{}, &purchaseLink.PurchaseLink{})

	rdb, err := cache.InitRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}

	services, _ := core.InitServices(cfg, db, rdb)

	authCtl := authController.New(services.Auth)
	bookCtl := bookController.New(services.Book)
	readCtl := readingController.New(services.Reading)
	wishCtl := wishlistController.New(services.Wishlist)
	userCtl := userController.New(services.Auth)
	mainCtl := mainController.New()

	app := fiber.New()
	app.Get("/health", mainCtl.Health)
	app.Get("/main/dashboard-summary", mainCtl.DashboardSummary)

	app.Post("/auth/register", authCtl.Register)
	app.Post("/auth/login", authCtl.Login)
	app.Post("/auth/refresh", authCtl.RefreshToken)

	protected := app.Group("", auth.Protected(cfg))
	protected.Post("/auth/logout", authCtl.Logout)
	protected.Get("/auth/me", authCtl.Me)

	protected.Get("/books", bookCtl.GetBooks)
	protected.Post("/books", bookCtl.CreateBook)
	protected.Get("/books/:id", bookCtl.GetBookByID)
	protected.Put("/books/:id", bookCtl.UpdateBook)
	protected.Delete("/books/:id", bookCtl.DeleteBook)
	protected.Patch("/books/:id/status", bookCtl.UpdateBookStatus)
	protected.Patch("/books/:id/bookmark", bookCtl.UpdateBookBookmark)

	protected.Get("/reading/current", readCtl.GetCurrentReadingBooks)
	protected.Patch("/reading/books/:id/progress", readCtl.UpdateReadingProgress)

	protected.Get("/wishlist", wishCtl.GetWishlist)
	protected.Post("/wishlist", wishCtl.CreateWishlistItem)
	protected.Get("/wishlist/:id", wishCtl.GetWishlistItemByID)
	protected.Put("/wishlist/:id", wishCtl.UpdateWishlistItem)
	protected.Delete("/wishlist/:id", wishCtl.DeleteWishlistItem)
	protected.Post("/wishlist/:id/links", wishCtl.AddPurchaseLink)
	protected.Put("/wishlist/:id/links/:linkId", wishCtl.UpdatePurchaseLink)
	protected.Delete("/wishlist/:id/links/:linkId", wishCtl.DeletePurchaseLink)

	protected.Get("/user/profile", userCtl.GetProfile)
	protected.Put("/user/profile", userCtl.UpdateProfile)
	protected.Patch("/user/password", userCtl.ChangePassword)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
