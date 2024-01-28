// Função simulada para obter produtos (substitua isso com uma chamada real ao seu servidor ou API)
function getProducts() {
    return [
        { name: 'Pizza Margherita', price: 15.99 },
        { name: 'Hamburguer Clássico', price: 8.99 },
        { name: 'Salada Caesar', price: 7.49 },
        { name: 'Sushi Misto', price: 18.99 },
        { name: 'Spaghetti Bolognese', price: 12.99 }
    ];
}

// Função para renderizar os produtos na página
function renderMenu() {
    const menuContainer = document.getElementById('menu-container');
    const products = getProducts();

    products.forEach(product => {
        const menuItem = document.createElement('div');
        menuItem.classList.add('menu-item');
        menuItem.innerHTML = `<h2>${product.name}</h2><p>R$ ${product.price.toFixed(2)}</p>`;
        menuContainer.appendChild(menuItem);
    });
}

// Chama a função de renderização ao carregar a página
window.onload = renderMenu;
