#include "Game.h"
#include "behaviours/Hovercraft.h"
#include "components/Camera.h"
#include "components/Renderable.h"
#include "components/Transform.h"
#include "input/HovercraftController.h"
#include "systems/RenderSystem.h"

#include "resources/MeshResource.h"
#include "resources/ResourceLoader.h"

namespace cp = Components;
namespace bh = Behaviours;

Game::Game() {
	systems.add<InputSystem>();
	systems.add<bh::HovercraftSystem>();
	systems.add<RenderSystem>();
	systems.configure();
}

void Game::init(Mesh* mesh) {
	// Camera
	auto camera = entities.create();
	camera.assign<cp::Transform>(cp::Transform({0, 3, 10}));
	camera.assign<cp::Camera>(cp::Camera());

	// Controller (for hovercraft)
	auto input = systems.system<InputSystem>();
	auto controller = std::make_shared<GCHovercraftController>(input->GetController(0));

	// Hovercraft
	auto hovercraft = entities.create();
	hovercraft.assign<cp::Transform>(cp::Transform({0, 0, 0}));
	hovercraft.assign<cp::Renderable>(cp::Renderable(mesh));
	hovercraft.assign<bh::Hovercraft>(bh::Hovercraft{controller, camera});

	// DEBUG: Load hardcoded model
	auto terrainres = ResourceLoader::Load<MeshResource>("models/plane.obj");
	Mesh* terrainmesh = terrainres->Load();

	// Static
	auto staticHovercraft = entities.create();
	staticHovercraft.assign<cp::Transform>(cp::Transform({0, 0, 0}));
	staticHovercraft.assign<cp::Renderable>(cp::Renderable(mesh));

	// Terrain
	auto terrain = entities.create();
	auto terrainTransform = terrain.assign<cp::Transform>(cp::Transform({0, 0, 0}));
	terrain.assign<cp::Renderable>(cp::Renderable(terrainmesh));
	terrainTransform->scale = {10, 10, 10};
}

void Game::update(ex::TimeDelta dt) { systems.update_all(dt); }