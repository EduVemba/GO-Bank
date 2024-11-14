document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.getElementById('login-form');
    const transferForm = document.getElementById('transfer-form');
    const userArea = document.getElementById('user-area');
    const loginSection = document.getElementById('login-section');
    
    // Verificar se já existe um token salvo
    const token = localStorage.getItem('token');
    if (token) {
        showUserArea();
        updateBalance();
    }

    // Gerenciar login
    loginForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;

        try {
            const response = await BankAPI.login(email, password);
            if (response.token) {
                localStorage.setItem('token', response.token);
                localStorage.setItem('accountId', response.accountId);
                showUserArea();
                updateBalance();
            }
        } catch (error) {
            alert('Erro no login. Verifique suas credenciais.');
        }
    });

    // Gerenciar transferências
    transferForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const destinationAccount = document.getElementById('destination-account').value;
        const amount = document.getElementById('amount').value;
        const token = localStorage.getItem('token');

        try {
            await BankAPI.makeTransfer({
                destination_account_id: destinationAccount,
                amount: parseFloat(amount)
            }, token);
            
            alert('Transferência realizada com sucesso!');
            updateBalance();
            transferForm.reset();
        } catch (error) {
            alert('Erro ao realizar transferência.');
        }
    });

    // Funções auxiliares
    async function updateBalance() {
        const token = localStorage.getItem('token');
        const accountId = localStorage.getItem('accountId');
        
        try {
            const balance = await BankAPI.getBalance(accountId, token);
            document.getElementById('current-balance').textContent = 
                `R$ ${balance.amount.toFixed(2)}`;
        } catch (error) {
            console.error('Erro ao atualizar saldo:', error);
        }
    }

    function showUserArea() {
        loginSection.classList.add('hidden');
        userArea.classList.remove('hidden');
    }
}); 