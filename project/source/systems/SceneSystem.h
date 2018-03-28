#pragma once
#include "../pchheader.h"

namespace ex = entityx;

class SceneSystem  {
public:
	static void initialize(ex::EntityX* manager) {
		SceneSystem::manager = manager;
	}
protected:
	template<typename T>
	friend class Scene;
	
	static ex::EntityX* manager;
};