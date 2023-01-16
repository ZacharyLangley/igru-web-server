package gardens_test

import (
	"log"
	"net/http"
	"testing"

	gardensv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1/gardensv1connect"
	"github.com/bufbuild/connect-go"
)

func initRecipesClient() gardensv1connect.RecipesServiceClient {
	return gardensv1connect.NewRecipesServiceClient(
		http.DefaultClient,
		"http://localhost:80/",
	)
}

func CreateRecipe() (*connect.Response[gardensv1.CreateRecipeResponse], error) {
	var fifteenMinutes float64 = 900000
	client := initRecipesClient()
	req := connect.NewRequest(&gardensv1.CreateRecipeRequest{
		Name:                "Recipe A_1",
		Comment:             "Recipe Mock Comment",
		Ingredients:         "[{name:\"Base Solvent\",comment:\"Reduces Tap Water to 5.8 ph\",amount:{value:2,metric:\"oz\"}},{name:\"Nutrient Blend\",comment:\"Provides Nutrients\",amount:{value:6,metric:\"oz\"}},{name:\"water\",comment:\"tapwater\",amount:{value:3,metric:\"l\"}}]",
		Instructions:        "[{step:1,name:\"Add Base Solvent to Water\",comment:\"Mix in small amounts and test until 5.8ph\",estimated_time:300000},{step:2,name:\"Add Nutrient Blend to Solution\",comment:\"You may need to mix in more Base Solvent in the case that the ph exceeds 7.08 ph\",estimated_time:600000},]",
		Ph:                  7.04,
		MixTimeMilliseconds: fifteenMinutes,
		Tags:                "[\"Tag A\", \"Tag B\"]",
	})
	ctx, done := initContext()
	defer done()
	return client.CreateRecipe(ctx, req)
}

func TestCreateRecipe(t *testing.T) {
	res, err := CreateRecipe()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}

func UpdateRecipe(id string) (*connect.Response[gardensv1.UpdateRecipeResponse], error) {
	var twentyMinutes float64 = 1200000
	client := initRecipesClient()
	req := connect.NewRequest(&gardensv1.UpdateRecipeRequest{
		Id:                  id,
		Name:                "Updated Recipe A_1",
		Comment:             "Updated Recipe Mock Comment",
		Ingredients:         "[{name:\"Base Solvent\",comment:\"Reduces Tap Water to 5.8 ph\",amount:{value:2,metric:\"oz\"}},{name:\"Nutrient Blend\",comment:\"Provides Nutrients\",amount:{value:6,metric:\"oz\"}},{name:\"water\",comment:\"tapwater\",amount:{value:3,metric:\"l\"}}]",
		Instructions:        "[{step:1,name:\"Add Base Solvent to Water\",comment:\"Mix in small amounts and test until 5.8ph\",estimated_time:300000},{step:2,name:\"Add Nutrient Blend to Solution\",comment:\"You may need to mix in more Base Solvent in the case that the ph exceeds 7.08 ph\",estimated_time:600000},]",
		Ph:                  6.8,
		MixTimeMilliseconds: twentyMinutes,
		Tags:                "[\"Updated Tag A\", \"Updated Tag B\"]",
	})
	ctx, done := initContext()
	defer done()
	return client.UpdateRecipe(ctx, req)
}

func TestUpdateRecipe(t *testing.T) {
	res, err := UpdateRecipe("")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}

func GetRecipe(id string) (*connect.Response[gardensv1.GetRecipeResponse], error) {
	client := initRecipesClient()
	req := connect.NewRequest(&gardensv1.GetRecipeRequest{
		Id: id,
	})
	ctx, done := initContext()
	defer done()
	return client.GetRecipe(ctx, req)
}

func TestGetRecipe(t *testing.T) {
	res, err := GetRecipe("")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}

func GetRecipes() (*connect.Response[gardensv1.GetRecipesResponse], error) {
	client := initRecipesClient()
	req := connect.NewRequest(&gardensv1.GetRecipesRequest{})
	ctx, done := initContext()
	defer done()
	return client.GetRecipes(ctx, req)
}

func TestGetRecipes(t *testing.T) {
	res, err := GetRecipes()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}

func DeleteRecipe(id string) (*connect.Response[gardensv1.DeleteRecipeResponse], error) {
	client := initRecipesClient()
	req := connect.NewRequest(&gardensv1.DeleteRecipeRequest{
		Id: id,
	})
	ctx, done := initContext()
	defer done()
	return client.DeleteRecipe(ctx, req)
}

func TestDeleteRecipe(t *testing.T) {
	res, err := DeleteRecipe("")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}
