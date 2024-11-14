// Configuração base da API
const API_BASE_URL = 'http://localhost:8080/api';

// Objeto para gerenciar as chamadas à API
const BankAPI = {
    // Autenticação
    async login(email, password) {
        try {
            const response = await fetch(`${API_BASE_URL}/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email, password })
            });
            return await response.json();
        } catch (error) {
            console.error('Erro no login:', error);
            throw error;
        }
    },

    // Criar nova conta
    async createAccount(userData) {
        try {
            const response = await fetch(`${API_BASE_URL}/accounts`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(userData)
            });
            return await response.json();
        } catch (error) {
            console.error('Erro ao criar conta:', error);
            throw error;
        }
    },

    // Obter saldo
    async getBalance(accountId, token) {
        try {
            const response = await fetch(`${API_BASE_URL}/accounts/${accountId}/balance`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });
            return await response.json();
        } catch (error) {
            console.error('Erro ao obter saldo:', error);
            throw error;
        }
    },

    // Realizar transferência
    async makeTransfer(transferData, token) {
        try {
            const response = await fetch(`${API_BASE_URL}/transfers`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify(transferData)
            });
            return await response.json();
        } catch (error) {
            console.error('Erro na transferência:', error);
            throw error;
        }
    }
}; 