//MIT License
//
//Copyright (c) 2017 rangertaha
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package main

import (
	"net/http"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
	"github.com/rangertaha/go-accounts/cmd"
	"github.com/rangertaha/go-accounts/accounts"
)

func LoginRequired (){

}

func main() {
	cmd.Execute()


	// Override default Configs
	// var CustomConfig =

	// Get accounts handler
	// accounts : = NewAccounts(CustomConfig)




	e := echo.New()

	g := e.Group("/account")


	g.POST("/signup", accounts.Signup)
	g.POST("/login", accounts.Login)
	g.GET("/logout", accounts.Logout)
	g.GET("/confirm/:Token", accounts.Confirm)
	g.GET("/delete", accounts.Delete, LoginRequired)

	g.GET("/password", accounts.Password, LoginRequired)
	g.GET("/password/reset", accounts.Reset)
	g.GET("/password/reset/:Token", accounts.ResetId)
	g.GET("/settings", accounts.Settings, LoginRequired)

	e.Logger.Fatal(e.Start(":8888"))

}
