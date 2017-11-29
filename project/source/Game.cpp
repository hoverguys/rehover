#include "Game.h"
#include "behaviours/Hovercraft.h"
#include "components/Camera.h"
#include "components/Light.h"
#include "components/Renderable.h"
#include "components/Transform.h"
#include "input/HovercraftController.h"

#include "systems/BehaviourSystem.h"
#include "systems/RenderSystem.h"
#include "systems/PhysicsSystem.h"

#include "rendering/Material.h"
#include "resources/MeshResource.h"
#include "resources/ResourceLoader.h"
#include "resources/ShaderResource.h"
#include "resources/TextureResource.h"

namespace cp = Components;
namespace bh = Behaviours;

Game::Game() {
	systems.add<InputSystem>();
	systems.add<BehaviourSystem<bh::Hovercraft>>();
	systems.add<PhysicsSystem>();
	systems.add<RenderSystem>();
	systems.configure();
}

void Game::init() {
	// Camera
	auto camera = entities.create();
	camera.assign<cp::Transform>(cp::Transform({0, 3, 10}));
	camera.assign<cp::Camera>(cp::Camera());

	// Controller (for hovercraft)
	auto input = systems.system<InputSystem>();
	auto controller = std::make_shared<GCHovercraftController>(input->GetController(0));

	// Hovercraft
	auto hovercraftRes = ResourceLoader::Load<MeshResource>("models/hovercraft.obj");
	auto hovercraftMesh = hovercraftRes->Load();

	auto hovercraftDiffRes = ResourceLoader::Load<TextureResource>("textures/hovercraftGlobal.png");
	auto hovercraftDiffTex = hovercraftDiffRes->Load();
	auto hovercraftShadeRes = ResourceLoader::Load<TextureResource>("textures/hovercraftShade.png");
	auto hovercraftShadeTex = hovercraftShadeRes->Load();
	auto hovercraftShaderBin = ResourceLoader::Load<ShaderResource>("shaders/hovercraft.bin");
	auto hovercraftShader = hovercraftShaderBin->Load();
	auto hovercraftMat = std::make_shared<Material>();
	hovercraftMat->textures = {hovercraftDiffTex, hovercraftShadeTex};
	hovercraftMat->shader = hovercraftShader;
	hovercraftMat->uniforms.color0 = {0xff, 0x29, 0x5b, 0xff};

	auto hovercraft = entities.create();
	hovercraft.assign<cp::Transform>(cp::Transform({10, 0, 0}));
	hovercraft.assign<cp::Renderable>(cp::Renderable(hovercraftMesh, hovercraftMat));
	hovercraft.assign<bh::Hovercraft>(bh::Hovercraft{controller, camera});

	// Terrain
	auto terrainRes = ResourceLoader::Load<MeshResource>("models/terrain.obj");
	auto terrainMesh = terrainRes->Load();
	auto terrainTexRes = ResourceLoader::Load<TextureResource>("textures/terrain.png");
	auto terrainTex = terrainTexRes->Load();
	auto terrainMat = std::make_shared<Material>();
	terrainMat->textures = {terrainTex};
	auto terrain = entities.create();
	terrain.assign<cp::Transform>(cp::Transform({0, -1, 0}));
	terrain.assign<cp::Renderable>(cp::Renderable(terrainMesh, terrainMat));

	// DEBUG: Load hardcoded model
	auto planeRes = ResourceLoader::Load<MeshResource>("models/plane.obj");
	auto planeMesh = planeRes->Load();

	auto checkerRes = ResourceLoader::Load<TextureResource>("textures/checkerboard.png");
	auto checkerTex = checkerRes->Load();
	auto checkerMat = std::make_shared<Material>();
	checkerMat->textures = {checkerTex};

	// Plane / Sea
	auto plane = entities.create();
	plane.assign<cp::Transform>(cp::Transform({0, 0, 0}))->scale = {10, 10, 10};
	plane.assign<cp::Renderable>(cp::Renderable(planeMesh, checkerMat));

	// Light
	auto light = entities.create();
	light.assign<cp::Transform>(cp::Transform({0, 0, 0}))->Lookat({0, -1, -0.5f});
	light.assign<cp::DirectionalLight>(cp::DirectionalLight({0xff, 0xee, 0xee, 0xff}, 0));
}

void Game::update(ex::TimeDelta dt) {
	systems.update<InputSystem>(dt);
	systems.update<BehaviourSystem<bh::Hovercraft>>(dt);
	systems.update<PhysicsSystem>(dt);
	systems.update<RenderSystem>(dt);
}