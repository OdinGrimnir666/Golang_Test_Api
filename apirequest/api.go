package apirequest

import (
	"encoding/json"
	"fmt"

	"io/ioutil"

	"net"
	"net/http"
	"package/DataBase"
	"strconv"

	"github.com/gin-gonic/gin"
)


func PostUrlPackage(c *gin.Context){
	var urlrequest RequestModel

	if err := c.BindJSON(&urlrequest); err != nil {
        return
    }

	if !Is_valid_ip(urlrequest.Ip) || len(urlrequest.Urlpackage) <0 {
		c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Status code 204"})
	}
	var urlpackage []DataBase.UrlPackage
	urlpackage =  DataBase.RequestPostArray(urlrequest.Urlpackage)
	fmt.Println(urlpackage)

	fmt.Println(len(urlpackage),"==",len(urlrequest.Urlpackage))

	if len(urlpackage) != len(urlrequest.Urlpackage){		
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Status code 204"}) 
	}else{
		c.IndentedJSON(200 ,MaxPrice(urlpackage))
	}



}

func GetUrlPackage(c *gin.Context) {
	
	url_package:= c.Query("url_package")
	ip:= c.Query("ip")

	fmt.Printf(url_package)

	if url_package == "" || ip == "" {
		c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Status code 204"})

	}else{

		if testInt, err := strconv.Atoi(url_package); err == nil {

			c.IndentedJSON(http.StatusOK,MakeRequest(DataBase.RequestGet(testInt).Url))
			
		}
		
	}
    
}

func Is_valid_ip(ip string) bool {
	res := net.ParseIP(ip)

	if res == nil {
	return false
}
return true
}

func MakeRequest(url string)(RequestPrice){
	resp, _ := http.Get(url)


		fmt.Println("url",url)

		body, _ := ioutil.ReadAll(resp.Body)


	fmt.Println(string(body))
	var price RequestPrice

	json.Unmarshal(body,&price)
	fmt.Println(price.Price)
    return price

}
func findMaxElement(arr []float64) float64 {
	max_num := arr[0]
	for i:=0; i<len(arr); i++{
   		if arr[i] > max_num {
	  	max_num = arr[i]
   		}
	}
	return max_num
}

func MaxPrice(arrayurlpackage []DataBase.UrlPackage)(pricemax RequestPrice){

	var arrayprice []float64
	for _,s := range arrayurlpackage {
		arrayprice = append(arrayprice,MakeRequest(s.Url).Price)
	}
	fmt.Println("max",findMaxElement(arrayprice))
	return RequestPrice{findMaxElement(arrayprice)}
}
