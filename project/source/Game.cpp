#include "Game.h"
#include "rendering/RenderSystem.h"

Game::Game() {
    systems.add<RenderSystem>();
    systems.configure();
}

void Game::init(Mesh* mesh) {
    ex::Entity hovercaft = entities.create();
    hovercaft.assign<Mesh*>(mesh);
}

void Game::update(ex::TimeDelta dt) {
    systems.update_all(dt);
}