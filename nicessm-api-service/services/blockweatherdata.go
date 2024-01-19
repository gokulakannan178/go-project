package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"nicessm-api-service/models"
	"strings"

	"github.com/jlaffaye/ftp"
)

func (s *Service) LoadBlockWeatherReport(ctx *models.Context) {
	client, err := ftp.Dial("103.215.208.49:21")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := client.Login("anonymous", "anonymous"); err != nil {
		fmt.Println(err)
		return
	}
	r, err := client.Retr("/pub/DIST_BLOCK_FT1534/2020/dfcst/andhra-pradesh/dfcst200z200101")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	// println(string(buf))
	//var lbwr []models.LoadBlockWeatherReport
	lines := strings.Split(string(buf), "\n")

	for _, v := range lines {
		words := strings.Fields(v)
		fmt.Println(words, len(words))
	}
	//	err = s.Daos.LoadBlockWeatherReport(ctx)
	if err != nil {
		return
	}
	//	fmt.Println("weather data===>", r)
	if err := client.Quit(); err != nil {
		log.Fatal(err)
	}
}
