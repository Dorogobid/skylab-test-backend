package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlerInterface interface {
	GetAllCategories(c echo.Context) error
	AddCategory(c echo.Context) error
	GetQuestionsByCategoryID(c echo.Context) error
	AddQuestion(c echo.Context) error
}

type Handler struct {
	db DBManagerInterface
}

// GetAllCategories ... Get all quiz categories
// @Summary Get all categories
// @Description Search all categories in database and return in JSON
// @Tags Quiz
// @Produce json
// @Success 200 {object} []Category
// @Failure 400 {object} ErrorResponse
// @Router /quiz/category [get]
func (h *Handler) GetAllCategories(c echo.Context) error {
	categories := h.db.GetAllCategoriesFromDB()
	return c.JSON(http.StatusOK, categories)
}

// AddCategory ... Add category
// @Summary Add category
// @Description Add category in database (JSON body) and return in JSON
// @Tags Quiz
// @Accept json
// @Produce json
// @Param request body CategoryCreate true "Request body example"
// @Success 200 {object} Category
// @Failure 400 {object} ErrorResponse
// @Router /quiz/category [post]
func (h *Handler) AddCategory(c echo.Context) error {
	var categoryCreate CategoryCreate
	err := c.Bind(&categoryCreate); 
	if err != nil {
		errResp := ErrorResponse{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, errResp)
	}
	category, err := h.db.AddCategoryToDB(categoryCreate)
	if err != nil {
		errResp := ErrorResponse{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, errResp)
	}
	return c.JSON(http.StatusOK, category)
}

// GetQuestionsByCategoryID ... Get questions by category ID
// @Summary Get questions by category ID
// @Description Search all questions by category ID in database and return in JSON
// @Tags Quiz
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} []Question
// @Failure 400 {object} ErrorResponse
// @Router /quiz/question/{id} [get]
func (h *Handler) GetQuestionsByCategoryID(c echo.Context) error {
	categoryID := c.Param("category_id")
	if categoryID == "" {
		errResp := ErrorResponse{Message: "Bad request"}
		return c.JSON(http.StatusBadRequest, errResp)
	} 
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		errResp := ErrorResponse{Message: "Bad request"}
		return c.JSON(http.StatusBadRequest, errResp)
	}
	questions := h.db.GetQuestionsByCategoryIDFromDB(id)
	return c.JSON(http.StatusOK, questions)
}

// AddQuestion ... Add question and answers
// @Summary Add question and answers
// @Description Add question and answers in database (JSON body) and return in JSON
// @Tags Quiz
// @Accept json
// @Produce json
// @Param request body QuestionCreate true "Request body example"
// @Success 200 {object} Question
// @Failure 400 {object} ErrorResponse
// @Router /quiz/question [post]
func (h *Handler) AddQuestion(c echo.Context) error {
	var questionCreate QuestionCreate
	err := c.Bind(&questionCreate); 
	if err != nil {
		errResp := ErrorResponse{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, errResp)
	}
	question, err := h.db.AddQuestionToDB(questionCreate)
	if err != nil {
		errResp := ErrorResponse{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, errResp)
	}
	return c.JSON(http.StatusOK, question)
}