#include "Game.h"
#include "systems/RenderSystem.h"
#include "components/Renderable.h"

Game::Game() {
    systems.add<RenderSystem>();
    systems.configure();
}

void Game::init(Mesh* mesh) {
    ex::Entity hovercaft = entities.create();
    hovercaft.assign<Components::Renderable>(Components::Renderable(mesh);
}

void Game::update(ex::TimeDelta dt) {
    systems.update_all(dt);
}