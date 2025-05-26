package routes 
import (
	"github.com/gin-gonic/gin"
	controller "RestaurantManagement/controllers"

)
func InvoiceRoutes(incominRoutes *gin.Engine){

	incominRoutes.GET("/invoices", controller.GetInvoicess())
	incominRoutes.GET("/invoices/:invoice_id",controller.GetInvoice())
	incominRoutes.POST("/invoices", controller.CreateInvoice())
	incominRoutes.PATCH("./invoices/invoice_id", controller.UpdateInvoice())
}