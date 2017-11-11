#include "Game.h"
#include "components/Camera.h"
#include "components/Renderable.h"
#include "components/Transform.h"
<<<<<<< HEAD
#include "systems/RenderSystem.h"
=======
#include "behaviours/Hovercraft.h"
>>>>>>> Move hovecraft behaviour into seperate class

namespace cp = Components;
namespace bh = Behaviours;

Game::Game() {
<<<<<<< HEAD
	systems.add<RenderSystem>();
	systems.configure();
}

void Game::init(Mesh* mesh) {
	// Hovercraft
	hovercraft = entities.create();
	hovercraft.assign<cp::Transform>(cp::Transform({0, 0, 0}));
	hovercraft.assign<cp::Renderable>(cp::Renderable(mesh));
=======
    systems.add<bh::HovercraftSystem>();
    systems.add<RenderSystem>();
    systems.configure();
}

void Game::init(Mesh* mesh) {
    // Hovercraft
    hovercraft = entities.create();
    hovercraft.assign<cp::Transform>(cp::Transform({0,0,0}));
    hovercraft.assign<cp::Renderable>(cp::Renderable(mesh));
    hovercraft.assign<bh::Hovercraft>(bh::Hovercraft());
>>>>>>> Move hovecraft behaviour into seperate class

	// Camera
	ex::Entity camera = entities.create();
	camera.assign<cp::Transform>(cp::Transform({0, 0, -10}));
	camera.assign<cp::Camera>(cp::Camera());
}

void Game::update(ex::TimeDelta dt) {
<<<<<<< HEAD
	systems.update_all(dt);

	ex::ComponentHandle<cp::Transform> position = hovercraft.component<cp::Transform>();
	position->position.x += 0.2f * dt;
=======
    systems.update_all(dt);
>>>>>>> Move hovecraft behaviour into seperate class
}