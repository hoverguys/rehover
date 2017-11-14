#include "Game.h"
#include "behaviours/Hovercraft.h"
#include "components/Camera.h"
#include "components/Renderable.h"
#include "components/Transform.h"
#include "input/HovercraftController.h"
#include "systems/BehaviourSystem.h"
#include "systems/RenderSystem.h"

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
	auto hovercraftShaderBin = ResourceLoader::Load<ShaderResource>("shader/hovercraft.bin");
	auto hovercraftShader = hovercraftShaderBin->Load();
	auto hovercraftMat = std::make_shared<Material>();
	hovercraftMat->textures = {hovercraftDiffTex, hovercraftShadeTex};
	hovercraftMat->shader = hovercraftShader;
	hovercraftMat->uniforms = {
	    {0xff, 0x29, 0x5b, 0xff} // Color 0
	};

	auto hovercraft = entities.create();
	hovercraft.assign<cp::Transform>(cp::Transform({0, 0, 0}));
	hovercraft.assign<cp::Renderable>(cp::Renderable(hovercraftMesh, hovercraftMat));
	hovercraft.assign<bh::Hovercraft>(bh::Hovercraft{controller, camera});

	// Static
	auto staticHovercraft = entities.create();
	staticHovercraft.assign<cp::Transform>(cp::Transform({0, 0, 0}));
	staticHovercraft.assign<cp::Renderable>(cp::Renderable(hovercraftMesh));
	auto statichovercraftMat = std::make_shared<Material>();
	statichovercraftMat->textures = {hovercraftDiffTex, hovercraftShadeTex};
	statichovercraftMat->shader = hovercraftShader;
	statichovercraftMat->uniforms = {
	    {0xac, 0xff, 0x29, 0xff} // Color 0
	};

	// DEBUG: Load hardcoded model
	auto terrainRes = ResourceLoader::Load<MeshResource>("models/plane.obj");
	auto terrainMesh = terrainRes->Load();

	auto checkerRes = ResourceLoader::Load<TextureResource>("textures/checkerboard.png");
	auto checkerTex = checkerRes->Load();
	auto checkerMat = std::make_shared<Material>();
	checkerMat->textures = {checkerTex};

	// Terrain
	auto terrain = entities.create();
	terrain.assign<cp::Transform>(cp::Transform({0, 0, 0}))->scale = {10, 10, 10};
	terrain.assign<cp::Renderable>(cp::Renderable(terrainMesh, checkerMat));
}

void Game::update(ex::TimeDelta dt) { systems.update_all(dt); }