package routes

import (
	constroller "github.com/atifali-pm/golang-resturant-management/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", constroller.GetInvoice())
	incomingRoutes.POST("/invoices", constroller.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", constroller.UpdateInvoice())

}
