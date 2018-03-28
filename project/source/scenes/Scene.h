#pragma once
#include "../pchheader.h"
#include "../systems/SceneSystem.h"

namespace ex = entityx;

template <typename T>
struct Scene {
public:
	static ex::Entity create(bool tag = true) {
		auto entity = SceneSystem::manager->entities.create();
		if (tag) {
			entity.assign<T>();
		}
		return entity;
	}

	template <typename S>
	static std::shared_ptr<S> system() {
		return SceneSystem::manager->systems.system<S>();
	}

	static void unload() {
		SceneSystem::manager->entities.each<T>([](ex::Entity entity, T &marker) {
			entity.destroy();
		});
	}
private:
	Scene() = delete;
};