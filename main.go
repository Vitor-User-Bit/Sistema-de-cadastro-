package main

import (
	"errors"
	"fmt"
	"strings"
)

type Usuario struct{
	Nome string
	Email string
	Idade int
}

func validarNome (nome string) error{

	//remove espaços no começo e no final
	nome = strings.TrimSpace(nome)

	if nome == ""{
		return errors.New("O nome não pode estar vazio")
	}
	if len(nome) < 3 {
		return errors.New("O nome deve possuir pelo menos 3 caracteres")
	}
	return nil
}

func validarEmail (email string) error{

	//remove espaços no começo e no final
	email = strings.TrimSpace(email)

	if email == ""{
		return errors.New("O email não pode estar vazio")
	}
	if !strings.Contains(email, "@") {
		return errors.New("O email deve possuir o caractere @")
	}
		if !strings.Contains(email, ".") {
		return errors.New("O email deve possuir um domínio valido")
	}
	return nil
}

func validarIdade (idade int) error{

	if idade <= 0{
		return errors.New("A idade não pode ser 0 ou menor!")
	}	
	if idade > 130{
		return errors.New("Idade invalida")
	}
	return nil
}
func validarUsuario (usuario Usuario) error{

	err := validarNome(usuario.Nome)
	if err != nil {
		return err
	}

		err = validarEmail(usuario.Email)
	if err != nil {
		return err
	}

		err = validarIdade(usuario.Idade)
	if err != nil {
		return err
	}
	return nil
}
func cadastrarUsuario(usuario Usuario) error{
	err := validarUsuario(usuario)
	if err != nil{
		return err
	}

	return nil
}

func main(){

	usuario := Usuario{
		Nome: "João",
		Email: "joao@email.com",
		Idade: 180,
	}
	err := cadastrarUsuario(usuario)

	if err != nil{
		fmt.Println("Erro no cadastro: ", err)
	}else{
		fmt.Println("Cadastro realizado")
	}

}
