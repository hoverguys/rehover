#pragma once
#include <entityx/entityx.h>
#include "rendering/Mesh.h"

namespace ex = entityx;

class Game : public ex::EntityX {
public:
  explicit Game();
  void init(Mesh* mesh);
  void update(ex::TimeDelta dt);
};