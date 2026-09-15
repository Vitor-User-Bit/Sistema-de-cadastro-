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

//---------------validando nome --------------------------------------

func validarNome (nome string) error{

	//remove espaços no começo e no final
	nome = strings.TrimSpace(nome)



	//se o usuario sem digitar nada ele vai retornar esse "errors.New" (pacote errors)
	if nome == ""{
		return errors.New("O nome não pode estar vazio")
	}
	// se o usuario tentar cadastrar um nome com menos de 3 caracteres e retorna o erro
	if len(nome) < 3 {									
		return errors.New("O nome deve possuir pelo menos 3 caracteres")
	}

	//se meu nome passou por essas condicoes ela vai me retornar o (return nil) = nao existe erro
	return nil
}

// ------------------------- validando email -----------------------------

func validarEmail (email string) error{

	//remove espaços no começo e no final
	email = strings.TrimSpace(email)

//se o usuario sem digitar nada ele vai retornar esse "errors.New" (pacote errors)
	if email == ""{
		return errors.New("O email não pode estar vazio")
	}

//se nao tiver string NAO conter o @ ele retorna o erro (o ! significa um NAO)
	if !strings.Contains(email, "@") {
		return errors.New("O email deve possuir o caractere @")
	}
		if !strings.Contains(email, ".") {
		return errors.New("O email deve possuir um domínio valido")
	}
	return nil
}

//------------------------ validando idade ----------------------------

func validarIdade (idade int) error{

	// se idade for menor ou igual a 0 ele retorna o erro
	if idade <= 0{
		return errors.New("A idade não pode ser 0 ou menor!")
	}	
	// se idade for maior que 130 ele retorna o erro
	if idade > 130{
		return errors.New("Idade invalida")
	}
	return nil
}

//-----------------------validando usuario---------------------------------
// ---------------- VAI VALIDAR OS 3 ITENS DE UMA VEZ ------------------
func validarUsuario (usuario Usuario) error{


	// VALIDAR O NOME, SE O ERRO FOR DIFERENTE DE NIL ELE VAI RETORNAR O ERRO ESCRITO NA FUNC PRINCIPAL 
	err := validarNome(usuario.Nome)
	if err != nil {
		return err
	}
// VALIDAR O EMAIL, SE O ERRO FOR DIFERENTE DE NIL ELE VAI RETORNAR O ERRO ESCRITO NA FUNC PRINCIPAL
		err = validarEmail(usuario.Email)
	if err != nil {
		return err
	}
// VALIDAR A IDADE, SE O ERRO FOR DIFERENTE DE NIL ELE VAI RETORNAR O ERRO ESCRITO NA FUNC PRINCIPAL
		err = validarIdade(usuario.Idade)
	if err != nil {
		return err
	}
	return nil
}

//---------------------- CADASTRANDO O USUARIO --------------------------
// VOU USAR A FUNC VALIDAR USUARIO PARA CONFERIR SE EXISTE ALGUM ERRO
func cadastrarUsuario(usuario Usuario) error{
	err := validarUsuario(usuario)
	if err != nil{
		return err
	}
	return nil
}

func main(){

//--------------------- CRIANDO O USUARIO ----------------------

	usuario := Usuario{
		Nome: "João",
		Email: "joao@email.com",
		Idade: 25,
	}

// --------------------- CADASTRANDO O USUARIO ------------------
// VAI VERIFICAR SE NAO HOUVE NENHUM ERRO
	err := cadastrarUsuario(usuario)

	if err != nil{
		fmt.Println("Erro no cadastro: ", err)
	}else{
		fmt.Println("Cadastro realizado")
	}

}
