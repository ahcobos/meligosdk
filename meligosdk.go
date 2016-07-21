package meligosdk

type Meli struct {
  clientId  string
  clientSecret string
  redirectUri string
  accesToken string
  refreshToken string
}

func NewMeli(ClientId string, ClientSecret string, AccesToken string, RefreshToken string) Meli{
  return Meli{
    clientId: ClientId,
    clientSecret: ClientSecret,
    accesToken: AccesToken,
    refreshToken: RefreshToken}
}

func NewMeliUsingAccesToken(ClientId string, ClientSecret string, AccesToken string) Meli{
  return Meli{
    clientId: ClientId,
    clientSecret: ClientSecret,
    accesToken: AccesToken}
}

func NewMeliUsingClientAndSecret(ClientId string, ClientSecret string) Meli{
  return Meli{
    clientId: ClientId,
    clientSecret: ClientSecret}
}
