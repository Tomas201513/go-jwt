package services

import "go-jwt/models"

type AuthService interface {
	SignUpUser(*models.SignUpInput) (*models.DBResponse, error)
	SignInUser(*models.SignInInput) (*models.DBResponse, error)
}