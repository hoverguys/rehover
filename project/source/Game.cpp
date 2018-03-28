#include "Game.h"

#include "resources/ResourceLoader.h"

#include "behaviours/Hovercraft.h"
#include "systems/SceneSystem.h"
#include "systems/BehaviourSystem.h"
#include "systems/PhysicsSystem.h"
#include "systems/RenderSystem.h"
#include "systems/UISystem.h"
#include "systems/InputSystem.h"

#include "scenes/GameScene.h"

namespace bh = Behaviours;

Game::Game() {
	systems.add<InputSystem>();
	systems.add<BehaviourSystem<bh::Hovercraft>>();
	systems.add<PhysicsSystem>();
	systems.add<RenderSystem>();
	systems.add<UISystem>();
	systems.configure();
}

void Game::init() {
	SceneSystem::initialize(this);
	std::printf("Loading Game scene\n");
	GameScene::load();
	ResourceLoader::PrintUsage();
}

void Game::update(ex::TimeDelta dt) {
	systems.update<InputSystem>(dt);
	systems.update<BehaviourSystem<bh::Hovercraft>>(dt);
	systems.update<PhysicsSystem>(dt);
	systems.update<RenderSystem>(dt);
	systems.update<UISystem>(dt);
}