#pragma once
#include "../pchheader.h"
#include "Scene.h"

struct GameSceneMarker {};
struct GameScene : public Scene<GameSceneMarker> {
public:
    static void load();
};