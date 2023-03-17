package car

import(
	"net/http"
	"se/jwt-api/orm"
	"github.com/gin-gonic/gin"
)

func CarAll(c *gin.Context){
	var cars []orm.Car
	orm.Db.Find(&cars)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Car Read Successful",
"cars": "cars"})
}

type CarBody struct{
	Carname string
	Detail string
	Image string
}

func RegisterCar(c*gin.Context){
	var json CarBody
	if err := c.ShouldBindJSON(&json); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	car := orm.Car{Carname: json.Carname, Detail: json.Detail,
		Image: json.Image}
	
	orm.Db.Create(&car)
	if car.ID > 0 {
		c.JSON(200, gin.H{"status":"ok", "message":"Car Create Successful",
	"Carname": car.ID})
	}else{
		c.JSON(200,gin.H{"status": "error","message": "Car Register Fail"})
	}
	
}