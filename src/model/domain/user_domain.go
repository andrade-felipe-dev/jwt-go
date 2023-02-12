package domain

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetEmail() string
	GetSenha() string
	GetTipo() string
	GetNome() string
	GetID() string

	SetID(string)

	EncryptPassword()
}

func NewUserDomain(email, senha, nome, tipo string) UserDomainInterface {
	return &userDomain{
		email: email,
		senha: senha,
		nome:  nome,
		tipo:  tipo,
	}
}

type userDomain struct {
	id    string
	email string
	senha string
	nome  string
	tipo  string
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetSenha() string {
	return ud.senha
}

func (ud *userDomain) GetNome() string {
	return ud.nome
}

func (ud *userDomain) GetTipo() string {
	return ud.tipo
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}

func (ud *userDomain) GetID() string {
	return ud.id
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.senha))
	ud.senha = hex.EncodeToString(hash.Sum(nil))
}
