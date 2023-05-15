package db_stuff

import "github.com/google/uuid"

type ServerMessage struct {
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Type    string `json:"type,omitempty" bson:"type,omitempty"`
}

type Card struct {
	Question string `json:"question,omitempty" bson:"question,omitempty"`
	Answer   string `json:"answer,omitempty" bson:"answer,omitempty"`
	Photo    string `json:"photo,omitempty" bson:"photo,omitempty"`
}

type CardCollection struct {
	Name                     string `json:"name,omitempty" bson:"name,omitempty"`
	Category                 string `json:"category,omitempty" bson:"category,omitempty"`
	Description              string `json:"description,omitempty" bson:"description,omitempty"`
	CardList                 []Card `json:"cardList,omitempty" bson:"cardList,omitempty"`
	CreationDate             string `json:"creationDate,omitempty" bson:"creationDate,omitempty"`
	LastView                 string `json:"lastView,omitempty" bson:"lastView,omitempty"`
	LastViewStudy            string `json:"lastViewStudy,omitempty" bson:"lastViewStudy,omitempty"`
	MostRecentScore          int    `json:"mostRecentScore,omitempty" bson:"mostRecentScore,omitempty"`
	TotalScores              []int  `json:"totalScores,omitempty" bson:"totalScores,omitempty"`
	HighScore                int    `json:"highScore,omitempty" bson:"highScore,omitempty"`
	LastViewEasy             string `json:"lastViewEasy,omitempty" bson:"lastViewEasy,omitempty"`
	MostRecentGradeEasy      int    `json:"mostRecentGradeEasy,omitempty" bson:"mostRecentGradeEasy,omitempty"`
	TotalGradesEasy          []int  `json:"totalGradesEasy,omitempty" bson:"totalGradesEasy,omitempty"`
	HighGradeEasy            int    `json:"highGradeEasy,omitempty" bson:"highGradeEasy,omitempty"`
	LastViewDifficult        string `json:"lastViewDifficult,omitempty" bson:"lastViewDifficult,omitempty"`
	MostRecentGradeDifficult int    `json:"mostRecentGradeDifficult,omitempty" bson:"mostRecentGradeDifficult,omitempty"`
	TotalGradesDifficult     []int  `json:"totalGradesDifficult,omitempty" bson:"totalGradesDifficult,omitempty"`
	HighGradeDifficult       int    `json:"highGradeDifficult,omitempty" bson:"highGradeDifficult,omitempty"`
}

type User struct {
	ID          uuid.UUID        `json:"ID,omitempty" bson:"ID,omitempty"`
	Name        string           `json:"name,omitempty" bson:"name,omitempty"`
	Email       string           `json:"email,omitempty" bson:"email,omitempty"`
	Username    string           `json:"username,omitempty" bson:"username,omitempty"`
	Password    string           `json:"password,omitempty" bson:"password,omitempty"`
	JWT         string           `json:"jwt,omitempty" bson:"jwt,omitempty"`
	Collections []CardCollection `json:"collections,omitempty" bson:"collections,omitempty"`
}

type UserDictionary struct {
	ID      uuid.UUID `json:"ID,omitempty" bson:"ID,omitempty"`
	Words   []Word    `json:"Words,omitempty" bson:"Words,omitempty"`
	Phrases []Phrase  `json:"Phrases,omitempty" bson:"Phrases,omitempty"`
}

type Word struct {
	NewWord         string `json:"NewWord,omitempty" bson:"NewWord,omitempty"`
	Translation     string `json:"Translation,omitempty" bson:"Translation,omitempty"`
	Meaning         string `json:"Meaning,omitempty" bson:"Meaning,omitempty"`
	CreatedAt       string `json:"CreatedAt,omitempty" bson:"CreatedAt,omitempty"`
	LastViewed      string `json:"LastViewed,omitempty" bson:"LastViewed,omitempty"`
	ConfidenceLevel string `json:"ConfidenceLevel,omitempty" bson:"ConfidenceLevel,omitempty"`
}

type Phrase struct {
	NewPhrase       string `json:"NewPhrase,omitempty" bson:"NewPhrase,omitempty"`
	Translation     string `json:"Translation,omitempty" bson:"Translation,omitempty"`
	CreatedAt       string `json:"CreatedAt,omitempty" bson:"CreatedAt,omitempty"`
	LastViewed      string `json:"LastViewed,omitempty" bson:"LastViewed,omitempty"`
	ConfidenceLevel string `json:"ConfidenceLevel,omitempty" bson:"ConfidenceLevel,omitempty"`
}

type AddCollectionRequest struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Category    string `json:"category,omitempty" bson:"category,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	CardList    []Card `json:"cardList,omitempty" bson:"cardList,omitempty"`
	Username    string `json:"username,omitempty" bson:"username,omitempty"`
}

type LoginResult struct {
	AuthResult bool
	JWT        string
}

type Session struct {
	JWT string
}

type SessionRequest struct {
	Email    string
	Name     string
	IsSignup bool
}
