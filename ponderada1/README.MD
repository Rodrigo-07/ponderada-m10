

internal/
├── handler/      ← camada de entrada (HTTP / gRPC)
├── service/      ← regras de negócio (lógica da aplicação)
├── client/       ← integrações externas (APIs, gRPC, etc.)
├── model/        ← structs e tipos do domínio


handler  →  service  →  client
     \           \       \
      →----------→--------→ model

sudo PATH=$PATH make dev

http://localhost:8080/swagger/index.html


--------

- Endpoint para começar jogo
  - Opção de singleplayer e multiplayer (2 players)
  - Chamo a api de criar um novo deck
  - Salvo no banco
    - Singleplayer
      - Id do jogo
      - Nome do jogador
      - Vitoria ou derrota
      - Id do deck
      - Cartas puxadas
      - Movimentos
    - Multplayers
      - Id do jogo
      - Nome do jogador 1
      - Nome do jogador 2
      - Quem venceu ou foi empate
      - Id do deck do player 1
      - Id do deck do player 2
      - Cartas puxadas player 1
      - Cartas puxadas player 2
      - Movimentos player 1
      - Movimentos player 2
- Endpoint para puxar uma carta singleplayer (Id do jogo)
  - Devolve:
    - Carta puxada
    - Jogo acabou ou não
    - Cartas que foram puxadas
    - Movimentos
    - Valor que falta para 21
- Endpoint para puxar uma carta multiplayer (Id do jogo, jogador)
  - Devolve:
    - Carta puxada
    - Jogo acabou ou não
    - Cartas que foram puxadas
    - Movimentos
    - Valor que falta para 21

Idea=buff de um jogador poder colocar uma carta na mão do oponente