package azuresas

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"html/template"
	"time"
	"strconv"
)

func GenerateToken(resourceURI string, keyName string, key string, expireSecs int) string {
	uri := template.URLQueryEscaper(resourceURI)
	expireTime := time.Now().Add(time.Duration(expireSecs) * time.Second)
	expireTimestamp := expireTime.Unix()
	expireStr := strconv.FormatInt(expireTimestamp, 10)
	toSign := uri + "\n" + expireStr
	sign := getHmac256(toSign, key)
	encodedSign := template.URLQueryEscaper(sign)
	token := "SharedAccessSignature sr=" + uri + "&sig=" + encodedSign + "&se=" + expireStr + "&skn=" + keyName
	return token
}

func getHmac256(str string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}