#include "Game.h"

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
	GameScene::load();
}

void Game::update(ex::TimeDelta dt) {
	systems.update<InputSystem>(dt);
	systems.update<BehaviourSystem<bh::Hovercraft>>(dt);
	systems.update<PhysicsSystem>(dt);
	systems.update<RenderSystem>(dt);
	systems.update<UISystem>(dt);
}