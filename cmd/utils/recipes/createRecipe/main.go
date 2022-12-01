package main

import (
	"context"
	"log"
	"net/http"
	"time"

	gardensv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1/gardensv1connect"
	"github.com/bufbuild/connect-go"
)

func main() {
	var fifteenMinutes float64 = 900000
	client := gardensv1connect.NewRecipesServiceClient(
		http.DefaultClient,
		"http://localhost:8082/",
	)
	req := connect.NewRequest(&gardensv1.CreateRecipeRequest{
		Name:         "Recipe A_1",
		Comment:      "Recipe Mock Comment",
		Ingredients:  "[{name:\"Base Solvent\",comment:\"Reduces Tap Water to 5.8 ph\",amount:{value:2,metric:\"oz\"}},{name:\"Nutrient Blend\",comment:\"Provides Nutrients\",amount:{value:6,metric:\"oz\"}},{name:\"water\",comment:\"tapwater\",amount:{value:3,metric:\"l\"}}]",
		Instructions: "[{step:1,name:\"Add Base Solvent to Water\",comment:\"Mix in small amounts and test until 5.8ph\",estimated_time:300000},{step:2,name:\"Add Nutrient Blend to Solution\",comment:\"You may need to mix in more Base Solvent in the case that the ph exceeds 7.08 ph\",estimated_time:600000},]",
		Ph:           7.04,
		MixTime:      fifteenMinutes,
		Tags:         "[\"Tag A\", \"Tag B\"]",
	})
	ctx, done := context.WithTimeout(context.Background(), time.Second*5)
	defer done()
	res, err := client.CreateRecipe(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}
