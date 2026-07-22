package webconfiguration

import "time"

type User struct{
	ID            uint   
	Name          string 
	Email         string 
	Password      string 
	Role          string 
	KYCStatus     bool	
	IsVerified    bool	
	IsBlocked     bool	
	ProfilePicURL string   
	CreatedAt     time.Time 
	UpdatedAt     time.Time 
	
	}