package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

type DBManagerInterface interface {
	ConnectToDb(cfg *DBConfig)
	GetAllCategoriesFromDB() []Category 
	GetQuestionsByCategoryIDFromDB(categoryID int) []Question
	AddCategoryToDB(categoryCreate CategoryCreate) (*Category, error)
	AddQuestionToDB(questionCreate QuestionCreate) (*Question, error)
	LoadQuestionsFromJSON(questionsCreate []QuestionCreate) error
}

type DBManager struct {
	db *gorm.DB
}

func (db *DBManager) ConnectToDb(cfg *DBConfig) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
	var err error
	for {
		db.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Connection to database failed. Trying to reconnect...")
		} else {
			fmt.Println("Connection to database done.")
			break
		}
		time.Sleep(time.Second * 5)
	}

	err = db.db.AutoMigrate(&Category{}, &Question{}, &Answer{})
	if err != nil {
		panic("Could not run migration.")
	}
}

func (db *DBManager) GetAllCategoriesFromDB() []Category {
	var categories []Category
	db.db.Find(&categories)
	return categories
}

func (db *DBManager) AddCategoryToDB(categoryCreate CategoryCreate) (*Category, error) {
	category := &Category{Name: categoryCreate.Name}
	result := db.db.Create(&category)
	return category, result.Error
}

func (db *DBManager) GetQuestionsByCategoryIDFromDB(categoryID int) []Question {
	var questions []Question
	db.db.Where("category_id = ?", categoryID).Preload("Answers").Find(&questions)
	return questions
}

func (db *DBManager) AddQuestionToDB(questionCreate QuestionCreate) (*Question, error) {
	var answers []Answer
	for _, answer := range questionCreate.Answers {
		answers = append(answers, Answer{AnswerText: answer.AnswerText, IsTrue: answer.IsTrue})
	}
	question := &Question{QuestionText: questionCreate.QuestionText, Answers: answers, CategoryID: questionCreate.CategoryID}
	result := db.db.Create(&question)
	return question, result.Error
}

func (db *DBManager) LoadQuestionsFromJSON(questionsCreate []QuestionCreate) error {
	var answers []Answer
	var questions []Question
	for _, question := range questionsCreate {
		answers = nil
		for _, answer := range question.Answers {
			answers = append(answers, Answer{AnswerText: answer.AnswerText, IsTrue: answer.IsTrue})
		}
		questions = append(questions, Question{QuestionText: question.QuestionText, Answers: answers, CategoryID: question.CategoryID})
	}

	result := db.db.Create(&questions)
	return result.Error
}