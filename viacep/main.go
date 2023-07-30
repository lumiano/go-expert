package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, url := range os.Args[1:] {

		req, err := http.Get(url)

		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "fetch: %v\n", err)

			if err != nil {
				return
			}

		}

		res, err := io.ReadAll(req.Body)

		if err != nil {
			_, err2 := fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			if err2 != nil {
				return
			}

		}

		_, err = fmt.Fprintf(os.Stdout, "fetch: %s\n", res)

		if err != nil {
			return
		}

		var viaCep ViaCep

		err = json.Unmarshal(res, &viaCep)

		if err != nil {
			_, err2 := fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			if err2 != nil {
				return
			}
		}

		_, err = fmt.Fprintf(os.Stdout, "fetch: %v\n", viaCep)

		if err != nil {
			return
		}

		file, err := os.Create("response.txt")

		_, err = file.WriteString(fmt.Sprintln(viaCep))

		err = req.Body.Close()

		err = file.Close()

	}

}
