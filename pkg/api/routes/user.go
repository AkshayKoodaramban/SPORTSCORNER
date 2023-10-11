package routes

import (
	"sportscorner/pkg/api/handler"
	"sportscorner/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup, userHandler *handler.UserHandler, ProductHandler *handler.ProductHandler, otpHandler *handler.OtpHandler, CartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler) {

	engine.POST("/signup", userHandler.UserSignUp)
	engine.GET("/login", userHandler.LoginHandler)
	engine.GET("/forgot-password", userHandler.ForgotPasswordSend)
	engine.POST("/forgot-password", userHandler.ForgotPasswordVerifyAndChange)

	engine.GET("/otplogin", otpHandler.SendOTP)
	engine.POST("/verifyotp", otpHandler.VerifyOTP)

	payment := engine.Group("/payment")
	{
		payment.GET("/razorpay", paymentHandler.MakePaymentRazorPay)
		payment.GET("/update_status", paymentHandler.VerifyPayment)
	}

	engine.Use(middleware.UserAuthMiddleware)
	{
		home := engine.Group("/home")
		{
			home.GET("/product", ProductHandler.ListProducts)
			home.GET("/product/details", ProductHandler.ShowIndividualProducts)
			home.POST("/add-to-cart", CartHandler.AddToCart)

		}

		profile := engine.Group("/profile")
		{
			profile.GET("/details", userHandler.GetUserDetails)
			profile.GET("/address", userHandler.GetAddresses)
			profile.POST("/address/add", userHandler.AddAddress)
		}
		edit := engine.Group("/edit")
		{
			edit.PUT("/name", userHandler.EditName)
			edit.PUT("/email", userHandler.EditEmail)
			edit.PUT("/phone", userHandler.EditPhone)
		}

		cart := engine.Group("/cart")
		{
			cart.GET("/", userHandler.GetCart)
			cart.DELETE("/remove", userHandler.RemoveFromCart)
			cart.PUT("/updateQuantity/plus", userHandler.UpdateQuantityAdd)
			cart.PUT("/updateQuantity/minus", userHandler.UpdateQuantityLess)
		}

		order := engine.Group("/order")
		{
			order.GET("", orderHandler.GetOrders)

		}

		checkout := engine.Group("/check-out")
		{
			checkout.GET("", CartHandler.CheckOut)
			checkout.POST("/order", orderHandler.OrderItemsFromCart)

		}

	}

}
