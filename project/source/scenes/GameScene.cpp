#include "GameScene.h"

#include "../behaviours/Hovercraft.h"
#include "../components/Camera.h"
#include "../components/Light.h"
#include "../components/MeshCollider.h"
#include "../components/Renderable.h"
#include "../components/Rigidbody.h"
#include "../components/Sprite.h"
#include "../components/Transform.h"
#include "../input/HovercraftController.h"

#include "../systems/BehaviourSystem.h"
#include "../systems/InputSystem.h"

#include "../rendering/Material.h"
#include "../resources/MeshResource.h"
#include "../resources/ResourceLoader.h"
#include "../resources/ShaderResource.h"
#include "../resources/TextureResource.h"

namespace cp = Components;
namespace bh = Behaviours;

void GameScene::load() {
	// Camera
	auto camera = create();
	camera.assign<cp::Transform>();
	camera.assign<cp::Camera>(cp::Camera());

	// Controller (for hovercraft)
	auto input = system<InputSystem>();
	auto controller = std::make_shared<GCHovercraftController>(input->GetController(0));

	// Hovercraft
	auto hovercraftRes = ResourceLoader::Load<MeshResource>("models/hovercraft.obj");
	auto hovercraftMesh = hovercraftRes->Load();

	auto hovercraftDiffRes = ResourceLoader::Load<TextureResource>("textures/hovercraftGlobal.png");
	auto hovercraftDiffTex = hovercraftDiffRes->Load();
	auto hovercraftShadeRes = ResourceLoader::Load<TextureResource>("textures/hovercraftShade.png");
	auto hovercraftShadeTex = hovercraftShadeRes->Load();
	auto hovercraftShaderBin = ResourceLoader::Load<ShaderResource>("shaders/hovercraft.tev");
	auto hovercraftShader = hovercraftShaderBin->Load();
	auto hovercraftMat = std::make_shared<Material>();
	hovercraftMat->textures = {hovercraftDiffTex, hovercraftShadeTex};
	hovercraftMat->shader = hovercraftShader;
	hovercraftMat->uniforms.color0 = {0xff, 0x29, 0x5b, 0xff};

	auto hovercraft = create();
	hovercraft.assign<cp::Transform>(cp::Transform({0, 1, 0}));
	hovercraft.assign<cp::Renderable>(cp::Renderable(hovercraftMesh, hovercraftMat));
	hovercraft.assign<bh::Hovercraft>(bh::Hovercraft{controller, camera});
	hovercraft.assign<cp::Rigidbody>();

	// Terrain
	auto terrainRes = ResourceLoader::Load<MeshResource>("models/testmap.obj");
	auto terrainMesh = terrainRes->Load();
	auto terrainTexRes = ResourceLoader::Load<TextureResource>("textures/testmap.png");
	auto terrainTex = terrainTexRes->Load();
	auto terrainMat = std::make_shared<Material>();
	terrainMat->textures = {terrainTex};

	auto terrain = create();
	terrain.assign<cp::Transform>(cp::Transform({0, 0, 0}));
	terrain.assign<cp::Renderable>(cp::Renderable(terrainMesh, terrainMat));
	terrain.assign<cp::MeshCollider>(cp::MeshCollider(terrainMesh));

	// Light
	auto light = create();
	light.assign<cp::Transform>()->Lookat({0, -1, -0.5f});
	light.assign<cp::DirectionalLight>(cp::DirectionalLight({0xff, 0xee, 0xee, 0xff}, 0));

	// Timer sprite
	auto timer = create();
	auto timerTexRes = ResourceLoader::Load<TextureResource>("generated/timerfont.png");
	auto timerTex = timerTexRes->Load();
	auto fontShaderBin = ResourceLoader::Load<ShaderResource>("shaders/font.tev");
	auto fontShader = fontShaderBin->Load();
	auto timerMat = std::make_shared<Material>();
	timerMat->textures = {timerTex};
	timerMat->shader = fontShader;
	timerMat->uniforms.color0 = GXColor{0xff, 0, 0, 0xff};
	// auto timerAtlasRes = ResourceLoader::Load<AtlasResource>("generated/timer.png.atlas");
	// auto timerAtlas = timerAtlasRes->Load();
	timer.assign<cp::Transform>(cp::Transform({10, 10, -1}));
	timer.assign<cp::Sprite>(cp::Sprite(Vector2D(128, 128), timerMat)); // TODO Make Atlas work (also nicer to use)
}