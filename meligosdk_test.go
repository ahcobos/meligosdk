package meligosdk

import (
  "testing"
)

func testMeliMethod(t *testing.T){
  clientId := "SomeClientId"
  clientSecret := "someClientSecret"
  clientAccessToken := "someAccessToken"
  clientRefreshToken := "someRefreshToken"

  baseMeli := Meli{
    clientId: clientId,
    clientSecret: clientSecret,
    accesToken: clientAccessToken,
    refreshToken: clientRefreshToken}

  meliClient := NewMeli(clientId,clientSecret,clientAccessToken,clientRefreshToken)

  if meliClient != baseMeli {
    t.Error("Meli Method is not does not assing attributes properly")
  }
}

func testMeliUsingAccesTokenMethod(t *testing.T){
  clientId := "SomeClientId"
  clientSecret := "someClientSecret"
  clientAccessToken := "someAccessToken"

  baseMeli := Meli{
    clientId: clientId,
    clientSecret: clientSecret,
    accesToken: clientAccessToken}

  meliClient := NewMeliUsingAccesToken(clientId,clientSecret,clientAccessToken)

  if meliClient != baseMeli {
    t.Error("Meli Method is not does not assing attributes properly")
  }
}

func testMeliUsingClientAndSecretMethod(t *testing.T){
  clientId := "SomeClientId"
  clientSecret := "someClientSecret"

  baseMeli := Meli{
    clientId: clientId,
    clientSecret: clientSecret}

  meliClient := NewMeliUsingClientAndSecret(clientId,clientSecret)

  if meliClient != baseMeli {
    t.Error("Meli Method is not does not assing attributes properly")
  }
}
