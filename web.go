package main

//importing the packages
import (
  "github.com/gorilla/mux"
  "net/http"
  "crypto/aes"
  "crypto/cipher"
  "crypto/rand"
  "io/ioutil"
  "io"
  "encoding/json"
  "encoding/hex"
  "log"
  "fmt"
)

//Random 32 byte key for AES-256 that needs to be stored in vault
const key string = "feb3924352093825ce234d7b9b671fe93a194628254b7da703a3e8d4d9292f62"

func main() {
    //create new mux router
    router := mux.NewRouter().StrictSlash(true)
    //redirct routes/calls to related handlers
    router.HandleFunc("/api/encrypt", encryptJson).Methods("POST")
    router.HandleFunc("/api/decrypt", decryptJson).Methods("POST")
    //run the service with port 8080
    log.Fatal(http.ListenAndServe(":8080", router))
}

func encryptJson(w http.ResponseWriter, r *http.Request) {
  //get the body of the request
  reqBody, _ := ioutil.ReadAll(r.Body)
  data := make(map[string]interface{})
  //Map the request with data variable
  err := json.Unmarshal(reqBody, &data)
  if err != nil {
    fmt.Println(err)
  }
  //read the varibale and get the input string
  if readdata, ok := data["data"].(string); ok {
    //call the encrypt function to encrypt the data
    encrypted := encrypt(readdata, key)
    //Write encrypted output to the response
    outputdata := map[string]string{"encrypted": encrypted}
    err := json.NewEncoder(w).Encode(outputdata)
    if err != nil {
      fmt.Println(err)
    }
	} else {
	   fmt.Println("Input json is not in proper format.")
	}
	fmt.Println("Endpoint Hit: encrypt api")
}

func decryptJson(w http.ResponseWriter, r *http.Request) {
  //get the body of the request
  reqBody, _ := ioutil.ReadAll(r.Body)
  data := make(map[string]interface{})
  //Map the request with data variable
  err := json.Unmarshal(reqBody, &data)
  if err != nil {
    fmt.Println(err)
  }
  //read the varibale and get the input string
  if readdata, ok := data["encrypted"].(string); ok {
    //call the decrypt function to decrypt the data
    decrypted := decrypt(readdata, key)
    //Write encrypted output to the response
    outputdata := map[string]string{"decrypted": decrypted}
    err := json.NewEncoder(w).Encode(outputdata)
    if err != nil {
      fmt.Println(err)
    }
	} else {
	   fmt.Println("Input json is not in proper format.")
	}
	fmt.Println("Endpoint Hit: decrypt api")
}

func encrypt(text, key string) string {
  //convert the string inputs to bytes
  textbyte := []byte(text)
  keybyte, _ := hex.DecodeString(key)

  //create the new cyper block from the key
  block, err := aes.NewCipher(keybyte)
  if err != nil {
      fmt.Println(err)
  }
  // create new gcm from using cyper block
  gcm, err := cipher.NewGCM(block)
  if err != nil {
      fmt.Println(err)
  }
  //create the nonce
  nonce := make([]byte, gcm.NonceSize())
  if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
      fmt.Println(err)
  }
  //encrypt the text using gcm
  ciphertext := gcm.Seal(nonce, nonce, textbyte, nil)
  return fmt.Sprintf("%x", ciphertext)
}

func decrypt(ciphertext, key string) string {
  //convert the string inputs to bytes
  keybyte, _ := hex.DecodeString(key)
  cipherbyte, _ := hex.DecodeString(ciphertext)

  //create the new cyper block from the key
  block, err := aes.NewCipher(keybyte)
  if err != nil {
      fmt.Println(err)
  }
  // create new gcm from using cyper block
  gcm, err := cipher.NewGCM(block)
  if err != nil {
      fmt.Println(err)
  }
  //get the nonce size
  nonceSize := gcm.NonceSize()
  nonce, text := cipherbyte[:nonceSize], cipherbyte[nonceSize:]
  //decrypt the data using gcm open
  plaintext, err := gcm.Open(nil, nonce, text, nil)
  if err != nil {
      fmt.Println(err)
  }
  return fmt.Sprintf("%s", plaintext)
}
