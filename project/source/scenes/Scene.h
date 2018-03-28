#pragma once
#include "../pchheader.h"
#include "../systems/SceneSystem.h"

namespace ex = entityx;

template <typename T>
struct Scene {
public:
	static ex::Entity create() {
        auto entity = SceneSystem::manager->entities.create();
        entity.assign<T>();
        return entity;
    }

    template <typename S>
    static std::shared_ptr<S> system() {
        return SceneSystem::manager->systems.system<S>();
    }
private:
    Scene() = delete;
};