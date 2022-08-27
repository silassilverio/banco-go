package contas

import "github.com/Banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
} // letras maiusculas tornam os atributos visiveis a outros arquivos

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	pordeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if pordeSacar {
		c.saldo -= valorDoSaque
		return "Saque Realizado com Sucesso!"
	} else {
		return "Saldo Insuficiente!"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito menor que zero", c.saldo
	}

}

func (c *ContaCorrente) Tranferir(valorDaTrasferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTrasferencia <= c.saldo && valorDaTrasferencia > 0 {
		c.saldo -= valorDaTrasferencia
		contaDestino.Depositar(valorDaTrasferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
