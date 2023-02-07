// @generated by protoc-gen-connect-web v0.6.0 with parameter "target=ts"
// @generated from file gardens/v1/recipes.proto (package gardens.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateRecipeRequest, CreateRecipeResponse, DeleteRecipeRequest, DeleteRecipeResponse, GetRecipeRequest, GetRecipeResponse, GetRecipesRequest, GetRecipesResponse, UpdateRecipeRequest, UpdateRecipeResponse } from "./recipes_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service gardens.v1.RecipesService
 */
export const RecipesService = {
  typeName: "gardens.v1.RecipesService",
  methods: {
    /**
     * @generated from rpc gardens.v1.RecipesService.CreateRecipe
     */
    createRecipe: {
      name: "CreateRecipe",
      I: CreateRecipeRequest,
      O: CreateRecipeResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc gardens.v1.RecipesService.DeleteRecipe
     */
    deleteRecipe: {
      name: "DeleteRecipe",
      I: DeleteRecipeRequest,
      O: DeleteRecipeResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc gardens.v1.RecipesService.UpdateRecipe
     */
    updateRecipe: {
      name: "UpdateRecipe",
      I: UpdateRecipeRequest,
      O: UpdateRecipeResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc gardens.v1.RecipesService.GetRecipes
     */
    getRecipes: {
      name: "GetRecipes",
      I: GetRecipesRequest,
      O: GetRecipesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc gardens.v1.RecipesService.GetRecipe
     */
    getRecipe: {
      name: "GetRecipe",
      I: GetRecipeRequest,
      O: GetRecipeResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

