package taranis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/vimiew/api-template/dbcontroller"
)

type Action struct {
	ActionID                int    `gorm:"primary_key;type:int;" json:"action_id"`
	ActionNumber            int    `gorm:"type:int;" json:"action_number"`
	LogDate                 string `gorm:"type:datetime2(7)" json:"log_date"`
	LogBy                   string `gorm:"type:varchar(25);" json:"log_by"`
	ActionTypeID            int    `gorm:"type:int;" json:"action_type_id"`
	Description             string `gorm:"type:nvarchar(max);" json:"description"`
	Investigator            string `gorm:"type:varchar(25);" json:"investigator"`
	TrainingStandard        int    `gorm:"type:int;" json:"training_standard"`
	EstimatedCompletionDate string `gorm:"type:varchar(10);" json:"estimated_completion_date"`
	FinalClosureBy          string `gorm:"type:varchar(25);" json:"final_closure_by"`
	FinalClosureDate        string `gorm:"type:datetime2(7);" json:"finalClosure_date"`
}

// GetActions - Select Command
func GetActions(c *gin.Context) {
	var actions []Action
	page, _ := strconv.Atoi(c.Query("page"))
	search := c.Query("search")

	db := dbcontroller.Database("Taranis")
	defer db.Close()

	if search != "" {
		isearch, _ := strconv.Atoi(search)
		if err := db.Raw(`select * from actions_view 
			where action_number = ? or description like '%' + ? + '%' or investigator like '%' + ? + '%'
			order by action_number desc`, isearch, search, search).
			Scan(&actions).Error; err != nil {
			panic(err)
		}
	} else {
		pageCount := 100
		p := (page - 1) * pageCount
		if err := db.Raw(`select * from actions_view order by action_number desc
			offset ? rows fetch next ? rows only`, p, pageCount).
			Scan(&actions).Error; err != nil {
			panic(err)
		}
	}

	if len(actions) <= 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "No actions found!"})
		return
	}

	c.JSON(http.StatusOK, actions)
}

// GetAction- Select Command
func GetAction(c *gin.Context) {
	var action Action
	action_no := c.Param("id")

	db := dbcontroller.Database("Taranis")
	defer db.Close()

	if err := db.Table("actions_view").Where("action_number = ?", action_no).Find(&action).Error; err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Action not found!"})
			return
		}
		panic(err)
	}

	c.JSON(http.StatusOK, action)
}

// // PostNote - Insert Command to Table
// func PostNote(c *gin.Context) {
// 	Note := Note{
// 		NoteHash:        dbcontroller.UniqueCode(11),
// 		ApplicationHash: c.PostForm("application"),
// 		Note:            c.PostForm("note"),
// 	}

// 	db := dbcontroller.Database("Taranis")
// 	defer db.Close()

// 	if err := db.Create(&Note).Error; err != nil {
// 		panic(err)
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated,
// 		"message": "Note logged successfully!"})
// }

// // PutNote - Update Command to Table
// func PutNote(c *gin.Context) {
// 	hash := c.Param("id")
// 	// var Note Note
// 	var rowCount int

// 	db := dbcontroller.Database("Taranis")
// 	defer db.Close()

// 	if err := db.Table("notes").Where("note_hash = ?", hash).Count(&rowCount).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":  http.StatusNotFound,
// 			"message": err.Error()})
// 		return
// 	}

// 	if rowCount == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":  http.StatusNotFound,
// 			"message": "Could not find standard"})
// 		return
// 	}

// 	affectedRows := db.Table("notes").Where("note_hash = ?", hash).Update(Note{
// 		ApplicationHash: c.PostForm("application"),
// 		Note:            c.PostForm("note"),
// 	}).RowsAffected

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":       http.StatusOK,
// 		"message":      "Note has been updated",
// 		"affectedRows": affectedRows})
// }

// // DeleteNote - Delete Command to Table
// func DeleteNote(c *gin.Context) {
// 	var Note Note
// 	hash := c.Param("id")

// 	db := dbcontroller.Database("Taranis")
// 	if err := db.Table("notes").Where("note_hash = ?", hash).Find(&Note).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNoContent, "message": err.Error()})
// 		return
// 	}
// 	defer db.Close()

// 	affectedRows := db.Table("notes").Where("note_hash = ?", hash).Delete(&Note).RowsAffected
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":       http.StatusOK,
// 		"message":      "Note has been removed",
// 		"affectedRows": affectedRows})
// }
