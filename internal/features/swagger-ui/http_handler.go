package swaggerui

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/GymSquad/archive-api/api"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

type SwaggerUIHandler struct {
	// PageTitle is the title of the swagger ui page
	PageTitle string
	// OpenAPIURL is the url to the openapi.json file
	OpenAPIURL string
	// SwaggerUIURL is the url to the swagger ui docs page
	SwaggerUIURL string

	// swagger is the openapi spec object
	swagger *openapi3.T
	// swaggerTemplate is the template for the swagger ui page
	swaggerTemplate *template.Template
}

type swaggerInfo struct {
	// Title is the page title of the swagger docs
	Title string
}

// NewHandler creates a new SwaggerUIHandler
func NewHandler(pageTitle, openapiURL, swaggerUIURL string) (*SwaggerUIHandler, error) {
	swagger, err := api.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("Failed to get swagger: %w", err)
	}

	// readin docs/swagger-ui.html
	bytes, err := os.ReadFile("docs/swagger-ui.html")
	if err != nil {
		return nil, fmt.Errorf("Failed to read swagger-ui.html: %w", err)
	}
	swaggerTemplate, err := template.New("swagger-ui").Parse(string(bytes))
	if err != nil {
		return nil, fmt.Errorf("Failed to parse swagger-ui.html: %w", err)
	}

	return &SwaggerUIHandler{
		PageTitle:       pageTitle,
		OpenAPIURL:      openapiURL,
		SwaggerUIURL:    swaggerUIURL,
		swagger:         swagger,
		swaggerTemplate: swaggerTemplate,
	}, nil
}

// DefaultHandler creates a new SwaggerUIHandler with default values
func DefaultHandler() (*SwaggerUIHandler, error) {
	return NewHandler("Archive API", "/openapi.json", "/docs")
}

// Register registers the swagger ui handlers to the router
func (h *SwaggerUIHandler) Register(router gin.IRouter) {
	router.GET(h.OpenAPIURL, h.HandleOpenAPI)
	router.GET(h.SwaggerUIURL, h.HandleSwaggerUI)
}

// HandleOpenAPI handles the openapi.json endpoint
func (h *SwaggerUIHandler) HandleOpenAPI(c *gin.Context) {
	c.JSON(200, h.swagger)
}

// HandleSwaggerUI handles the swagger ui endpoint
func (h *SwaggerUIHandler) HandleSwaggerUI(c *gin.Context) {
	err := h.swaggerTemplate.Execute(c.Writer, &swaggerInfo{
		Title: h.PageTitle,
	})
	if err != nil {
		slog.Error("failed to execute swagger template", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
	}
}
