package http

import (
	"net/http"
	"src/internal/dto"
	"src/internal/model"
	"src/internal/service"
	"src/pkg/logging"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SupplierController struct {
	l               logging.Interface
	supplierService service.ISupplierService
	userService     service.IUserService
}

func NewSupplierController(
	l logging.Interface,
	supplierService service.ISupplierService,
	userService service.IUserService) *SupplierController {
	return &SupplierController{
		l:               l,
		supplierService: supplierService,
		userService:     userService,
	}
}

func (s *SupplierController) ListsAllSuppliers(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		s.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.userService.GetUserByID(c, userID)

	if user.Role != model.UserRoleAdmin {
		s.l.Errorf("%s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	suppliers, err := s.supplierService.GetAllSuppliers(c)
	if err != nil {
		s.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"suppliers": suppliers,
	})
}

func (s *SupplierController) AddSupplier(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		s.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.userService.GetUserByID(c, userID)

	if user.Role != model.UserRoleAdmin {
		s.l.Errorf("%s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var req dto.CreateSupplierReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		s.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplier, err := s.supplierService.CreateSupplier(c, &req)
	if err != nil {
		s.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"supplier": supplier,
	})
}

func (s *SupplierController) RemoveSupplier(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		s.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.userService.GetUserByID(c, userID)

	if user.Role != model.UserRoleAdmin {
		s.l.Errorf("%s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	supplierID, _ := strconv.Atoi(c.Param("id"))

	supplier, err := s.supplierService.GetSupplierByID(c, supplierID)
	if err != nil {
		s.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = s.supplierService.RemoveSupplier(c, supplier.Email)
	if err != nil {
		s.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
