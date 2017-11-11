#include "Game.h"
#include "components/Camera.h"
#include "components/Renderable.h"
#include "components/Transform.h"
#include "systems/RenderSystem.h"

namespace cp = Components;

Game::Game() {
	systems.add<RenderSystem>();
	systems.configure();
}

void Game::init(Mesh* mesh) {
	// Hovercraft
	hovercraft = entities.create();
	hovercraft.assign<cp::Transform>(cp::Transform({0, 0, 0}));
	hovercraft.assign<cp::Renderable>(cp::Renderable(mesh));

	// Camera
	ex::Entity camera = entities.create();
	camera.assign<cp::Transform>(cp::Transform({0, 0, -10}));
	camera.assign<cp::Camera>(cp::Camera());
}

void Game::update(ex::TimeDelta dt) {
	systems.update_all(dt);

	ex::ComponentHandle<cp::Transform> position = hovercraft.component<cp::Transform>();
	position->position.x += 0.2f * dt;
}