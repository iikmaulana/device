package token

import (
	"strings"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

type Initialize struct {
	Alg                    string
	method                 jwt.SigningMethod
	RefreshTokenExpiration int64
	AccesstokenExpiration  int64
	SecretKey              string
	parser                 jwt.Parser
}

type MapClaims struct {
	params   jwt.MapClaims
	register Initialize
}

func (intz Initialize) New() *MapClaims {
	intz.method = jwt.GetSigningMethod(intz.Alg)
	intz.parser = jwt.Parser{}
	return &MapClaims{register: intz, params: map[string]interface{}{}}
}

func (intz Initialize) keyFunc(t *jwt.Token) (interface{}, error) {
	return intz.SecretKey, nil
}

func (intz Initialize) decorder(tokenString string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(intz.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func (intz Initialize) Decode(tokenString string) (map[string]interface{}, error) {
	claim := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		claim["exp"] = time.Now().Add(time.Hour * time.Duration(intz.AccesstokenExpiration)).Unix()
		return []byte(intz.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return claim, nil
}

func (cmap *MapClaims) Issuer(stringOrURI string) *MapClaims {
	if strings.TrimSpace(stringOrURI) != "" {
		cmap.params["iss"] = stringOrURI
	}
	return cmap
}

func (cmap *MapClaims) Subject(stringOrURI string) *MapClaims {
	if strings.TrimSpace(stringOrURI) != "" {
		cmap.params["sub"] = stringOrURI
	}
	return cmap
}

func (cmap *MapClaims) Audience(stringOrURI string) *MapClaims {
	if strings.TrimSpace(stringOrURI) != "" {
		cmap.params["aud"] = stringOrURI
	}
	return cmap
}

func (cmap *MapClaims) ExpirationTime(numericDate int64) *MapClaims {
	if numericDate > 0 {
		cmap.params["exp"] = numericDate
	}
	return cmap
}

func (cmap *MapClaims) NotBefore(numericDate int64) *MapClaims {
	if numericDate > 0 {
		cmap.params["nbf"] = numericDate
	}
	return cmap
}

func (cmap *MapClaims) IssuedAt(numericDate int64) *MapClaims {
	if numericDate > 0 {
		cmap.params["iat"] = numericDate
	}
	return cmap
}

func (cmap *MapClaims) JwtID(id string) *MapClaims {
	if strings.TrimSpace(id) != "" {
		cmap.params["jti"] = id
	}
	return cmap
}

func (cmap *MapClaims) Claim() (string, error) {
	return cmap.ClaimWithMap(map[string]interface{}{})
}

func (cmap *MapClaims) ClaimWithMapAndCustomKey(sk string, maps map[string]interface{}) (string, error) {
	if len(maps) > 0 {
		for k, v := range maps {
			if cmap.params[k] == nil {
				cmap.params[k] = v
			}
		}
	}

	clm := jwt.NewWithClaims(cmap.register.method, cmap.params)
	return clm.SignedString([]byte(sk))
}

func (cmap *MapClaims) ClaimWithMap(custom map[string]interface{}) (string, error) {
	if len(custom) > 0 {
		for key, val := range custom {
			if cmap.params[key] == nil {
				cmap.params[key] = val
			}
		}
	}
	claim := jwt.NewWithClaims(cmap.register.method, cmap.params)
	return claim.SignedString([]byte(cmap.register.SecretKey))
}

func (cmap *MapClaims) ClaimAccessToken(uniq string, userId string, username string, isOrgAdmin int64, isActive int64, organizationId string, appId string) (string, error) {
	cmap.params = jwt.MapClaims{}
	if cmap.params["exp"] == nil {
		cmap.params["exp"] = time.Now().Add(time.Hour * time.Duration(cmap.register.AccesstokenExpiration)).Unix()
	}
	if cmap.params["userid"] == nil {
		cmap.params["userid"] = userId
	}
	if cmap.params["username"] == nil {
		cmap.params["username"] = username
	}
	if cmap.params["isOrgAdmin"] == nil {
		cmap.params["isOrgAdmin"] = isOrgAdmin
	}
	if cmap.params["isActive"] == nil {
		cmap.params["isActive"] = isActive
	}
	if cmap.params["organization_id"] == nil {
		cmap.params["organization_id"] = organizationId
	}
	if cmap.params["app_id"] == nil {
		cmap.params["app_id"] = appId
	}
	if cmap.params["iat"] == nil {
		cmap.params["iat"] = time.Now().Unix()
	}
	if cmap.params["uniq"] == nil {
		cmap.params["uniq"] = uniq
	}

	claim := jwt.NewWithClaims(cmap.register.method, cmap.params)
	return claim.SignedString([]byte("!token!key!"))
}

func (cmap *MapClaims) ClaimRefreshToken(uniq string) (string, error) {
	cmap.params = jwt.MapClaims{}
	if cmap.params["exp"] == nil {
		cmap.params["exp"] = time.Now().Add(time.Hour * time.Duration(cmap.register.RefreshTokenExpiration)).Unix()
	}
	if cmap.params["iat"] == nil {
		cmap.params["iat"] = time.Now().Unix()
	}
	if cmap.params["uniq"] == nil {
		cmap.params["uniq"] = uniq
	}
	claim := jwt.NewWithClaims(cmap.register.method, cmap.params)
	return claim.SignedString([]byte("tester-tester.1@"))
}

func (cmap *MapClaims) ClaimPairToken(uniq string, userId string, username string, isOrgAdmin int64, isActive int64, organizationId string, appId string) (result map[string]string, err error) {
	access, err := cmap.ClaimAccessToken(uniq, userId, username, isOrgAdmin, isActive, organizationId, appId)
	if err != nil {
		return result, err
	}
	refresh, err := cmap.ClaimRefreshToken(uniq)
	if err != nil {
		return result, err
	}

	result = map[string]string{
		"access_token":  access,
		"refresh_token": refresh,
	}

	return result, nil
}
