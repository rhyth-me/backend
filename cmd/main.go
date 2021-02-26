package main

import (
	"fmt"
	"os"

	"github.com/ScoreMarket/backend/interfaces"
	"github.com/ScoreMarket/backend/interfaces/props"
	"github.com/ScoreMarket/backend/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	middlewareList := make([]*interfaces.MiddlewareSet, 0)
	mid := &interfaces.MiddlewareSet{
		Path: "/api/",
		MiddlewareFunc: []echo.MiddlewareFunc{
			middleware.JWTWithConfig(middleware.JWTConfig{
				ContextKey:     "jwt",
				SuccessHandler: nil,
				SigningKey:     []byte("key"),
				SigningMethod:  jwt.SigningMethodHS512.Name,
				Claims:         new(jwt.StandardClaims),
				TokenLookup:    "cookie:ApiGenSession",
			}),
		},
	}
	middlewareList = append(middlewareList, mid)

	e.Debug = true
	e.Use(middleware.Recover())

	p := new(props.ControllerProps)

	// firebase init
	p.Firestore = utils.InitFirestore()

	interfaces.Bootstrap(p, e, nil, os.Stdout)

	fmt.Println("All routes are...")
	for _, r := range e.Routes() {
		fmt.Printf("%s %s: %s\n", r.Method, r.Path, r.Name)
	}

	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
