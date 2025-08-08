package main

import (
    "log"
    "net/http"

    "vietceramics-apis/config"
    "vietceramics-apis/infrastructure/db"
    "vietceramics-apis/infrastructure/jwt"
    "vietceramics-apis/infrastructure/repository"
    "vietceramics-apis/interfaces/handler"
    "vietceramics-apis/interfaces/middleware"
    "vietceramics-apis/interfaces/router"
    "vietceramics-apis/usecase"
)

// main is the entry point of the application.
// It wires up the configuration, database connection, repositories, use cases,
// handlers and starts the HTTP server.
func main() {
    // Load application configuration from environment variables or defaults.
    cfg := config.New()

    // Initialize database connection. In a real application you would handle
    // connection errors and perform migrations. Here we connect to Postgres
    // using configuration values from cfg.
    pgDB, err := db.NewPostgresDB(cfg)
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    // Initialize repositories with the database connection.
    userRepo := repository.NewUserRepository(pgDB)
    permissionRepo := repository.NewPermissionRepository(pgDB)

    // Initialize JWT service to issue and validate tokens.
    jwtService := jwt.NewService(cfg)

    // Initialize use cases which contain the business logic.
    userUC := usecase.NewUserUseCase(userRepo, permissionRepo, jwtService, cfg)

    // Initialize HTTP handlers which will receive requests and call use cases.
    userHandler := handler.NewUserHandler(userUC)

    // Setup router and register routes and middleware.
    r := router.NewRouter()
    // Public endpoints
    r.POST("/login", userHandler.Login)
    // Group protected endpoints with authentication middleware.
    api := r.Group("/users")
    api.Use(middleware.AuthMiddleware(jwtService))
    {
        api.GET("/me", userHandler.GetMe)
        api.GET("", middleware.RBACMiddleware("SA"), userHandler.ListUsers)
        api.POST("", middleware.RBACMiddleware("SA"), userHandler.CreateUser)
    }

    // Start the HTTP server.
    addr := cfg.ServerAddress()
    log.Printf("starting server on %s", addr)
    if err := http.ListenAndServe(addr, r); err != nil {
        log.Fatalf("server error: %v", err)
    }
}
