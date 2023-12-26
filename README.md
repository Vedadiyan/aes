# AES 
AES is a simple AES encryptor for Go which supports IV out of the box and does not require any Padding mechansim. It comes with a simple API which only includes `Encrypt` and `Decrypt`. 

Example:

    package main 

    import (
        "log"
        "github.com/vedadiyan/aes"
    )

    func main() {
        aes, err := aes.New(YOUR KEY, YOUR IV)
        if err != nil {
            log.Fatalln(err)
        }
        encrypted := aes.Encrypt([]byte {1,2,3})
        decrypted := aes.Decrypt(encrypted)
        _ = decrypted 
    }
